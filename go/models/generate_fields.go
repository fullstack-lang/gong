package models

import (
	"go/types"
	"log"
	"strings"
)

//
// GenerateFields appends to modelPkg the fields of the __struct
//
// isCompositeField if it is a field from a composition with another struct
// compositeTypeStructName is the name of the composite struct if it is the case
//
func GenerateFields(structName string, __struct *types.Struct,
	modelPkg *ModelPkg,
	isCompositeField bool,
	compositeTypeStructName string) {

	for fieldIndex := 0; fieldIndex < __struct.NumFields(); fieldIndex++ {
		// log.Printf("field #%d\n", fieldIndex)

		_typesField := __struct.Field(fieldIndex)
		// log.Printf("field name %s\n", _typesField.Name())
		// log.Printf("field type name %s\n", _typesField.Type().String())

		// only deal with exported fields
		if _typesField.Name() != strings.Title(_typesField.Name()) {
			continue
		}

		// check if it is an Enum
		fieldType := _typesField.Type().String()
		gongEnum := modelPkg.GongEnums[fieldType]

		// fetch the field
		// we have to process pointers to spinosa type & array of of pointers
		fieldUnderlying := _typesField.Type().Underlying()
		switch t2 := fieldUnderlying.(type) {
		case *types.Basic:
			var kindAsString string
			// log.Printf("field is a basic pointer of type %s\n", t2.Underlying().String())
			switch t2.Kind() {
			case types.Bool:
				kindAsString = "bool"
			case types.Int64:
				kindAsString = "int64"
			case types.Int:
				kindAsString = "int"
			case types.Uint:
				kindAsString = "uint"
			case types.String:
				kindAsString = "string"
			case types.Float64:
				kindAsString = "float64"
			default:
				log.Panic("not suitable kind for gong")
			}
			modelPkg.GongStructs[structName].Fields =
				append(modelPkg.GongStructs[structName].Fields, //
					&GongBasicField{
						Name:          _typesField.Name(),
						basicKind:     t2.Kind(),
						BasicKindName: kindAsString,
						GongEnum:      gongEnum,
						DeclaredType:  fieldType,
						Index:         len(modelPkg.GongStructs[structName].Fields),
					})
		case *types.Pointer:
			// if pointer, field should be of form *Struct

			// or a pointer on a basic type
			typesBehindPointer := t2.Elem().String()

			if typesBehindPointer == "bool" {
				log.Fatal("Field is pointer to bool")
			}
			// log.Printf("field is a pointer of type %s\n", typesBehindPointer)

			__struct, ok := modelPkg.GongStructs[t2.Elem().String()]

			// if type of field is unkown, do not take into account
			if !ok {
				continue
			}

			modelPkg.GongStructs[structName].Fields =
				append(modelPkg.GongStructs[structName].Fields, //
					&PointerToGongStructField{
						Name:                _typesField.Name(),
						GongStruct:          __struct,
						Index:               len(modelPkg.GongStructs[structName].Fields),
						CompositeStructName: compositeTypeStructName,
					},
				)

		case *types.Slice:
			// if slice, field should be of form []*TypeSpinosa

			// first, fetch type of array, should be *TypeSpinosa
			assocPtr := t2.Elem()

			// then, fetch the type after the pointer, should be TypeSpinosa
			switch t3 := assocPtr.(type) {
			case *types.Pointer:
				assocLongName := t3.Elem().String()
				// log.Printf("field is a slice of type %s\n", assocLongName)
				modelPkg.GongStructs[structName].Fields =
					append(modelPkg.GongStructs[structName].Fields, //
						&SliceOfPointerToGongStructField{
							Name:                _typesField.Name(),
							GongStruct:          modelPkg.GongStructs[assocLongName],
							Index:               len(modelPkg.GongStructs[structName].Fields),
							CompositeStructName: compositeTypeStructName,
						},
					)

			default:
				// log.Printf("unkown type in slice" + assocPtr.String())
			}

		//
		// composition of another struct
		//
		case *types.Struct:
			compositeTypeStruct := fieldUnderlying.(*types.Struct)

			compositeTypeStructName := _typesField.Type().String()
			_ = compositeTypeStructName
			// log.Printf("field is a composite struct of type %s\n", compositeTypeStructName)

			if _typesField.Type().String() == "time.Time" {
				modelPkg.GongStructs[structName].Fields =
					append(modelPkg.GongStructs[structName].Fields, //
						&GongTimeField{
							Name:  _typesField.Name(),
							Index: len(modelPkg.GongStructs[structName].Fields),
						},
					)
			} else {
				localIdentifiers := strings.Split(compositeTypeStructName, ".")
				GenerateFields(structName, compositeTypeStruct, modelPkg, true, localIdentifiers[len(localIdentifiers)-1])
			}

		default:
		}
	}
}
