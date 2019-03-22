package goidoit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReport(t *testing.T) {
	server := idoitAPIStub()
	api, err := NewLogin(server.URL, "apikey", "username", "password")
	assert.Nil(t, err)

	res, err := api.GetReport(1)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
