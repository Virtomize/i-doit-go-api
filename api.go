/*
	Go library for simple i-doit api usage

	Copyright (C) 2017 Carsten Seeger

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.

	@author Carsten Seeger
	@copyright Copyright (C) 2017 Carsten Seeger
	@license http://www.gnu.org/licenses/gpl-3.0 GNU General Public License 3
	@link https://github.com/cseeger-epages/i-doit-go-api
*/

package goidoit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// request id
var id int = 0
var debug bool = false

// basic api interface
type ApiMethods interface {
	// i-doit api request structure
	// as defined in https://kb.i-doit.com/pages/viewpage.action?pageId=7831613
	// also there is a list of methods available
	Request(string, interface{}) (Response, error)

	// search CMDB using a string
	//
	// The search function does handle type assertions
	// for simple output usage
	Search(string) (GenericResponse, error)

	// get object(s) data,
	// for input an id can be used an array of ids can be used or the title string can be used
	GetObject(interface{}) (GenericResponse, error)

	// fast delete option where archive, delete and purge will be executed one after another
	// accepts id or []id as input
	Quickpurge(interface{}) (GenericResponse, error)
	/*
		Login()
		Logout()
		IsLoggedIn()
	*/
}

// api struct used for implementing the apiMethods interface
type Api struct {
	Url, Apikey string
}

type Request struct {
	Version string      `json:"version"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      int         `json:"id"`
}

// i-doit api response structure
type Response struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   IdoitError  `json:"error"`
}

// i-doit api response structure used for search requests
//
// the map is used to handle type assertions
type GenericResponse struct {
	Jsonrpc string
	Result  []map[string]interface{}
	Error   IdoitError
}

// i-doit api error structure
type IdoitError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Apikey used for requests
type Apikey struct {
	Apikey string `json:"apikey"`
}

func Debug(v bool) {
	debug = v
}

// api constructor
func NewApi(url string, apikey string) (*Api, error) {
	if len(url) != 0 && len(apikey) != 0 {
		a := Api{url, apikey}
		return &a, nil
	}
	return nil, errors.New("url or apikey empty")
}

func (a Api) Request(method string, parameters interface{}) (Response, error) {

	var params = GetParams(a, parameters)
	id = getID()

	data := Request{
		Version: "2.0",
		Method:  method,
		Params:  params,
		Id:      id,
	}

	dataJson, err := json.Marshal(data)

	// logging tbd
	debugPrint("----> # Request # <----\n%s\n", string(dataJson))

	req, err := http.NewRequest("POST", a.Url, bytes.NewBuffer(dataJson))
	if err != nil {
		fmt.Println("REQUEST ERROR: ", err)
		return Response{}, err
	}
	req.Header.Add("content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("REQUEST ERROR: ", err)
		return Response{}, err
	}
	var ret = ParseResponse(resp)
	return ret, nil
}

func (a *Api) Search(query string) (GenericResponse, error) {
	params := struct {
		Query string `json:"q"`
	}{query}
	data, err := a.Request("idoit.search", &params)
	if err != nil {
		return GenericResponse{}, err
	}

	return TypeAssertResult(data)
}

// Object filter type int or []int
type F1 struct {
	Data []int `json:"ids"`
}

// Object filter type string
type F2 struct {
	Data string `json:"title"`
}

func (a *Api) GetObject(query interface{}) (GenericResponse, error) {

	var Params interface{}
	switch query.(type) {
	case int:
		Params = struct {
			Filter F1 `json:"filter"`
		}{F1{[]int{query.(int)}}}
	case []int:
		Params = struct {
			Filter F1 `json:"filter"`
		}{F1{query.([]int)}}
	case string:
		Params = struct {
			Filter F2 `json:"filter"`
		}{F2{query.(string)}}
	default:
		return GenericResponse{}, errors.New("Input type is not int, []int or string")
	}

	data, err := a.Request("cmdb.objects.read", &Params)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}

// Quickpurge ftw
func (a *Api) Quickpurge(ids interface{}) (GenericResponse, error) {

	var Params interface{}
	switch ids.(type) {
	case int:
		Params = struct {
			Filter F1 `json:"filter"`
		}{F1{[]int{ids.(int)}}}
	case []int:
		Params = struct {
			Filter F1 `json:"filter"`
		}{F1{ids.([]int)}}
	default:
		return GenericResponse{}, errors.New("Input type is not int or []int")
	}

	data, err := a.Request("cmdb.objects.quickpurge", &Params)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}

// generic Type Assert function
func TypeAssertResult(data Response) (GenericResponse, error) {
	ret := GenericResponse{Jsonrpc: data.Jsonrpc, Error: data.Error}

	ret.Error.Data = ""
	if data.Error.Data != nil {
		ret.Error.Data = data.Error.Data.(string)
	}

	results := data.Result.([]interface{})
	for i := range results {
		ret.Result = append(ret.Result, results[i].(map[string]interface{}))
	}

	return ret, nil
}

// increment request id's
func getID() int {
	id++
	return id
}

// append nessesary parameters to user provided one
func GetParams(a Api, parameters interface{}) interface{} {

	var params map[string]interface{}
	apikey := Apikey{a.Apikey}

	jsonParameters, err := json.Marshal(parameters)

	if err != nil {
		log.Fatal("JSON ERROR: ", err)
	}

	json.Unmarshal(jsonParameters, &params)
	jsonApikey, err := json.Marshal(apikey)

	if err != nil {
		log.Fatal("JSON ERROR: ", err)
	}

	json.Unmarshal(jsonApikey, &params)

	return params
}

// parse json response
func ParseResponse(resp *http.Response) Response {
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("PARSING ERROR: ", err)
	}

	// logging tbd
	debugPrint("----> # Response # <----\n%s\n", string(data))

	var ret Response
	_ = json.Unmarshal(data, &ret)

	return ret
}

// used for Request/Response debugging
func debugPrint(format string, a ...interface{}) (n int, err error) {
	if debug {
		return fmt.Printf(format, a)
	} else {
		return 0, nil
	}
}
