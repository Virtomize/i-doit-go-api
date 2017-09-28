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
	viaString, _ := a.GetObjTypeCat("C__OBJTYPE__PERSON")
	fmt.Printf("%#v\n\n", viaString.Result[0]["catg"])

	// select via objectid
	viaInt, _ := a.GetObjTypeCat(53)
	fmt.Printf("%#v\n", viaInt.Result[0]["catg"])

}
