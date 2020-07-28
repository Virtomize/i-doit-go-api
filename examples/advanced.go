package main

import (
	"fmt"
	"log"

	"github.com/virtomize/i-doit-go-api"
)

/* this time we will do some more advanced requests using the Request Method
we have seen in some examples before
first we will create a simple read method for object_type_categories similar to the library implementation
*/

// GetObjTypeCat for getting object type categories
func GetObjTypeCat(a *goidoit.API, objType string) (goidoit.GenericResponse, error) {

	// lets create our request parameters, in our case we simple query by type
	p := struct {
		Type string `json:"type"`
	}{objType}

	// then we do our request
	data, err := a.Request("cmdb.object_type_categories.read", &p)
	if err != nil {
		return goidoit.GenericResponse{}, err
	}

	// our Request gives us only an interface back so we need to type assert it
	// thankfully our api client package provides a generic TypeAssertResult function that takes care
	// of all forms of type assertions
	return goidoit.TypeAssertResult(data), nil

}

func main() {
	// Debug and SSL Skip
	//goidoit.Debug(true)
	//goidoit.SkipTLSVerify(true)

	// create api object using NewLogin for X-RPC-Auth
	a, err := goidoit.NewLogin("https://example.com/src/jsonrpc.php", "yourapikey", "username", "password")
	if err != nil {
		log.Fatal(err)
	}

	// now lets requests some object_type_categories like C__OBJTYPE__VIRTUAL_SERVER
	servers, err := GetObjTypeCat(a, "C__OBJTYPE__VIRTUAL_SERVER")
	if err != nil {
		log.Fatal(err)
	}

	// and we got a list back
	fmt.Println(servers)
}
