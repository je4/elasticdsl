package elastic

type Shard struct {
	Total      int `json:"total,omitempty"`
	Successful int `json:"successful,omitempty"`
	Skipped    int `json:"skipped,omitempty"`
	Failed     int `json:"failed,omitempty"`
}

type Total struct {
	Value    int            `json:"value,omitempty"`
	Relation ResultRelation `json:"relation,omitempty"`
}

type Hit struct {
	Index  string         `json:"_index"`
	ID     string         `json:"_id"`
	Score  float32        `json:"_score,omitempty"`
	Source any            `json:"_source,omitempty"`
	Fields map[string]any `json:"fields,omitempty"`
}

type Hits struct {
	Total    *Total  `json:"total,omitempty"`
	MaxScore float64 `json:"max_score,omitempty"`
	Hits     []*Hit  `json:"hits"`
}

type SearchResult struct {
	ScrollID string `json:"_scroll_id,omitempty"`
	Took     int    `json:"took,omitempty"`
	TimedOut bool   `json:"timed_out,omitempty"`
	Shards   *Shard `json:"_shards,omitempty"`
	Hits     *Hits  `json:"hits,omitempty"`
}
