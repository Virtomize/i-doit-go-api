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

	// or you provide your own filter struct as defined in i-doit api reference page 27f.
	//
	// create a struct with some specified filterparameter
	type S struct {
		FN string `json:"first_name"`
	}
	// initialise another struct providing your filterparameter as filter
	CustomStruct := struct {
		Filter S   `json:"filter"`
		L      int `json:"limit"`
	}{S{"Max"}, 1} // <-- struct initialisation using our parameter struct

	// get your filtered objects
	viaStruct, _ := a.GetObject(CustomStruct)

	for i := range viaStruct.Result {
		fmt.Printf("%#v\n", viaStruct.Result[i]["title"])
	}
}
