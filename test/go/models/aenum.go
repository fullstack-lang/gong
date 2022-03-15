package models

import "time"

const CONST_OUTSIDE_GONG = time.Minute * 80

// swagger:enum AEnumType
type AEnumType string

// values for EnumType
const (
	ENUM_VAL1 AEnumType = "ENUM_VAL1_NOT_THE_SAME"
	ENUM_VAL2 AEnumType = "ENUM_VAL2"
)

func (aEnum AEnumType) ToString() (res string) {

	// migration of former implementation of enum
	switch aEnum {
	case ENUM_VAL1:
		res = "ENUM_VAL1_NOT_THE_SAME"
	case ENUM_VAL2:
		res = "ENUM_VAL2"
	}
	return
}

func (aEnum *AEnumType) FromString(input string) {

	switch input {
	case "ENUM_VAL1_NOT_THE_SAME":
		*aEnum = ENUM_VAL1
	case "ENUM_VAL2":
		*aEnum = ENUM_VAL2

	// migration of former implementation of enum
	case "ENUM_VAL1":
		*aEnum = ENUM_VAL1
	}
}
