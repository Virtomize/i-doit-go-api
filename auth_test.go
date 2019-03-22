package goidoit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testValuesNewAPI struct {
	URL, APIKEY string
	Error       error
}

func TestNewAPI(t *testing.T) {

	testValues := []testValuesNewAPI{
		{"", "test", fmt.Errorf("url or apikey empty")},
		{"test", "", fmt.Errorf("url or apikey empty")},
		{"test", "test", nil},
	}

	for _, test := range testValues {
		api, err := NewAPI(test.URL, test.APIKEY)
		if test.Error == nil {
			assert.Nil(t, err)
		} else {
			assert.Equal(t, test.Error.Error(), err.Error())
		}
		if err == nil {
			assert.Equal(t, test.URL, api.URL)
			assert.Equal(t, test.APIKEY, api.APIkey)
		}
	}
}

func TestLogin(t *testing.T) {
	server := idoitAPIStub()
	api, err := NewAPI(server.URL, "apikey")
	assert.Nil(t, err)

	test := api.IsLoggedIn()
	assert.False(t, test)

	err = api.Login()
	assert.Nil(t, err)

	api, err = NewLogin(server.URL, "apikey", "username", "password")
	assert.Nil(t, err)
	assert.NotNil(t, api)

	test = api.IsLoggedIn()
	assert.True(t, test)

	err = api.Logout()
	assert.Nil(t, err)
}
