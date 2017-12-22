package slices

// https://play.golang.org/p/gh-3fngSx4

import (
	"reflect"
)

func BytesEqual(slices [][]byte) bool {
	li := len(slices[0])
	lj := len(slices)
	var b byte
	for i := 0; i < li; i++ {
		b = slices[0][i]
		for j := 1; j < lj; j++ {
			if slices[j][i] != b {
				return false
			}
		}
	}
	return true
}

func IntsEqual(slices [][]int) bool {
	li := len(slices[0])
	lj := len(slices)
	var b int
	for i := 0; i < li; i++ {
		b = slices[0][i]
		for j := 1; j < lj; j++ {
			if slices[j][i] != b {
				return false
			}
		}
	}
	return true
}

func StringsEqual(slices []string) bool {
	li := len(slices[0])
	lj := len(slices)
	var b byte
	for i := 0; i < li; i++ {
		b = slices[0][i]
		for j := 1; j < lj; j++ {
			if slices[j][i] != b {
				return false
			}
		}
	}
	return true
}

func Equal(slices ...interface{}) bool {
	if len(slices) > 1 {
		val0 := reflect.ValueOf(slices[0])
		slicetype := val0.Type()
		slicelen := val0.Len()

		for _, slice := range slices[1:] {
			val := reflect.ValueOf(slice)
			if val.Type() != slicetype || val.Len() != slicelen {
				return false
			}
		}

		switch slices[0].(type) {
		case []byte:
			byteslices := make([][]byte, len(slices))
			for i := 0; i < len(slices); i++ {
				byteslices[i] = slices[i].([]byte)
			}
			return BytesEqual(byteslices)
		case []int:
			intslices := make([][]int, len(slices))
			for i := 0; i < len(slices); i++ {
				intslices[i] = slices[i].([]int)
			}
			return IntsEqual(intslices)
		case string:
			strings := make([]string, len(slices))
			for i := 0; i < len(slices); i++ {
				strings[i] = slices[i].(string)
			}
			return StringsEqual(strings)
		}
	}
	return true
}
