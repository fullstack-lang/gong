package models

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var dummy_strconv_import strconv.NumError

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFile(pathToFile string) {

	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		log.Panic("Path does not exist %s ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	startParser := time.Now()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		log.Panic("Unable to parser ", errParser.Error())
	}

	// astCoordinate := "File "
	// log.Println(// astCoordinate)
	for _, decl := range inFile.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			funcDecl := decl
			// astCoordinate := // astCoordinate + "\tFunction " + funcDecl.Name.Name
			if name := funcDecl.Name; name != nil {
				isOfInterest := strings.Contains(funcDecl.Name.Name, "Injection")
				if !isOfInterest {
					continue
				}
				// log.Println(// astCoordinate)
			}
			if doc := funcDecl.Doc; doc != nil {
				// astCoordinate := // astCoordinate + "\tDoc"
				for _, comment := range doc.List {
					_ = comment
					// astCoordinate := // astCoordinate + "\tComment: " + comment.Text
					// log.Println(// astCoordinate)
				}
			}
			if body := funcDecl.Body; body != nil {
				// astCoordinate := // astCoordinate + "\tBody: "
				for _, stmt := range body.List {
					switch stmt := stmt.(type) {
					case *ast.ExprStmt:
						exprStmt := stmt
						// astCoordinate := // astCoordinate + "\tExprStmt: "
						switch expr := exprStmt.X.(type) {
						case *ast.CallExpr:
							// astCoordinate := // astCoordinate + "\tCallExpr: "
							callExpr := expr
							switch fun := callExpr.Fun.(type) {
							case *ast.Ident:
								ident := fun
								_ = ident
								// astCoordinate := // astCoordinate + "\tIdent: " + ident.Name
								// log.Println(// astCoordinate)
							}
						}
					case *ast.AssignStmt:
						astCoordinate := "\tAssignStmt: "
						// log.Println(// astCoordinate)
						assignStmt := stmt
						instance, id, gongstruct, fieldName := UnmarshallGongstructStaging(assignStmt, astCoordinate)
						_ = instance
						_ = id
						_ = gongstruct
						_ = fieldName
					}
				}
			}
		case *ast.GenDecl:
			genDecl := decl
			// log.Println("\tAST GenDecl: ")
			if doc := genDecl.Doc; doc != nil {
				for _, comment := range doc.List {
					_ = comment
					// log.Println("\tAST Comment: ", comment.Text)
				}
			}
			for _, spec := range genDecl.Specs {
				switch spec := spec.(type) {
				case *ast.ImportSpec:
					importSpec := spec
					if path := importSpec.Path; path != nil {
						// log.Println("\t\tAST Path: ", path.Value)
					}
				}
			}
		}

	}
}

var __gong__map_Indentifiers_gongstructName = make(map[string]string)

