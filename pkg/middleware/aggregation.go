package middleware

import "github.com/je4/elasticdsl/v2/pkg/dsl"

type Aggregation interface {
	GetName() string
	GetAgg(api *dsl.API) dsl.BaseAgg
	UnmarshalJSON([]byte) error
}
