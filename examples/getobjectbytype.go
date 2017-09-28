package main

import (
	"fmt"
	"github.com/cseeger-epages/i-doit-go-api"
)

func main() {
	// Debug and SSL Skip
	//goidoit.Debug(true)
	//goidoit.SkipTLSVerify(true)

	// create api object using api url and your api key
	a, _ := goidoit.NewApi("https://example.com/src/jsonrpc.php", "yourapikey")

	// select via string
	viaString, _ := a.GetObjectByType("C__OBJTYPE__LAYER3_NET", "Global v4")
	fmt.Printf("%#v\n", viaString)

	// select via objectid
	viaInt, _ := a.GetObjectByType("C__OBJTYPE__LAYER3_NET", 20)
	fmt.Printf("%#v\n", viaInt)

	// select via slice of objectid's
	viaIntSlice, _ := a.GetObjectByType("C__OBJTYPE__LAYER3_NET", []int{20, 21})
	for _, v := range viaIntSlice.Result {
		fmt.Printf("%#v\n", v)
	}
}
