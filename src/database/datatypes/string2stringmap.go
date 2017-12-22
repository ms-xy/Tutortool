package datatypes

import (
	"database/sql/driver"
)

type String2StringMap map[string]string

func (this String2StringMap) Value() (driver.Value, error) {
	return toJSON(this)
}

func (this *String2StringMap) Scan(value interface{}) error {
	return fromJSON(value, this)
}
