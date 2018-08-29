package main

import (
	"fmt"

	"github.com/cseeger-epages/i-doit-go-api"
)

func main() {

	// Debug and SSL Skip
	goidoit.Debug(true)
	//goidoit.SkipTLSVerify(true)

	// create api object using NewLogin for X-RPC-Auth
	a, err := goidoit.NewLogin("https://example.com/src/jsonrpc.php", "yourapikey", "username", "password")

	// create your parameters as a struct, that gets marshalled to json
	p := struct {
		Query string `json:"q"`
	}{"test"}

	// define request method and request by referencing to your defined parameters
	data, _ := a.Request("idoit.search", &p)

	// output
	fmt.Println(data)
}
