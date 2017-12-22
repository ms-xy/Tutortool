package reflecttools

import (
	"reflect"
)

func DereferenceIfPointer(t reflect.Type) (result reflect.Type) {
	defer func() {
		if r := recover(); r != nil {
			result = t
		}
	}()
	result = t.Elem()
	return
}

func GetStructFields(t reflect.Type) <-chan reflect.StructField {
	channel := make(chan reflect.StructField)
	go func() {
		for i := 0; i < t.NumField(); i++ {
			channel <- t.Field(i)
		}
		close(channel)
	}()
	return channel
}

func MapStructRecursively(t reflect.Type, idx []int) map[string][]int {
	m := map[string][]int{}
	t = DereferenceIfPointer(t)
	// only deal with structs, everything else should be a panic
	if t.Kind() == reflect.Struct {
		for field := range GetStructFields(t) {
			cidx := append(idx, field.Index...)
			if field.Anonymous && field.Type.Kind() == reflect.Struct {
				// recursive approach
				_m := MapStructRecursively(field.Type, cidx)
				for _key, _idx := range _m {
					if midx, exists := m[_key]; !exists || len(_idx) < len(midx) {
						m[_key] = _idx
					}
				}
			} else {
				// add field
				if midx, exists := m[field.Name]; !exists || len(cidx) < len(midx) {
					m[field.Name] = cidx
				}
			}
		}
	} else {
		panic("Input type is not a struct")
	}
	return m
}
