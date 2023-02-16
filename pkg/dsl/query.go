package dsl

import "encoding/json"

/*
Query
https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl.html
*/
type tQuery interface {
	GetQueryName() string
}

type tDSL struct {
	Query tQuery `json:"-"`
}

func (d *tDSL) MarshalJSON() ([]byte, error) {
	type _tInvervalsQuery tIntervalsQuery
	data := map[string]tQuery{
		d.Query.GetQueryName(): d.Query,
	}
	return json.Marshal(data)
}
