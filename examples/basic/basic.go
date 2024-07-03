package main

import (
	"fmt"

	hawapi "github.com/HawAPI/go-sdk"
)

func main() {
	client := hawapi.NewClient()

	actors, err := client.ListActors()
	if err != nil {
		panic(err)
	}

	fmt.Println(actors)
}
