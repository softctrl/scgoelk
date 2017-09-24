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
	"fmt"
)

const (
	GEO_FILTER_JSON_FMT = `{"geo_distance":{"distance":"%s","%s":[%f,%f]}}`
)

type GeoFilter struct {
	Distance  string
	Field     string
	Latitude  float32
	Longitude float32
}

//
//
//
func (__obj *GeoFilter) ToJson() []byte {
	return []byte(fmt.Sprintf(GEO_FILTER_JSON_FMT, __obj.Distance, __obj.Field, __obj.Longitude, __obj.Latitude))
}
