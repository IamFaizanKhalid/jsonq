package jsonq

type Value struct {
	val any
}

func (v Value) Str() string {
	return v.val.(string)
}

func (v Value) Int() int {
	return int(v.val.(float64))
}

func (v Value) Float() float64 {
	return v.val.(float64)
}

func (v Value) Bool() bool {
	return v.val.(bool)
}

func (v Value) OptStr() (string, bool) {
	val, ok := v.val.(string)
	return val, ok
}

func (v Value) OptInt() (int, bool) {
	val, ok := v.val.(float64)
	return int(val), ok
}

func (v Value) OptFloat() (float64, bool) {
	val, ok := v.val.(float64)
	return val, ok
}

func (v Value) OptBool() (bool, bool) {
	val, ok := v.val.(bool)
	return val, ok
}
