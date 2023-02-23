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

func (*AggTerms) WithDocCountErrorUpperBound(DocCountErrorUpperBound int) func(all *tAggTerms) {
	return func(all *tAggTerms) {
		all.DocCountErrorUpperBound = DocCountErrorUpperBound
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

func (*AggTerms) WithShardMinDocCount(ShardMinDocCount int) func(all *tAggTerms) {
	return func(all *tAggTerms) {
		all.ShardMinDocCount = ShardMinDocCount
	}
}

func (*AggTerms) WithInclude(Include []string) func(all *tAggTerms) {
	return func(all *tAggTerms) {
		all.Include = Include
	}
}

func (*AggTerms) WithExclude(Exclude []string) func(all *tAggTerms) {
	return func(all *tAggTerms) {
		all.Exclude = Exclude
	}
}

func (*AggTerms) WithExecutionHint(executionHint ExecutionHint) func(all *tAggTerms) {
	return func(all *tAggTerms) {
		all.ExecutionHint = executionHint
	}
}

func (*AggTerms) WithMissing(missing string) func(all *tAggTerms) {
	return func(all *tAggTerms) {
		all.Missing = missing
	}
}

func (*AggTerms) WithValueType(valueType string) func(all *tAggTerms) {
	return func(all *tAggTerms) {
		all.ValueType = valueType
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
	Name                    string         `json:"-"`
	Field                   string         `json:"field"`
	Size                    int            `json:"size,omitempty"`
	ShardSize               int            `json:"shard_size,omitempty"`
	SumOtherDocCount        int            `json:"sum_other_doc_count,omitempty"`
	ShowTermDocCountError   bool           `json:"show_term_doc_count_error,omitempty"`
	DocCountErrorUpperBound int            `json:"doc_count_error_upper_bound,omitempty"`
	Order                   *AggTermsOrder `json:"order,omitempty"`
	MinDocCount             int            `json:"min_doc_count,omitempty"`       // default: 1
	ShardMinDocCount        int            `json:"shard_min_doc_count,omitempty"` // default: 0
	Include                 StringOrList   `json:"include,omitempty"`
	Exclude                 StringOrList   `json:"exclude,omitempty"`
	CollectMode             CollectMode    `json:"collect_mode,omitempty"`
	ExecutionHint           ExecutionHint  `json:"execution_hint,omitempty"`
	Missing                 string         `json:"missing,omitempty"`
	ValueType               string         `json:"value_type,omitempty"`
}

func (AggTerms *tAggTerms) GetAggName() string { return AggTerms.Name }

func (a *tAggTerms) MarshalJSON() ([]byte, error) {
	type _tAggTerms tAggTerms
	var data = map[string]*_tAggTerms{
		"terms": (*_tAggTerms)(a),
	}
	return json.Marshal(data)
}
