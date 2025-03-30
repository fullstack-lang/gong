// generated code - do not edit
package orm

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type IntSlice []int

// Value makes IntSlice implement the driver.Valuer interface.
func (s IntSlice) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Scan makes IntSlice implement the sql.Scanner interface.
func (s *IntSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return error(fmt.Errorf("Failed to unmarshal IntSlice value: %v", value))
	}

	return json.Unmarshal(bytes, s)
}
