package main

import (
	"context"
	"github.com/xitongsys/parquet-go-source/s3"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"
	"log"
	"time"
)

type user struct {
	ID        string    `parquet:"name=id, type=UTF8, encoding=PLAIN_DICTIONARY"`
	FirstName string    `parquet:"name=firstname, type=UTF8, encoding=PLAIN_DICTIONARY"`
	LastName  string    `parquet:"name=lastname, type=UTF8, encoding=PLAIN_DICTIONARY"`
	Email     string    `parquet:"name=email, type=UTF8, encoding=PLAIN_DICTIONARY"`
	Phone     string    `parquet:"name=phone, type=UTF8, encoding=PLAIN_DICTIONARY"`
	Blog      string    `parquet:"name=blog, type=UTF8, encoding=PLAIN_DICTIONARY"`
	Username  string    `parquet:"name=username, type=UTF8, encoding=PLAIN_DICTIONARY"`
	Score     float64   `parquet:"name=score, type=DOUBLE"`
	CreatedAt time.Time //wont be saved in the parquet file
}

const recordNumber = 10000

func main() {
	var data []*user
	//create fake data
	for i := 0; i < recordNumber; i++ {
		u := &user{
			//ID:        faker.UUIDDigit(),
			//FirstName: faker.FirstName(),
			//LastName:  faker.LastName(),
			//Email:     faker.Email(),
			//Phone:     faker.Phonenumber(),
			//Blog:      faker.URL(),
			//Username:  faker.Username(),
			//Score:     float64(i),
			//CreatedAt: time.Now(),
		}
		data = append(data, u)
	}
	err := generateParquet(data)
	if err != nil {
		log.Fatal(err)
	}

}

func generateParquet(data []*user) error {
	log.Println("generating parquet file")

	ctx := context.Background()
	bucket := "bdr-go-blog"
	key := "foobar.parquet"
	fw, err := s3.NewS3FileWriter(ctx, bucket, key, nil)

	pw, err := writer.NewParquetWriter(fw, new(user), int64(len(data)))
	if err != nil {
		log.Fatal(err)
	}
	pw.CompressionType = parquet.CompressionCodec_SNAPPY

	for _, d := range data {
		if err = pw.Write(d); err != nil {
			return err
		}
	}
	if err = pw.WriteStop(); err != nil {
		return err
	}
	err = fw.Close()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
