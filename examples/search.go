package main

import (
	"fmt"
	"log"

	"github.com/cseeger-epages/i-doit-go-api"
)

func main() {

	// Debug and SSL skip
	//goidoit.Debug(true)
	//goidoit.SkipTLSVerify(true)

	// create api object using NewLogin for X-RPC-Auth
	a, err := goidoit.NewLogin("https://example.com/src/jsonrpc.php", "yourapikey", "username", "password")
	if err != nil {
		log.Fatal(err)
	}

	// search for sht like test
	s, err := a.Search("test")
	if err != nil {
		log.Fatal(err)
	}

	/* show output
	the SearchResponse type is nearly the same like Response type
	except that it does handle the type assertions for you
	*/
	for i, v := range s.Result {
		fmt.Println("----", i, "----")
		for k, d := range v {
			fmt.Println(k, d)
		}
	}

	// or a specific value
	fmt.Println("----", "Single Output", "----")
	fmt.Println(s.Result[0]["value"])
}
