package dsl

// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
type Search func(o ...func(all *tSearch)) *tSearch

func (s *Search) WithQuery(query *tQuery) func(all *tSearch) {
	return func(all *tSearch) {
		all.Query = query
	}
}

func (s *Search) WithAggs(aggs *tAggs) func(all *tSearch) {
	return func(all *tSearch) {
		all.Aggs = aggs
	}
}

func (s *Search) WithIndicesBoost(ib tIndicesBoost) func(all *tSearch) {
	return func(all *tSearch) {
		all.IndicesBoost = ib
	}
}

func NewSearch() Search {
	return func(o ...func(search *tSearch)) *tSearch {
		var r = &tSearch{}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tSearch struct {
	Query        *tQuery       `json:"query,omitempty"`
	Aggs         *tAggs        `json:"aggs,omitempty"`
	IndicesBoost tIndicesBoost `json:"indices_boost,omitempty"`
}

type SearchQueryParams struct {
	AllowNoIndices             bool                 `json:"allow_no_indices,omitempty"`             // default: true
	AllowPartialSearchResults  bool                 `json:"allow_partial_search_results,omitempty"` // default: true
	Analyzer                   string               `json:"analyzer,omitempty"`
	AnalyzeWildcard            bool                 `json:"analyze_wildcard,omitempty"`        // default: false
	BatchedReduceSize          int                  `json:"batched_reduce_size,omitempty"`     // default: 512
	CSSMinimizeRoundtrips      bool                 `json:"css_minimize_roundtrips,omitempty"` // default: true
	DefaultOperator            QueryOperator        `json:"default_operator,omitempty"`        // default: QueryOperatorOR
	DF                         string               `json:"df,omitempty"`
	DocvalueFields             []string             `json:"docvalue_fields,omitempty"`
	ExpandWildcards            QueryExpandWildcards `json:"expand_wildcards,omitempty"`              // default: QueryExpandWildcardsOpen
	Explain                    bool                 `json:"explain,omitempty"`                       // default: false
	From                       int                  `json:"from,omitempty"`                          // default: 0
	IgnoreThrottled            bool                 `json:"ignore_throttled,omitempty"`              // default: false
	Lenient                    bool                 `json:"lenient,omitempty"`                       // default: false
	MaxConcurrentShardRequests int                  `json:"max_concurrent_shard_requests,omitempty"` // default: 5
	PreFilterShardSize         int                  `json:"pre_filter_shard_size,omitempty"`
	Preference                 string               `json:"preference,omitempty"`
	Q                          string               `json:"q,omitempty"`
	RequestCache               bool                 `json:"request_cache,omitempty"`
	RestTotalHitsAsInt         bool                 `json:"rest_total_hits_as_int,omitempty"` // default: false
	Routing                    string               `json:"routing,omitempty"`
	Scroll                     string               `json:"scroll,omitempty"`
	SearchType                 QuerySearchType      `json:"search_type,omitempty"` // default: QuerySearchTypeQueryThenFetch
	SeqNoPrimaryTerm           bool                 `json:"seq_no_primary_term,omitempty"`
	Size                       int                  `json:"size,omitempty"` // default: 10
	Sort                       SearchSortVal        `json:"sort,omitempty"`
	Source                     *SearchSourceVal     `json:"_source,omitempty"` // default: true
	SourceExcludes             StringList           `json:"_source_excludes,omitempty"`
	SourceIncludes             StringList           `json:"_source_includes,omitempty"`
	Stats                      string               `json:"stats,omitempty"`
	StoredFields               StringList           `json:"stored_fields,omitempty"`
	SuggestField               string               `json:"suggest_field,omitempty"`
	SuggestMode                SearchSuggestMode    `json:"suggest_mode,omitempty"` // default: SearchSuggestModeMissing
	SuggestSize                int                  `json:"suggest_size,omitempty"`
	SuggestText                string               `json:"suggest_text,omitempty"`
	TerminateAfter             int                  `json:"terminate_after,omitempty"` // default: 0
	Timeout                    string               `json:"timeout,omitempty"`
	TrackScores                bool                 `json:"track_scores,omitempty"`     // default: false
	TrackTotalHits             *TrackTotalHitsVal   `json:"track_total_hits,omitempty"` // default: 10000
	TypedKeys                  bool                 `json:"typed_keys,omitempty"`       // default: true
	Version                    bool                 `json:"version,omitempty"`          // default: false
}
