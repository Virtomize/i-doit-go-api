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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Debug(v bool) {
	debug = v
}

// disable SSL/TLS verification for self signed certificates
func SkipTLSVerify(v bool) {
	insecure = v
}

// generic Type Assert function
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

	// logging
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
