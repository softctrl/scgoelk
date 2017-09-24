//
// Author:
//  Carlos Timoshenko
//  carlostimoshenkorodrigueslopes@gmail.com
//
//  https://github.com/softctrl
//
// This project is free software; you can redistribute it and/or
// modify it under the terms of the GNU Lesser General Public
// License as published by the Free Software Foundation; either
// version 3 of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Lesser General Public License for more details.
//
package query

import (
	jsn "encoding/json"
	"fmt"
)

const (
	_BOOL_FMT = `{"bool": %s}`
)

//
// Based on:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
//
type Bool struct {
	Must    []jsn.RawMessage `json:"must,omitempty"`     //
	Should  []jsn.RawMessage `json:"should,omitempty"`   //
	MustNot []jsn.RawMessage `json:"must_not,omitempty"` //
	Filter  jsn.RawMessage   `json:"filter,omitempty"`   //
}

//
//
// =============================================================================
func NewBool() *Bool {

	_res := Bool{}
	return &_res

}

//
//
//
func (__obj *Bool) Bool_ToJson() ([]byte, error) {
	if _jsn, _err := jsn.Marshal(__obj); _err != nil {
		return nil, _err
	} else {
		return []byte(fmt.Sprintf(_BOOL_FMT, string(_jsn))), nil
	}
}

//
//
//
func (__obj *Bool) MustTermsStrings(__field string, __values []string) *Bool {
	__obj.Must = append(__obj.Must, _TermsStrings_ToJson(__field, __values))
	return __obj
}

//
//
//
func (__obj *Bool) ShouldTermsStrings(__field string, __values []string) *Bool {
	__obj.Should = append(__obj.Should, _TermsStrings_ToJson(__field, __values))
	return __obj
}

//
//
//
func (__obj *Bool) MustBool(__bool Bool) *Bool {
	_json, _ := __bool.Bool_ToJson()
	__obj.Must = append(__obj.Must, _json)
	return __obj
}

//
// Add a Match clausule string for this Bool instance.
//
func (__obj *Bool) MustMatchString(__field, __value string) *Bool {
	__obj.Must = append(__obj.Must, _MatchString_ToJson(__field, __value))
	return __obj
}

//
// Add a Match clausule bool for this Bool instance.
//
func (__obj *Bool) MustMatchBool(__field string, __value bool) *Bool {
	__obj.Must = append(__obj.Must, _MatchBool_ToJson(__field, __value))
	return __obj
}

//
// Add a Match clausule int for this Bool instance.
//
func (__obj *Bool) MustMatchInt(__field string, __value int) *Bool {
	__obj.Must = append(__obj.Must, _MatchInt_ToJson(__field, __value))
	return __obj
}

//
// Add a Range clausule for this Bool instance.
//
func (__obj *Bool) MustRangeInt(__field string, __from, __to *int) *Bool {
	__obj.Must = append(__obj.Must, _MustRangeInt(__field, __from, __to).ToJson())
	return __obj
}
