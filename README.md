# HawAPI - go-sdk

HawAPI SDK for Golang

- [API Docs](https://hawapi.theproject.id/docs/)
- [SDK Docs](https://pkg.go.dev/github.com/HawAPI/go-sdk)

## Topics

- [Installation](#installation)
- [Usage](#usage)
    - [Init client](#init-client)
    - [Fetch information](#fetch-information)
    - [Error handling](#error-handling)

## Installation

```
go get github.com/HawAPI/go-sdk/hawapi@latest
```

## Usage

- [See examples](./_examples)

### Init client

```go
package main

import (
  "fmt"

  "github.com/HawAPI/go-sdk/hawapi"
)

func main() {
    // Create a new client with default options
    client := hawapi.NewClient()
    
    // Create client with custom options
    client = hawapi.NewClientWithOpts(hawapi.Options{
        Endpoint: "http://localhost:8080/api",
        // When using 'WithOpts' or 'NewClientWithOpts' the value of
        // 'UseInMemoryCache' will be set to false
        UseInMemoryCache: true,
        // Version
        // Language
        // Token
        // ...
    })
	
    // You can also change the options later
    client.WithOpts(hawapi.Options{
        Language: "pt-BR",
        // When using 'WithOpts' or 'NewClientWithOpts' the value of
        // 'UseInMemoryCache' will be set to false
        UseInMemoryCache: true,
    })
}
```

### Fetch information

```go
package main

import (
  "fmt"

  "github.com/HawAPI/go-sdk/hawapi"
)

func main() {
    client := hawapi.NewClient()
    
    res, err := client.ListActors()
    if err != nil {
        panic(err)
    }
    
    fmt.Println(res)
}
```

### Error handling

- Check out the [hawapi.ErrorResponse](hawapi/error.go)

```go
package main

import (
	"fmt"

	"github.com/HawAPI/go-sdk/hawapi"
	"github.com/google/uuid"
)

func main() {
    client := hawapi.NewClient()
    
    id, _ := uuid.Parse("<unknown uuid>")
    res, err := client.FindActor(id)
    if err != nil {
        // If the error is coming from the API request, 
        // it'll be of type hawapi.ErrorResponse.
        if resErr, ok := err.(hawapi.ErrorResponse); ok {
            fmt.Printf("API error %d Message: %s\n", resErr.Code, resErr.Message)
        } else {
            fmt.Println("SDK error:", err)
        }
    }
    
    fmt.Println(res)
}
```