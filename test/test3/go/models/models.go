package models

import "time"

type EnumTypeString string

const (
	EnumTypeString_Value1 EnumTypeString = "Value1"
	EnumTypeString_Value2 EnumTypeString = "Value2"
)

type EnumTypeInt int

const (
	EnumTypeInt_Value1 EnumTypeInt = 1
	EnumTypeInt_Value2 EnumTypeInt = 2
)

type A struct {

	//gong:text gong:width 600 gong:height 300
	Name string

	//gong:accordion-start "Time fields"
	Date time.Time

	//gong:accordion-end
	Duration time.Duration

	FloatValue float64
	IntValue   int

	EnumString EnumTypeString
	EnumInt    EnumTypeInt

	B *B

	Bs []*B

	UUID string
}
type B struct {
	Name string
}
