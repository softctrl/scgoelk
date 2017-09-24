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
package elk

import (
	b "bytes"
	f "fmt"

	scgu "github.com/softctrl/gutils/schttp"
)

const (
	_UPDATE_ALIAS_JSON_FMT = `{"actions": [{ "remove": { "index": "%s", "alias": "%s" }},{ "add":    { "index": "%s", "alias": "%s" }}]}`
)

//
// Format the index Alias URL.
//
func (__obj *SCElkClient) _MakeIndexAliasUrl(__index, __alias string) string {

	var _buff b.Buffer
	_buff.WriteString(MakeCommandUrlWithIndex(__obj._Server, __index, ALIAS))
	_buff.WriteString(SLASH)
	_buff.WriteString(__alias)
	return _buff.String()

}

//
// Create a new alias for the informed index.
// Based on:
// https://www.elastic.co/guide/en/elasticsearch/guide/master/index-aliases.html
//
func (__obj *SCElkClient) CreateAlias(__index, __alias string) ([]byte, error) {

	return scgu.Put(__obj._MakeIndexAliasUrl(__index, __alias))

}

//
// Delete an alias for the informed index.
//
func (__obj *SCElkClient) RemoveAlias(__index, __alias string) ([]byte, error) {

	return scgu.Delete(__obj._MakeIndexAliasUrl(__index, __alias))

}

//
//  POST /_aliases
//  {
//      "actions": [
//          { "remove": { "index": "my_index_v1", "alias": "my_index" }},
//          { "add":    { "index": "my_index_v2", "alias": "my_index" }}
//      ]
//  }
//
func (__obj *SCElkClient) ChangeAlias(__index_old, __index_new, __alias string) ([]byte, error) {
	return scgu.PostBody(MakeCommandUrl(__obj._Server, ALIASES),
		[]byte(f.Sprintf(_UPDATE_ALIAS_JSON_FMT, __index_old, __alias, __index_new, __alias)))
}

//
// Get all aliases into an informed index.
//
func (__obj *SCElkClient) GetAliases(__index string) ([]byte, error) {

	return scgu.Get(MakeCommandUrlWithIndex(__obj._Server, __index, ALIAS))

}
