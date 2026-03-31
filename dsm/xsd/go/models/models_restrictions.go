package models

type Enumeration struct {
	Name string
	Annotated
	ElementWithValueAttribute
}

type MinInclusive struct {
	Name string
	Annotated
	ElementWithValueAttribute
}

type MaxInclusive struct {
	Name string
	Annotated
	ElementWithValueAttribute
}

type Pattern struct {
	Name string
	Annotated
	ElementWithValueAttribute
}

type WhiteSpace struct {
	Name string
	Annotated
	ElementWithValueAttribute
}

type MinLength struct {
	Name string
	Annotated
	ElementWithValueAttribute
}

type MaxLength struct {
	Name string
	Annotated
	ElementWithValueAttribute
}

type Length struct {
	Name string
	Annotated
	ElementWithValueAttribute
}

type TotalDigit struct {
	Name string
	Annotated
	ElementWithValueAttribute
}
