package main

import (
	"fmt"
	"github.com/cseeger-epages/i-doit-go-api"
)

/* dummy file for advanced example 2 */
func dummy(a *goidoit.Api) (goidoit.GenericResponse, error) {

}

func main() {
	// Debug and SSL Skip
	//goidoit.Debug(true)
	//goidoit.SkipTLSVerify(true)

	// create api object using api url and your api key
	a, _ := goidoit.NewApi("https://example.com/src/jsonrpc.php", "yourapikey")
}
