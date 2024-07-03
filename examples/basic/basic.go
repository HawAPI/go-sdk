package main

import (
	"fmt"

	"github.com/HawAPI/go-sdk/pkg/hawapi"
)

func main() {
	client := hawapi.NewClient()

	actors, err := client.ListActors()
	if err != nil {
		panic(err)
	}

	fmt.Println(actors)
}
