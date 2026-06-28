// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for Level0
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (level0 Level0) ToInt() (res int) {

	// migration of former implementation of enum
	switch level0 {
	// insertion code per enum code
	case Level0AllGongstructsCode:
		res = 0
	case Level0AllGongenumsCode:
		res = 1
	case Level0Nb:
		res = 2
	}
	return
}

func (level0 *Level0) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*level0 = Level0AllGongstructsCode
		return
	case 1:
		*level0 = Level0AllGongenumsCode
		return
	case 2:
		*level0 = Level0Nb
		return
	default:
		return errUnkownEnum
	}
}

func (level0 *Level0) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Level0AllGongstructsCode":
		*level0 = Level0AllGongstructsCode
	case "Level0AllGongenumsCode":
		*level0 = Level0AllGongenumsCode
	case "Level0Nb":
		*level0 = Level0Nb
	default:
		err = errUnkownEnum
	}
	return
}

func (level0 *Level0) ToCodeString() (res string) {

	switch *level0 {
	// insertion code per enum code
	case Level0AllGongstructsCode:
		res = "Level0AllGongstructsCode"
	case Level0AllGongenumsCode:
		res = "Level0AllGongenumsCode"
	case Level0Nb:
		res = "Level0Nb"
	}
	return
}

func (level0 Level0) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Level0AllGongstructsCode")
	res = append(res, "Level0AllGongenumsCode")
	res = append(res, "Level0Nb")

	return
}

func (level0 Level0) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)
	res = append(res, 2)

	return
}

// Utility function for Level1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (level1 Level1) ToInt() (res int) {

	// migration of former implementation of enum
	switch level1 {
	// insertion code per enum code
	case Level1NamedStructCode:
		res = 0
	case Level1UnNamedStructCode:
		res = 1
	case Level1NamedEnumCode:
		res = 2
	case Level1Nb:
		res = 3
	}
	return
}

func (level1 *Level1) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*level1 = Level1NamedStructCode
		return
	case 1:
		*level1 = Level1UnNamedStructCode
		return
	case 2:
		*level1 = Level1NamedEnumCode
		return
	case 3:
		*level1 = Level1Nb
		return
	default:
		return errUnkownEnum
	}
}

func (level1 *Level1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Level1NamedStructCode":
		*level1 = Level1NamedStructCode
	case "Level1UnNamedStructCode":
		*level1 = Level1UnNamedStructCode
	case "Level1NamedEnumCode":
		*level1 = Level1NamedEnumCode
	case "Level1Nb":
		*level1 = Level1Nb
	default:
		err = errUnkownEnum
	}
	return
}

func (level1 *Level1) ToCodeString() (res string) {

	switch *level1 {
	// insertion code per enum code
	case Level1NamedStructCode:
		res = "Level1NamedStructCode"
	case Level1UnNamedStructCode:
		res = "Level1UnNamedStructCode"
	case Level1NamedEnumCode:
		res = "Level1NamedEnumCode"
	case Level1Nb:
		res = "Level1Nb"
	}
	return
}

func (level1 Level1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Level1NamedStructCode")
	res = append(res, "Level1UnNamedStructCode")
	res = append(res, "Level1NamedEnumCode")
	res = append(res, "Level1Nb")

	return
}

func (level1 Level1) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)
	res = append(res, 2)
	res = append(res, 3)

	return
}

// Utility function for Level2
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (level2 Level2) ToInt() (res int) {

	// migration of former implementation of enum
	switch level2 {
	// insertion code per enum code
	case Level2Structname:
		res = 0
	case Level2Comment:
		res = 1
	case Level2Source:
		res = 2
	case Level2Fields:
		res = 3
	case Level2Enumname:
		res = 4
	case Level2EnumComment:
		res = 5
	case Level2EnumValues:
		res = 6
	case Level2Nb:
		res = 7
	}
	return
}

