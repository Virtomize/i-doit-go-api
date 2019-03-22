package goidoit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategories(t *testing.T) {
	server := idoitAPIStub()
	api, err := NewLogin(server.URL, "apikey", "username", "password")
	assert.Nil(t, err)

	res, err := api.GetCat(1, "test")
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.GetCat(1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.GetCat(1, struct{ Foo string }{"bar"})
	assert.Equal(t, fmt.Errorf("query should be of type int or string"), err)
	assert.Equal(t, GenericResponse{}, res)

	res, err = api.CreateCat(1, "C__CAT_G_CUSTOM_FOO", nil)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.CreateCat(1, "C__CAT_G_TEST", nil)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.UpdateCat(1, "C__CAT_G_CUSTOM_FOO", nil)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.UpdateCat(1, "C__CAT_G_TEST", nil)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.DelCatObj(1, "C__CAT_G_CUSTOM_FOO", "C__CAT_E_BAR")
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = api.DelCatObj(1, "C__CAT_G_TEST", "C__CAT_E_BAR")
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
