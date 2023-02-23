package dsl

import "encoding/json"

// AggScript returns an AggScript body.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/8.6/search-AggScriptregations.html
type AggScript func(name string, bucketsPath map[string]string, o ...func(AggScript *tAggScript)) *tAggScript

func (*AggScript) WithGapPolicy(gapPolicy GapPolicy) func(all *tAggScript) {
	return func(all *tAggScript) {
		all.GapPolicy = gapPolicy
	}
}

func (*AggScript) WithFormat(format string) func(all *tAggScript) {
	return func(all *tAggScript) {
		all.Format = format
	}
}

func NewAggScript() AggScript {
	return func(name string, bucketsPath map[string]string, o ...func(AggScript *tAggScript)) *tAggScript {
		var r = &tAggScript{
			Name:        name,
			BucketsPath: bucketsPath,
		}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tAggScript struct {
	Name        string            `json:"-"`
	Script      string            `json:"script"`
	BucketsPath map[string]string `json:"buckets_path,omitempty"`
	GapPolicy   GapPolicy         `json:"gap_policy,omitempty"` // default: GapPolicySkip
	Format      string            `json:"format,omitempty"`
}

func (AggScript *tAggScript) GetAggName() string { return AggScript.Name }

func (a *tAggScript) MarshalJSON() ([]byte, error) {
	type _tAggScript tAggScript
	var data = map[string]*_tAggScript{
		"bucket_script": (*_tAggScript)(a),
	}
	return json.Marshal(data)
}
