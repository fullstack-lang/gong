// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for AEnumType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (aenumtype AEnumType) ToString() (res string) {

	// migration of former implementation of enum
	switch aenumtype {
	// insertion code per enum code
	case ENUM_VAL1:
		res = "ENUM_VAL1_NOT_THE_SAME"
	case ENUM_VAL2:
		res = "ENUM_VAL2"
	}
	return
}

func (aenumtype *AEnumType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ENUM_VAL1_NOT_THE_SAME":
		*aenumtype = ENUM_VAL1
		return
	case "ENUM_VAL2":
		*aenumtype = ENUM_VAL2
		return
	default:
		return errUnkownEnum
	}
}

func (aenumtype *AEnumType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ENUM_VAL1":
		*aenumtype = ENUM_VAL1
	case "ENUM_VAL2":
		*aenumtype = ENUM_VAL2
	default:
		err = errUnkownEnum
	}
	return
}

func (aenumtype *AEnumType) ToCodeString() (res string) {

	switch *aenumtype {
	// insertion code per enum code
	case ENUM_VAL1:
		res = "ENUM_VAL1"
	case ENUM_VAL2:
		res = "ENUM_VAL2"
	}
	return
}

func (aenumtype AEnumType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ENUM_VAL1")
	res = append(res, "ENUM_VAL2")

	return
}

func (aenumtype AEnumType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ENUM_VAL1_NOT_THE_SAME")
	res = append(res, "ENUM_VAL2")

	return
}

// Utility function for BEnumType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (benumtype BEnumType) ToString() (res string) {

	// migration of former implementation of enum
	switch benumtype {
	// insertion code per enum code
	case BENUM_VAL1:
		res = "BENUM_VAL1_NOT_THE_SAME"
	case BENUM_VAL2:
		res = "BENUM_VAL2"
	}
	return
}

func (benumtype *BEnumType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "BENUM_VAL1_NOT_THE_SAME":
		*benumtype = BENUM_VAL1
		return
	case "BENUM_VAL2":
		*benumtype = BENUM_VAL2
		return
	default:
		return errUnkownEnum
	}
}

func (benumtype *BEnumType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "BENUM_VAL1":
		*benumtype = BENUM_VAL1
	case "BENUM_VAL2":
		*benumtype = BENUM_VAL2
	default:
		err = errUnkownEnum
	}
	return
}

func (benumtype *BEnumType) ToCodeString() (res string) {

	switch *benumtype {
	// insertion code per enum code
	case BENUM_VAL1:
		res = "BENUM_VAL1"
	case BENUM_VAL2:
		res = "BENUM_VAL2"
	}
	return
}

func (benumtype BEnumType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "BENUM_VAL1")
	res = append(res, "BENUM_VAL2")

	return
}

func (benumtype BEnumType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "BENUM_VAL1_NOT_THE_SAME")
	res = append(res, "BENUM_VAL2")

	return
}

// Utility function for CEnumTypeInt
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (cenumtypeint CEnumTypeInt) ToInt() (res int) {

	// migration of former implementation of enum
	switch cenumtypeint {
	// insertion code per enum code
	case CENUM_VAL1:
		res = 0
	case CENUM_VAL2:
		res = 1
	}
	return
}

func (cenumtypeint *CEnumTypeInt) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*cenumtypeint = CENUM_VAL1
		return
	case 1:
		*cenumtypeint = CENUM_VAL2
		return
	default:
		return errUnkownEnum
	}
}

func (cenumtypeint *CEnumTypeInt) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "CENUM_VAL1":
		*cenumtypeint = CENUM_VAL1
	case "CENUM_VAL2":
		*cenumtypeint = CENUM_VAL2
	default:
		err = errUnkownEnum
	}
	return
}

func (cenumtypeint *CEnumTypeInt) ToCodeString() (res string) {

	switch *cenumtypeint {
	// insertion code per enum code
	case CENUM_VAL1:
		res = "CENUM_VAL1"
	case CENUM_VAL2:
		res = "CENUM_VAL2"
	}
	return
}

func (cenumtypeint CEnumTypeInt) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "CENUM_VAL1")
	res = append(res, "CENUM_VAL2")

	return
}

func (cenumtypeint CEnumTypeInt) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)

	return
}

// Utility function for FEnumType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (fenumtype FEnumType) ToString() (res string) {

	// migration of former implementation of enum
	switch fenumtype {
	// insertion code per enum code
	}
	return
}

func (fenumtype *FEnumType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (fenumtype *FEnumType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		err = errUnkownEnum
	}
	return
}

func (fenumtype *FEnumType) ToCodeString() (res string) {

	switch *fenumtype {
	// insertion code per enum code
	}
	return
}

func (fenumtype FEnumType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (fenumtype FEnumType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for StacksNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (stacksnames StacksNames) ToString() (res string) {

	// migration of former implementation of enum
	switch stacksnames {
	// insertion code per enum code
	case Test:
		res = "test"
	}
	return
}

func (stacksnames *StacksNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "test":
		*stacksnames = Test
		return
	default:
		return errUnkownEnum
	}
}

func (stacksnames *StacksNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Test":
		*stacksnames = Test
	default:
		err = errUnkownEnum
	}
	return
}

func (stacksnames *StacksNames) ToCodeString() (res string) {

	switch *stacksnames {
	// insertion code per enum code
	case Test:
		res = "Test"
	}
	return
}

func (stacksnames StacksNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Test")

	return
}

func (stacksnames StacksNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "test")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	Codes() []string
	CodeValues() []string
	ToString() string
}

type PointerToGongstructEnumStringField interface {
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int | CEnumTypeInt
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	*CEnumTypeInt
	FromCodeString(input string) (err error)
}

// Last line of the template
