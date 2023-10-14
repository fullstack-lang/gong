package orm

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type StringSlice []int

// Value makes StringSlice implement the driver.Valuer interface.
func (s StringSlice) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Scan makes StringSlice implement the sql.Scanner interface.
func (s *StringSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return error(fmt.Errorf("Failed to unmarshal StringSlice value: %v", value))
	}

	return json.Unmarshal(bytes, s)
}