// insertion point for identifiers maps
var __gong__map_GongBasicField = make(map[string]*GongBasicField)
var __gong__map_GongEnum = make(map[string]*GongEnum)
var __gong__map_GongEnumValue = make(map[string]*GongEnumValue)
var __gong__map_GongNote = make(map[string]*GongNote)
var __gong__map_GongStruct = make(map[string]*GongStruct)
var __gong__map_GongTimeField = make(map[string]*GongTimeField)
var __gong__map_ModelPkg = make(map[string]*ModelPkg)
var __gong__map_PointerToGongStructField = make(map[string]*PointerToGongStructField)
var __gong__map_SliceOfPointerToGongStructField = make(map[string]*SliceOfPointerToGongStructField)

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging(assignStmt *ast.AssignStmt, astCoordinate_ string) (
	instance any,
	identifier string,
	gongstructName string,
	fieldName string) {
	astCoordinate := "\tAssignStmt: "
	for rank, expr := range assignStmt.Lhs {
		if rank > 0 {
			continue
		}
		switch expr := expr.(type) {
		case *ast.Ident:
			// we are on a variable declaration
			ident := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + ident.Name
			// log.Println(astCoordinate)
			identifier = ident.Name
		case *ast.SelectorExpr:
			// we are on a variable assignement
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + selectorExpr.X.(*ast.Ident).Name + "." + selectorExpr.Sel.Name
			// log.Println(astCoordinate)
			identifier = selectorExpr.X.(*ast.Ident).Name
			fieldName = selectorExpr.Sel.Name
		}
	}
	for _, expr := range assignStmt.Rhs {
		// astCoordinate := astCoordinate + "\tRhs"
		switch expr := expr.(type) {
		case *ast.CallExpr:
			callExpr := expr
			// astCoordinate := astCoordinate + "\tFun"
			switch fun := callExpr.Fun.(type) {
			// the is Fun      Expr
			// function expression xxx.Stage()
			case *ast.SelectorExpr:
				selectorExpr := fun
				// astCoordinate := astCoordinate + "\tSelectorExpr"
				switch x := selectorExpr.X.(type) {
				case *ast.ParenExpr:
					// A ParenExpr node represents a parenthesized expression.
					// the is the
					//   { Name : "A1"}
					// astCoordinate := astCoordinate + "\tX"
					parenExpr := x
					switch x := parenExpr.X.(type) {
					case *ast.UnaryExpr:
						unaryExpr := x
						// astCoordinate := astCoordinate + "\tUnaryExpr"
						switch x := unaryExpr.X.(type) {
						case *ast.CompositeLit:
							instanceName := "NoName yet"
							compositeLit := x
							// astCoordinate := astCoordinate + "\tX(CompositeLit)"
							for _, elt := range compositeLit.Elts {
								// astCoordinate := astCoordinate + "\tElt"
								switch elt := elt.(type) {
								case *ast.KeyValueExpr:
									// This is expression
									//     Name: "A1"
									keyValueExpr := elt
									// astCoordinate := astCoordinate + "\tKeyValueExpr"
									switch key := keyValueExpr.Key.(type) {
									case *ast.Ident:
										ident := key
										_ = ident
										// astCoordinate := astCoordinate + "\tKey" + "." + ident.Name
										// log.Println(astCoordinate)
									}
									switch value := keyValueExpr.Value.(type) {
									case *ast.BasicLit:
										basicLit := value
										// astCoordinate := astCoordinate + "\tBasicLit Value" + "." + basicLit.Value
										// log.Println(astCoordinate)
										instanceName = basicLit.Value

										// remove first and last char
										instanceName = instanceName[1 : len(instanceName)-1]
									}
								}
							}
							astCoordinate2 := astCoordinate + "\tType"
							_ = astCoordinate2
							switch type_ := compositeLit.Type.(type) {
							case *ast.SelectorExpr:
								slcExpr := type_
								// astCoordinate := astCoordinate2 + "\tSelectorExpr"
								switch X := slcExpr.X.(type) {
								case *ast.Ident:
									ident := X
									_ = ident
									// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
									// log.Println(astCoordinate)
								}
								if Sel := slcExpr.Sel; Sel != nil {
									// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
									// log.Println(astCoordinate)

									gongstructName = Sel.Name
									// this is the place where an instance is created
									switch gongstructName {
									// insertion point for identifiers
									case "GongBasicField":
										instanceGongBasicField := (&GongBasicField{Name: instanceName}).Stage()
										instance = any(instanceGongBasicField)
										__gong__map_GongBasicField[identifier] = instanceGongBasicField
									case "GongEnum":
										instanceGongEnum := (&GongEnum{Name: instanceName}).Stage()
										instance = any(instanceGongEnum)
										__gong__map_GongEnum[identifier] = instanceGongEnum
									case "GongEnumValue":
										instanceGongEnumValue := (&GongEnumValue{Name: instanceName}).Stage()
										instance = any(instanceGongEnumValue)
										__gong__map_GongEnumValue[identifier] = instanceGongEnumValue
									case "GongNote":
										instanceGongNote := (&GongNote{Name: instanceName}).Stage()
										instance = any(instanceGongNote)
										__gong__map_GongNote[identifier] = instanceGongNote
									case "GongStruct":
										instanceGongStruct := (&GongStruct{Name: instanceName}).Stage()
										instance = any(instanceGongStruct)
										__gong__map_GongStruct[identifier] = instanceGongStruct
									case "GongTimeField":
										instanceGongTimeField := (&GongTimeField{Name: instanceName}).Stage()
										instance = any(instanceGongTimeField)
										__gong__map_GongTimeField[identifier] = instanceGongTimeField
									case "ModelPkg":
										instanceModelPkg := (&ModelPkg{Name: instanceName}).Stage()
										instance = any(instanceModelPkg)
										__gong__map_ModelPkg[identifier] = instanceModelPkg
									case "PointerToGongStructField":
										instancePointerToGongStructField := (&PointerToGongStructField{Name: instanceName}).Stage()
										instance = any(instancePointerToGongStructField)
										__gong__map_PointerToGongStructField[identifier] = instancePointerToGongStructField
									case "SliceOfPointerToGongStructField":
										instanceSliceOfPointerToGongStructField := (&SliceOfPointerToGongStructField{Name: instanceName}).Stage()
										instance = any(instanceSliceOfPointerToGongStructField)
										__gong__map_SliceOfPointerToGongStructField[identifier] = instanceSliceOfPointerToGongStructField
									}
									__gong__map_Indentifiers_gongstructName[identifier] = gongstructName
									return
								}
							}
						}
					}
				}
				if sel := selectorExpr.Sel; sel != nil {
					// astCoordinate := astCoordinate + "\tSel" + "." + sel.Name
					// log.Println(astCoordinate)
				}
				for iteration, arg := range callExpr.Args {
					// astCoordinate := astCoordinate + "\tArg"
					switch arg := arg.(type) {
					case *ast.BasicLit:
						basicLit := arg
						// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
						// log.Println(astCoordinate)

						// first iteration should be ignored
						if iteration == 0 {
							continue
						}

						// remove first and last char
						date := basicLit.Value[1 : len(basicLit.Value)-1]
						_ = date

						var ok bool
						gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
						if !ok {
							log.Fatalln("gongstructName not found for identifier", identifier)
						}
						switch gongstructName {
						// insertion point for basic lit assignments
						case "GongBasicField":
							switch fieldName {
							// insertion point for date assign code
							}
						case "GongEnum":
							switch fieldName {
							// insertion point for date assign code
							}
						case "GongEnumValue":
							switch fieldName {
							// insertion point for date assign code
							}
						case "GongNote":
							switch fieldName {
							// insertion point for date assign code
							}
						case "GongStruct":
							switch fieldName {
							// insertion point for date assign code
							}
						case "GongTimeField":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ModelPkg":
							switch fieldName {
							// insertion point for date assign code
							}
						case "PointerToGongStructField":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SliceOfPointerToGongStructField":
							switch fieldName {
							// insertion point for date assign code
							}
						}
					}
				}
			case *ast.Ident:
				// append function
				ident := fun
				_ = ident
				// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			for _, arg := range callExpr.Args {
				// astCoordinate := astCoordinate + "\tArg"
				switch arg := arg.(type) {
				case *ast.Ident:
					ident := arg
					_ = ident
					// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
					// log.Println(astCoordinate)
					var ok bool
					gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
					if !ok {
						log.Fatalln("gongstructName not found for identifier", identifier)
					}
					switch gongstructName {
					// insertion point for slice of pointers assignments
					case "GongBasicField":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "GongEnum":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "GongEnumValues":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_GongEnumValue[targetIdentifier]
							__gong__map_GongEnum[identifier].GongEnumValues =
								append(__gong__map_GongEnum[identifier].GongEnumValues, target)
						}
					case "GongEnumValue":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "GongNote":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "GongStruct":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "GongBasicFields":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_GongBasicField[targetIdentifier]
							__gong__map_GongStruct[identifier].GongBasicFields =
								append(__gong__map_GongStruct[identifier].GongBasicFields, target)
						case "GongTimeFields":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_GongTimeField[targetIdentifier]
							__gong__map_GongStruct[identifier].GongTimeFields =
								append(__gong__map_GongStruct[identifier].GongTimeFields, target)
						case "PointerToGongStructFields":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_PointerToGongStructField[targetIdentifier]
							__gong__map_GongStruct[identifier].PointerToGongStructFields =
								append(__gong__map_GongStruct[identifier].PointerToGongStructFields, target)
						case "SliceOfPointerToGongStructFields":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_SliceOfPointerToGongStructField[targetIdentifier]
							__gong__map_GongStruct[identifier].SliceOfPointerToGongStructFields =
								append(__gong__map_GongStruct[identifier].SliceOfPointerToGongStructFields, target)
						}
					case "GongTimeField":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ModelPkg":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "PointerToGongStructField":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SliceOfPointerToGongStructField":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					}
				case *ast.SelectorExpr:
					slcExpr := arg
					// astCoordinate := astCoordinate + "\tSelectorExpr"
					switch X := slcExpr.X.(type) {
					case *ast.Ident:
						ident := X
						_ = ident
						// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
						// log.Println(astCoordinate)

					}
					if Sel := slcExpr.Sel; Sel != nil {
						// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
						// log.Println(astCoordinate)
					}
				}
			}
		case *ast.BasicLit:
			// assignment to string field
			basicLit := expr
			// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}

			switch gongstructName {
			// insertion point for basic lit assignments
			case "GongBasicField":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongBasicField[identifier].Name = fielValue
				case "BasicKindName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongBasicField[identifier].BasicKindName = fielValue
				case "DeclaredType":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongBasicField[identifier].DeclaredType = fielValue
				case "CompositeStructName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongBasicField[identifier].CompositeStructName = fielValue
				case "Index":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_GongBasicField[identifier].Index = int(fielValue)
				}
			case "GongEnum":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongEnum[identifier].Name = fielValue
				}
			case "GongEnumValue":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongEnumValue[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongEnumValue[identifier].Value = fielValue
				}
			case "GongNote":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongNote[identifier].Name = fielValue
				case "Body":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongNote[identifier].Body = fielValue
				}
			case "GongStruct":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongStruct[identifier].Name = fielValue
				}
			case "GongTimeField":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongTimeField[identifier].Name = fielValue
				case "Index":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_GongTimeField[identifier].Index = int(fielValue)
				case "CompositeStructName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongTimeField[identifier].CompositeStructName = fielValue
				}
			case "ModelPkg":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ModelPkg[identifier].Name = fielValue
				case "PkgPath":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ModelPkg[identifier].PkgPath = fielValue
				}
			case "PointerToGongStructField":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_PointerToGongStructField[identifier].Name = fielValue
				case "Index":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_PointerToGongStructField[identifier].Index = int(fielValue)
				case "CompositeStructName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_PointerToGongStructField[identifier].CompositeStructName = fielValue
				}
			case "SliceOfPointerToGongStructField":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SliceOfPointerToGongStructField[identifier].Name = fielValue
				case "Index":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SliceOfPointerToGongStructField[identifier].Index = int(fielValue)
				case "CompositeStructName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SliceOfPointerToGongStructField[identifier].CompositeStructName = fielValue
				}
			}
		case *ast.Ident:
			// assignment to boolean field ?
			ident := expr
			_ = ident
			// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}
			switch gongstructName {
			// insertion point for bool & pointers assignments
			case "GongBasicField":
				switch fieldName {
				// insertion point for field dependant code
				case "GongEnum":
					targetIdentifier := ident.Name
					__gong__map_GongBasicField[identifier].GongEnum = __gong__map_GongEnum[targetIdentifier]
				}
			case "GongEnum":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "GongEnumValue":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "GongNote":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "GongStruct":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "GongTimeField":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "ModelPkg":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "PointerToGongStructField":
				switch fieldName {
				// insertion point for field dependant code
				case "GongStruct":
					targetIdentifier := ident.Name
					__gong__map_PointerToGongStructField[identifier].GongStruct = __gong__map_GongStruct[targetIdentifier]
				}
			case "SliceOfPointerToGongStructField":
				switch fieldName {
				// insertion point for field dependant code
				case "GongStruct":
					targetIdentifier := ident.Name
					__gong__map_SliceOfPointerToGongStructField[identifier].GongStruct = __gong__map_GongStruct[targetIdentifier]
				}
			}
		case *ast.SelectorExpr:
			// assignment to enum field
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tSelectorExpr"
			switch X := selectorExpr.X.(type) {
			case *ast.Ident:
				ident := X
				_ = ident
				// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			if Sel := selectorExpr.Sel; Sel != nil {
				// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
				// log.Println(astCoordinate)

				// enum field
				var ok bool
				gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
				if !ok {
					log.Fatalln("gongstructName not found for identifier", identifier)
				}

				// remove first and last char
				enumValue := Sel.Name
				_ = enumValue
				switch gongstructName {
				// insertion point for enums assignments
				case "GongBasicField":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "GongEnum":
					switch fieldName {
					// insertion point for enum assign code
					case "Type":
						var val GongEnumType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_GongEnum[identifier].Type = GongEnumType(val)
					}
				case "GongEnumValue":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "GongNote":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "GongStruct":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "GongTimeField":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ModelPkg":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "PointerToGongStructField":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SliceOfPointerToGongStructField":
					switch fieldName {
					// insertion point for enum assign code
					}
				}
			}
		}
	}
	return
}
