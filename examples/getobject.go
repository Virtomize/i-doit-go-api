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
		Filter S `json:"filter"`
	}{S{"Max"}} // <-- struct initialisation using our parameter struct

	// get your filtered objects
	viaStruct, _ := a.GetObject(CustomStruct)

	for i := range viaStruct.Result {
		fmt.Printf("%#v\n", viaStruct.Result[i]["title"])
	}
}
