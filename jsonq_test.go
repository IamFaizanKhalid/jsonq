package jsonq

import (
	"testing"
)

const TestData = `{
	"foo": 1,
	"bar": 2,
	"test": "Hello, world!",
	"baz": 123.1,
	"numstring": "42",
	"floatstring": "42.1",
	"array": [
		{"foo": 1},
		{"bar": 2},
		{"baz": 3}
	],
	"subobj": {
		"foo": 1,
		"subarray": [1,2,3],
		"subsubobj": {
			"bar": 2,
			"baz": 3,
			"array": ["hello", "world"]
		}
	},
	"collections": {
		"bools": [false, true, false],
		"strings": ["hello", "strings"],
		"numbers": [1,2,3,4],
		"arrays": [[1.0,2.0],[2.0,3.0],[4.0,3.0]],
		"objects": [
			{"obj1": 1},
			{"obj2": 2}
		]
	},
	"bool": true
}`

func tErr(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
}

func TestQuery(t *testing.T) {
	q, err := ParseObject([]byte(TestData))
	tErr(t, err)

	ival := q.Int("foo")
	if ival != 1 {
		t.Errorf("Expecting 1, got %v\n", ival)
	}

	ival = q.Int("bar")
	if ival != 2 {
		t.Errorf("Expecting 2, got %v\n", ival)
	}

	ival = q.Obj("subobj").Int("foo")
	if ival != 1 {
		t.Errorf("Expecting 1, got %v\n", ival)
	}

	//// test that strings can get int-ed
	//ival = q.Int("numstring")
	//if ival != 42 {
	//	t.Errorf("Expecting 42, got %v\n", ival)
	//}

	x := q.Obj("subobj").(object)["subarray"]
	t.Log(x.([]interface{}))
	arr := q.Obj("subobj").IntArr("subarray")
	for i := 0; i < 3; i++ {
		ival := arr[i]
		if ival != i+1 {
			t.Errorf("Expecting %d, got %v\n", i+1, ival)
		}
	}

	fval := q.Float("baz")
	if fval != 123.1 {
		t.Errorf("Expecting 123.1, got %f\n", fval)
	}

	//// test that strings can get float-ed
	//fval = q.Float("floatstring")
	//if fval != 42.1 {
	//	t.Errorf("Expecting 42.1, got %v\n", fval)
	//}

	sval := q.Str("test")
	if sval != "Hello, world!" {
		t.Errorf("Expecting \"Hello, World!\", got \"%v\"\n", sval)
	}

	sval = q.Obj("subobj").Obj("subsubobj").StrArr("array")[0]
	if sval != "hello" {
		t.Errorf("Expecting \"hello\", got \"%s\"\n", sval)
	}

	bval := q.Bool("bool")
	if !bval {
		t.Errorf("Expecting true, got %v\n", bval)
	}

	obj := q.Obj("subobj").Obj("subsubobj")

	sval = obj.StrArr("array")[1]
	if sval != "world" {
		t.Errorf("Expecting \"world\", got \"%s\"\n", sval)
	}

	aobj := q.Obj("subobj").FloatArr("subarray")

	if aobj[0] != 1 {
		t.Errorf("Expecting 1, got %v\n", aobj[0])
	}

	iobj := q.OptStr("numstring")
	if iobj == nil {
		t.Errorf("Expecting type string got: %v", iobj)
	}

	/*
		Test Extraction of typed slices
	*/

	//test array of strings
	astrings := q.Obj("collections").StrArr("strings")

	if astrings[0] != "hello" {
		t.Errorf("Expecting hello, got %v\n", astrings[0])
	}

	//test array of ints
	aints := q.Obj("collections").IntArr("numbers")

	if aints[0] != 1 {
		t.Errorf("Expecting 1, got %v\n", aints[0])
	}

	//test array of floats
	afloats := q.Obj("collections").FloatArr("numbers")

	if afloats[0] != 1.0 {
		t.Errorf("Expecting 1.0, got %v\n", afloats[0])
	}

	//test array of bools
	abools := q.Obj("collections").BoolArr("bools")

	if abools[0] {
		t.Errorf("Expecting true, got %v\n", abools[0])
	}

	////test array of arrays
	//aa := q.Obj("collections").ArrArr("arrays")
	//
	//if aa[0][0].(float64) != 1 {
	//	t.Errorf("Expecting 1, got %v\n", aa[0][0])
	//}

	//test array of objs
	aobjs := q.Obj("collections").ObjArr("objects")

	if aobjs[0].Float("obj1") != 1 {
		t.Errorf("Expecting 1, got %v\n", aobjs[0])
	}
}
