package slices

import (
	"errors"
	"reflect"
)

// https://play.golang.org/p/xygZoflaUi
func Insert(slice interface{}, pos int, item interface{}) (interface{}, error) {
	oslice := reflect.ValueOf(slice)
	if oslice.Type().Kind() != reflect.Slice {
		return slice, errors.New("first parameter (slice) expected to be a slice")
	}
	len := oslice.Len()
	if pos > len {
		return slice, errors.New("second parameter (pos) must be at most len(slice) (==appending)")
	}
	if oslice.Type() != reflect.SliceOf(reflect.TypeOf(item)) {
		return slice, errors.New("third parameter (item) must be of the type contained in the slice given by the first parameter (slice)")
	}

	nslice := reflect.MakeSlice(oslice.Type(), len+1, len+1)
	i := 0
	for ; i < pos; i++ {
		nslice.Index(i).Set(oslice.Index(i))
	}
	nslice.Index(i).Set(reflect.ValueOf(item))
	i++
	for ; i < len+1; i++ {
		nslice.Index(i).Set(oslice.Index(i - 1))
	}
	return nslice.Interface(), nil
}
