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
// TODO Under development
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-sort.html#_track_scores
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-post-filter.html
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-highlighting.html
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-rescore.html
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-scroll.html
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-preference.html
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-index-boost.html
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-named-queries-and-filters.html
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-inner-hits.html
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-search-after.html
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-template.html
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/multi-search-template.html
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters-term.html
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters-phrase.html
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters-completion.html
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/suggester-context.html
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html;
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-count.html
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-validate.html
// TODO https://www.elastic.co/guide/en/elasticsearch/reference/current/search-explain.html
//
package query

// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-uri-request.html

import (
	jsn "encoding/json"

	scgu "github.com/softctrl/gutils"
)

type ElkQuery struct {
	// Source   []string `json:"_source,omitempty"` // TODO
	Source         jsn.RawMessage   `json:"_source,omitempty"`         // https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-source-filtering.html#search-request-source-filtering
	Explain        *bool            `json:"explain,omitempty"`         // Enables explanation for each hit on how its score was computed - https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-explain.html#search-request-explain
	Version        *bool            `json:"version,omitempty"`         // Returns a version for each search hit - https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-version.html#search-request-version
	MinScore       float32          `json:"min_score,omitempty"`       // https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-min-score.html
	StoredFields   jsn.RawMessage   `json:"stored_fields,omitempty"`   // https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-stored-fields.html
	From           int              `json:"from,omitempty"`            // "from" : 0
	Size           int              `json:"size,omitempty"`            // "size" : 10,
	Query          Query            `json:"query,omitempty"`           //
	Sort           []jsn.RawMessage `json:"sort,omitempty"`            //
	ScriptFields   jsn.RawMessage   `json:"script_fields,omitempty"`   // https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-script-fields.html#search-request-script-fields
	DocValueFields []string         `json:"docvalue_fields,omitempty"` // https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-docvalue-fields.html#search-request-docvalue-fields
}

//
// =============================================================================
//

//
// To disable _source retrieval set to false.
//
func (__obj *ElkQuery) DisableSource() *ElkQuery {
	__obj.Source = []byte("false")
	return __obj
}

//
// Set a wildcard patterns to control what parts of the _source should be returned.
//
func (__obj *ElkQuery) SourcePattern(__pattern string) *ElkQuery {
	__obj.Source = []byte(__pattern)
	return __obj
}

//
// Set many wildcard patterns to control what parts of the _source should be returned.
//
func (__obj *ElkQuery) SourcePatterns(__patterns []string) *ElkQuery {

	__obj.Source = []byte(scgu.JoinQuoted(__patterns, ","))
	return __obj

}

//
// For complete control specify both includes and excludes patterns.
// TODO
func (__obj *ElkQuery) SourceControl(__source Source) (*ElkQuery, error) {

	if _rsc, _err := jsn.Marshal(__source); _err != nil {
		__obj.Source = _rsc
		return __obj, _err
	} else {
		return __obj, nil
	}

}

//
// =============================================================================
//

//
// Selectively load specific stored fields for each document represented by a search hit.
//
func (__obj *ElkQuery) SetStoredFields(__fields []string) *ElkQuery {
	/*__obj.StoredFields = []byte("false")*/ // TODO under development
	return __obj
}

//
// Cause only the _id and _type for each hit to be returned.
//
func (__obj *ElkQuery) EmptyStoredFields() *ElkQuery {
	/*__obj.StoredFields = []byte("false")*/ // TODO under development
	return __obj
}

//
// Cause only the _id and _type for each hit to be returned.
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-stored-fields.html#_disable_stored_fields_entirely
// TODO
//func (__obj *ElkQuery) DisabledStoredFields() *ElkQuery {
//	__obj.StoredFields = []byte("_none_")
//	return __obj
//}

//
// =============================================================================
//
