package main

import (
	"fmt"
	"github.com/cseeger-epages/i-doit-go-api"
)

func main() {
	// create api object
	a, _ := goidoit.NewApi("https://example.com/src/jsonrpc.php", "yourapikey")

	// enable debug
	//goidoit.Debug(true)

	data, _ := a.GetReport(1)

	for _, v := range data.Result {
		/*
			do sth with you report data
		*/
	}
}
