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

type FEnumType string
