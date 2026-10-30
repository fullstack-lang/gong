// generated code - do not edit
package refs

// insertion point of enum utility functions
// Utility function for Enum_StatusEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_statusenum Enum_StatusEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_statusenum {
	// insertion code per enum code
	case Enum_StatusEnum_ACTIVE:
		res = "ACTIVE"
	case Enum_StatusEnum_INACTIVE:
		res = "INACTIVE"
	case Enum_StatusEnum_DEPRECATED:
		res = "DEPRECATED"
	}
	return
}

func (enum_statusenum *Enum_StatusEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ACTIVE":
		*enum_statusenum = Enum_StatusEnum_ACTIVE
		return
	case "INACTIVE":
		*enum_statusenum = Enum_StatusEnum_INACTIVE
		return
	case "DEPRECATED":
		*enum_statusenum = Enum_StatusEnum_DEPRECATED
		return
	default:
		return errUnkownEnum
	}
}

func (enum_statusenum *Enum_StatusEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_StatusEnum_ACTIVE":
		*enum_statusenum = Enum_StatusEnum_ACTIVE
	case "Enum_StatusEnum_INACTIVE":
		*enum_statusenum = Enum_StatusEnum_INACTIVE
	case "Enum_StatusEnum_DEPRECATED":
		*enum_statusenum = Enum_StatusEnum_DEPRECATED
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_statusenum *Enum_StatusEnum) ToCodeString() (res string) {

	switch *enum_statusenum {
	// insertion code per enum code
	case Enum_StatusEnum_ACTIVE:
		res = "Enum_StatusEnum_ACTIVE"
	case Enum_StatusEnum_INACTIVE:
		res = "Enum_StatusEnum_INACTIVE"
	case Enum_StatusEnum_DEPRECATED:
		res = "Enum_StatusEnum_DEPRECATED"
	}
	return
}

func (enum_statusenum Enum_StatusEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_StatusEnum_ACTIVE")
	res = append(res, "Enum_StatusEnum_INACTIVE")
	res = append(res, "Enum_StatusEnum_DEPRECATED")

	return
}

func (enum_statusenum Enum_StatusEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ACTIVE")
	res = append(res, "INACTIVE")
	res = append(res, "DEPRECATED")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	Codes() []string
	CodeValues() []string
	ToString() string
}

type GongEnumStringPointer interface {
	FromCodeString(input string) (err error)
}

type PointerToGongstructEnumStringField = GongEnumStringPointer

type GongstructEnumIntField interface {
	int
	Codes() []string
	CodeValues() []int
}

type GongEnumIntPointer interface {
	//insertion point for pointers to enum int types
	FromCodeString(input string) (err error)
}

type PointerToGongstructEnumIntField = GongEnumIntPointer

// Last line of the template
