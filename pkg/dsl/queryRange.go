package dsl

import (
	"encoding/json"
	"math"
	"time"
)

type MultiType struct {
	FloatVal  float64
	StringVal string
	DateVal   time.Time
}

func (mt *MultiType) IsDate() bool {
	return !mt.DateVal.IsZero() && mt.StringVal == ""
}

func (mt *MultiType) TimeFormat() string {
	return "yyyy-MM-ddTHH:mm:ss"
}

func (mt *MultiType) MarshalJSON() ([]byte, error) {
	if mt.StringVal != "" {
		return json.Marshal(mt.StringVal)
	}
	if !mt.DateVal.IsZero() {
		return json.Marshal(mt.DateVal.Format("2006-01-02T15:04:05"))
	}
	if math.Trunc(mt.FloatVal) == mt.FloatVal {
		return json.Marshal(int64(mt.FloatVal))
	} else {
		return json.Marshal(mt.FloatVal)
	}
}

type RangeQuery struct {
	Field    string        `json:"-"`
	GT       *MultiType    `json:"gt,omitempty"`
	GTE      *MultiType    `json:"gte,omitempty"`
	LT       *MultiType    `json:"lt,omitempty"`
	LTE      *MultiType    `json:"lte,omitempty"`
	Format   string        `json:"format,omitempty"`
	Relation QueryRelation `json:"relation,omitempty"`
	TimeZone string        `json:"time_zone,omitempty"`
	Boost    float64       `json:"boost,omitempty"` // default: 1.0
}

func (*RangeQuery) GetQueryName() string { return "range" }

func (r *RangeQuery) MarshalJSON() ([]byte, error) {
	type _RangeQuery RangeQuery
	if r.GT != nil && r.GT.IsDate() {
		r.Format = r.GT.TimeFormat()
	}
	if r.GTE != nil && r.GTE.IsDate() {
		r.Format = r.GTE.TimeFormat()
	}
	if r.LT != nil && r.LT.IsDate() {
		r.Format = r.LT.TimeFormat()
	}
	if r.LTE != nil && r.LTE.IsDate() {
		r.Format = r.LTE.TimeFormat()
	}

	data := map[string]*_RangeQuery{
		r.Field: (*_RangeQuery)(r),
	}
	return json.Marshal(data)
}

var (
	_ Query = (*RangeQuery)(nil)
)
