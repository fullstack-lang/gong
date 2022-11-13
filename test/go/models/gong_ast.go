package models

import (
	"go/ast"
	"log"
	"strconv"
	"time"
)

var __gong__map_Indentifiers_gongstructName = make(map[string]string)

// insertion point for identifiers maps
var __gong__map_Astruct = make(map[string]*Astruct)
var __gong__map_AstructBstruct2Use = make(map[string]*AstructBstruct2Use)
var __gong__map_AstructBstructUse = make(map[string]*AstructBstructUse)
var __gong__map_Bstruct = make(map[string]*Bstruct)
var __gong__map_Dstruct = make(map[string]*Dstruct)

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
									case "Astruct":
										instanceAstruct := (&Astruct{Name: instanceName}).Stage()
										instance = any(instanceAstruct)
										__gong__map_Astruct[identifier] = instanceAstruct
									case "Bstruct":
										instanceBstruct := (&Bstruct{Name: instanceName}).Stage()
										instance = any(instanceBstruct)
										__gong__map_Bstruct[identifier] = instanceBstruct
									case "Dstruct":
										instanceDstruct := (&Dstruct{Name: instanceName}).Stage()
										instance = any(instanceDstruct)
										__gong__map_Dstruct[identifier] = instanceDstruct
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
			case "Astruct":
				switch fieldName {
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Astruct[identifier].Name = fielValue
				case "CName":
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Astruct[identifier].CName = fielValue
				case "CFloatfield":
					// conevert string to float
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Astruct[identifier].CFloatfield = fielValue
				case "Intfield":
					// conevert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Astruct[identifier].Intfield = int(fielValue)
				case "Duration1":
					// conevert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Astruct[identifier].Duration1 = time.Duration(fielValue)
				}
			case "Bstruct":
				switch fieldName {
				case "Name":
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bstruct[identifier].Name = fielValue
				}
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
			case "Astruct":
				switch fieldName {
				case "Booleanfield":
					// conevert string to bool
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Astruct[identifier].Booleanfield = fielValue
				case "Anotherbooleanfield":
					// conevert string to bool
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Astruct[identifier].Anotherbooleanfield = fielValue
				case "Associationtob":
					targetIdentifier := ident.Name
					__gong__map_Astruct[identifier].Associationtob = __gong__map_Bstruct[targetIdentifier]
				case "Anotherassociationtob_2":
					targetIdentifier := ident.Name
					__gong__map_Astruct[identifier].Anotherassociationtob_2 = __gong__map_Bstruct[targetIdentifier]
				}
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
