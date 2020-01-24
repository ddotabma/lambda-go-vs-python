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

type MyEvent struct {
	//Name string `json:"name"`
}

type Data struct {
	Datetime string `json: datetime` // name conversion in struct :D
	Values   []int  `json: values`
}

func Deserialize(r http.Response) Data {
	defer r.Body.Close()
	data := Data{}
	_ = json.NewDecoder(r.Body).Decode(&data) // Deserialize response json
	return data
}

func Get(wg *sync.WaitGroup) {
	r, _ := http.Get(os.Getenv("API"))

	fmt.Println(Deserialize(*r))
	println("Done")
	wg.Done()
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	wg := sync.WaitGroup{}
	numRequests := 100
	for i := 0; i < numRequests; i++ { // make 1000 requests
		wg.Add(1)   // increment counter
		go Get(&wg) // <<<=== Run Get concurrently
	}
	wg.Wait() // Wait till all requests to be finished
	return "Processed requests", nil
}

func main() {
	//HandleRequest(nil, MyEvent{})
	lambda.Start(HandleRequest)
}
