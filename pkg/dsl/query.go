package dsl

/*
Query
https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl.html
*/
type Query interface {
	GetQueryName() string
}
