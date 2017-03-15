package main

import (
	"fmt"
	"github.com/cseeger-epages/i-doit-go-api"
)

func main() {

	// create api object using api url and your api key
	a, _ := goidoit.Newapi("https://example.com/src/jsonrpc.php", "myapikey")

	// do a simple request  like getting the idoit version
	json, _ := a.Request("idoit.version", nil)
	fmt.Printf("\njson Response: %s\n\n", json)

	/* parse data
	json is of type Response having a Result and Error type
	the Result will auto maped by go and needs a type assertion
	the correct type assertion here is map[string]interface{}
	*/
	var data, err = json.Result.(map[string]interface{})
	fmt.Printf("the mapped result data: %s\n\n", data)
	if !err {
		fmt.Println("sth is going wrong here")
	}

	// after parsing you can use the data
	fmt.Printf("i-doit version: %s %s\n", data["version"], data["type"])

	// if you want to go deeper into your Result data you need type assertion again
	var data2, _ = data["login"].(map[string]interface{})
	fmt.Println("i am user: ", data2["name"])

}
