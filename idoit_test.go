package goidoit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	server := idoitAPIStub()
	api, err := NewLogin(server.URL, "apikey", "username", "password")
	assert.Nil(t, err)

	res, err := api.Request("test", nil)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestSearch(t *testing.T) {
	server := idoitAPIStub()
	api, err := NewLogin(server.URL, "apikey", "username", "password")
	assert.Nil(t, err)

	res, err := api.Search("test")
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestVersion(t *testing.T) {
	server := idoitAPIStub()
	api, err := NewLogin(server.URL, "apikey", "username", "password")
	assert.Nil(t, err)

	res, err := api.Version()
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
