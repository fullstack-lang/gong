package models

type MultiplicityType string

// values for EnumType
const (
	ZERO_ONE MultiplicityType = "0..1"
	ONE      MultiplicityType = "1"
	MANY     MultiplicityType = "*"
)
