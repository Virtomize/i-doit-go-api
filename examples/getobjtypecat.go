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

	// select via string
	viaString, _ := a.GetObjTypeCat("C__OBJTYPE__PERSON")
	fmt.Printf("%#v\n\n", viaString.Result[0]["catg"])

	// select via objectid
	viaInt, _ := a.GetObjTypeCat(53)
	fmt.Printf("%#v\n", viaInt.Result[0]["catg"])

}
