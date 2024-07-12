package main

import (
	"fmt"

	"github.com/HawAPI/go-sdk/pkg/hawapi"
)

func main() {
	client := hawapi.NewClient()
	client.WithOpts(hawapi.Options{
		Endpoint: "http://localhost:8080/api",
	})

	res, err := client.ListActors()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
