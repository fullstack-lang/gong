// generated code - do not edit
package spatialrepresentation

// insertion point of enum utility functions
// Utility function for Enum_MD_PixelOrientationCode_Type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enum_md_pixelorientationcode_type Enum_MD_PixelOrientationCode_Type) ToString() (res string) {

	// migration of former implementation of enum
	switch enum_md_pixelorientationcode_type {
	// insertion code per enum code
	case Enum_MD_PixelOrientationCode_Type_Center:
		res = "center"
	case Enum_MD_PixelOrientationCode_Type_LowerLeft:
		res = "lowerLeft"
	case Enum_MD_PixelOrientationCode_Type_LowerRight:
		res = "lowerRight"
	case Enum_MD_PixelOrientationCode_Type_UpperRight:
		res = "upperRight"
	case Enum_MD_PixelOrientationCode_Type_UpperLeft:
		res = "upperLeft"
	}
	return
}

func (enum_md_pixelorientationcode_type *Enum_MD_PixelOrientationCode_Type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "center":
		*enum_md_pixelorientationcode_type = Enum_MD_PixelOrientationCode_Type_Center
		return
	case "lowerLeft":
		*enum_md_pixelorientationcode_type = Enum_MD_PixelOrientationCode_Type_LowerLeft
		return
	case "lowerRight":
		*enum_md_pixelorientationcode_type = Enum_MD_PixelOrientationCode_Type_LowerRight
		return
	case "upperRight":
		*enum_md_pixelorientationcode_type = Enum_MD_PixelOrientationCode_Type_UpperRight
		return
	case "upperLeft":
		*enum_md_pixelorientationcode_type = Enum_MD_PixelOrientationCode_Type_UpperLeft
		return
	default:
		return errUnkownEnum
	}
}

func (enum_md_pixelorientationcode_type *Enum_MD_PixelOrientationCode_Type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enum_MD_PixelOrientationCode_Type_Center":
		*enum_md_pixelorientationcode_type = Enum_MD_PixelOrientationCode_Type_Center
	case "Enum_MD_PixelOrientationCode_Type_LowerLeft":
		*enum_md_pixelorientationcode_type = Enum_MD_PixelOrientationCode_Type_LowerLeft
	case "Enum_MD_PixelOrientationCode_Type_LowerRight":
		*enum_md_pixelorientationcode_type = Enum_MD_PixelOrientationCode_Type_LowerRight
	case "Enum_MD_PixelOrientationCode_Type_UpperRight":
		*enum_md_pixelorientationcode_type = Enum_MD_PixelOrientationCode_Type_UpperRight
	case "Enum_MD_PixelOrientationCode_Type_UpperLeft":
		*enum_md_pixelorientationcode_type = Enum_MD_PixelOrientationCode_Type_UpperLeft
	default:
		err = errUnkownEnum
	}
	return
}

func (enum_md_pixelorientationcode_type *Enum_MD_PixelOrientationCode_Type) ToCodeString() (res string) {

	switch *enum_md_pixelorientationcode_type {
	// insertion code per enum code
	case Enum_MD_PixelOrientationCode_Type_Center:
		res = "Enum_MD_PixelOrientationCode_Type_Center"
	case Enum_MD_PixelOrientationCode_Type_LowerLeft:
		res = "Enum_MD_PixelOrientationCode_Type_LowerLeft"
	case Enum_MD_PixelOrientationCode_Type_LowerRight:
		res = "Enum_MD_PixelOrientationCode_Type_LowerRight"
	case Enum_MD_PixelOrientationCode_Type_UpperRight:
		res = "Enum_MD_PixelOrientationCode_Type_UpperRight"
	case Enum_MD_PixelOrientationCode_Type_UpperLeft:
		res = "Enum_MD_PixelOrientationCode_Type_UpperLeft"
	}
	return
}

func (enum_md_pixelorientationcode_type Enum_MD_PixelOrientationCode_Type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enum_MD_PixelOrientationCode_Type_Center")
	res = append(res, "Enum_MD_PixelOrientationCode_Type_LowerLeft")
	res = append(res, "Enum_MD_PixelOrientationCode_Type_LowerRight")
	res = append(res, "Enum_MD_PixelOrientationCode_Type_UpperRight")
	res = append(res, "Enum_MD_PixelOrientationCode_Type_UpperLeft")

	return
}

func (enum_md_pixelorientationcode_type Enum_MD_PixelOrientationCode_Type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "center")
	res = append(res, "lowerLeft")
	res = append(res, "lowerRight")
	res = append(res, "upperRight")
	res = append(res, "upperLeft")

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
