package datatypes

import (
	"database/sql/driver"
	"encoding/json"
)

type JSONB struct {
	Data interface{}
}

func (this JSONB) Value() (driver.Value, error) {
	bytes, err := json.Marshal(this.Data)
	return string(bytes), err
}

func (this *JSONB) Scan(value interface{}) error {
	if sv, err := driver.String.ConvertValue(value); err != nil {
		return err
	} else {
		*this = JSONB{}
		return json.Unmarshal(sv.([]byte), this.Data)
	}
}

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- //
// helper functions for other types

func toJSON(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	return string(bytes), err
}

func fromJSON(dbval interface{}, v interface{}) error {
	if sv, err := driver.String.ConvertValue(dbval); err != nil {
		return err
	} else {
		return json.Unmarshal(sv.([]byte), v)
	}
}
