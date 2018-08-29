package main

import (
	"fmt"
	"log"

	"github.com/cseeger-epages/i-doit-go-api"
)

func main() {
	// Debug and SSL Skip
	//goidoit.Debug(true)
	//goidoit.SkipTLSVerify(true)

	// create api object using NewLogin for X-RPC-Auth
	a, err := goidoit.NewLogin("https://example.com/src/jsonrpc.php", "yourapikey", "username", "password")
	if err != nil {
		log.Fatal(err)
	}

	// request constants
	data, err := a.Request("idoit.constants", nil)
	if err != nil {
		log.Fatal(err)
	}

	// idk what to use it for :P until it responses some id's to the names
	fmt.Printf("%#v\n", data)
}
