package dsl

import "encoding/json"

// AggTerms returns an AggTerms body.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/8.6/search-AggTermsregations.html
type AggTerms func(name string, field string, o ...func(AggTerms *tAggTerms)) *tAggTerms

func (*AggTerms) WithSize(size int) func(all *tAggTerms) {
	return func(all *tAggTerms) {
		all.Size = size
	}
}

func (*AggTerms) WithShardSize(shardSize int) func(all *tAggTerms) {
	return func(all *tAggTerms) {
		all.Size = shardSize
	}
}

func (*AggTerms) WithSumOtherDocCount(SumOtherDocCount int) func(all *tAggTerms) {
	return func(all *tAggTerms) {
		all.SumOtherDocCount = SumOtherDocCount
	}
}

func (*AggTerms) WithShowTermDocCountError(ShowTermDocCountError bool) func(all *tAggTerms) {
	return func(all *tAggTerms) {
		all.ShowTermDocCountError = ShowTermDocCountError
	}
}

func (*AggTerms) WithOrder(Order *AggTermsOrder) func(all *tAggTerms) {
	return func(all *tAggTerms) {
		all.Order = Order
	}
}

func NewAggTerms() AggTerms {
	return func(name string, field string, o ...func(AggTerms *tAggTerms)) *tAggTerms {
		var r = &tAggTerms{
			Name:  name,
			Field: field,
		}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tAggTerms struct {
	Name                  string         `json:"-"`
	Field                 string         `json:"field"`
	Size                  int            `json:"size,omitempty"`
	ShardSize             int            `json:"shard_size,omitempty"`
	SumOtherDocCount      int            `json:"sum_other_doc_count,omitempty"`
	ShowTermDocCountError bool           `json:"show_term_doc_count_error,omitempty"`
	Order                 *AggTermsOrder `json:"order,omitempty"`
}

func (AggTerms *tAggTerms) GetAggName() string { return AggTerms.Name }

func (a *tAggTerms) MarshalJSON() ([]byte, error) {
	type _tAggTerms tAggTerms
	var data = map[string]*_tAggTerms{
		"terms": (*_tAggTerms)(a),
	}
	return json.Marshal(data)
}
