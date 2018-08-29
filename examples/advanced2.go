package main

import (
	"github.com/cseeger-epages/i-doit-go-api"
)

/* dummy file for advanced example 2 */
func dummy(a *goidoit.API) (goidoit.GenericResponse, error) {

}

func main() {
	// Debug and SSL Skip
	//goidoit.Debug(true)
	//goidoit.SkipTLSVerify(true)

	// create api object using NewLogin for X-RPC-Auth
	_, _ := goidoit.NewLogin("https://example.com/src/jsonrpc.php", "yourapikey", "username", "password")
}
