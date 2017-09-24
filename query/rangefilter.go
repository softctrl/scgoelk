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
// Based on:
// https://www.elastic.co/guide/en/elasticsearch/guide/master/_ranges.html
//
package query

import (
	b "bytes"
	jsn "encoding/json"
	"fmt"
)

/*
{
    "query" : {
        "constant_score" : {
            "filter" : {
                "range" : {
                    "price" : {
                        "gte" : 20,
                        "lt"  : 40
                    }
                }
            }
        }
    }
}
*/

type Option string

const (
	GT  = Option("gt")
	LT  = Option("lt")
	GTE = Option("gte")
	LTE = Option("lte")
)

const (
	_OPTION_AND_VALUE_FMT = `"%s":%s`
	_VALUE_INT            = "%d"
	_VALUE_STR            = "%s"
	_VALUE_BOOL           = "%t"
	_RANGE_INT_FMT        = `{"range":{"%s":{%s}}}`
)

//
//
//
type _From struct {
	Option Option
	Value  jsn.RawMessage
}

//
//
//
func (__obj *_From) ToJson() []byte {
	return []byte(fmt.Sprintf(_OPTION_AND_VALUE_FMT, __obj.Option, string(__obj.Value)))
}

//
//
//
type _To _From

//
//
//
func (__obj *_To) ToJson() []byte {
	return []byte(fmt.Sprintf(_OPTION_AND_VALUE_FMT, __obj.Option, string(__obj.Value)))
}

//
//
//
type Range struct {
	Field string
	From  *_From
	To    *_To
}

//
//
//
func _MustRangeInt(__field string, __from, __to *int) *Range {

	_res := Range{
		Field: __field,
	}
	if __from != nil {
		_res.From = &_From{
			Option: GTE,
			Value:  []byte(fmt.Sprintf(_VALUE_INT, *__from)),
		}
	}
	if __to != nil {
		_res.To = &_To{
			Option: LTE,
			Value:  []byte(fmt.Sprintf(_VALUE_INT, *__to)),
		}
	}
	return &_res

}

//
// Based on:
//  {
//      "range": {
//          "__field": {
//              "Option": __from,
//              "Option": __to
//          }
//      }
//  }
//
func (__obj *Range) ToJson() jsn.RawMessage {

	var _buff b.Buffer
	if __obj.From != nil {
		_buff.Write(__obj.From.ToJson())
		if __obj.To != nil {
			_buff.WriteByte(',')
			_buff.Write(__obj.To.ToJson())
		}
	} else if __obj.To != nil {
		_buff.Write(__obj.To.ToJson())
	}
	return []byte(fmt.Sprintf(_RANGE_INT_FMT, __obj.Field, _buff.String()))

}
