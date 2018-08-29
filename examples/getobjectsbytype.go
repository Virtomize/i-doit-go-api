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

	// enable debug
	// goidoit.Debug(true)

	// gives you all objects of a sepcific type
	// eg. give me all layer 3 networks
	Net3, err := a.GetObjectsByType("C__OBJTYPE__LAYER3_NET")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Net3)
}
