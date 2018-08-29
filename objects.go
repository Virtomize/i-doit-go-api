package goidoit

import (
	"errors"
	"strings"
)

// GetObject function returns an object matching the given query
func (a *API) GetObject(query interface{}) (GenericResponse, error) {

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

// GetObjectByType function return an object of a specific type
func (a *API) GetObjectByType(objType string, obj interface{}) (GenericResponse, error) {
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

// GetObjectsByType returns multiple objects of the given type
func (a *API) GetObjectsByType(objType string) (GenericResponse, error) {
	var Params interface{}
	if strings.Contains(objType, "C__CATG") {
		Params = struct {
			Type string `json:"type_title"`
		}{objType}
	} else {
		Params = struct {
			Filter OSF1 `json:"filter"`
		}{OSF1{objType}}
	}

	data, err := a.Request("cmdb.objects.read", &Params)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}

// GetObjTypeCat returns the objecttype category
func (a *API) GetObjTypeCat(query interface{}) (GenericResponse, error) {

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

// CreateObj creates an object
func (a *API) CreateObj(Params interface{}) (GenericResponse, error) {

	data, err := a.Request("cmdb.object.create", &Params)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}

// UpdateObj updates an object
func (a *API) UpdateObj(Params interface{}) (GenericResponse, error) {

	data, err := a.Request("cmdb.object.update", &Params)
	if err != nil {
		return GenericResponse{}, err
	}
	return TypeAssertResult(data)
}

// DeleteObj delete/archive/purge an object
func (a *API) DeleteObj(deleteMe interface{}) (GenericResponse, error) {

	var Params interface{}
	switch deleteMe.(type) {
	case int:
		Params = struct {
			ID int `json:"id"`
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
func (a *API) Quickpurge(ids interface{}) (GenericResponse, error) {

	var Params interface{}
	switch ids.(type) {
	case int:
		Params = struct {
			ID int `json:"id"`
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
