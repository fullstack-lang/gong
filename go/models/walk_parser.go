package models

import (
	"fmt"
	"go/ast"
	"log"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func WalkParser(parserPkgs map[string]*ast.Package, modelPkg *ModelPkg) {

	var pkg *ast.Package
	var ok bool
	if pkg, ok = parserPkgs["models"]; !ok {
		log.Fatal("No package models")
	}

	modelPkg.GongEnums = make(map[string]*GongEnum)
	modelPkg.GongStructs = make(map[string]*GongStruct)

	if len(pkg.Files) == 0 {
		log.Fatal("No go file to parse")
	}

	// first pass : get "type" definition for enum & struct
	//
	// search all files
	for filePath, file := range pkg.Files {

		var fileName string
		if strings.Contains(filePath, "\\") {
			fileNames := strings.Split(filePath, "\\")
			fileName = fileNames[len(fileNames)-1]

			// package name is three steps behind
			modelPkg.Name = fileNames[len(fileNames)-4]
			modelPkg.PkgPath =
				fileNames[len(fileNames)-7] + "/" +
					fileNames[len(fileNames)-6] + "/" +
					fileNames[len(fileNames)-5] + "/" +
					fileNames[len(fileNames)-4] + "/" +
					fileNames[len(fileNames)-3] + "/" +
					fileNames[len(fileNames)-2]
		}

		if fileName == "gong.go" {
			continue
		}

		// for exploration
		if fileName != "astruct.go" {
			continue
		}
		// if fileName != "aenum.go" {
		// 	continue
		// }

		log.Println("Parsing file ", fileName)

		// for structName, scope := range file.Scope.Objects {

		// 	log.Println("Parsing structName ", structName)
		// 	_ = scope
		// }

		// first pass for gathering the "type" definitions
		for _, decl := range file.Decls {
			_ = decl

			switch genDecl := decl.(type) {
			case *ast.GenDecl:
				// log.Println("Length of Spec", len(genDecl.Specs))

				for _, spec := range genDecl.Specs {

					switch typeSpec := spec.(type) {
					case *ast.TypeSpec:

						// TypeSpec is for declaration with the "type" keyword
						// in gong, it can be either a GongStruct or a GongEnum

						// we are only interested in exported symbols
						titler := cases.Title(language.English)

						if titler.String(typeSpec.Name.Name) != typeSpec.Name.Name {
							continue
						}

						log.Println("Type spec name is ", typeSpec.Name.Name)

						// If it is a GongEnum, the typeSpec Type is a int or a string
						switch _type := typeSpec.Type.(type) {
						case *ast.Ident:
							log.Println("Type spec type is ", _type.Name)

							var enumType GongEnumType
							switch _type.Name {
							case "int":
								enumType = Int
							case "string":
								enumType = String
							default:
								log.Fatal("Invalid definition of a Gongstruct ", typeSpec.Name.Name, " type ", _type.Name)
							}

							modelEnum := &GongEnum{
								Name: typeSpec.Name.Name,
								Type: enumType,
							}
							modelPkg.GongEnums[modelPkg.PkgPath+"."+typeSpec.Name.Name] = modelEnum

							// If it is a GongStruct, the ast is a StructType
						case *ast.StructType:

							// fetch the name of the Gongstruct by identifying if there is a field with name "Name"
							var isGongStruct bool
							for _, field := range _type.Fields.List {
								if len(field.Names) > 0 && field.Names[0].Name == "Name" {
									isGongStruct = true
								}
							}

							gongstruct := GongStruct{Name: typeSpec.Name.Name}
							if isGongStruct {
								modelPkg.GongStructs[modelPkg.PkgPath+"."+typeSpec.Name.Name] = &gongstruct
							}
						default:
						}
					default:
					}

				}

			default:
				continue
			}
		}

		// pass for gathering the "const" definitions of enums
		for _, decl := range file.Decls {
			_ = decl

			switch genDecl := decl.(type) {
			case *ast.GenDecl:

				var gongEnum *GongEnum

				for enumValue, spec := range genDecl.Specs {

					_ = spec
					switch valueSpec := spec.(type) {
					case *ast.ValueSpec:
						if len(valueSpec.Names) > 0 {
							if len(valueSpec.Names) > 1 {
								log.Fatal("To many Names for value spec ", valueSpec.Names[0].Name)
							}
						}

						// for an Enum, the type is only defined at the first element
						switch _type := valueSpec.Type.(type) {
						case *ast.Ident:
							gongEnum, ok = modelPkg.GongEnums[modelPkg.PkgPath+"."+_type.Name]
							if !ok {
								log.Fatalln("Unkown GongEnum Type")
							}
							log.Println("Const ", valueSpec.Names[0].Name, " of type ", gongEnum.Name)

							if gongEnum.Type == Int {
								gongEnumValue := (&GongEnumValue{
									Name:  valueSpec.Names[0].Name,
									Value: fmt.Sprintf("%d", enumValue)})
								gongEnum.GongEnumValues = append(gongEnum.GongEnumValues, gongEnumValue)
							} else { // string
								// the value is in the BasicLit expression
								if len(valueSpec.Values) != 1 {
									log.Fatal("Wrong def of const ", valueSpec.Names[0].Name)
								}
								_value := valueSpec.Values[0]

								switch value := _value.(type) {
								case *ast.BasicLit:
									gongEnumValue := (&GongEnumValue{
										Name:  valueSpec.Names[0].Name,
										Value: value.Value})
									gongEnum.GongEnumValues = append(gongEnum.GongEnumValues, gongEnumValue)
								default:
									log.Fatal("Wrong def of const ", valueSpec.Names[0].Name)
								}
							}
						default:
						}

						// if gongEnum is nil, this is because of a non Gongstruct const
						if valueSpec.Type == nil && gongEnum != nil {
							gongEnumValue := (&GongEnumValue{
								Name:  valueSpec.Names[0].Name,
								Value: fmt.Sprintf("%d", enumValue)})
							gongEnum.GongEnumValues = append(gongEnum.GongEnumValues, gongEnumValue)
						}
					default:
					}

				}

			default:
				continue
			}
		}
	}

}
