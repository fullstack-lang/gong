// generated code - do not edit
package xlink

// insertion point of enum utility functions
// Utility function for EnumActuateType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumactuatetype EnumActuateType) ToString() (res string) {

	// migration of former implementation of enum
	switch enumactuatetype {
	// insertion code per enum code
	case EnumActuateTypeOnLoad:
		res = "onLoad"
	case EnumActuateTypeOnRequest:
		res = "onRequest"
	case EnumActuateTypeOther:
		res = "other"
	case EnumActuateTypeNone:
		res = "none"
	}
	return
}

func (enumactuatetype *EnumActuateType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "onLoad":
		*enumactuatetype = EnumActuateTypeOnLoad
		return
	case "onRequest":
		*enumactuatetype = EnumActuateTypeOnRequest
		return
	case "other":
		*enumactuatetype = EnumActuateTypeOther
		return
	case "none":
		*enumactuatetype = EnumActuateTypeNone
		return
	default:
		return errUnkownEnum
	}
}

func (enumactuatetype *EnumActuateType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumActuateTypeOnLoad":
		*enumactuatetype = EnumActuateTypeOnLoad
	case "EnumActuateTypeOnRequest":
		*enumactuatetype = EnumActuateTypeOnRequest
	case "EnumActuateTypeOther":
		*enumactuatetype = EnumActuateTypeOther
	case "EnumActuateTypeNone":
		*enumactuatetype = EnumActuateTypeNone
	default:
		err = errUnkownEnum
	}
	return
}

func (enumactuatetype *EnumActuateType) ToCodeString() (res string) {

	switch *enumactuatetype {
	// insertion code per enum code
	case EnumActuateTypeOnLoad:
		res = "EnumActuateTypeOnLoad"
	case EnumActuateTypeOnRequest:
		res = "EnumActuateTypeOnRequest"
	case EnumActuateTypeOther:
		res = "EnumActuateTypeOther"
	case EnumActuateTypeNone:
		res = "EnumActuateTypeNone"
	}
	return
}

func (enumactuatetype EnumActuateType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumActuateTypeOnLoad")
	res = append(res, "EnumActuateTypeOnRequest")
	res = append(res, "EnumActuateTypeOther")
	res = append(res, "EnumActuateTypeNone")

	return
}

func (enumactuatetype EnumActuateType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "onLoad")
	res = append(res, "onRequest")
	res = append(res, "other")
	res = append(res, "none")

	return
}

// Utility function for EnumShowType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumshowtype EnumShowType) ToString() (res string) {

	// migration of former implementation of enum
	switch enumshowtype {
	// insertion code per enum code
	case EnumShowTypeNew:
		res = "new"
	case EnumShowTypeReplace:
		res = "replace"
	case EnumShowTypeEmbed:
		res = "embed"
	case EnumShowTypeOther:
		res = "other"
	case EnumShowTypeNone:
		res = "none"
	}
	return
}

func (enumshowtype *EnumShowType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "new":
		*enumshowtype = EnumShowTypeNew
		return
	case "replace":
		*enumshowtype = EnumShowTypeReplace
		return
	case "embed":
		*enumshowtype = EnumShowTypeEmbed
		return
	case "other":
		*enumshowtype = EnumShowTypeOther
		return
	case "none":
		*enumshowtype = EnumShowTypeNone
		return
	default:
		return errUnkownEnum
	}
}

func (enumshowtype *EnumShowType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumShowTypeNew":
		*enumshowtype = EnumShowTypeNew
	case "EnumShowTypeReplace":
		*enumshowtype = EnumShowTypeReplace
	case "EnumShowTypeEmbed":
		*enumshowtype = EnumShowTypeEmbed
	case "EnumShowTypeOther":
		*enumshowtype = EnumShowTypeOther
	case "EnumShowTypeNone":
		*enumshowtype = EnumShowTypeNone
	default:
		err = errUnkownEnum
	}
	return
}

