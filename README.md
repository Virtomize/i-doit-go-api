**__ This project is still under development and not fully featured __**

# i-doit-go-api/goidoit 

is an I-doit API client implementation written in [GOLANG](https://golang.org).

Its focused on simplicity to create easy writable code for your projects.

## Installation

If you already installed GO on your system and configured it properly than its simply:

```
go get github.com/cseeger-epages/i-doit-go-api
```

If not follow [these instructions](https://nats.io/documentation/tutorials/go-install/).

## Usage 

### Simple example

```
package main

import (
	"fmt"
	"github.com/cseeger-epages/i-doit-go-api"
)

func main() {
	// create api object
	a, _ := goidoit.NewApi("https://example.com/src/jsonrpc.php", "yourapikey")

	// search for sht like test
	s, _ := a.Search("test")

  // thats it
  fmt.Println(s)
}
```

### Advanced examples

There are more advanced [examples documented](https://github.com/cseeger-epages/i-doit-go-api/tree/master/examples) in the repo for common use cases.

## Code Documentation

You find the full [code documentation here](https://godoc.org/github.com/cseeger-epages/i-doit-go-api)

## Additional Informations

- [i-doit API documentation](https://kb.i-doit.com/pages/viewpage.action?pageId=37355644)
- [golang](https://golang.org/)
- [godoc](https://godoc.org/)
- [golang playground](https://play.golang.org/)

## Other client implementations

- Benjamin Heisigs [i-doit-api-client-php](https://github.com/bheisig/i-doit-api-client-php/)
- Dennis Stückens [i-doit-api-clients](https://bitbucket.org/dstuecken/i-doit-api-clients/)
