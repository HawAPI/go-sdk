package main

import (
	"fmt"

	"github.com/HawAPI/go-sdk/hawapi"
)

func main() {
	client := hawapi.NewClient()
	client.WithOpts(hawapi.Options{
		// JWT auth is required when performing any POST, PATCH and DELETE requests
		Token: "<JWT>",
	})

	actor := hawapi.CreateActor{
		// ...
	}
	res, err := client.CreateActor(actor)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.FirstName)
}
