package elastic

// JSONDummy reflects json value as []byte
type JSONDummy []byte

func (a *JSONDummy) UnmarshalJSON(data []byte) error {
	*a = data
	return nil
}
