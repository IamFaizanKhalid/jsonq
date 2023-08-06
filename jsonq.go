package jsonq

type Object interface {
	Has(key string) bool
	Str(key string) *string
	Int(key string) *int64
	Float(key string) *float64
	Bool(key string) *bool
	Obj(key string) Object
	StrArr(key string) []string
	IntArr(key string) []int64
	FloatArr(key string) []float64
	BoolArr(key string) []bool
	ObjArr(key string) []Object
}
