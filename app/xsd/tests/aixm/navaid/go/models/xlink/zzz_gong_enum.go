// generated code - do not edit
package xlink

// insertion point of enum utility functions
// Utility function for Enum_ActuateType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_actuatetype Enum_ActuateType) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_actuatetype {
	// insertion code per enum code
	case Enum_ActuateType_OnLoad:
		res = "onLoad"
	case Enum_ActuateType_OnRequest:
		res = "onRequest"
	case Enum_ActuateType_Other:
		res = "other"
	case Enum_ActuateType_None:
		res = "none"
	}
	return
}

func (enum_actuatetype *Enum_ActuateType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "onLoad":
		*enum_actuatetype = Enum_ActuateType_OnLoad
		return
	case "onRequest":
		*enum_actuatetype = Enum_ActuateType_OnRequest
		return
	case "other":
		*enum_actuatetype = Enum_ActuateType_Other
		return
	case "none":
		*enum_actuatetype = Enum_ActuateType_None
		return
	default:
		return errUnkownEnum
	}
}

func (enum_actuatetype *Enum_ActuateType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_ActuateType_OnLoad":
		*enum_actuatetype = Enum_ActuateType_OnLoad
	case "Enum_ActuateType_OnRequest":
		*enum_actuatetype = Enum_ActuateType_OnRequest
	case "Enum_ActuateType_Other":
		*enum_actuatetype = Enum_ActuateType_Other
	case "Enum_ActuateType_None":
		*enum_actuatetype = Enum_ActuateType_None
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_actuatetype *Enum_ActuateType) ToCodeString() (res string) {

	switch *enum_actuatetype {
	// insertion code per enum code
	case Enum_ActuateType_OnLoad:
		res = "Enum_ActuateType_OnLoad"
	case Enum_ActuateType_OnRequest:
		res = "Enum_ActuateType_OnRequest"
	case Enum_ActuateType_Other:
		res = "Enum_ActuateType_Other"
	case Enum_ActuateType_None:
		res = "Enum_ActuateType_None"
	}
	return
}

func (enum_actuatetype Enum_ActuateType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_ActuateType_OnLoad")
	res = append(res, "Enum_ActuateType_OnRequest")
	res = append(res, "Enum_ActuateType_Other")
	res = append(res, "Enum_ActuateType_None")

	return
}

func (enum_actuatetype Enum_ActuateType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "onLoad")
	res = append(res, "onRequest")
	res = append(res, "other")
	res = append(res, "none")

	return
}

// Utility function for Enum_ShowType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_showtype Enum_ShowType) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_showtype {
	// insertion code per enum code
	case Enum_ShowType_New_XSD:
		res = "new"
	case Enum_ShowType_Replace:
		res = "replace"
	case Enum_ShowType_Embed:
		res = "embed"
	case Enum_ShowType_Other:
		res = "other"
	case Enum_ShowType_None:
		res = "none"
	}
	return
}

func (enum_showtype *Enum_ShowType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "new":
		*enum_showtype = Enum_ShowType_New_XSD
		return
	case "replace":
		*enum_showtype = Enum_ShowType_Replace
		return
	case "embed":
		*enum_showtype = Enum_ShowType_Embed
		return
	case "other":
		*enum_showtype = Enum_ShowType_Other
		return
	case "none":
		*enum_showtype = Enum_ShowType_None
		return
	default:
		return errUnkownEnum
	}
}

