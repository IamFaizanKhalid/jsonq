package jsonq

type object map[string]interface{}

func (o object) Has(key string) bool {
	_, ok := o[key]
	return ok
}

func (o object) Str(key string) string {
	return o[key].(string)
}

func (o object) Int(key string) int {
	return o[key].(int)
}

func (o object) Float(key string) float64 {
	return o[key].(float64)
}

func (o object) Bool(key string) bool {
	return o[key].(bool)
}

func (o object) Obj(key string) Object {
	return o[key].(object)
}

func (o object) StrArr(key string) []string {
	val, ok := o[key]
	if !ok {
		return []string{}
	}

	arr, ok := val.([]string)
	if !ok {
		return []string{}
	}

	return arr
}

func (o object) IntArr(key string) []int {
	val, ok := o[key]
	if !ok {
		return []int{}
	}

	arr, ok := val.([]int)
	if !ok {
		return []int{}
	}

	return arr
}

func (o object) FloatArr(key string) []float64 {
	val, ok := o[key]
	if !ok {
		return []float64{}
	}

	arr, ok := val.([]float64)
	if !ok {
		return []float64{}
	}

	return arr
}

func (o object) BoolArr(key string) []bool {
	val, ok := o[key]
	if !ok {
		return []bool{}
	}

	arr, ok := val.([]bool)
	if !ok {
		return []bool{}
	}

	return arr
}

func (o object) ObjArr(key string) []Object {
	val, ok := o[key]
	if !ok {
		return []Object{}
	}

	arr, ok := val.([]Object)
	if !ok {
		return []Object{}
	}

	return arr
}

func (o object) OptStr(key string) *string {
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

func (o object) OptInt(key string) *int {
	val, ok := o[key]
	if !ok {
		return nil
	}

	n, ok := val.(int)
	if !ok {
		return nil
	}

	return &n
}

func (o object) OptFloat(key string) *float64 {
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

func (o object) OptBool(key string) *bool {
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

func (o object) OptObj(key string) Object {
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
