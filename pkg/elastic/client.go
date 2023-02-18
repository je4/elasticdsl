package elastic

import "github.com/je4/elasticdsl/v2/pkg/dsl"

type Client interface {
	Info() (*ResultInfo, error)
	GetDSLAPI() *dsl.API
	Search(index string, srch any) (*SearchResult, error)
}
