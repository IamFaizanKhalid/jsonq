package jsonq

type Object map[string]interface{}

func (o Object) Has(key string) bool {
	_, ok := o[key]
	return ok
}

func (o Object) Val(key string) Value {
	return Value{o[key]}
}

func (o Object) OptVal(key string) (Value, bool) {
	val, ok := o[key]
	return Value{val}, ok
}

func (o Object) Obj(key string) Object {
	return o[key].(map[string]interface{})
}

func (o Object) OptObj(key string) (Object, bool) {
	val, ok := o[key].(map[string]interface{})
	return val, ok
}

func (o Object) Arr(key string) Array {
	return o[key].([]any)
}

func (o Object) OptArr(key string) (Array, bool) {
	val, ok := o[key].([]any)
	return val, ok
}
