package main

import (
	"fmt"
	"github.com/cseeger-epages/i-doit-go-api"
)

func main() {
	// create api object
	a, _ := goidoit.NewApi("https://example.com/src/jsonrpc.php", "yourapikey")

	// enable debug
	//goidoit.Debug(true)

	// request constants
	data, _ := a.Request("idoit.constants", nil)

	// idk what to use it for :P until it responses some id's to the names
	fmt.Printf("%#v\n", data)
}
