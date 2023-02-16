package elastic

type Client interface {
	Info() (*ResultInfo, error)
	Search(cfg *SearchConfig) ([]map[string][]string, map[string]any, int64, ResultFacet, error)
}
