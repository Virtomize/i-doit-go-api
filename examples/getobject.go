package main

import (
	"fmt"
	"github.com/cseeger-epages/i-doit-go-api"
)

func main() {
	// create api object
	a, _ := goidoit.Newapi("http://example.com/src/jsonrpc.php", "yourapikey")

	// select via string
	viaString, _ := a.GetObject("Root location")
	fmt.Printf("%#v\n", viaString.Result[0]["title"])

	// select via objectid
	viaInt, _ := a.GetObject(4)
	fmt.Printf("%#v\n", viaInt.Result[0]["title"])

	// select via slice of objectid's
	viaIntSlice, _ := a.GetObject([]int{1, 4, 5})
	for i := range viaIntSlice.Result {
		fmt.Printf("%#v\n", viaIntSlice.Result[i]["title"])
	}
}
