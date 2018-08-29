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

	// select via string
	viaString, err := a.GetObjTypeCat("C__OBJTYPE__PERSON")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n\n", viaString.Result[0]["catg"])

	// select via objectid
	viaInt, err := a.GetObjTypeCat(53)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", viaInt.Result[0]["catg"])
}