func (enum_showtype *Enum_ShowType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_ShowType_New_XSD":
		*enum_showtype = Enum_ShowType_New_XSD
	case "Enum_ShowType_Replace":
		*enum_showtype = Enum_ShowType_Replace
	case "Enum_ShowType_Embed":
		*enum_showtype = Enum_ShowType_Embed
	case "Enum_ShowType_Other":
		*enum_showtype = Enum_ShowType_Other
	case "Enum_ShowType_None":
		*enum_showtype = Enum_ShowType_None
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_showtype *Enum_ShowType) ToCodeString() (res string) {

	switch *enum_showtype {
	// insertion code per enum code
	case Enum_ShowType_New_XSD:
		res = "Enum_ShowType_New_XSD"
	case Enum_ShowType_Replace:
		res = "Enum_ShowType_Replace"
	case Enum_ShowType_Embed:
		res = "Enum_ShowType_Embed"
	case Enum_ShowType_Other:
		res = "Enum_ShowType_Other"
	case Enum_ShowType_None:
		res = "Enum_ShowType_None"
	}
	return
}

func (enum_showtype Enum_ShowType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_ShowType_New_XSD")
	res = append(res, "Enum_ShowType_Replace")
	res = append(res, "Enum_ShowType_Embed")
	res = append(res, "Enum_ShowType_Other")
	res = append(res, "Enum_ShowType_None")

	return
}

func (enum_showtype Enum_ShowType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "new")
	res = append(res, "replace")
	res = append(res, "embed")
	res = append(res, "other")
	res = append(res, "none")

	return
}

// Utility function for Enum_TypeType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_typetype Enum_TypeType) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_typetype {
	// insertion code per enum code
	case Enum_TypeType_Simple:
		res = "simple"
	case Enum_TypeType_Extended:
		res = "extended"
	case Enum_TypeType_Title:
		res = "title"
	case Enum_TypeType_Resource:
		res = "resource"
	case Enum_TypeType_Locator:
		res = "locator"
	case Enum_TypeType_Arc:
		res = "arc"
	}
	return
}

func (enum_typetype *Enum_TypeType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "simple":
		*enum_typetype = Enum_TypeType_Simple
		return
	case "extended":
		*enum_typetype = Enum_TypeType_Extended
		return
	case "title":
		*enum_typetype = Enum_TypeType_Title
		return
	case "resource":
		*enum_typetype = Enum_TypeType_Resource
		return
	case "locator":
		*enum_typetype = Enum_TypeType_Locator
		return
	case "arc":
		*enum_typetype = Enum_TypeType_Arc
		return
	default:
		return errUnkownEnum
	}
}

func (enum_typetype *Enum_TypeType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_TypeType_Simple":
		*enum_typetype = Enum_TypeType_Simple
	case "Enum_TypeType_Extended":
		*enum_typetype = Enum_TypeType_Extended
	case "Enum_TypeType_Title":
		*enum_typetype = Enum_TypeType_Title
	case "Enum_TypeType_Resource":
		*enum_typetype = Enum_TypeType_Resource
	case "Enum_TypeType_Locator":
		*enum_typetype = Enum_TypeType_Locator
	case "Enum_TypeType_Arc":
		*enum_typetype = Enum_TypeType_Arc
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_typetype *Enum_TypeType) ToCodeString() (res string) {

	switch *enum_typetype {
	// insertion code per enum code
	case Enum_TypeType_Simple:
		res = "Enum_TypeType_Simple"
	case Enum_TypeType_Extended:
		res = "Enum_TypeType_Extended"
	case Enum_TypeType_Title:
		res = "Enum_TypeType_Title"
	case Enum_TypeType_Resource:
		res = "Enum_TypeType_Resource"
	case Enum_TypeType_Locator:
		res = "Enum_TypeType_Locator"
	case Enum_TypeType_Arc:
		res = "Enum_TypeType_Arc"
	}
	return
}

func (enum_typetype Enum_TypeType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_TypeType_Simple")
	res = append(res, "Enum_TypeType_Extended")
	res = append(res, "Enum_TypeType_Title")
	res = append(res, "Enum_TypeType_Resource")
	res = append(res, "Enum_TypeType_Locator")
	res = append(res, "Enum_TypeType_Arc")

	return
}

func (enum_typetype Enum_TypeType) CodeValues() (res []string) {

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
