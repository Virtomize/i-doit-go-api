package main

import (
	"fmt"
	"github.com/cseeger-epages/i-doit-go-api"
	"strconv"
)

func main() {
	// create api object
	a, _ := goidoit.NewApi("https://example.com/src/jsonrpc.php", "yourapikey")

	// enable debug
	goidoit.Debug(true)

	// create a struct defining your object
	fmt.Println("#### create #####")
	data := struct {
		Type  string `json:"type"`
		Title string `json:"title"`
	}{"C__OBJTYPE__VIRTUAL_SERVER", "test-vm"}

	// create it
	a.Create(data)

	// read data back using GetObject
	fmt.Println("#### read #####")
	obj, _ := a.GetObject("test-vm")
	// we need it to be int
	id, _ := strconv.Atoi(obj.Result[0]["id"].(string))

	// update object
	fmt.Println("#### update #####")
	updata := struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}{id, "test-neu"}

	a.Update(updata)

	// hopefully theres only one object or we will delete maybe the wrong one :P
	fmt.Println("#### purge #####")
	a.Quickpurge(id)

	// create another object
	fmt.Println("#### create net #####")
	dataNet := struct {
		Type  string `json:"type"`
		Title string `json:"title"`
	}{"C__OBJTYPE__LAYER3_NET", "test-net"}
	a.Create(dataNet)

	// get our id
	fmt.Println("#### read net #####")
	net, _ := a.GetObject("test-net")
	idnet, _ := strconv.Atoi(net.Result[0]["id"].(string))

	// and now lets archive it
	fmt.Println("#### archive net #####")
	archive := struct {
		ID     int    `json:"id"`
		Status string `json:"status"`
	}{idnet, "C__RECORD_STATUS__ARCHIVED"}
	a.Delete(archive)
}
