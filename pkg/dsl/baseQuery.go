package dsl

/*
BaseQuery
https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl.html
*/
type BaseQuery interface {
	GetQueryName() string
}
