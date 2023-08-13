package jsonq

import "encoding/json"

type Object interface {
	Has(key string) bool
	Str(key string) string
	Int(key string) int
	Float(key string) float64
	Bool(key string) bool
	Obj(key string) Object
	StrArr(key string) []string
	IntArr(key string) []int
	FloatArr(key string) []float64
	BoolArr(key string) []bool
	ObjArr(key string) []Object
	OptStr(key string) (string, bool)
	OptInt(key string) (int, bool)
	OptFloat(key string) (float64, bool)
	OptBool(key string) (bool, bool)
	OptObj(key string) (Object, bool)
}

func ParseObject(b []byte) (Object, error) {
	var obj object
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func ParseArray(b []byte) ([]Object, error) {
	var arr []object
	err := json.Unmarshal(b, &arr)
	if err != nil {
		return nil, err
	}

	objArr := make([]Object, len(arr))
	for i := range arr {
		objArr[i] = arr[i]
	}

	return objArr, nil
}
