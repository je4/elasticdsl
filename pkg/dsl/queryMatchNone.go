package dsl

// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-all-query.html

type MatchNoneQuery struct {
}

func (*MatchNoneQuery) GetQueryName() string { return "match_none" }

var (
	_ Query = (*MatchNoneQuery)(nil)
)
