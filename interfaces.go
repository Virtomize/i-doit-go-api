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

// basic api interface
type ApiMethods interface {
	// idoit.login using X-RPC-AUTH
	// goidoit.NewLogin wraps this function
	// eg. idoit.NewLogin(<url>, <api-key>, <username>, <password>)
	Login() error

	// check login state
	IsLoggedIn() bool

	// Destroy X-RPC-Session
	Logout() error

	// i-doit api request structure
	// as defined in https://kb.i-doit.com/pages/viewpage.action?pageId=7831613
	// also there is a list of methods available
	Request(string, interface{}) (Response, error)

	// search CMDB using a string
	//
	// The search function does handle type assertions
	// for simple output usage
	Search(string) (GenericResponse, error)

	// version information
	Version() (GenericResponse, error)

	// get object(s) data,
	// input can be of type int, []int, string or a custom filter struct
	GetObject(interface{}) (GenericResponse, error)

	// shortcut function for Type Filter using C__OBJTYPE const
	GetObjectByType(string, interface{}) (GenericResponse, error)

	// shortcut function for get all objects of type <C__OBJTYPE__>
	GetObjectsByType(string) (GenericResponse, error)

	// get Object Type Categories using category id or constant
	// eg. GetObjTypeCat("C__OBJTYPE__PERSON")
	// or GetObjTypeCat(50)
	GetObjTypeCat(interface{}) (GenericResponse, error)

	// create objects using struct
	CreateObj(interface{}) (GenericResponse, error)

	// update object
	UpdateObj(interface{}) (GenericResponse, error)

	/* Delete/Archive/Purge object, input can be int (using the object id) or
	data := struct {
		Id int `json:"id"`
		Status string `json:"status"`
	}

	where Id represents the object id
	and Status can be
	C__RECORD_STATUS__ARCHIVED
	C__RECORD_STATUS__DELETED
	C__RECORD_STATUS__PURGE
	*/
	DeleteObj(interface{}) (GenericResponse, error)

	// fast delete option where archive, delete and purge will be executed one after another
	// accepts id or []id as input
	Quickpurge(interface{}) (GenericResponse, error)

	// get categorys for object using object id and category id or category constant
	// eg. GetCat(20,50)
	// or GetCat(20,"C__CATG__CUSTOM_FIELD_TEST")
	GetCat(int, interface{}) (GenericResponse, error)

	// create category using struct
	CreateCat(int, string, interface{}) (GenericResponse, error)

	// update category using struct
	UpdateCat(int, string, interface{}) (GenericResponse, error)

	// delete entry from category
	DelCatObj(int, string, string) (GenericResponse, error)

	// get report data via id
	GetReport(int) (GenericResponse, error)
}
