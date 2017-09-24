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

// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-uri-request.html

import (
	jsn "encoding/json"
)

//
// =============================================================================
//

//
// Based on:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-filter-context.html
//
type Query struct {
	Term jsn.RawMessage `json:"term,omitempty"` // { "term" : { "user" : "kimchy" } }
	Bool Bool           `json:"bool,omitempty"`
}

//
// =============================================================================
//
