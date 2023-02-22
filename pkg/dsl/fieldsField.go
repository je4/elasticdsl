package dsl

import "encoding/json"

// FieldsField returns an FieldsField body.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/8.6/search-FieldsFieldregations.html
type FieldsField func(name string, o ...func(FieldsField *tFieldsField)) *tFieldsField

func (*FieldsField) WithFormat(format string) func(all *tFieldsField) {
	return func(all *tFieldsField) {
		all.Format = format
	}
}
func NewFieldsField() FieldsField {
	return func(name string, o ...func(FieldsField *tFieldsField)) *tFieldsField {
		var r = &tFieldsField{
			Name: name,
		}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tFieldsField struct {
	Name   string `json:"-"`
	Format string `json:"-"`
}

func (FieldsField *tFieldsField) GetFieldName() string { return FieldsField.Name }

func (a *tFieldsField) MarshalJSON() ([]byte, error) {
	if a.Format == "" {
		return json.Marshal(a.Name)
	}
	var val = &struct {
		Field  string `json:"field"`
		Format string `json:"format"`
	}{
		Field:  a.Name,
		Format: a.Format,
	}
	return json.Marshal(val)
}
