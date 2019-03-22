package goidoit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetObject(t *testing.T) {
	server := idoitAPIStub()
	api, err := NewLogin(server.URL, "apikey", "username", "password")
	assert.Nil(t, err)

	res, err := api.GetObject("test")
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.GetObject([]int{1, 2})
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.GetObject(1)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.GetObject(struct{ Foo string }{"bar"})
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestGetObjectByType(t *testing.T) {
	server := idoitAPIStub()
	api, err := NewLogin(server.URL, "apikey", "username", "password")
	assert.Nil(t, err)

	res, err := api.GetObjectByType("VirtualMachines", "test")
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.GetObjectByType("VirtualMachines", 1)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.GetObjectByType("VirtualMachines", []int{1, 2})
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.GetObjectByType("VirtualMachines", struct{ Foo string }{"bar"})
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestGetObjectsByType(t *testing.T) {
	server := idoitAPIStub()
	api, err := NewLogin(server.URL, "apikey", "username", "password")
	assert.Nil(t, err)

	res, err := api.GetObjectsByType("C__CATG_TEST")
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.GetObjectsByType("VirtualMachines")
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestGetObjTypeCat(t *testing.T) {
	server := idoitAPIStub()
	api, err := NewLogin(server.URL, "apikey", "username", "password")
	assert.Nil(t, err)

	res, err := api.GetObjTypeCat("test")
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.GetObjTypeCat(1)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestAlterObj(t *testing.T) {
	server := idoitAPIStub()
	api, err := NewLogin(server.URL, "apikey", "username", "password")
	assert.Nil(t, err)

	res, err := api.CreateObj("test")
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.UpdateObj("test")
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.DeleteObj("test")
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.DeleteObj(1)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.Quickpurge(1)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.Quickpurge([]int{1, 2})
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.Quickpurge("test")
	assert.Equal(t, fmt.Errorf("Input type is not int or []int"), err)
	assert.Equal(t, GenericResponse{}, res)
}
