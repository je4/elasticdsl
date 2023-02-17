package dsl

// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-all-query.html

type MatchAllQuery func(o ...func(all *queryMatchAll)) Query

func (qma *MatchAllQuery) WithBoost(boost float64) func(all *queryMatchAll) {
	return func(all *queryMatchAll) {
		all.Boost = boost
	}
}

func NewMatchAllQuery() MatchAllQuery {
	return func(o ...func(*queryMatchAll)) Query {
		var r = &queryMatchAll{}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type queryMatchAll struct {
	Boost float64 `json:"boost,omitempty"`
}

func (*queryMatchAll) GetQueryName() string { return "match_all" }

var (
	_ Query = (*queryMatchAll)(nil)
)
