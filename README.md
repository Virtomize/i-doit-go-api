# i-doit-go-api/goidoit 

[![GoDoc](https://img.shields.io/badge/godoc-reference-green.svg)](https://godoc.org/github.com/cseeger-epages/i-doit-go-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/cseeger-epages/i-doit-go-api)](https://goreportcard.com/report/github.com/cseeger-epages/i-doit-go-api)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/cseeger-epages/i-doit-go-api/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/cseeger-epages/i-doit-go-api.svg?branch=master)](https://travis-ci.org/cseeger-epages/i-doit-go-api)
[![Built with Mage](https://magefile.org/badge.svg)](https://magefile.org)

is an [I-doit](https://www.i-doit.com/) API client implementation written in [GOLANG](https://golang.org).

Its focused on simplicity to create easy writable code for your projects.

the package supports:
- Login
- Requests
- Searching
- Getting object information by id, id slice, string or custom field
- creating, altering, deleting, archiving, quickpurging objects
- reports
- categories (only parts implemented) (wip)
- object\_type\_categories
- dialog

Other types of requests like
- object\_types (coming soon)
- object\_type\_groups (coming soon)
- objects\_by\_relation (coming soon)
- categories (category, category\_info) (wip)
- location\_tree
- workstation\_components
- logbook
- impact

are not implemented yet (and maybe never will) but can easily be implemented using the goidoit.Request function
there are some [advanced](https://github.com/cseeger-epages/i-doit-go-api/blob/master/examples/advanced.go) [examples](https://github.com/cseeger-epages/i-doit-go-api/blob/master/examples/advanced2.go) of how to do this.

There is also a second repo [i-doit-go-tools](https://github.com/cseeger-epages/i-doit-go-tools) where you can find some more scripts for special use-cases that can be used as an
even more advanced example :P.

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
  a, _ := goidoit.NewAPI("https://example.com/src/jsonrpc.php", "yourapikey")

  // or login using X-RPC-Auth (recommended)
  // a, _ := goidoit.NewLogin("https://example.com/src/jsonrpc.php", "yourapikey", "username", "password")

  // search for sth like test
  s, _ := a.Search("test")

  // thats it
  fmt.Println(s)
}
```

### Advanced examples

There are more advanced [examples](https://github.com/cseeger-epages/i-doit-go-api/tree/master/examples) documented in the repo for common use cases.

## Code Documentation

You find the full [code documentation here](https://godoc.org/github.com/cseeger-epages/i-doit-go-api)

## Additional Informations

- [i-doit API documentation](https://kb.i-doit.com/pages/viewpage.action?pageId=37355644)
- [golang](https://golang.org/)
- [godoc](https://godoc.org/)
- [golang playground](https://play.golang.org/)

## Other great client implementations

- Benjamin Heisigs [i-doit-api-client-php](https://github.com/bheisig/i-doit-api-client-php/)
- Dennis Stückens [i-doit-api-clients](https://bitbucket.org/dstuecken/i-doit-api-clients/)

Thanks for your code inspirations!
