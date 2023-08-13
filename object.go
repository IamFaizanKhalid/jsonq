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

func (o object) OptStr(key string) (string, bool) {
	val, ok := o[key].(string)
	return val, ok
}

func (o object) OptInt(key string) (int, bool) {
	val, ok := o[key].(float64)
	return int(val), ok
}

func (o object) OptFloat(key string) (float64, bool) {
	val, ok := o[key].(float64)
	return val, ok
}

func (o object) OptBool(key string) (bool, bool) {
	val, ok := o[key].(bool)
	return val, ok
}

func (o object) OptObj(key string) (Object, bool) {
	val, ok := o[key].(map[string]interface{})
	return object(val), ok
}
