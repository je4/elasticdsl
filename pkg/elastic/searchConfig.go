package elastic

type TermFacet struct {
	Selected map[string]bool
	Prefix   string
	Limit    int64
}

type SearchConfig struct {
	Index          string
	Fields         map[string][]string
	QStr           string
	FiltersFields  map[string][]string
	Facets         map[string]TermFacet
	Groups         []string
	ContentVisible bool
	Start          int
	Rows           int
	IsAdmin        bool
}
