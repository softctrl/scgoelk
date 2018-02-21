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
package scgoelk

import (
	"bytes"
	"net/url"
	scnv "strconv"
)

const (
	SLASH                    = "/"
	COLON                    = ":"
	HEALTH_PATH              = "_cluster/health"
	ALIAS                    = "_alias"
	ALIASES                  = "_aliases"
	BULK_COMMAND             = "_bulk"
	EMPTY_SEARCH_RESULT_JSON = `{"st":0, "body":{"took":1,"timed_out":false,"_shards":{"total":0,"successful":0,"failed":0},"hits":{"total":0,"max_score":null,"hits":[]}}}`
)

var EMPTY_SEARCH_RESULT_JSON_BYTES = []byte(EMPTY_SEARCH_RESULT_JSON)

//
//
//
func MakeIndexUrl(__server string, __port int, __index string) string {

	var buf bytes.Buffer
	buf.WriteString(__server)
	buf.WriteString(COLON)
	buf.WriteString(scnv.Itoa(__port))
	buf.WriteString(SLASH)
	buf.WriteString(__index)
	return buf.String()

}

//
//
//
func MakeIndexTypeUrl(__server string, __port int, __index, __type string) string {

	var buf bytes.Buffer
	buf.WriteString(__server)
	buf.WriteString(COLON)
	buf.WriteString(scnv.Itoa(__port))
	buf.WriteString(SLASH)
	buf.WriteString(__index)
	buf.WriteString(SLASH)
	buf.WriteString(__type)
	return buf.String()

}

//
//
//
func MakeQueryFilterUrl(__elk, __query string) string {

	var buf bytes.Buffer
	buf.WriteString(__elk)
	buf.WriteString(SLASH)
	buf.WriteString("_search?q=\"")
	buf.WriteString(url.QueryEscape(__query))
	buf.WriteString("\"")
	return buf.String()

}

//
//
//
func MakeQueryUrl(__elk string) string {

	var buf bytes.Buffer
	buf.WriteString(__elk)
	buf.WriteString(SLASH)
	buf.WriteString("_search")
	return buf.String()

}

//
//
//
func MakeCommandUrlWithIndex(__server string, __port int, __index, __command string) string {

	var buf bytes.Buffer
	buf.WriteString(__server)
	buf.WriteString(COLON)
	buf.WriteString(scnv.Itoa(__port))
	buf.WriteString(SLASH)
	buf.WriteString(__index)
	buf.WriteString(SLASH)
	buf.WriteString(__command)
	return buf.String()

}

//
//
//
func MakeCommandUrl(__server string, __port int, __command string) string {

	var buf bytes.Buffer
	buf.WriteString(__server)
	buf.WriteString(COLON)
	buf.WriteString(scnv.Itoa(__port))
	buf.WriteString(SLASH)
	buf.WriteString(__command)
	return buf.String()

}
