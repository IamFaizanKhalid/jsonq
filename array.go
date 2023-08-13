package jsonq

type Array []any

func (a Array) Str() []string {
	output := make([]string, len(a))
	for i, val := range a {
		output[i] = val.(string)
	}
	return output
}

func (a Array) Int() []int {
	output := make([]int, len(a))
	for i, val := range a {
		output[i] = int(val.(float64))
	}
	return output
}

func (a Array) Float() []float64 {
	output := make([]float64, len(a))
	for i, val := range a {
		output[i] = val.(float64)
	}
	return output
}

func (a Array) Bool() []bool {
	output := make([]bool, len(a))
	for i, val := range a {
		output[i] = val.(bool)
	}
	return output
}

func (a Array) Arr() []Array {
	output := make([]Array, len(a))
	for i, val := range a {
		output[i] = val.([]interface{})
	}
	return output
}

func (a Array) Obj() []Object {
	output := make([]Object, len(a))
	for i, val := range a {
		output[i] = val.(map[string]interface{})
	}
	return output
}
