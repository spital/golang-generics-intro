package main

import (
	"reflect"
	"testing"
)

/* rather DO NOT USE global vars neither init functions !!

 */

var (
	expected_int   = int64(46)
	expected_float = 62.97 // float64 default
)

// TODO Try Fuzzy testing too

// checks if values are equal
func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	// debug.PrintStack()
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

func TestNonGenericSums(t *testing.T) {
	ints, floats := declare_variables()
	assertEqual(t, SumInts(ints), expected_int)
	assertEqual(t, SumFloats(floats), expected_float)
}

func TestGenericSums(t *testing.T) {
	ints, floats := declare_variables()
	assertEqual(t, SumIntsOrFloats[string, int64](ints), expected_int)
	assertEqual(t, SumIntsOrFloats[string, float64](floats), expected_float)
}

func TestGenericSumsTypeParametersInferred(t *testing.T) {
	ints, floats := declare_variables()
	assertEqual(t, SumIntsOrFloats(ints), expected_int)
	assertEqual(t, SumIntsOrFloats(floats), expected_float)

}

func TestGenericSumsWithConstraint(t *testing.T) {
	ints, floats := declare_variables()
	assertEqual(t, SumNumbers(ints), expected_int)
	assertEqual(t, SumNumbers(floats), expected_float)
}
