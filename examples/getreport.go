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

	data, _ := a.GetReport(1)

	for _, v := range data.Result {
		fmt.Println(v)
		/*
			do sth with you report data
		*/
	}
}
