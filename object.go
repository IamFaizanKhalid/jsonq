package jsonq

import "encoding/json"

type object map[string]interface{}

func ParseObject(b []byte) (Object, error) {
	var obj object
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (o object) Has(key string) bool {
	_, ok := o[key]
	return ok
}

func (o object) Str(key string) *string {
	val, ok := o[key]
	if !ok {
		return nil
	}

	str, ok := val.(string)
	if !ok {
		return nil
	}

	return &str
}

func (o object) Int(key string) *int64 {
	val, ok := o[key]
	if !ok {
		return nil
	}

	n, ok := val.(int64)
	if !ok {
		return nil
	}

	return &n
}

func (o object) Float(key string) *float64 {
	val, ok := o[key]
	if !ok {
		return nil
	}

	f, ok := val.(float64)
	if !ok {
		return nil
	}

	return &f
}

func (o object) Bool(key string) *bool {
	val, ok := o[key]
	if !ok {
		return nil
	}

	b, ok := val.(bool)
	if !ok {
		return nil
	}

	return &b
}

func (o object) Obj(key string) Object {
	val, ok := o[key]
	if !ok {
		return nil
	}

	obj, ok := val.(object)
	if !ok {
		return nil
	}

	return obj
}

func (o object) StrArr(key string) []string {
	val, ok := o[key]
	if !ok {
		return nil
	}

	arr, ok := val.([]string)
	if !ok {
		return nil
	}

	return arr
}

func (o object) IntArr(key string) []int64 {
	val, ok := o[key]
	if !ok {
		return nil
	}

	arr, ok := val.([]int64)
	if !ok {
		return nil
	}

	return arr
}

func (o object) FloatArr(key string) []float64 {
	val, ok := o[key]
	if !ok {
		return nil
	}

	arr, ok := val.([]float64)
	if !ok {
		return nil
	}

	return arr
}

func (o object) BoolArr(key string) []bool {
	val, ok := o[key]
	if !ok {
		return nil
	}

	arr, ok := val.([]bool)
	if !ok {
		return nil
	}

	return arr
}

func (o object) ObjArr(key string) []Object {
	val, ok := o[key]
	if !ok {
		return nil
	}

	arr, ok := val.([]Object)
	if !ok {
		return nil
	}

	return arr
}
