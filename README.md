# jsonq

<img align="right" width="150" src="https://user-images.githubusercontent.com/209884/92572337-42b42900-f2bf-11ea-973a-c74a359553a5.png" alt="jsonq">

[![Go Report Card](https://goreportcard.com/badge/github.com/IamFaizanKhalid/jsonq)](https://goreportcard.com/report/github.com/IamFaizanKhalid/jsonq) [![Release](https://img.shields.io/github/v/release/IamFaizanKhalid/jsonq.svg?style=flat-square)](https://github.com/IamFaizanKhalid/jsonq/releases) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE) [![GoDoc](https://pkg.go.dev/badge/github.com/IamFaizanKhalid/jsonq)](https://pkg.go.dev/github.com/IamFaizanKhalid/jsonq)


Easily traverse through the JSON data with no fixed structure.

## Installation

```shell
go get -u github.com/IamFaizanKhalid/jsonq
```

## Usage Example

```go
package main

import (
	"fmt"

	"github.com/IamFaizanKhalid/jsonq"
)

func main() {
	apiResponse := []byte(`{
		"code": 200,
		"status": "success",
		"data": {
			"user": {
				"id": 123,
				"name": "John Doe",
				"email": "john@example.com"
			},
			"posts": [
				{
					"id": 1,
					"title": "First Post"
				},
				{
					"id": 2,
					"title": "Second Post"
				}
			]
		}
	}`)

	resp, err := jsonq.ParseObject(apiResponse)
	if err != nil {
		panic(err)
	}

	code := resp.Val("code").Int()
	status := resp.Val("status").Str()

	fmt.Println(code, status)

	if !resp.Has("data") {
		return
	}
	data := resp.Obj("data")

	userInfo, ok := data.OptObj("user")
	if ok {
		fmt.Println("-- User Info --")

		fmt.Println("ID:", userInfo.Val("id").Int())
		fmt.Println("Name:", userInfo.Val("name").Str())
		fmt.Println("Email:", userInfo.Val("email").Str())
	}

	fmt.Println("-- Posts --")

	posts := data.Arr("posts").Obj()

	for _, post := range posts {
		id := post.Val("id").Int()
		title := post.Val("title").Str()

		fmt.Println(id, title)
	}
}
```


## License

[MIT License](./LICENSE)
