package dsl

import "encoding/json"

// Fields returns an fields body.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/8.6/search-fieldregations.html
type Fields func(o ...func(fields *tFields)) *tFields

func (a *Fields) WithFields(fields ...*tFieldsField) func(all *tFields) {
	return func(all *tFields) {
		all.Fields = fields
	}
}

func NewFields() Fields {
	return func(o ...func(fields *tFields)) *tFields {
		var r = &tFields{}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tFields struct {
	Fields []*tFieldsField `json:"-"`
}

func (*tFields) GetFieldsName() string { return "fields" }

func (a *tFields) MarshalJSON() ([]byte, error) {
	var data = []any{}
	for _, field := range a.Fields {
		data = append(data, field)
	}

	return json.Marshal(data)
}
