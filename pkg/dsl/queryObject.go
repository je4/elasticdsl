package dsl

import "encoding/json"

type QueryObject func(o ...func(*queryRequest)) ([]byte, error)

func (qo QueryObject) WithQuery(q Query) func(request *queryRequest) {
	return func(qo *queryRequest) {
		qo.Query = q
	}
}

func NewQueryObject() QueryObject {
	return func(o ...func(*queryRequest)) ([]byte, error) {
		var r = &queryRequest{}
		for _, f := range o {
			f(r)
		}
		return json.Marshal(r)
	}
}

type queryRequest struct {
	Query Query `json:"-"`
}

func (d *queryRequest) MarshalJSON() ([]byte, error) {
	data := map[string]Query{
		d.Query.GetQueryName(): d.Query,
	}
	query := map[string]any{
		"query": data,
	}
	return json.Marshal(query)
}
