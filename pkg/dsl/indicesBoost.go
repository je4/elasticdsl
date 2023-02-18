package dsl

import "encoding/json"

// IndicesBoost results a indices_boost struct
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multiple-indices.html
type IndicesBoost func(o ...func(all tIndicesBoost)) tIndicesBoost

func (ib *IndicesBoost) AppendIndex(index string, boost float64) func(all tIndicesBoost) {
	return func(all tIndicesBoost) {
		all[index] = boost
	}
}

func NewIndicesBoost() IndicesBoost {
	return func(o ...func(all tIndicesBoost)) tIndicesBoost {
		var r = tIndicesBoost{}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tIndicesBoost map[string]float64

func (ib tIndicesBoost) MarshalJSON() ([]byte, error) {
	var data = []map[string]float64{}
	for index, boost := range ib {
		data = append(data, map[string]float64{index: boost})
	}
	return json.Marshal(data)
}
