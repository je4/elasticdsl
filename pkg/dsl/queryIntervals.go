package dsl

import "encoding/json"

// https://www.elastic.co/guide/en/elasticsearch/reference/8.6/query-dsl-intervals-query.html

type Script struct {
	Lang   string         `json:"lang,omitempty"` // default: "painless"
	Source string         `json:"source,omitempty"`
	Id     string         `json:"id,omitempty"`
	Params map[string]any `json:"params,omitempty"`
}

type InvervalsQueryFilter struct {
	After          *BaseQuery `json:"after,omitempty"`
	Before         *BaseQuery `json:"before,omitempty"`
	ContainedBy    *BaseQuery `json:"contained_by,omitempty"`
	Containing     *BaseQuery `json:"containing,omitempty"`
	NotContainedBy *BaseQuery `json:"not_contained_by,omitempty"`
	NotContaining  *BaseQuery `json:"not_containing,omitempty"`
	NotOverlapping *BaseQuery `json:"not_overlapping,omitempty"`
	Overlapping    *BaseQuery `json:"overlapping,omitempty"`
	Script         *Script    `json:"script,omitempty"`
}

type InvervalsQueryMatch struct {
	Query    string                `json:"query"`
	MaxGaps  int                   `json:"max_gaps,omitempty"` // default: -1
	Ordered  bool                  `json:"ordered,omitempty"`  // default: false
	Analyzer string                `json:"analyzer,omitempty"`
	Filter   *InvervalsQueryFilter `json:"filter,omitempty"`
	UseField string                `json:"use_field,omitempty"`
}

type InvervalsQueryPrefix struct {
	Prefix   string `json:"prefix"`
	Analyzer string `json:"analyzer,omitempty"`
	UseField string `json:"use_field,omitempty"`
}

type InvervalsQueryWildcard struct {
	Pattern  string `json:"pattern"`
	Analyzer string `json:"analyzer,omitempty"`
	UseField string `json:"use_field,omitempty"`
}

type InvervalsQueryFuzzy struct {
	Term           string `json:"term"`
	PrefixLength   int    `json:"prefix_length"`            // default: 0
	Transpositions bool   `json:"transpositions,omitempty"` // default: false
	Fuzziness      string `json:"fuzziness,omitempty"`      // default: "auto"
	Analyzer       string `json:"analyzer,omitempty"`
	UseField       string `json:"use_field,omitempty"`
}

type InvervalsQueryAllOf struct {
	Intervals []any                 `json:"intervals"`
	MaxGaps   int                   `json:"max_gaps,omitempty"` // default: -1
	Ordered   bool                  `json:"ordered,omitempty"`  // default: false
	Filter    *InvervalsQueryFilter `json:"filter,omitempty"`
}

type InvervalsQueryAnyOf struct {
	Intervals []any                 `json:"intervals"`
	Filter    *InvervalsQueryFilter `json:"filter,omitempty"`
}

type IntervalsQuery struct {
	Field    string                  `json:"-"`
	Match    *InvervalsQueryMatch    `json:"match,omitempty"`
	Prefix   *InvervalsQueryPrefix   `json:"prefix,omitempty"`
	Wildcard *InvervalsQueryWildcard `json:"wildcard,omitempty"`
	Fuzzy    *InvervalsQueryFuzzy    `json:"fuzzy,omitempty"`
	AllOf    *InvervalsQueryAllOf    `json:"all_of,omitempty"`
	AnyOf    *InvervalsQueryAnyOf    `json:"any_of,omitempty"`
}

func (i *IntervalsQuery) GetQueryName() string { return "intervals" }

func (i *IntervalsQuery) MarshalJSON() ([]byte, error) {
	type _tInvervalsQuery IntervalsQuery
	data := map[string]*_tInvervalsQuery{
		i.Field: (*_tInvervalsQuery)(i),
	}
	return json.Marshal(data)
}

var (
	_ BaseQuery = (*IntervalsQuery)(nil)
)
