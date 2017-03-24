package main

import (
	"fmt"
	"github.com/cseeger-epages/i-doit-go-api"
)

/* this time we will do some more advanced requests using the Request Method
we have seen in some examples before
first we will create a simple read method for object_type_categories since our api client only implements this for standard i-doit objects
*/
func GetObjTypeCat(a *goidoit.Api, objType string) (goidoit.GenericResponse, error) {

	// lets create our request parameters, in our case we simple query by type
	p := struct {
		Type string `json:"type"`
	}{objType}

	// then we do our request
	data, _ := a.Request("cmdb.object_type_categories.read", &p)

	// our Request gives us only an interface back so we need to type assert it
	// thankfully our api client package provides a generic TypeAssertResult function that takes care
	// of all forms of type assertions
	return goidoit.TypeAssertResult(data)

}

func main() {
	// create api object
	a, _ := goidoit.NewApi("https://example.com/src/jsonrpc.php", "yourapikey")

	// enable debug
	goidoit.Debug(true)

	// now lets requests some object_type_categories like C__OBJTYPE__VIRTUAL_SERVER
	servers, _ := GetObjTypeCat(a, "C__OBJTYPE__VIRTUAL_SERVER")

	// and we got a list back
	fmt.Println(servers)
}
