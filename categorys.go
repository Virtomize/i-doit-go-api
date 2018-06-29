package goidoit

import (
	"strconv"
	"strings"
)

// GetCat returns the category defined by query for an object given by its object ID
func (a *Api) GetCat(objID int, query interface{}) (GenericResponse, error) {

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

// CreateCat creates a category given by its category ID and the defined content in params for a given object ID
func (a *Api) CreateCat(ObjId int, CatgId string, Params interface{}) (GenericResponse, error) {

	var CustomStruct interface{}

	if strings.Contains(CatgId, "_CUSTOM_") {
		CustomStruct = struct {
			ObjId  int         `json:"objID"`
			CatgId string      `json:"category"`
			Data   interface{} `json:"data"`
		}{ObjId, CatgId, Params}
	} else {
		CustomStruct = struct {
			ObjId  int         `json:"objID"`
			CatgId string      `json:"catgID"`
			Data   interface{} `json:"data"`
		}{ObjId, CatgId, Params}
	}

	data, err := a.Request("cmdb.category.create", CustomStruct)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}

// UpdateCat updates a category given by its category ID and the defined content in params for a given object ID
func (a *Api) UpdateCat(ObjId int, CatgId string, Params interface{}) (GenericResponse, error) {

	var CustomStruct interface{}

	if strings.Contains(CatgId, "_CUSTOM_") {
		CustomStruct = struct {
			ObjId  int         `json:"objID"`
			CatgId string      `json:"category"`
			Data   interface{} `json:"data"`
		}{ObjId, CatgId, Params}
	} else {
		CustomStruct = struct {
			ObjId  int         `json:"objID"`
			CatgId string      `json:"catgID"`
			Data   interface{} `json:"data"`
		}{ObjId, CatgId, Params}
	}

	data, err := a.Request("cmdb.category.update", CustomStruct)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}

// DelCatObj deletes a given category for a given object
func (a *Api) DelCatObj(ObjId int, CatgId string, CateId string) (GenericResponse, error) {

	var CustomStruct interface{}

	if strings.Contains(CatgId, "_CUSTOM_") {
		CustomStruct = struct {
			ObjId  int    `json:"objID"`
			CatgId string `json:"category"`
			CateId string `json:"id"`
		}{ObjId, CatgId, CateId}
	} else {
		CustomStruct = struct {
			ObjId  int    `json:"objID"`
			CatgId string `json:"catgID"`
			CateId string `json:"id"`
		}{ObjId, CatgId, CateId}
	}

	data, err := a.Request("cmdb.category.delete", CustomStruct)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}
