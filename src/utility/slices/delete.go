package slices

import (
	"errors"
	"reflect"
)

// https://play.golang.org/p/lJhWBESVyB
func Delete(slice interface{}, pos int) (interface{}, error) {
	if pos < 0 {
		return slice, errors.New("second parameter (pos) must be greater or equal to 0")
	}
	oslice := reflect.ValueOf(slice)
	if oslice.Type().Kind() != reflect.Slice {
		return slice, errors.New("first parameter (slice) expected to be a slice")
	}
	len := oslice.Len()
	if pos > len-1 {
		return slice, errors.New("second parameter (pos) must be at most len(slice)-1")
	}

	nslice := reflect.MakeSlice(oslice.Type(), len-1, len-1)
	i := 0
	for ; i < pos; i++ {
		nslice.Index(i).Set(oslice.Index(i))
	}
	for ; i < len-1; i++ {
		nslice.Index(i).Set(oslice.Index(i + 1))
	}
	return nslice.Interface(), nil
}
