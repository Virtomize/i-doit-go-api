package main

import (
	"github.com/cseeger-epages/i-doit-go-api"
	"log"
)

func main() {
	// enable Debug
	goidoit.Debug(true)

	// create api object using NewLogin for X-RPC-Auth
	a, err := goidoit.NewLogin("https://example.com/src/jsonrpc.php", "yourapikey", "username", "password")
	// check if our login was successfull
	if err != nil {
		log.Fatal(err)
	}

	// do sth like a simple search
	a.Search("test")

	// logout when your done
	a.Logout()
}
