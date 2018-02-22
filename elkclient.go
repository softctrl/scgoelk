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
package scgoelk

import (
	b "bytes"
	err "errors"

	"net/http"

	"github.com/softctrl/scgotils/schttp"
)

const HTTP_MESSAGE_ERROR = "One error ocurred. if it persistis, please contact support."

//
// Elk client struct
//
type SCElkClient struct {
	_Server string
	_Port   int
}

type SCData struct {
	__Index string
	__Type  string
}

//
// Create a instance object of SCElkClient.
//
func NewSCElkClient() *SCElkClient {
	_res := SCElkClient{_Server: "http://localhost", _Port: 9200}
	return &_res
}

//
// Create a instance object of SCElkClient with the informed parameters.
//
func NewSCElkClientWithValues(__server string, __port int) *SCElkClient {
	_res := SCElkClient{_Server: __server, _Port: __port}
	return &_res
}

//
// Just return a string representation of the SCElkClient instance.
//
func (__obj *SCElkClient) ToString() string {

	var _buff b.Buffer
	_buff.WriteString("Server:")
	_buff.WriteString(__obj._Server)
	_buff.WriteString("Port:")
	_buff.WriteString(string(__obj._Port))
	return _buff.String()

}

//
// Set the Elastic server address.
//
func (__obj *SCElkClient) Server(__server string) *SCElkClient {
	__obj._Server = __server
	return __obj
}

//
// Set the Elastic server port.
//
func (__obj *SCElkClient) Port(__port int) *SCElkClient {
	__obj._Port = __port
	return __obj
}

//
// Perform a Elastic search by url query path.
//
func (__obj *SCElkClient) FindByQuery(__index, __type, __query string) ([]byte, error) {

	_bytes, _err := schttp.Get(MakeQueryFilterUrl(MakeIndexTypeUrl(__obj._Server,
		__obj._Port, __index, __type), __query))
	if _err != nil {
		return nil, _err // err.New(HTTP_MESSAGE_ERROR)
	} else {
		return _bytes, nil
	}

}

//
// Perform a Elastic search by Json query.
//
func (__obj *SCElkClient) FindByJson(__index, __type string, __json []byte) ([]byte, error) {

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	_, _code, _bytes, _err := schttp.Perform(schttp.GET, MakeQueryUrl(
		MakeIndexTypeUrl(__obj._Server, __obj._Port, __index, __type)),
		__json, header)
	if _code != 200 {
		if _err == nil {
			return _bytes, err.New(HTTP_MESSAGE_ERROR)
		} else {
			return _bytes, _err
		}
	} else {
		return _bytes, nil
	}

}

//
// Do whathever you want with this awesome elastisc server.
// Use with love and respect or i will kill you after night comes.
//                                :D i'm messing wiht you, or not >:(
//
func (__obj *SCElkClient) Perform(__command schttp.Method, __url string, __body []byte, __header http.Header) (string, int, []byte, error) {
	return schttp.Perform(__command, MakeCommandUrl(__obj._Server, __obj._Port, __url), __body, __header)
}

//
// Query for the cluster healt of the Elastic server.
//
func (__obj *SCElkClient) ClusterHealt() ([]byte, error) {
	return schttp.Get(MakeCommandUrl(__obj._Server, __obj._Port, HEALTH_PATH))
}

//
// Execute a bulk operation into the Elastic server.
//
func (__obj *SCElkClient) BulkOperation(__document []byte) ([]byte, error) {

	return schttp.PostBody(MakeCommandUrl(__obj._Server, __obj._Port, BULK_COMMAND), __document)

}

//
// Insert a document into a informed index type
//
func (__obj *SCElkClient) Insert(__index, __type string, __document []byte) ([]byte, error) {

	return schttp.PostBody(MakeIndexTypeUrl(__obj._Server, __obj._Port, __index,
		__type), __document)

}

//
// Update a document into a informed index type
//
func (__obj *SCElkClient) Update(__index, __type string, __document []byte) ([]byte, error) {
	return __obj.Insert(__index, __type, __document)
}

//
// Remove a document into a informed index type
//
// TODO under development
func (__obj *SCElkClient) Remove(__index, __type string, __document []byte) ([]byte, error) {

	return nil, err.New("(Remove)under development")

}
