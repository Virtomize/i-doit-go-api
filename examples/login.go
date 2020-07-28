package main

import (
	"log"

	"github.com/virtomize/i-doit-go-api"
)

func main() {
	// enable Debug
	goidoit.Debug(true)
	//goidoit.SkipTLSVerify(true)

	// create api object using NewLogin for X-RPC-Auth
	a, err := goidoit.NewLogin("https://example.com/src/jsonrpc.php", "yourapikey", "username", "password")

	// check if our login was successful
	if err != nil {
		log.Fatal(err)
	}

	// do sth like a simple search
	a.Search("test")

	// logout when your done
	a.Logout()

	// another way is to reuse the session key
	a, err = goidoit.NewLogin("https://example.com/src/jsonrpc.php", "yourapikey", "username", "password")

	// check if our login was successful
	if err != nil {
		log.Fatal(err)
	}

	// here is the session key
	sessionKey := a.SessionId

	// you can reuse this key without needing the user password anymore as long as the sessionKey is valid
	b := goidoit.API{"https://example.com/src/jsonrpc.php", "yourapikey", "username", "", sessionKey}

	b.Search("test")

	// destroys the session
	a.Logout()
}
