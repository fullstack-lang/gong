package golang

const GongAstTemplate = `package models

import (
	"go/ast"
	"log"
	"strconv"
	"time"
)

var __gong__map_Indentifiers_gongstructName = make(map[string]string)

// insertion point for identifiers maps{{` + string(rune(ModelGongAstGenericMaps)) + `}}

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
			astCoordinate := astCoordinate + "\tLhs" + "." + ident.Name
			log.Println(astCoordinate)
			identifier = ident.Name
		case *ast.SelectorExpr:
			// we are on a variable assignement
			selectorExpr := expr
			astCoordinate := astCoordinate + "\tLhs" + "." + selectorExpr.X.(*ast.Ident).Name + "." + selectorExpr.Sel.Name
			log.Println(astCoordinate)
			identifier = selectorExpr.X.(*ast.Ident).Name
			fieldName = selectorExpr.Sel.Name
		}
	}
	for _, expr := range assignStmt.Rhs {
		astCoordinate := astCoordinate + "\tRhs"
		switch expr := expr.(type) {
		case *ast.CallExpr:
			callExpr := expr
			astCoordinate := astCoordinate + "\tFun"
			switch fun := callExpr.Fun.(type) {
			// the is Fun      Expr
			// function expression xxx.Stage()
			case *ast.SelectorExpr:
				selectorExpr := fun
				astCoordinate := astCoordinate + "\tSelectorExpr"
				switch x := selectorExpr.X.(type) {
				case *ast.ParenExpr:
					// A ParenExpr node represents a parenthesized expression.
					// the is the
					//   { Name : "A1"}
					astCoordinate := astCoordinate + "\tX"
					parenExpr := x
					switch x := parenExpr.X.(type) {
					case *ast.UnaryExpr:
						unaryExpr := x
						astCoordinate := astCoordinate + "\tUnaryExpr"
						switch x := unaryExpr.X.(type) {
						case *ast.CompositeLit:
							instanceName := "NoName yet"
							compositeLit := x
							astCoordinate := astCoordinate + "\tX(CompositeLit)"
							for _, elt := range compositeLit.Elts {
								astCoordinate := astCoordinate + "\tElt"
								switch elt := elt.(type) {
								case *ast.KeyValueExpr:
									// This is expression
									//     Name: "A1"
									keyValueExpr := elt
									astCoordinate := astCoordinate + "\tKeyValueExpr"
									switch key := keyValueExpr.Key.(type) {
									case *ast.Ident:
										ident := key
										astCoordinate := astCoordinate + "\tKey" + "." + ident.Name
										log.Println(astCoordinate)
									}
									switch value := keyValueExpr.Value.(type) {
									case *ast.BasicLit:
										basicLit := value
										astCoordinate := astCoordinate + "\tBasicLit Value" + "." + basicLit.Value
										log.Println(astCoordinate)
										instanceName = basicLit.Value

										// remove first and last char
										instanceName = instanceName[1 : len(instanceName)-1]
									}
								}
							}
							astCoordinate2 := astCoordinate + "\tType"
							switch type_ := compositeLit.Type.(type) {
							case *ast.SelectorExpr:
								slcExpr := type_
								astCoordinate := astCoordinate2 + "\tSelectorExpr"
								switch X := slcExpr.X.(type) {
								case *ast.Ident:
									ident := X
									astCoordinate := astCoordinate + "\tX" + "." + ident.Name
									log.Println(astCoordinate)
								}
								if Sel := slcExpr.Sel; Sel != nil {
									astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
									log.Println(astCoordinate)

									gongstructName = Sel.Name
									// this is the place where an instance is created
									switch gongstructName {
									// insertion point for identifiers {{` + string(rune(ModelGongAstStageProcessing)) + `}}
									}
									__gong__map_Indentifiers_gongstructName[identifier] = gongstructName
									return
								}
							}
						}
					}
				}
				if sel := selectorExpr.Sel; sel != nil {
					astCoordinate := astCoordinate + "\tSel" + "." + sel.Name
					log.Println(astCoordinate)
				}
				for iteration, arg := range callExpr.Args {
					astCoordinate := astCoordinate + "\tArg"
					switch arg := arg.(type) {
					case *ast.BasicLit:
						basicLit := arg
						astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
						log.Println(astCoordinate)

						// first iteration should be ignored
						if iteration == 0 {
							continue
						}

						// remove first and last char
						date := basicLit.Value[1 : len(basicLit.Value)-1]

						var ok bool
						gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
						if !ok {
							log.Fatalln("gongstructName not found for identifier", identifier)
						}
						switch gongstructName {
						case "Astruct":
							switch fieldName {
							case "Date":
								__gong__map_Astruct[identifier].Date, _ = time.Parse(
									"2006-01-02 15:04:05.999999999 -0700 MST",
									date)
							}
						}
					}
				}
			case *ast.Ident:
				// append function
				ident := fun
				astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
				log.Println(astCoordinate)
			}
			for _, arg := range callExpr.Args {
				astCoordinate := astCoordinate + "\tArg"
				switch arg := arg.(type) {
				case *ast.Ident:
					ident := arg
					astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
					log.Println(astCoordinate)
					var ok bool
					gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
					if !ok {
						log.Fatalln("gongstructName not found for identifier", identifier)
					}

					switch gongstructName {
					case "Astruct":
						switch fieldName {
						case "Anarrayofb":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Bstruct[targetIdentifier]
							__gong__map_Astruct[identifier].Anarrayofb =
								append(__gong__map_Astruct[identifier].Anarrayofb, target)
						}
					}
				case *ast.SelectorExpr:
					slcExpr := arg
					astCoordinate := astCoordinate + "\tSelectorExpr"
					switch X := slcExpr.X.(type) {
					case *ast.Ident:
						ident := X
						astCoordinate := astCoordinate + "\tX" + "." + ident.Name
						log.Println(astCoordinate)

					}
					if Sel := slcExpr.Sel; Sel != nil {
						astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
						log.Println(astCoordinate)
					}
				}
			}
		case *ast.BasicLit:
			// assignment to string field
			basicLit := expr
			astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
			log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}

			switch gongstructName {
			// insertion point for basic lit assignments{{` + string(rune(ModelGongAstBasicLitAssignment)) + `}}
			}
		case *ast.Ident:
			// assignment to boolean field ?
			ident := expr
			astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
			log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}
			switch gongstructName {
			// insertion point for bool & pointers assignments{{` + string(rune(ModelGongAstIdentBooleanAndPointerAssignment)) + `}}
			}
		case *ast.SelectorExpr:
			// assignment to enum field
			selectorExpr := expr
			astCoordinate := astCoordinate + "\tSelectorExpr"
			switch X := selectorExpr.X.(type) {
			case *ast.Ident:
				ident := X
				astCoordinate := astCoordinate + "\tX" + "." + ident.Name
				log.Println(astCoordinate)
			}
			if Sel := selectorExpr.Sel; Sel != nil {
				astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
				log.Println(astCoordinate)

				// enum field
				var ok bool
				gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
				if !ok {
					log.Fatalln("gongstructName not found for identifier", identifier)
				}

				// remove first and last char
				enumValue := Sel.Name
				switch gongstructName {
				// insertion point for enums assignments{{` + string(rune(ModelGongAstIdentEnumAssignment)) + `}}
				}
				switch gongstructName {
				case "Astruct":
					switch fieldName {
					case "Aenum":
						var val AEnumType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Astruct[identifier].Aenum = AEnumType(val)
					case "Aenum_2":
						var val AEnumType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Astruct[identifier].Aenum_2 = AEnumType(val)
					case "Benum":
						var val BEnumType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Astruct[identifier].Benum = BEnumType(val)
					}
				}

			}
		}
	}
	return
}
`
