package main

import (
	"fmt"
	"github.com/cseeger-epages/i-doit-go-api"
)

/* dummy file for advanced example 2 */
func dummy(a *goidoit.Api) (goidoit.GenericResponse, error) {

}

func main() {
	// create api object
	a, _ := goidoit.NewApi("https://example.com/src/jsonrpc.php", "yourapikey")

	// enable debug
	goidoit.Debug(true)
}
