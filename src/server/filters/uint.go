package filters

import (
	// templating engine
	"github.com/flosch/pongo2"

	// utility
	"strconv"
)

func init() {
	pongo2.RegisterFilter("uint2string", filterUint2String)
}

/*
Convert an uint to a string
*/
func filterUint2String(_v *pongo2.Value, p *pongo2.Value) (*pongo2.Value, *pongo2.Error) {

	var _uint64 uint64

	switch v := _v.Interface().(type) {
	case uint64:
		_uint64 = v
	case uint:
		_uint64 = uint64(v)
	case uint32:
		_uint64 = uint64(v)
	case uint16:
		_uint64 = uint64(v)
	case uint8:
		_uint64 = uint64(v)

	default:
		return nil, &pongo2.Error{
			OrigError: EINVALIDTYPE,
			Sender:    "filter:uint2string",
		}
	}

	return pongo2.AsSafeValue(strconv.FormatUint(_uint64, 10)), nil
}
