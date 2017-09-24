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
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-sort.html
//
package query

import (
	"bytes"
	jsn "encoding/json"
	"fmt"
)

const (
	SCORE_JSON = "{ \"%s\" : {\"order\":\"%s\"}}"
)

//
// Based on:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-sort.html#_sort_values
// =============================================================================
type Order string

const (
	ASCENDING  = Order("asc")  // Sort in ascending order
	DESCENDING = Order("desc") // Sort in descending order
)

// =============================================================================

//
// Based on:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-sort.html#_sort_mode_option
// =============================================================================
type Mode string

const (
	MAXIMUM   = Mode("max")    // Pick the highest value.
	MINIMUM   = Mode("min")    // Pick the lowest value.
	SUMMARISE = Mode("sum")    // Use the sum of all values as sort value. Only applicable for number based array fields.
	AVERAGE   = Mode("avg")    // Use the average of all values as sort value. Only applicable for number based array fields.
	MEDIAN    = Mode("median") // Use the median of all values as sort value. Only applicable for number based array fields.
)

// =============================================================================

//
// Based on:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-sort.html#_missing_values
// =============================================================================
type Missing string

const (
	FIRST = Missing("_first")
	LAST  = Missing("_last")
)

// =============================================================================

//
// Based on:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-sort.html#_ignoring_unmapped_fields
// =============================================================================
type UnmappedType string

const (
	LONG = UnmappedType("long")
	// TODO put all others here
)

// =============================================================================

//
// Based on:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-sort.html#geo-sorting
// =============================================================================

type DistanceType string

const (
	SLOPPY_ARC = DistanceType("sloppy_arc")
	ARC        = DistanceType("arc")
	PLANE      = DistanceType("plane")
)

// =============================================================================

//
// Based on:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-sort.html#geo-sorting
// =============================================================================

type Unit string

const (
	UNIT_KM = Unit("km")
	UNIT_M  = Unit("m")
)

// =============================================================================

//
// Based on:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-sort.html#geo-sorting
// =============================================================================

type GeoSort struct {
	Field        string
	Latitude     float32
	Longitude    float32
	Order        Order
	Unit         Unit
	DistanceType DistanceType
}

func (__obj *GeoSort) GeoSort_ToJson() []byte {
	return []byte(fmt.Sprintf(`{"_geo_distance":{ "%s":[%f,%f], "order":"%s", "unit":"%s" } }`, __obj.Field, __obj.Longitude, __obj.Latitude, __obj.Order, __obj.Unit))
}

// =============================================================================

//
// Based on:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-sort.html#_script_based_sorting
// =============================================================================

type SortScript struct {
	Type   string `json:"type,omitempty"`
	Script struct {
		Lang   string         `json:"lang,omitempty"`
		Inline string         `json:"inline,omitempty"`
		Params jsn.RawMessage `json:"params,omitempty"`
	}
	Order Order `json:"order,omitempty"`
}

// =============================================================================

//https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-sort.html

type _Config struct {
	Order        Order          `json:"order,omitempty"`
	Mode         Mode           `json:"mode,omitempty"`
	NestedPath   string         `json:"nested_path,omitempty"`
	NestedFilter jsn.RawMessage `json:"nested_filter,omitempty"` // TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-sort.html#_nested_sorting_example
	Missing      Missing        `json:"missing,omitempty"`
	UnmappedType UnmappedType   `json:"unmapped_type,omitempty"`
	GeoSort      jsn.RawMessage `json:"_geo_distance,omitempty"`
}

type Sort struct {
	Field  string
	Config _Config
}

func (__obj *_Config) _Sort_Config_ToJson() ([]byte, error) {
	return jsn.Marshal(__obj)
}

func NewSort(__field string, __order Order) *Sort {
	_res := Sort{Field: __field}
	_res.Config.Order = __order
	return &_res

}

//
// Based on: { "_score" : {"order":"desc"} }
//
//func (__obj *SortBy) ToJson() []byte { tests
const SORT_JSON = `{"%s":%s}`

func (__obj *Sort) Sort_ToJson() ([]byte, error) {

	if _cfg, _err := __obj.Config._Sort_Config_ToJson(); _err != nil {
		return nil, _err
	} else {
		var _buff bytes.Buffer
		_buff.WriteString(fmt.Sprintf(SORT_JSON, __obj.Field, string(_cfg)))
		return _buff.Bytes(), nil
	}

}
