package models

import (
	"strings"
)

// GongStruct is a go struct that is selected by the gongc compiler
// swagger:model
type GongStruct struct {
	Name string

	// Fields (stored according to the source file order)
	// swagger:ignore
	Fields []FieldInterface `gorm:"-"`

	// Slice of Fields by their type (not populated by the gongc)
	GongBasicFields                  []*GongBasicField
	GongTimeFields                   []*GongTimeField
	PointerToGongStructFields        []*PointerToGongStructField
	SliceOfPointerToGongStructFields []*SliceOfPointerToGongStructField

	// HasOnAfterUpdateSignature is used to generate orchestrator code
	HasOnAfterUpdateSignature bool

	// IsIgnoredForFront, if true, let the compiler avoid generation of code
	// in the angular front
	// .frontignore file provides a pattern for parsing go files similar to the gitingore syntax "*"
	// every GongStruct defined in matched go file has IsIgnoredForFront set to true
	IsIgnoredForFront bool
}

// HasNameField indicates wether the gong struct has a field with Name "Name"
//
// This is important since
// - only Gong Struct with Name field can be stored in DB
// - only Gong Struct without file
func (gongStruct *GongStruct) HasNameField() (hasNameField bool) {

	// hasNameField default value is false
	for _, field := range gongStruct.Fields {
		switch field := field.(type) {
		case *GongBasicField:
			if field.Name == "Name" {
				hasNameField = true
				continue
			}
		}
	}

	return
}

// ComputeFielProloguesEpilogues computes wether the field shall
// be preceded by the prologue declaring the containing anonymous struct
func (gongStruct *GongStruct) ComputeFielProloguesEpilogues(field FieldInterface) (
	prologueDB, epilogueDB,
	prologuePointerEncoding, epiloguePointerEncoding,
	prologueWOP, epilogueWOP string) {

	// is the field within an anomous struct
	fieldName := field.GetName()
	fieldNameSplitted := strings.Split(fieldName, ".")
	isWithinAnAnonymousStruct := len(fieldNameSplitted) > 1

	if !isWithinAnAnonymousStruct {
		return "", "", "", "", "", ""
	}

	prefix := fieldNameSplitted[0]

	// among all fields, records the index of fields with this prefix
	var indexesOfFieldsWithPrefix []int
	var indexOfTheFieldOfInterest int
	var indexesOfFieldsWithPrefixPointerEncoding []int
	var indexOfTheFieldOfInterestPointerEncoding int
	var indexesOfFieldsWithPrefixWOP []int
	var indexOfTheFieldOfInterestWOP int
	for idx, _field := range gongStruct.Fields {

		_fieldName := _field.GetName()
		_fieldNameSplitted := strings.Split(_fieldName, ".")
		_isWithinAnAnonymousStruct := len(_fieldNameSplitted) > 1
		if _isWithinAnAnonymousStruct && _fieldNameSplitted[0] == prefix {
			indexesOfFieldsWithPrefix = append(indexesOfFieldsWithPrefix, idx)
		}
		if _field == field {
			indexOfTheFieldOfInterest = idx
		}

		switch _field.(type) {
		case *PointerToGongStructField, *SliceOfPointerToGongStructField:
			if _isWithinAnAnonymousStruct && _fieldNameSplitted[0] == prefix {
				indexesOfFieldsWithPrefixPointerEncoding = append(indexesOfFieldsWithPrefixPointerEncoding, idx)
			}
			if _field == field {
				indexOfTheFieldOfInterestPointerEncoding = idx
			}
		default:
			if _isWithinAnAnonymousStruct && _fieldNameSplitted[0] == prefix {
				indexesOfFieldsWithPrefixWOP = append(indexesOfFieldsWithPrefixWOP, idx)
			}
			if _field == field {
				indexOfTheFieldOfInterestWOP = idx
			}
		}
	}

	// compute within all fields wether this field is the first with this prefix

	if indexOfTheFieldOfInterest == indexesOfFieldsWithPrefix[0] {
		prologueDB = "\n\n	" + prefix + " struct {"
	}
	if indexOfTheFieldOfInterest == indexesOfFieldsWithPrefix[len(indexesOfFieldsWithPrefix)-1] {
		epilogueDB = "\n\n	} `gorm:\"embedded\"`"
	}

	switch field.(type) {
	case *PointerToGongStructField, *SliceOfPointerToGongStructField:
		if indexOfTheFieldOfInterestPointerEncoding == indexesOfFieldsWithPrefixPointerEncoding[0] {
			prologuePointerEncoding = "\n\n	" + prefix + " struct {"
		}
		if indexOfTheFieldOfInterestPointerEncoding == indexesOfFieldsWithPrefixPointerEncoding[len(indexesOfFieldsWithPrefixPointerEncoding)-1] {
			epiloguePointerEncoding = "\n\n	} `gorm:\"embedded\"`"
		}
	default:
		if indexOfTheFieldOfInterestWOP == indexesOfFieldsWithPrefixWOP[0] {
			prologueWOP = "\n\n	" + prefix + " struct {"
		}
		if indexOfTheFieldOfInterestWOP == indexesOfFieldsWithPrefixWOP[len(indexesOfFieldsWithPrefixWOP)-1] {
			epilogueWOP = "\n\n	} `gorm:\"embedded\"`"
		}
	}

	// log.Println("ComputeFielProloguesEpilogues",
	// 	"\nField", field.GetName(), field.GetIndex(),
	// 	"\nDB", prologueDB, epilogueDB,
	// 	"\nPointerEncoding", prologuePointerEncoding, epiloguePointerEncoding,
	// 	"\nWOP", prologueWOP, epilogueWOP,
	// )

	return
}
