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

	dialog, err := a.GetDialog("C__CATG__MODEL", "manufacturer")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dialog)
}
