package main

import (
	"fmt"
	"github.com/cseeger-epages/i-doit-go-api"
)

func main() {
	// create api object
	a, _ := goidoit.NewApi("http://example.com/src/jsonrpc.php", "yourapikey")

	// enable debug
	// goidoit.Debug(true)

	// gives you all objects of a sepcific type
	// eg. give me all layer 3 networks
	Net3, _ := a.GetObjectsByType("C__OBJTYPE__LAYER3_NET")
	fmt.Println(Net3)
}
