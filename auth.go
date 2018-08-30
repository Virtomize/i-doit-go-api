package goidoit

import (
	"errors"
)

// NewAPI implements API constructor
func NewAPI(url string, apikey string) (*API, error) {
	if len(url) != 0 && len(apikey) != 0 {
		a := API{url, apikey, "", "", ""}
		return &a, nil
	}
	return nil, errors.New("url or apikey empty")
}

// NewLogin implements API constructor using login x-rpc auth
func NewLogin(url string, apikey string, username string, password string) (*API, error) {
	if len(url) != 0 && len(apikey) != 0 && len(username) != 0 && len(password) != 0 {
		a := API{url, apikey, username, password, ""}
		err := a.Login()
		if err != nil {
			return nil, err
		}
		return &a, nil
	}
	return nil, errors.New("empty parameter")
}

// Login implements idoit.login method
func (a *API) Login() error {
	data, err := a.Request("idoit.login", nil)
	if err != nil {
		return err
	}
	res, err := TypeAssertResult(data)
	if err != nil {
		return err
	}

	if len(res.Result) != 0 {
		a.SessionID = res.Result[0]["session-id"].(string)
	} else {
		return errors.New(res.Error.Message)
	}
	return nil
}

// IsLoggedIn checks the login status
func (a API) IsLoggedIn() bool {
	if len(a.SessionID) != 0 {
		return true
	}
	return false
}

// Logout implements idoit.logout method
func (a API) Logout() error {
	_, err := a.Request("idoit.logout", nil)
	if err != nil {
		return err
	}
	return nil
}
