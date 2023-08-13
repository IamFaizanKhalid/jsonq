package jsonq

import "github.com/goccy/go-json"

func ParseObject(b []byte) (Object, error) {
	var obj Object
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func ParseArray(b []byte) (Array, error) {
	var arr Array
	err := json.Unmarshal(b, &arr)
	if err != nil {
		return nil, err
	}

	return arr, nil
}
