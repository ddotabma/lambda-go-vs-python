package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/xitongsys/parquet-go-source/s3"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/source"
	"github.com/xitongsys/parquet-go/writer"
	"log"
	"net/http"
	"os"
	"sync"
)

var api = os.Getenv("API")

type Data struct {
	Datetime string  `json:"datetime" parquet:"name=datetime, type=UTF8"` // Name conversion
	Values   []int64 `json:"values" parquet:"name=values, type=LIST, valuetype=INT64"`
}

func DataParquetWriter(datas chan Data) (*writer.ParquetWriter, source.ParquetFile, error) {
	log.Println("generating parquet file")
	ctx := context.Background()
	bucket := "bdr-go-blog"
	key := "foobar.parquet"
	fileWriter, err := s3.NewS3FileWriter(ctx, bucket, key, nil)
	parquetWriter, err := writer.NewParquetWriter(fileWriter, new(Data), int64(cap(datas)))
	if err != nil {
		log.Fatal(err)
	}
	parquetWriter.CompressionType = parquet.CompressionCodec_SNAPPY

	return parquetWriter, fileWriter, nil
}

func AddDataToParquet(wg *sync.WaitGroup, pw *writer.ParquetWriter, datas chan Data) {
	data := <-datas
	fmt.Println(data)
	if err := pw.Write(data); err != nil {
		log.Fatal(err)
	}
	wg.Done()
}

func CloseParquetWrite(pw *writer.ParquetWriter, fw source.ParquetFile) error {
	if err := pw.WriteStop(); err != nil {
		return err
	}

	err := fw.Close()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func Deserialize(r http.Response, datas chan<- Data) Data {
	data := Data{}
	_ = json.NewDecoder(r.Body).Decode(&data) // Deserialize body
	datas <- data
	_ = r.Body.Close()
	return data
}

func Get(wg *sync.WaitGroup, datas chan<- Data) {
	defer wg.Done()       // Communicate end of goroutine
	r, _ := http.Get(api) // Make get request
	fmt.Println(r.Status)
	Deserialize(*r, datas) // Print body

}

func Handler(_ context.Context, _ interface{}) (string, error) {
	numberOfRequests := 100
	wgGet := sync.WaitGroup{}

	datas := make(chan Data, 100)

	for i := 0; i < numberOfRequests; i++ { // make 1000 requests
		wgGet.Add(1)
		go Get(&wgGet, datas) // <<<=== Run Get concurrently
	}
	wgGet.Wait() // Wait until all responses obtained

	wgParquet := sync.WaitGroup{}
	parquetWriter, fileWriter, err := DataParquetWriter(datas)
	for i := 0; i < numberOfRequests; i++ {
		wgParquet.Add(1)
		go AddDataToParquet(&wgParquet, parquetWriter, datas)
	}
	wgParquet.Wait()
	err = CloseParquetWrite(parquetWriter, fileWriter)

	if err != nil {
		log.Fatal(err)
	}

	return "Processed requests", nil
}

func main() {
	//_, _ = Handler(nil, nil)
	lambda.Start(Handler)
}
