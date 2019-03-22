package goidoit

import (
	"fmt"
	"strconv"
	"strings"
)

// GetCat returns the category defined by query for an object given by its object ID
func (a *API) GetCat(objID int, query interface{}) (GenericResponse, error) {

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
	default:
		return GenericResponse{}, fmt.Errorf("query should be of type int or string")
	}

	data, err := a.Request("cmdb.category.read", CustomStruct)
	if err != nil {
		return GenericResponse{}, err
	}

	return TypeAssertResult(data)
}

// CreateCat creates a category given by its category ID and the defined content in params for a given object ID
func (a *API) CreateCat(ObjID int, CatgID string, Params interface{}) (GenericResponse, error) {

	var CustomStruct interface{}

	if strings.Contains(CatgID, "_CUSTOM_") {
		CustomStruct = struct {
			ObjID  int         `json:"objID"`
			CatgID string      `json:"category"`
			Data   interface{} `json:"data"`
		}{ObjID, CatgID, Params}
	} else {
		CustomStruct = struct {
			ObjID  int         `json:"objID"`
			CatgID string      `json:"catgID"`
			Data   interface{} `json:"data"`
		}{ObjID, CatgID, Params}
	}

	data, err := a.Request("cmdb.category.create", CustomStruct)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}

// UpdateCat updates a category given by its category ID and the defined content in params for a given object ID
func (a *API) UpdateCat(ObjID int, CatgID string, Params interface{}) (GenericResponse, error) {

	var CustomStruct interface{}

	if strings.Contains(CatgID, "_CUSTOM_") {
		CustomStruct = struct {
			ObjID  int         `json:"objID"`
			CatgID string      `json:"category"`
			Data   interface{} `json:"data"`
		}{ObjID, CatgID, Params}
	} else {
		CustomStruct = struct {
			ObjID  int         `json:"objID"`
			CatgID string      `json:"catgID"`
			Data   interface{} `json:"data"`
		}{ObjID, CatgID, Params}
	}

	data, err := a.Request("cmdb.category.update", CustomStruct)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}

// DelCatObj deletes a given category for a given object
func (a *API) DelCatObj(ObjID int, CatgID string, CateID string) (GenericResponse, error) {

	var CustomStruct interface{}

	if strings.Contains(CatgID, "_CUSTOM_") {
		CustomStruct = struct {
			ObjID  int    `json:"objID"`
			CatgID string `json:"category"`
			CateID string `json:"id"`
		}{ObjID, CatgID, CateID}
	} else {
		CustomStruct = struct {
			ObjID  int    `json:"objID"`
			CatgID string `json:"catgID"`
			CateID string `json:"id"`
		}{ObjID, CatgID, CateID}
	}

	data, err := a.Request("cmdb.category.delete", CustomStruct)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}
