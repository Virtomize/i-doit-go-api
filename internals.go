package goidoit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Debug function to enable debug output
func Debug(v bool) {
	debug = v
}

// SkipTLSVerify disable SSL/TLS verification for self signed certificates
func SkipTLSVerify(v bool) {
	insecure = v
}

// TypeAssertResult is a generic type assert function
func TypeAssertResult(data Response) (GenericResponse, error) {
	ret := GenericResponse{Jsonrpc: data.Jsonrpc, Error: data.Error}

	ret.Error.Data = ""
	if data.Error.Data != nil {
		ret.Error.Data = data.Error.Data.(map[string]interface{})
	}

	if data.Result != nil {
		switch data.Result.(type) {
		case []interface{}:
			results := data.Result.([]interface{})
			for i := range results {
				ret.Result = append(ret.Result, results[i].(map[string]interface{}))
			}
		case interface{}:
			ret.Result = append(ret.Result, data.Result.(map[string]interface{}))
		}
	}

	return ret, nil
}

// getID increments request id's
func getID() int {
	id++
	return id
}

// GetParams is used to append nessesary parameters to user provided one
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

// ParseResponse parses json response
func ParseResponse(resp *http.Response) Response {
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("PARSING ERROR: ", err)
	}

	// logging
	debugPrint("----> # Response # <----\n%s\n", string(data))

	var ret Response
	_ = json.Unmarshal(data, &ret)

	return ret
}

// debugPrint used for request/response debugging
func debugPrint(format string, a ...interface{}) (n int, err error) {
	if debug {
		return fmt.Printf(format, a)
	}
	return 0, nil
}
