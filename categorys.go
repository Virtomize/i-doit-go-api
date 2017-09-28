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
	"strconv"
)

func (a *Api) GetCategory(objID int, query interface{}) (GenericResponse, error) {

	var CustomStruct interface{}
	switch query.(type) {
	case int:
		CustomStruct = struct {
			ObjID  string `json:"objID"`
			CatgID int    `json:"catgID"`
		}{strconv.Itoa(objID), query.(int)}
	case string:
		CustomStruct = struct {
			ObjID    string `json:"objID"`
			Category string `json:"category"`
		}{strconv.Itoa(objID), query.(string)}
	}

	data, err := a.Request("cmdb.category.read", CustomStruct)
	if err != nil {
		return GenericResponse{}, err
	}

	return TypeAssertResult(data)
}

func (a *Api) CreateCat(ObjId int, CatgId string, Params interface{}) (GenericResponse, error) {

	CustomStruct := struct {
		ObjId  int         `json:"objID"`
		CatgId string      `json:"catgID"`
		Data   interface{} `json:"data"`
	}{ObjId, CatgId, Params}

	data, err := a.Request("cmdb.category.create", CustomStruct)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}

func (a *Api) UpdateCat(ObjId int, CatgId string, Params interface{}) (GenericResponse, error) {

	CustomStruct := struct {
		ObjId  int         `json:"objID"`
		CatgId string      `json:"catgID"`
		Data   interface{} `json:"data"`
	}{ObjId, CatgId, Params}

	data, err := a.Request("cmdb.category.update", CustomStruct)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}
