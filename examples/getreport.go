package main

import (
	"fmt"
	"log"

	"github.com/virtomize/i-doit-go-api"
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

	data, err := a.GetReport(1)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range data.Result {
		fmt.Println(v)
		/*
			do sth with you report data
		*/
	}
}
