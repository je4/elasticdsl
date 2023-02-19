package middleware

import "github.com/je4/elasticdsl/v2/pkg/dsl"

type Facet interface {
	GetName() string
	GetAgg(api *dsl.API) dsl.Aggregation
}
