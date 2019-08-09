package main

import (
	"fmt"
	"log"

	"github.com/cseeger-epages/i-doit-go-api"
)

func main() {
	// Debug and SSL Skip
	//goidoit.Debug(true)
	//goidoit.SkipTLSVerify(true)

	// create api object using NewLogin for X-RPC-Auth
	a, err := goidoit.NewLogin("https://example.com/src/jsonrpc.php", "yourapikey", "username", "password")
	if err != nil {
		log.Fatal(err)
	}

	// select via string
	viaString, err := a.GetObject("Root location")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", viaString.Result[0]["title"])

	// select via objectid
	viaInt, err := a.GetObject(4)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", viaInt.Result[0]["title"])

	// select via slice of objectid's
	viaIntSlice, err := a.GetObject([]int{1, 4, 5})
	if err != nil {
		log.Fatal(err)
	}

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
	}{S{"Max"}, 1}

	// get your filtered objects
	viaStruct, err := a.GetObject(CustomStruct)
	if err != nil {
		log.Fatal(err)
	}

	for i := range viaStruct.Result {
		fmt.Printf("%#v\n", viaStruct.Result[i]["title"])
	}

	// you can optional request categorys using the following CustomStruct
	CustomStruct := struct {
		Filter     S    `json:"filter"`
		L          int  `json:"limit"`
		Categories bool `json:"categories"`
	}{S{"Max"}, 1, true}

	// get your filtered objects
	viaStruct, err := a.GetObject(CustomStruct)
	if err != nil {
		log.Fatal(err)
	}

	for i := range viaStruct.Result {
		fmt.Printf("%#v\n", viaStruct.Result[i]["title"])
	}

	// or specify what Categories you want to see
	CustomStruct := struct {
		Filter     S        `json:"filter"`
		L          int      `json:"limit"`
		Categories []string `json:"categories"`
	}{S{"Max"}, 1, []string{"C__CATG__...", "C__CATG__..."}}

	// get your filtered objects
	viaStruct, err := a.GetObject(CustomStruct)
	if err != nil {
		log.Fatal(err)
	}

	for i := range viaStruct.Result {
		fmt.Printf("%#v\n", viaStruct.Result[i]["title"])
	}
}
