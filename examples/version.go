package main

import (
	"fmt"
	"log"

	"github.com/virtomize/i-doit-go-api"
)

func main() {

	// Debug and SSL Skip
	goidoit.Debug(true)
	//goidoit.SkipTLSVerify(true)

	// create api object using NewLogin for X-RPC-Auth
	a, err := goidoit.NewLogin("https://example.com/src/jsonrpc.php", "yourapikey", "username", "password")
	if err != nil {
		log.Fatal(err)
	}

	// do a simple request  like getting the idoit version
	json, err := a.Request("idoit.version", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\njson Response: %s\n\n", json)

	/* parse data
	json is of type Response having a Result and Error type
	the Result will auto maped by go and needs a type assertion
	the correct type assertion here is map[string]interface{}
	*/
	data, err := json.Result.(map[string]interface{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("the mapped result data: %s\n\n", data)

	// after parsing you can use the data
	fmt.Printf("i-doit version: %s %s\n", data["version"], data["type"])

	// if you want to go deeper into your Result data you need type assertion again
	data2, err := data["login"].(map[string]interface{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("i am user: ", data2["name"])

}