func (level2 *Level2) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*level2 = Level2Structname
		return
	case 1:
		*level2 = Level2Comment
		return
	case 2:
		*level2 = Level2Source
		return
	case 3:
		*level2 = Level2Fields
		return
	case 4:
		*level2 = Level2Enumname
		return
	case 5:
		*level2 = Level2EnumComment
		return
	case 6:
		*level2 = Level2EnumValues
		return
	case 7:
		*level2 = Level2Nb
		return
	default:
		return errUnkownEnum
	}
}

func (level2 *Level2) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Level2Structname":
		*level2 = Level2Structname
	case "Level2Comment":
		*level2 = Level2Comment
	case "Level2Source":
		*level2 = Level2Source
	case "Level2Fields":
		*level2 = Level2Fields
	case "Level2Enumname":
		*level2 = Level2Enumname
	case "Level2EnumComment":
		*level2 = Level2EnumComment
	case "Level2EnumValues":
		*level2 = Level2EnumValues
	case "Level2Nb":
		*level2 = Level2Nb
	default:
		err = errUnkownEnum
	}
	return
}

func (level2 *Level2) ToCodeString() (res string) {

	switch *level2 {
	// insertion code per enum code
	case Level2Structname:
		res = "Level2Structname"
	case Level2Comment:
		res = "Level2Comment"
	case Level2Source:
		res = "Level2Source"
	case Level2Fields:
		res = "Level2Fields"
	case Level2Enumname:
		res = "Level2Enumname"
	case Level2EnumComment:
		res = "Level2EnumComment"
	case Level2EnumValues:
		res = "Level2EnumValues"
	case Level2Nb:
		res = "Level2Nb"
	}
	return
}

func (level2 Level2) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Level2Structname")
	res = append(res, "Level2Comment")
	res = append(res, "Level2Source")
	res = append(res, "Level2Fields")
	res = append(res, "Level2Enumname")
	res = append(res, "Level2EnumComment")
	res = append(res, "Level2EnumValues")
	res = append(res, "Level2Nb")

	return
}

func (level2 Level2) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)
	res = append(res, 2)
	res = append(res, 3)
	res = append(res, 4)
	res = append(res, 5)
	res = append(res, 6)
	res = append(res, 7)

	return
}

// Utility function for Level3
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (level3 Level3) ToInt() (res int) {

	// migration of former implementation of enum
	switch level3 {
	// insertion code per enum code
	case Level3EnumValueIdentifier:
		res = 0
	case Level3EnumValueString:
		res = 1
	case Level3Nb:
		res = 2
	}
	return
}

func (level3 *Level3) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*level3 = Level3EnumValueIdentifier
		return
	case 1:
		*level3 = Level3EnumValueString
		return
	case 2:
		*level3 = Level3Nb
		return
	default:
		return errUnkownEnum
	}
}

func (level3 *Level3) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Level3EnumValueIdentifier":
		*level3 = Level3EnumValueIdentifier
	case "Level3EnumValueString":
		*level3 = Level3EnumValueString
	case "Level3Nb":
		*level3 = Level3Nb
	default:
		err = errUnkownEnum
	}
	return
}

func (level3 *Level3) ToCodeString() (res string) {

	switch *level3 {
	// insertion code per enum code
	case Level3EnumValueIdentifier:
		res = "Level3EnumValueIdentifier"
	case Level3EnumValueString:
		res = "Level3EnumValueString"
	case Level3Nb:
		res = "Level3Nb"
	}
	return
}

func (level3 Level3) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Level3EnumValueIdentifier")
	res = append(res, "Level3EnumValueString")
	res = append(res, "Level3Nb")

	return
}

func (level3 Level3) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)
	res = append(res, 2)

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
	int | Level0 | Level1 | Level2 | Level3
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	//insertion point for pointers to enum int types | *Level0 | *Level1 | *Level2 | *Level3
	FromCodeString(input string) (err error)
}

// Last line of the template
