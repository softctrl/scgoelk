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
// Type of text analizers of the Elasticsearch.
//
type Analizer string

const (
	SNOWBALL_ANALYZER = Analizer("snowball")
	STANDARD_ANALYZER = Analizer("standard")
)

//
// Type of operations of the Elasticsearch.
//
type Operation string

const (
	AND_OPERATION = Operation("and")
	OR_OPERATION  = Operation("OR")
)

const (
	COMMA        = ','
	DEFUALT_SIZE = 20
)
