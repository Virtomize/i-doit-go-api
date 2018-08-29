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
	viaString, err := a.GetObjectByType("C__OBJTYPE__LAYER3_NET", "Global v4")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", viaString)

	// select via objectid
	viaInt, err := a.GetObjectByType("C__OBJTYPE__LAYER3_NET", 20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", viaInt)

	// select via slice of objectid's
	viaIntSlice, err := a.GetObjectByType("C__OBJTYPE__LAYER3_NET", []int{20, 21})
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range viaIntSlice.Result {
		fmt.Printf("%#v\n", v)
	}
}
