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
	"errors"
)

// Get Object by everything
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
		Params = query
	}

	data, err := a.Request("cmdb.objects.read", &Params)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}

func (a *Api) GetObjectByType(objType string, obj interface{}) (GenericResponse, error) {
	var Params interface{}
	switch obj.(type) {
	case int:
		Params = struct {
			Filter OF1 `json:"filter"`
		}{OF1{[]int{obj.(int)}, objType}}
	case []int:
		Params = struct {
			Filter OF1 `json:"filter"`
		}{OF1{obj.([]int), objType}}
	case string:
		Params = struct {
			Filter OF2 `json:"filter"`
		}{OF2{obj.(string), objType}}
	}

	data, err := a.Request("cmdb.objects.read", &Params)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}

func (a *Api) GetObjectsByType(objType string) (GenericResponse, error) {
	var Params interface{}
	Params = struct {
		Filter OSF1 `json:"filter"`
	}{OSF1{objType}}

	data, err := a.Request("cmdb.objects.read", &Params)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}

func (a *Api) GetObjTypeCat(query interface{}) (GenericResponse, error) {

	var CustomStruct interface{}
	switch query.(type) {
	case int:
		CustomStruct = struct {
			Type int `json:"type"`
		}{query.(int)}
	case string:
		CustomStruct = struct {
			Type string `json:"type"`
		}{query.(string)}
	}

	data, err := a.Request("cmdb.object_type_categories.read", CustomStruct)
	if err != nil {
		return GenericResponse{}, err
	}

	return TypeAssertResult(data)
}

// Create Object
func (a *Api) CreateObj(Params interface{}) (GenericResponse, error) {

	data, err := a.Request("cmdb.object.create", &Params)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}

// Update Object
func (a *Api) UpdateObj(Params interface{}) (GenericResponse, error) {

	data, err := a.Request("cmdb.object.update", &Params)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}

// Delete/Archive/Purge Object
func (a *Api) DeleteObj(deleteMe interface{}) (GenericResponse, error) {

	var Params interface{}
	switch deleteMe.(type) {
	case int:
		Params = struct {
			Id int `json:"id"`
		}{deleteMe.(int)}
	default:
		Params = deleteMe
	}

	data, err := a.Request("cmdb.object.delete", &Params)
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
			Id int `json:"id"`
		}{ids.(int)}
	case []int:
		Params = struct {
			Ids []int `json:"ids"`
		}{ids.([]int)}
	default:
		return GenericResponse{}, errors.New("Input type is not int or []int")
	}

	data, err := a.Request("cmdb.object.quickpurge", &Params)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}
