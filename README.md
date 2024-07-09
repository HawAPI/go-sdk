# HawAPI - go-sdk

HawAPI SDK for Golang

- [API Docs](https://hawapi.theproject.id/docs/)
- [SDK Docs](https://pkg.go.dev/github.com/HawAPI/go-sdk)

## Topics

- [Installation](#installation)
- [Usage](#usage)
    - [Init client with default options](#init-client-with-default-options)
    - [Init client with custom options](#init-client-with-custom-options)
      - [NewClient](#newclient)
      - [NewClientWithOpts](#newclientwithopts)

## Installation

```
go get github.com/HawAPI/go-sdk
```

## Usage

- [See examples](./examples)

### Init client with default options

```go
package main

import (
  "fmt"

  "github.com/HawAPI/go-sdk"
)

func main() {
  client := hawapi.NewClient()
}
```

### Init client with custom options

#### NewClient

```go
package main

import (
  "fmt"

  "github.com/HawAPI/go-sdk"
)

func main() {
  client := hawapi.NewClient()
  client.WithOpts(hawapi.Options{
    Endpoint: "http://localhost:8080/api",
	// Version
	// Language
	// Token
	// ...
  })
}
```

#### NewClientWithOpts

```go
package main

import (
  "fmt"

  "github.com/HawAPI/go-sdk"
)

func main() {
  client := hawapi.NewClientWithOpts(hawapi.Options{
    Endpoint: "http://localhost:8080/api",
    // Version
    // Language
    // Token
    // ...
  })
}
```
