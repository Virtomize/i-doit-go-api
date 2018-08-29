package main

import (
	"fmt"

	"github.com/cseeger-epages/i-doit-go-api"
)

func main() {
	// Debug and SSL Skip
	//goidoit.Debug(true)
	//goidoit.SkipTLSVerify(true)

	// create api object using NewLogin for X-RPC-Auth
	a, err := goidoit.NewLogin("https://example.com/src/jsonrpc.php", "yourapikey", "username", "password")

	// request constants
	data, _ := a.Request("idoit.constants", nil)

	// idk what to use it for :P until it responses some id's to the names
	fmt.Printf("%#v\n", data)
}
