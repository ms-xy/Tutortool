package datatypes

import (
	"database/sql/driver"
)

type Strings []string

func (this Strings) Value() (driver.Value, error) {
	return toJSON(this)
}

func (this *Strings) Scan(value interface{}) error {
	return fromJSON(value, this)
}
