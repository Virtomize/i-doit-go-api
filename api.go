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
	Search(string) (SearchResponse, error)

	// get object(s) data,
	// single id or slice of id's can be used
	GetObjectsByID([]int) (Response, error)
	/*
		Login()
		Logout()
		IsLoggedIn()
	*/
}

// api struct used for implementing the apiMethods interface
type Api struct {
	url, apikey string
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
type SearchResponse struct {
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

// api constructor
func Newapi(url string, apikey string) (*Api, error) {
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
	//fmt.Println("Request: ", string(dataJson))

	req, err := http.NewRequest("POST", a.url, bytes.NewBuffer(dataJson))
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

func (a *Api) Search(query string) (SearchResponse, error) {
	params := struct {
		Query string `json:"q"`
	}{query}
	data, err := a.Request("idoit.search", &params)
	if err != nil {
		return SearchResponse{}, err
	}

	// do type assertions for easy output handling
	ret := SearchResponse{Jsonrpc: data.Jsonrpc, Error: data.Error}

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

	var params map[string]string
	apikey := Apikey{a.apikey}

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
	//fmt.Println("Response: ", string(data))

	var ret Response
	_ = json.Unmarshal(data, &ret)

	return ret
}
