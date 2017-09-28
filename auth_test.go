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
	"testing"
)

func TestNewApi(t *testing.T) {
	var tests = []constructortest{
		{[]string{"", "test"}, false},
		{[]string{"test", ""}, false},
		{[]string{"test", "test"}, true},
	}
	for i := range tests {
		_, err := NewApi(tests[i].input[0], tests[i].input[1])
		if tests[i].expected {
			if err != nil {
				t.Errorf("Expected nil got %v\n", err)
			}
		} else {
			if err == nil {
				t.Errorf("Expected some error got %v\n", err)
			}
		}
	}
}
