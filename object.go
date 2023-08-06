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
	return int(o[key].(float64))
}

func (o object) Float(key string) float64 {
	return o[key].(float64)
}

func (o object) Bool(key string) bool {
	return o[key].(bool)
}

func (o object) Obj(key string) Object {
	return object(o[key].(map[string]interface{}))
}

func (o object) getArr(key string) []interface{} {
	val, ok := o[key]
	if !ok {
		return []interface{}{}
	}

	arr, ok := val.([]interface{})
	if !ok {
		return []interface{}{}
	}

	return arr
}

func (o object) StrArr(key string) []string {
	arr := o.getArr(key)
	if arr == nil {
		return []string{}
	}

	respArr := make([]string, len(arr))
	for i := range arr {
		respArr[i] = arr[i].(string)
	}

	return respArr
}

func (o object) IntArr(key string) []int {
	arr := o.getArr(key)
	if arr == nil {
		return []int{}
	}

	respArr := make([]int, len(arr))
	for i := range arr {
		respArr[i] = int(arr[i].(float64))
	}

	return respArr
}

func (o object) FloatArr(key string) []float64 {
	arr := o.getArr(key)
	if arr == nil {
		return []float64{}
	}

	respArr := make([]float64, len(arr))
	for i := range arr {
		respArr[i] = arr[i].(float64)
	}

	return respArr
}

func (o object) BoolArr(key string) []bool {
	arr := o.getArr(key)
	if arr == nil {
		return []bool{}
	}

	respArr := make([]bool, len(arr))
	for i := range arr {
		respArr[i] = arr[i].(bool)
	}

	return respArr
}

func (o object) ObjArr(key string) []Object {
	arr := o.getArr(key)
	if arr == nil {
		return []Object{}
	}

	respArr := make([]Object, len(arr))
	for i := range arr {
		respArr[i] = object(arr[i].(map[string]interface{}))
	}

	return respArr
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
