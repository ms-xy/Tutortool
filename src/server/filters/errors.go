package filters

import (
	"errors"
)

var (
	EINVALIDTYPE = errors.New("invalid input parameter type")
	ENOTFOUND    = errors.New("key/value/item not found")
)
