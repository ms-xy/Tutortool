package requesttools

import (
	// utility
	"net/http"
	"strconv"

	// assertion tools and helpers
	"github.com/ms-xy/Tutortool/src/utility/assert"
)

func FormValueNotEmpty(r *http.Request, key string) string {
	value := r.FormValue(key)
	assert.NotEmpty(value, "missing parameter: "+key)
	return value
}

func FormValueAs_int64(r *http.Request, key string) int64 {
	num, err := strconv.ParseInt(FormValueNotEmpty(r, key), 10, 64)
	assert.ErrorNil(err, "could not parse parameter '"+key+"': ")
	return num
}

func FormValueAs_int(r *http.Request, key string) int {
	num, err := strconv.ParseInt(FormValueNotEmpty(r, key), 10, 32)
	assert.ErrorNil(err, "could not parse parameter '"+key+"': ")
	return int(num)
}

func FormValueAs_uint(r *http.Request, key string) uint {
	num, err := strconv.ParseUint(FormValueNotEmpty(r, key), 10, 32)
	assert.ErrorNil(err, "could not parse parameter '"+key+"': ")
	return uint(num)
}