func (enumshowtype *EnumShowType) ToCodeString() (res string) {

	switch *enumshowtype {
	// insertion code per enum code
	case EnumShowTypeNew:
		res = "EnumShowTypeNew"
	case EnumShowTypeReplace:
		res = "EnumShowTypeReplace"
	case EnumShowTypeEmbed:
		res = "EnumShowTypeEmbed"
	case EnumShowTypeOther:
		res = "EnumShowTypeOther"
	case EnumShowTypeNone:
		res = "EnumShowTypeNone"
	}
	return
}

func (enumshowtype EnumShowType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumShowTypeNew")
	res = append(res, "EnumShowTypeReplace")
	res = append(res, "EnumShowTypeEmbed")
	res = append(res, "EnumShowTypeOther")
	res = append(res, "EnumShowTypeNone")

	return
}

func (enumshowtype EnumShowType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "new")
	res = append(res, "replace")
	res = append(res, "embed")
	res = append(res, "other")
	res = append(res, "none")

	return
}

// Utility function for EnumTypeType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enumtypetype EnumTypeType) ToString() (res string) {

	// migration of former implementation of enum
	switch enumtypetype {
	// insertion code per enum code
	case EnumTypeTypeSimple:
		res = "simple"
	case EnumTypeTypeExtended:
		res = "extended"
	case EnumTypeTypeTitle:
		res = "title"
	case EnumTypeTypeResource:
		res = "resource"
	case EnumTypeTypeLocator:
		res = "locator"
	case EnumTypeTypeArc:
		res = "arc"
	}
	return
}

func (enumtypetype *EnumTypeType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "simple":
		*enumtypetype = EnumTypeTypeSimple
		return
	case "extended":
		*enumtypetype = EnumTypeTypeExtended
		return
	case "title":
		*enumtypetype = EnumTypeTypeTitle
		return
	case "resource":
		*enumtypetype = EnumTypeTypeResource
		return
	case "locator":
		*enumtypetype = EnumTypeTypeLocator
		return
	case "arc":
		*enumtypetype = EnumTypeTypeArc
		return
	default:
		return errUnkownEnum
	}
}

func (enumtypetype *EnumTypeType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "EnumTypeTypeSimple":
		*enumtypetype = EnumTypeTypeSimple
	case "EnumTypeTypeExtended":
		*enumtypetype = EnumTypeTypeExtended
	case "EnumTypeTypeTitle":
		*enumtypetype = EnumTypeTypeTitle
	case "EnumTypeTypeResource":
		*enumtypetype = EnumTypeTypeResource
	case "EnumTypeTypeLocator":
		*enumtypetype = EnumTypeTypeLocator
	case "EnumTypeTypeArc":
		*enumtypetype = EnumTypeTypeArc
	default:
		err = errUnkownEnum
	}
	return
}

func (enumtypetype *EnumTypeType) ToCodeString() (res string) {

	switch *enumtypetype {
	// insertion code per enum code
	case EnumTypeTypeSimple:
		res = "EnumTypeTypeSimple"
	case EnumTypeTypeExtended:
		res = "EnumTypeTypeExtended"
	case EnumTypeTypeTitle:
		res = "EnumTypeTypeTitle"
	case EnumTypeTypeResource:
		res = "EnumTypeTypeResource"
	case EnumTypeTypeLocator:
		res = "EnumTypeTypeLocator"
	case EnumTypeTypeArc:
		res = "EnumTypeTypeArc"
	}
	return
}

func (enumtypetype EnumTypeType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "EnumTypeTypeSimple")
	res = append(res, "EnumTypeTypeExtended")
	res = append(res, "EnumTypeTypeTitle")
	res = append(res, "EnumTypeTypeResource")
	res = append(res, "EnumTypeTypeLocator")
	res = append(res, "EnumTypeTypeArc")

	return
}

func (enumtypetype EnumTypeType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "simple")
	res = append(res, "extended")
	res = append(res, "title")
	res = append(res, "resource")
	res = append(res, "locator")
	res = append(res, "arc")

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
