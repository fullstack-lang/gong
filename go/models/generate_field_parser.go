package models

import (
	"go/ast"
	"go/types"
	"log"
	"strings"
)

// GenerateFieldParser generates Gongfields of owningGongstruct
//
// by using the map_Structname_fieldList for embedded struct fields
// and modelPkg for access existing gongstructs and gongenums
func GenerateFieldParser(
	fieldList *[]*ast.Field,
	owningGongstruct *GongStruct,
	map_Structname_fieldList *map[string]*[]*ast.Field,
	modelPkg *ModelPkg,
	compositeTypeStructName string, // when it is composed
	prefix string, // used in anonymous struct
) {

	for _, field := range *fieldList {

		// get the comment group and check wether it is "swagger:ignore" or "gong:ignore"
		var isIgnoredField bool
		var isTextArea bool
		var isBespokeWidth bool
		var bespokeWidth int
		var isBespokeHeight bool
		var bespokeHeight int
		if field.Comment != nil {
			for _, comment := range field.Comment.List {
				if strings.Contains(comment.Text, "swagger:ignore") || strings.Contains(comment.Text, "gong:ignore") {
					isIgnoredField = true
				}
				if strings.Contains(comment.Text, "gong:text") {
					isTextArea = true
				}
				if strings.Contains(comment.Text, "gong:width") {
					width, err := extractWidthNumber(comment.Text)
					if err == nil {
						isBespokeWidth = true
						bespokeWidth = width
					}
				}
				if strings.Contains(comment.Text, "gong:height") {
					height, err := extractHeightNumber(comment.Text)
					if err == nil {
						isBespokeHeight = true
						bespokeHeight = height
					}
				}
			}
		}
		if field.Doc != nil {
			for _, comment := range field.Doc.List {
				if strings.Contains(comment.Text, "swagger:ignore") || strings.Contains(comment.Text, "gong:ignore") {
					isIgnoredField = true
				}
				if strings.Contains(comment.Text, "gong:text") {
					isTextArea = true
				}
				if strings.Contains(comment.Text, "gong:width") {
					width, err := extractWidthNumber(comment.Text)
					if err == nil {
						isBespokeWidth = true
						bespokeWidth = width
					}
				}
				if strings.Contains(comment.Text, "gong:height") {
					height, err := extractHeightNumber(comment.Text)
					if err == nil {
						isBespokeHeight = true
						bespokeHeight = height
					}
				}
			}
		}
		if isIgnoredField {
			continue
		}

		if len(field.Names) == 0 {
			// This is the case for struct embeding
			switch embedType := field.Type.(type) {
			case *ast.Ident:
				// log.Println("processing embedded struct ", embedType.Name)
				if _fieldList, ok := (*map_Structname_fieldList)[embedType.Name]; ok {
					GenerateFieldParser(_fieldList,
						owningGongstruct,
						map_Structname_fieldList,
						modelPkg,
						embedType.Name,
						"")
				} else {
					log.Fatalln("Unknown embedded type", embedType.Name)
				}

			default:
			}
			continue
		}

		// parses all field names
		// This can be the case with
		// type Hello struct {
		//    X, Y int
		// }
		for _, fieldNameIdent := range field.Names {
			fieldName := fieldNameIdent.Name
			if prefix != "" {
				fieldName = prefix + "." + fieldName
			}
			switch __fieldType := field.Type.(type) {
			case *ast.Ident:
				switch __fieldType.Name {
				case "string":
					gongField :=
						&GongBasicField{
							Name:                fieldName,
							basicKind:           types.String,
							BasicKindName:       "string",
							DeclaredType:        "string",
							Index:               len(owningGongstruct.Fields),
							CompositeStructName: compositeTypeStructName,
							IsTextArea:          isTextArea,
							IsBespokeWidth:      isBespokeWidth,
							BespokeWidth:        bespokeWidth,
							IsBespokeHeight:     isBespokeHeight,
							BespokeHeight:       bespokeHeight,
						}

					if field.Doc != nil {
						for _, comment := range field.Doc.List {
							text := comment.Text
							if strings.HasPrefix(text, "//gong:ident") {
								gongField.IsDocLink = true
							}
						}
					}

					owningGongstruct.Fields = append(owningGongstruct.Fields, gongField)

				case "int":
					gongField :=
						&GongBasicField{
							Name:                fieldName,
							basicKind:           types.Int,
							BasicKindName:       "int",
							DeclaredType:        "int",
							Index:               len(owningGongstruct.Fields),
							CompositeStructName: compositeTypeStructName,
						}
					owningGongstruct.Fields = append(owningGongstruct.Fields, gongField)
				case "float64":
					gongField :=
						&GongBasicField{
							Name:                fieldName,
							basicKind:           types.Float64,
							BasicKindName:       "float64",
							DeclaredType:        "float64",
							Index:               len(owningGongstruct.Fields),
							CompositeStructName: compositeTypeStructName,
						}
					owningGongstruct.Fields = append(owningGongstruct.Fields, gongField)
				case "bool":
					gongField :=
						&GongBasicField{
							Name:                fieldName,
							basicKind:           types.Bool,
							BasicKindName:       "bool",
							DeclaredType:        "bool",
							Index:               len(owningGongstruct.Fields),
							CompositeStructName: compositeTypeStructName,
						}
					owningGongstruct.Fields = append(owningGongstruct.Fields, gongField)
				default:
					// Check if type is an enum
					if gongEnum, ok := modelPkg.GongEnums[modelPkg.PkgPath+"."+__fieldType.Name]; ok {
						if gongEnum.Type == Int {
							gongField :=
								&GongBasicField{
									Name:                fieldName,
									basicKind:           types.Int,
									BasicKindName:       "int",
									GongEnum:            gongEnum,
									DeclaredType:        modelPkg.PkgPath + "." + __fieldType.Name,
									Index:               len(owningGongstruct.Fields),
									CompositeStructName: compositeTypeStructName,
								}
							owningGongstruct.Fields = append(owningGongstruct.Fields, gongField)
						}
						if gongEnum.Type == String {
							gongField :=
								&GongBasicField{
									Name:                fieldName,
									basicKind:           types.String,
									BasicKindName:       "string",
									GongEnum:            gongEnum,
									DeclaredType:        modelPkg.PkgPath + "." + __fieldType.Name,
									Index:               len(owningGongstruct.Fields),
									CompositeStructName: compositeTypeStructName,
								}
							owningGongstruct.Fields = append(owningGongstruct.Fields, gongField)
						}

					} else {
						// log.Println("Cannot parse field of type ", __fieldType.Name)
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
									Name:                fieldName,
									Index:               len(owningGongstruct.Fields),
									CompositeStructName: compositeTypeStructName,
								}
							owningGongstruct.Fields = append(owningGongstruct.Fields, gongField)
						case "Duration":
							gongField :=
								&GongBasicField{
									Name:                fieldName,
									basicKind:           types.Int,
									BasicKindName:       "int",
									GongEnum:            nil,
									DeclaredType:        "time.Duration",
									Index:               len(owningGongstruct.Fields),
									CompositeStructName: compositeTypeStructName,
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
								Name:                fieldName,
								GongStruct:          targetGongstruct,
								Index:               len(owningGongstruct.Fields),
								CompositeStructName: compositeTypeStructName,
							}

						if field.Doc != nil {
							for _, comment := range field.Doc.List {
								text := comment.Text
								if strings.HasPrefix(text, "//gong:type") {
									gongField.IsType = true
								}
							}
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
									Name:                fieldName,
									GongStruct:          targetGongstruct,
									Index:               len(owningGongstruct.Fields),
									CompositeStructName: compositeTypeStructName,
								}
							owningGongstruct.Fields = append(owningGongstruct.Fields, gongField)
						}
					}
				}
			case *ast.StructType:
				// we are in an anonymous struct
				// log.Println("Anonymous struct", fieldName)
				list := __fieldType.Fields.List
				GenerateFieldParser(&list,
					owningGongstruct,
					map_Structname_fieldList,
					modelPkg,
					"",
					fieldName)

			default:
				// log.Println("Field ", fieldName, " of struct ", owningGongstruct.Name, " is not a gong type")
			}
		}
	}

}
