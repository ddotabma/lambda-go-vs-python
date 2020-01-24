package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"os"
	"sync"
)

var api = os.Getenv("API")


type Data struct {
	Datetime string `json: datetime` // Name conversion
	Values   []int  `json: values`
}

func Deserialize(r http.Response) Data {
	data := Data{}
	_ = json.NewDecoder(r.Body).Decode(&data) // Deserialize body
	_ = r.Body.Close()
	return data
}

func Get(wg *sync.WaitGroup) {
	r, _ := http.Get(api) 			// Make get request
	fmt.Println(Deserialize(*r))    // Print body
	wg.Done() 						// Communicate end of goroutine
}

func Handler(_ context.Context, _ interface{}) (string, error) {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {		// make 1000 requests
	wg.Add(1) 				// increment counter
		go Get(&wg)					// <<<=== Run Get concurrently
	}
	wg.Wait() 						// Wait untill all responses obtained
	return "Processed requests", nil
}

func main() {
	//Handler(nil, nil)
	lambda.Start(Handler)
}
