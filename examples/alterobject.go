package main

import (
	"fmt"
	"strconv"

	"github.com/cseeger-epages/i-doit-go-api"
)

func main() {
	// Debug and SSL Skip
	//goidoit.Debug(true)
	//goidoit.SkipTLSVerify(true)

	// create api object using NewLogin for X-RPC-Auth
	a, err := goidoit.NewLogin("https://example.com/src/jsonrpc.php", "yourapikey", "username", "password")

	// create a struct defining your object
	fmt.Println("#### create #####")
	data := struct {
		Type  string `json:"type"`
		Title string `json:"title"`
	}{"C__OBJTYPE__VIRTUAL_SERVER", "test-vm"}

	// create it
	res, _ := a.CreateObj(data)

	// now lets create a Hostname Category for oure new Server
	id, _ := strconv.Atoi(res.Result[0]["id"].(string))

	IPData := struct {
		Hostname       string `json:"hostname"`
		IP             string `json:"ipv4_address"`
		Ipv4Assingment int    `json:"ipv4_assignment"`
		NetType        int    `json:"net_type"`
		Domain         string `json:"domain"`
	}{"test-vm", "192.168.0.22", 1, 1, "example.de"}

	a.CreateCat(id, "C__CATG__IP", IPData)

	IPData = struct {
		Hostname       string `json:"hostname"`
		IP             string `json:"ipv4_address"`
		Ipv4Assingment int    `json:"ipv4_assignment"`
		NetType        int    `json:"net_type"`
		Domain         string `json:"domain"`
	}{"test-vm", "192.168.0.23", 1, 1, "example.de"}

	a.UpdateCat(id, "C__CATG__IP", IPData)

	// read data back using GetObject
	fmt.Println("#### read #####")
	obj, _ := a.GetObject("test-vm")
	// we need it to be int
	id, _ = strconv.Atoi(obj.Result[0]["id"].(string))

	// update object
	fmt.Println("#### update #####")
	updata := struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}{id, "test-neu"}

	a.UpdateObj(updata)

	// hopefully theres only one object or we will delete maybe the wrong one :P
	fmt.Println("#### purge #####")
	a.Quickpurge(id)

	// create another object
	fmt.Println("#### create net #####")
	dataNet := struct {
		Type  string `json:"type"`
		Title string `json:"title"`
	}{"C__OBJTYPE__LAYER3_NET", "test-net"}
	a.CreateObj(dataNet)

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
	a.DeleteObj(archive)
}
