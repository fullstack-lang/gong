package models

import (
	"go/ast"
	"go/types"
	"log"
)

// GenerateFieldParser generates Gongfields of owningGongstruct
//
// by using the map_Structname_fieldList for embedded struct fields
// and modelPkg for access existing gongstructs and gongenums
func GenerateFieldParser(fieldList *[]*ast.Field, owningGongstruct *GongStruct,
	map_Structname_fieldList *map[string]*[]*ast.Field,
	modelPkg *ModelPkg) {

	for _, field := range *fieldList {

		if len(field.Names) > 1 {
			log.Fatal("too many names for field", field.Names[0].Name)
		}

		if len(field.Names) == 0 {
			// This is the case for struct embeding
			switch embedType := field.Type.(type) {
			case *ast.Ident:
				log.Println("processing embedded struct ", embedType.Name)
				GenerateFieldParser((*map_Structname_fieldList)[embedType.Name], owningGongstruct, map_Structname_fieldList, modelPkg)
			default:
			}
			continue
		}

		fieldName := field.Names[0].Name

		switch __fieldType := field.Type.(type) {
		case *ast.Ident:
			switch __fieldType.Name {
			case "string":
				gongField :=
					&GongBasicField{
						Name: fieldName,

						// this field is only used for code generation
						Type:          &types.Basic{},
						basicKind:     types.String,
						BasicKindName: "string",
						Index:         len(owningGongstruct.Fields),
					}
				owningGongstruct.Fields = append(owningGongstruct.Fields, gongField)

			case "int":
				gongField :=
					&GongBasicField{
						Name:          fieldName,
						Type:          &types.Basic{},
						basicKind:     types.Int,
						BasicKindName: "int",
						Index:         len(owningGongstruct.Fields),
					}
				owningGongstruct.Fields = append(owningGongstruct.Fields, gongField)
			case "float64":
				gongField :=
					&GongBasicField{
						Name:          fieldName,
						Type:          &types.Basic{},
						basicKind:     types.Float64,
						BasicKindName: "float64",
						Index:         len(owningGongstruct.Fields),
					}
				owningGongstruct.Fields = append(owningGongstruct.Fields, gongField)
			case "bool":
				gongField :=
					&GongBasicField{
						Name:          fieldName,
						Type:          &types.Basic{},
						basicKind:     types.Bool,
						BasicKindName: "bool",
						Index:         len(owningGongstruct.Fields),
					}
				owningGongstruct.Fields = append(owningGongstruct.Fields, gongField)
			default:
				// Check if type is an enum
				if gongEnum, ok := modelPkg.GongEnums[modelPkg.PkgPath+"."+__fieldType.Name]; ok {
					if gongEnum.Type == Int {
						gongField :=
							&GongBasicField{
								Name:          fieldName,
								Type:          &types.Basic{},
								basicKind:     types.Int,
								BasicKindName: "int",
								GongEnum:      gongEnum,
								DeclaredType:  __fieldType.Name,
								Index:         len(owningGongstruct.Fields),
							}
						owningGongstruct.Fields = append(owningGongstruct.Fields, gongField)
					}
					if gongEnum.Type == String {
						gongField :=
							&GongBasicField{
								Name:          fieldName,
								Type:          &types.Basic{},
								basicKind:     types.String,
								BasicKindName: "string",
								GongEnum:      gongEnum,
								DeclaredType:  __fieldType.Name,
								Index:         len(owningGongstruct.Fields),
							}
						owningGongstruct.Fields = append(owningGongstruct.Fields, gongField)
					}

				} else {
					log.Println("Cannot parse field of type ", __fieldType.Name)
				}

			}
		case *ast.SelectorExpr:
			// tries to match "time.Time" // "time.Duration"
			switch __selX := __fieldType.X.(type) {
			case *ast.Ident:
				if __selX.Name == "time" {
					switch __fieldType.Sel.Name {
					case "Time":
						gongField :=
							&GongTimeField{
								Name:  fieldName,
								Index: len(owningGongstruct.Fields),
							}
						owningGongstruct.Fields = append(owningGongstruct.Fields, gongField)
					case "Duration":
						gongField :=
							&GongBasicField{
								Name:          fieldName,
								Type:          &types.Basic{},
								basicKind:     types.Int,
								BasicKindName: "int",
								GongEnum:      nil,
								DeclaredType:  "time.Duration",
								Index:         len(owningGongstruct.Fields),
							}
						owningGongstruct.Fields = append(owningGongstruct.Fields, gongField)
					}
				}
			}

		case *ast.StarExpr:
			// for Pointer to struct
			switch X := __fieldType.X.(type) {
			case *ast.Ident:
				if targetGongstruct, ok := modelPkg.GongStructs[modelPkg.PkgPath+"."+X.Name]; ok {
					gongField :=
						&PointerToGongStructField{
							Name:       fieldName,
							GongStruct: targetGongstruct,
							Index:      len(owningGongstruct.Fields),
						}
					owningGongstruct.Fields = append(owningGongstruct.Fields, gongField)
				}
			}
		case *ast.ArrayType:
			// for slice of pointers to struct
			switch elt := __fieldType.Elt.(type) {
			case *ast.StarExpr:
				switch X := elt.X.(type) {
				case *ast.Ident:
					if targetGongstruct, ok := modelPkg.GongStructs[modelPkg.PkgPath+"."+X.Name]; ok {
						gongField :=
							&SliceOfPointerToGongStructField{
								Name:       fieldName,
								GongStruct: targetGongstruct,
								Index:      len(owningGongstruct.Fields),
							}
						owningGongstruct.Fields = append(owningGongstruct.Fields, gongField)
					}
				}
			}
		default:
			log.Println("Cannot parse field named ", fieldName)
		}
	}

}
