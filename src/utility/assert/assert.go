package assert

import (
	"reflect"
	"strings"
)

/*
Some assertion helper functions.
*/
func ErrorNil(err error, msgs ...string) {
	if err != nil {
		if len(msgs) > 0 {
			prefix := strings.Join(msgs, ": ")
			panic(prefix + ": " + err.Error())
		} else {
			panic(err)
		}
	}
}

func True(b bool, msg string) {
	if !b {
		panic(msg)
	}
}

func NotEmpty(i interface{}, msg string) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	_t := t.Kind()

	if _t == reflect.Array ||
		_t == reflect.String {

		if v.Len() <= 0 {
			panic(msg)
		}

	} else if _t == reflect.Map ||
		_t == reflect.Slice {

		if v.IsNil() || v.Len() <= 0 {
			panic(msg)
		}

	} else if _t == reflect.Ptr ||
		_t == reflect.Chan ||
		_t == reflect.UnsafePointer {

		if v.IsNil() {
			panic(msg)
		}
	}
}
