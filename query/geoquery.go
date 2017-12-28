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
	jsn "encoding/json"
)

const (
	_GEO_DISTANCE = "_geo_distance"
)

//
// Based on:
//
//	{
//		"query":{
//			"bool":{
//				"must":[
//					{"match":{"fn":"Jhon Snow"}},
//					{"match":{"mn":"Sueli"}}
//				],
//				"filter":{
//					"geo_distance":{
//						"distance":"5km",
//						"ns.hs.lc":[-43.8786466,-19.9167531]
//					}
//				}
//			}
//		},
//		"sort":[
//			{"_geo_distance":{ "ns.hs.lc":[-43.8786466,-19.9167531], "order":"asc", "unit":"km" } },
//			{"cs":{"order":"desc"}},
//			{"_score":{"order":"desc"}}
//		]
//	}
//

//
// Constructor
//
func NewGeoQuery(__field string, __latitude float32, __longitude float32,
	__distance int, __unit Unit) *ElkQuery {
	return &ElkQuery{
		Query: Query{
			Bool: Bool{
				Filter: (&GeoFilter{
					Distance:  MakeDistance(__distance, __unit),
					Field:     __field,
					Latitude:  __latitude,
					Longitude: __longitude}).ToJson()}}}

}

//
// Constructor
//
func NewGeoQueryKmSorted(__field string, __latitude float32, __longitude float32,
	__distance int) *ElkQuery {

	return NewGeoQuery(__field, __latitude, __longitude, __distance, UNIT_KM).
		AddGeoSortWithValues(__field, __latitude, __longitude, ASCENDING, UNIT_KM)

}

//
// Add a geo filter for the Elasticsearch query instance.
//
func (__obj *ElkQuery) AddGeoFilter(__filter *GeoFilter) *ElkQuery {

	__obj.Query.Bool.Filter = __filter.ToJson()
	return __obj

}

//
// Add a geo filter for the Elasticsearch query instance and then make a sort.
//
func (__obj *ElkQuery) AddGeoFilterAndSort(__filter *GeoFilter) *ElkQuery {
	__obj.AddGeoFilter(__filter).AddGeoSortWithValues(__filter.Field, __filter.Latitude,
		__filter.Longitude, DESCENDING, UNIT_KM)
	return __obj
}

//
// Add a sort for the Elasticsearch instance query.
//
func (__obj *ElkQuery) AddSort(__sort *Sort) *ElkQuery {

	if _json, _err := __sort.Sort_ToJson(); _err == nil { // TODO log here
		__obj.Sort = append(__obj.Sort, _json)
	}
	return __obj
}

//
// Add a sort for the Elasticsearch instance query.
//
func (__obj *ElkQuery) AddSortWithValues(__field string, __order Order) *ElkQuery {
	return __obj.AddSort(NewSort(__field, __order))
}

//
// Add a GEO sort for the Elasticsearch instance query.
//
func (__obj *ElkQuery) AddGeoSort(__sort *GeoSort) *ElkQuery {
	__obj.Sort = append(__obj.Sort, __sort.GeoSort_ToJson())
	return __obj
}

//
// Add a GEO sort for the Elasticsearch instance query.
//
func (__obj *ElkQuery) AddGeoSortWithValues(__field string, __latitude float32,
	__longitude float32, __order Order, __unit Unit) *ElkQuery {

	return __obj.AddGeoSort(&GeoSort{
		Field:     __field,
		Latitude:  __latitude,
		Longitude: __longitude,
		Order:     __order,
		Unit:      __unit})

}

//
// Return the Json representation of this Elasticsearch query instance.
//
func (__obj *ElkQuery) ToJson() ([]byte, error) {
	return jsn.Marshal(__obj)
}
