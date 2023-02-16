package dsl

import "encoding/json"

// https://www.elastic.co/guide/en/elasticsearch/reference/8.6/query-dsl-intervals-query.html

type tScript struct {
	Lang   string         `json:"lang,omitempty"` // default: "painless"
	Source string         `json:"source,omitempty"`
	Id     string         `json:"id,omitempty"`
	Params map[string]any `json:"params,omitempty"`
}

type tInvervalsQueryFilter struct {
	After          *tQuery  `json:"after,omitempty"`
	Before         *tQuery  `json:"before,omitempty"`
	ContainedBy    *tQuery  `json:"contained_by,omitempty"`
	Containing     *tQuery  `json:"containing,omitempty"`
	NotContainedBy *tQuery  `json:"not_contained_by,omitempty"`
	NotContaining  *tQuery  `json:"not_containing,omitempty"`
	NotOverlapping *tQuery  `json:"not_overlapping,omitempty"`
	Overlapping    *tQuery  `json:"overlapping,omitempty"`
	Script         *tScript `json:"script,omitempty"`
}

type tInvervalsQueryMatch struct {
	Query    string                 `json:"query"`
	MaxGaps  int                    `json:"max_gaps,omitempty"` // default: -1
	Ordered  bool                   `json:"ordered,omitempty"`  // default: false
	Analyzer string                 `json:"analyzer,omitempty"`
	Filter   *tInvervalsQueryFilter `json:"filter,omitempty"`
	UseField string                 `json:"use_field,omitempty"`
}

type tInvervalsQueryPrefix struct {
	Prefix   string `json:"prefix"`
	Analyzer string `json:"analyzer,omitempty"`
	UseField string `json:"use_field,omitempty"`
}

type tInvervalsQueryWildcard struct {
	Pattern  string `json:"pattern"`
	Analyzer string `json:"analyzer,omitempty"`
	UseField string `json:"use_field,omitempty"`
}

type tInvervalsQueryFuzzy struct {
	Term           string `json:"term"`
	PrefixLength   int    `json:"prefix_length"`            // default: 0
	Transpositions bool   `json:"transpositions,omitempty"` // default: false
	Fuzziness      string `json:"fuzziness,omitempty"`      // default: "auto"
	Analyzer       string `json:"analyzer,omitempty"`
	UseField       string `json:"use_field,omitempty"`
}

type tInvervalsQueryAllOf struct {
	Intervals []any                  `json:"intervals"`
	MaxGaps   int                    `json:"max_gaps,omitempty"` // default: -1
	Ordered   bool                   `json:"ordered,omitempty"`  // default: false
	Filter    *tInvervalsQueryFilter `json:"filter,omitempty"`
}

type tInvervalsQueryAnyOf struct {
	Intervals []any                  `json:"intervals"`
	Filter    *tInvervalsQueryFilter `json:"filter,omitempty"`
}

type tIntervalsQuery struct {
	Field    string                   `json:"-"`
	Match    *tInvervalsQueryMatch    `json:"match,omitempty"`
	Prefix   *tInvervalsQueryPrefix   `json:"prefix,omitempty"`
	Wildcard *tInvervalsQueryWildcard `json:"wildcard,omitempty"`
	Fuzzy    *tInvervalsQueryFuzzy    `json:"fuzzy,omitempty"`
	AllOf    *tInvervalsQueryAllOf    `json:"all_of,omitempty"`
	AnyOf    *tInvervalsQueryAnyOf    `json:"any_of,omitempty"`
}

func (i *tIntervalsQuery) GetQueryName() string { return "intervals" }

func (i *tIntervalsQuery) MarshalJSON() ([]byte, error) {
	type _tInvervalsQuery tIntervalsQuery
	data := map[string]*_tInvervalsQuery{
		i.Field: (*_tInvervalsQuery)(i),
	}
	return json.Marshal(data)
}
