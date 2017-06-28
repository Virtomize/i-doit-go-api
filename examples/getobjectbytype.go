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
	viaString, _ := a.GetObjectByType("C__BOJTYPE__LAYER3_NET", "Global v4")
	fmt.Printf("%#v\n", viaString)

	// select via objectid
	viaInt, _ := a.GetObjectByType("C__BOJTYPE__LAYER3_NET", 20)
	fmt.Printf("%#v\n", viaInt)

	// select via slice of objectid's
	viaInt, _ := a.GetObjectByType("C__BOJTYPE__LAYER3_NET", []int{20, 21})
	for i := range viaIntSlice.Result {
		fmt.Printf("%#v\n", viaIntSlice)
	}
}
