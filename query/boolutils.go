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
	f "fmt"

	scs "github.com/softctrl/gutils"
)

const (
	_TERMS_STR_FMT  = `{"terms":{"%s":%s}}`
	_MATCH_BOOL_FMT = `{"match":{"%s":%t}}`
	_MATCH_STR_FMT  = `{"match":{"%s":"%s"}}`
	_MATCH_INT_FMT  = `{"match":{"%s":%d}}`
)

//
//
//
func _TermsStrings_ToJson(__field string, __values []string) []byte {
	return []byte(f.Sprintf(_TERMS_STR_FMT, __field, scs.JoinQuoted(__values, ",")))
}

//
//
//
func _MatchBool_ToJson(__field string, __value bool) []byte {
	return []byte(f.Sprintf(_MATCH_BOOL_FMT, __field, __value))
}

//
//
//
func _MatchString_ToJson(__field, __value string) []byte {
	return []byte(f.Sprintf(_MATCH_STR_FMT, __field, __value))
}

//
//
//
func _MatchInt_ToJson(__field string, __value int) []byte {
	return []byte(f.Sprintf(_MATCH_INT_FMT, __field, __value))
}
