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
	"errors"
)

// NewApi implements Api constructor
func NewApi(url string, apikey string) (*Api, error) {
	if len(url) != 0 && len(apikey) != 0 {
		a := Api{url, apikey, "", "", ""}
		return &a, nil
	}
	return nil, errors.New("url or apikey empty")
}

// NewLogin implements Api constructor using login x-rpc auth
func NewLogin(url string, apikey string, username string, password string) (*Api, error) {
	if len(url) != 0 && len(apikey) != 0 && len(username) != 0 && len(password) != 0 {
		a := Api{url, apikey, username, password, ""}
		err := a.Login()
		if err != nil {
			return nil, err
		}
		return &a, nil
	}
	return nil, errors.New("empty parameter")
}

// Login implements idoit.login method
func (a *Api) Login() error {
	data, err := a.Request("idoit.login", nil)
	if err != nil {
		return err
	}
	res, err := TypeAssertResult(data)
	if err != nil {
		return err
	}

	if len(res.Result) != 0 {
		a.SessionId = res.Result[0]["session-id"].(string)
	} else {
		return errors.New(res.Error.Data.(map[string]interface{})["error"].(string))
	}
	return nil
}

// IsLoggedIn checks the login status
func (a Api) IsLoggedIn() bool {
	if len(a.SessionId) != 0 {
		return true
	}
	return false
}

// Logout implements idoit.logout method
func (a Api) Logout() error {
	_, err := a.Request("idoit.logout", nil)
	if err != nil {
		return err
	}
	return nil
}
