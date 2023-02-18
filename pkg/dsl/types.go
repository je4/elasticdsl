package dsl

import (
	"encoding/json"
	"strings"
)

type StringList []string

func (sl StringList) MarshalJSON() ([]byte, error) {
	str := strings.Join(sl, ",")
	return json.Marshal(str)
}

type SearchSourceVal struct {
	BoolValue bool
	Fields    StringList
}

func (ssv *SearchSourceVal) MarshalJSON() ([]byte, error) {
	if len(ssv.Fields) > 0 {
		return json.Marshal(ssv.Fields)
	}
	return json.Marshal(ssv.BoolValue)
}

type SearchSortVal map[string]string

func (ssv SearchSortVal) MarshalJSON() ([]byte, error) {
	var strs = []string{}
	for field, direction := range ssv {
		strs = append(strs, field+":"+direction)
	}
	str := strings.Join(strs, ",")
	return json.Marshal(str)
}

type TrackTotalHitsVal struct {
	IntVal  int
	BoolVal bool
}

func (tthv *TrackTotalHitsVal) MarshalJSON() ([]byte, error) {
	if tthv.IntVal > 0 {
		return json.Marshal(tthv.IntVal)
	}
	return json.Marshal(tthv.BoolVal)
}

type MinimumShouldMatch struct {
	StringValue string
	IntValue    int
}

func (msm *MinimumShouldMatch) MarshalJSON() ([]byte, error) {
	if msm.StringValue != "" {
		return json.Marshal(msm.StringValue)
	}
	return json.Marshal(msm.IntValue)
}

type SimpleQueryStringQueryFlags []QuerySimpleQueryStringFlag

func (flags SimpleQueryStringQueryFlags) MarshalJSON() ([]byte, error) {
	var strFlags = []string{}
	for _, f := range flags {
		strFlags = append(strFlags, string(f))
	}
	data := strings.Join(strFlags, "|")
	return json.Marshal(data)
}
