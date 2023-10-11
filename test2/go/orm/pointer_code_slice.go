package orm

import (
	"database/sql/driver"
	"encoding/json"
)

type PointerCodeSlice []int

func (ss *PointerCodeSlice) Scan(value interface{}) error {
	return json.Unmarshal([]byte(value.(string)), ss)
}

func (ss PointerCodeSlice) Value() (driver.Value, error) {
	b, err := json.Marshal(ss)
	return string(b), err
}
