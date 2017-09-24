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

import (
	"bytes"
	jsn "encoding/json"
	"fmt"

	scelkqry "github.com/softctrl/elk/query"
)

type QuerySQS struct {
	Query SimpleQueryString `json:"simple_query_string,omitempty"`
}

type FtsElkQuery struct {
	Query  QuerySQS
	SortBy []scelkqry.Sort
	From   int
	Size   int
}

type FtsElkQueries []FtsElkQuery

func NewFtsElkQuery(__query string, __analyzer Analizer, __default_op Operation, __fields []string) *FtsElkQuery {
	_res := FtsElkQuery{}
	{
		_res.Query = QuerySQS{}
		{
			_res.Query.Query = SimpleQueryString{}
			{
				_res.Query.Query.Query = __query
				_res.Query.Query.Analyzer = __analyzer
				_res.Query.Query.Fields = __fields
				_res.Query.Query.DefaultOp = __default_op
			}
		}
		_res.Size = DEFUALT_SIZE
	}
	return &_res
}

func NewFtsElkQueryWithSort(__query string, __analyzer Analizer, __default_op Operation, __fields []string, __sorts []scelkqry.Sort) *FtsElkQuery {
	_res := NewFtsElkQuery(__query, __analyzer, __default_op, __fields)
	if __sorts != nil && len(__sorts) > 0 {
		_res.SortBy = __sorts
	}
	return _res
}

func NewFtsElkQueryDefault(__query string, __fields []string) *FtsElkQuery {
	return NewFtsElkQuery(__query, SNOWBALL_ANALYZER, AND_OPERATION, __fields)
}

func NewFtsElkQueryWithSortDefault(__query string, __fields []string, __sorts []scelkqry.Sort) *FtsElkQuery {
	return NewFtsElkQueryWithSort(__query, STANDARD_ANALYZER, AND_OPERATION, __fields, __sorts)
}

func NewFtsElkQueryWithSortDefaultSnowball(__query string, __fields []string, __sorts []scelkqry.Sort) *FtsElkQuery {
	return NewFtsElkQueryWithSort(__query, SNOWBALL_ANALYZER, AND_OPERATION, __fields, __sorts)
}

//
// Parse a FtsElkQuery object instance to a json []byte
//
func (__obj FtsElkQuery) ToJson() ([]byte, error) {

	var _buff bytes.Buffer
	_buff.WriteByte('{')
	{
		// Query marshal:
		_query, _err := jsn.Marshal(__obj.Query)
		var _comma bool
		if _err != nil {
			return nil, _err
		} else {
			_buff.WriteString(`"query":`)
			_buff.Write(_query)
			_comma = true
		}
		// =====================================================================

		// Sort marshal:
		if __obj.SortBy != nil && len(__obj.SortBy) > 0 {
			if _comma {
				_buff.WriteByte(',')
			}
			_comma = false
			_buff.WriteString(`"sort":[`)
			{
				for _, _sort := range __obj.SortBy {
					if _err != nil {
						return nil, _err
					} else {
						if _comma {
							_buff.WriteByte(COMMA)
						}
						if _json_sort, _err := _sort.Sort_ToJson(); _err == nil {
							_buff.Write(_json_sort)
							_comma = true
						}
					}
				}

			}
			_buff.WriteByte(']')
			_comma = true
		}
		// =====================================================================

		// From marshal:
		if __obj.From > 0 {
			if _comma {
				_buff.WriteByte(',')
			}
			_buff.WriteString(fmt.Sprintf(`"from":%d`, __obj.Size))
		}
		// =====================================================================

		// Size marshal:
		if __obj.Size > 0 {
			if _comma {
				_buff.WriteByte(',')
			}
			_buff.WriteString(fmt.Sprintf(`"size":%d`, __obj.Size))
		}
		// =====================================================================

	}
	_buff.WriteByte('}')

	return _buff.Bytes(), nil
}
