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
package fts

//
// Struct of the simple query string for the full text search pattern.
//
type SimpleQueryString struct {
	Query     string    `json:"query,omitempty"`
	Analyzer  Analizer  `json:"analyzer,omitempty"`
	Fields    []string  `json:"fields,omitempty"`
	DefaultOp Operation `json:"default_operator,omitempty"`
}

//
// Slice type of SimpleQueryString.
//
type SimpleQueryStrings []SimpleQueryString
