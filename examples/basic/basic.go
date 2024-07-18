package main

import (
	"fmt"

	"github.com/HawAPI/go-sdk/pkg/hawapi"
)

func main() {
	// Create a new client with default options
	client := hawapi.NewClient()

	// Override options
	client.WithOpts(hawapi.Options{
		Endpoint: "http://localhost:8080/api",
		// When using 'WithOpts' or 'NewClientWithOpts' the value of
		// 'UseInMemoryCache' will be set to false
		UseInMemoryCache: true,
	})

	res, err := client.ListActors()
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
	fmt.Println(len(res.Data))
}
