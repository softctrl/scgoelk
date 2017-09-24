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
//
// Based on the elasticsearch result
// {
// 	"error": {
// 		"root_cause": [
// 			{
// 			"type": "illegal_argument_exception",
// 			"reason": "request [/redis/pessoa/_searc] contains unrecognized parameters: [q], [size]"
// 			}
// 		],
// 		"type": "illegal_argument_exception",
// 		"reason": "request [/redis/pessoa/_searc] contains unrecognized parameters: [q], [size]"
// 	},
// 	"status": 400
// }
//
package query

type Error struct {
	RootCause []Error `json:"root_cause,omitempyt"`
	Type      string  `json:"type,omitempyt"`
	Reason    string  `json:"reason,omitempyt"`
	Line      int     `json:"reason,omitempyt"`
	Column    int     `json:"reason,omitempyt"`
}
