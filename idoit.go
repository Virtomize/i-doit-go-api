package goidoit

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

// Request implements basic idoit requests
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

	// logging
	debugPrint("----> # Request # <----\n%s\n", string(dataJson))

	req, err := http.NewRequest("POST", a.Url, bytes.NewBuffer(dataJson))
	if err != nil {
		fmt.Println("REQUEST ERROR: ", err)
		return Response{}, err
	}
	req.Header.Add("content-type", "application/json")
	tr := &http.Transport{}
	if insecure {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		}
	}
	client := &http.Client{Transport: tr}

	// use X-RPC auth or fallback to API-Key only
	if method == "idoit.login" {
		req.Header["X-RPC-Auth-Username"] = []string{a.Username}
		req.Header["X-RPC-Auth-Password"] = []string{a.Password}
	} else {
		if a.IsLoggedIn() {
			req.Header["X-RPC-Auth-Session"] = []string{a.SessionId}
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("REQUEST ERROR: ", err)
		return Response{}, err
	}
	var ret = ParseResponse(resp)
	return ret, nil
}

// Search implements the idoit.search method
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

// Version implements idoit.version method
func (a *Api) Version() (GenericResponse, error) {
	data, err := a.Request("idoit.version", nil)
	if err != nil {
		return GenericResponse{}, err
	}

	return TypeAssertResult(data)
}
