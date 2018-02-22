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
package res

import (
	"bytes"
	jsn "encoding/json"

	sc "github.com/softctrl/scgoelk/query"
)

const (
	OPEN_BRACKET  = '['
	CLOSE_BRACKET = ']'
	COMMA         = ','
)

type Result struct {
	Took     int      `json:"took,omitempyt"`
	Timedout bool     `json:"timed_out,omitempyt"`
	Shards   Shard    `json:"_shards,omitempyt"`
	Hits     Data     `json:"hits,omitempyt"`
	Error    sc.Error `json:"error,omitempyt"`
	Status   int      `json:"status,omitempyt"`
}

type Results []Result

//
// Transform a Result instance to the []byte Json
//
func (__obj *Result) ToJson() ([]byte, error) {

	_bytes, _err := jsn.Marshal(__obj)
	if _err != nil {
		return nil, _err
	}
	return _bytes, nil

}

//
// Parses the Json []byte to a valid Result instance
//
func ResultFromJson(__json []byte) (*Result, error) {

	var _res Result
	_err := jsn.Unmarshal(__json, &_res)
	if _err != nil {
		return nil, _err
	}
	return &_res, nil

}

//
// Remove all data from hits, and return a []byte json with this data
//
func (__obj *Result) RemoveData() []byte {

	if __obj.Hits.Total > 0 {
		var _buff bytes.Buffer
		_buff.WriteByte(OPEN_BRACKET)

		_hits := __obj.Hits.Hits

		_buff.Write(_hits[0].Source)

		var _len int = len(__obj.Hits.Hits)

		for _idx := 1; _idx < _len; _idx++ {
			_buff.WriteByte(COMMA)
			_buff.Write(_hits[_idx].Source)
		}

		_buff.WriteByte(CLOSE_BRACKET)

		return _buff.Bytes()

	} else {
		return nil
	}

}

//
// TODO change here to accept a 'Stream' like
//
func (__obj *Result) RemoveData2() []byte {

	if __obj.Hits.Total > 0 {
		var _buff bytes.Buffer
		_buff.WriteByte(OPEN_BRACKET)

		_hits := __obj.Hits.Hits

		_buff.Write(_hits[0].Source)

		var _len int = len(__obj.Hits.Hits)

		for _idx := 1; _idx < _len; _idx++ {
			_buff.WriteByte(COMMA)
			_buff.Write(_hits[_idx].Source)
		}

		_buff.WriteByte(CLOSE_BRACKET)

		return _buff.Bytes()

	} else {
		return nil
	}

}
