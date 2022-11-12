package models

import (
	"go/ast"
	"log"
)

var __gong__map_Astruct = make(map[string]*Astruct)
var __gong__map_Bstruct = make(map[string]*Bstruct)

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging[Type Gongstruct](assignStmt *ast.AssignStmt, astCoordinate_ string) (instance *Type, identifier string) {
	astCoordinate := "\tAssignStmt: "
	for _, expr := range assignStmt.Lhs {
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

									// this is the place where an instance is created
									switch Sel.Name {
									case "Astruct":
										instanceAstruct := (&Astruct{Name: instanceName}).Stage()
										instance = any(instanceAstruct).(*Type)
										__gong__map_Astruct[instanceName] = instanceAstruct
										return instance, identifier
									}
								}
							}
						}
					}
				}
				if sel := selectorExpr.Sel; sel != nil {
					astCoordinate := astCoordinate + "\tSel" + "." + sel.Name
					log.Println(astCoordinate)
				}
				for _, arg := range callExpr.Args {
					astCoordinate := astCoordinate + "\tArg"
					switch arg := arg.(type) {
					case *ast.BasicLit:
						basicLit := arg
						astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
						log.Println(astCoordinate)
					}
				}
			case *ast.Ident:
				// assignment to boolean field ?
				ident := fun
				astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
				log.Println(astCoordinate)
			}
		case *ast.BasicLit:
			// assignment to string field
			basicLit := expr
			astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
			log.Println(astCoordinate)
		case *ast.Ident:
			// assignment to boolean field ?
			ident := expr
			astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
			log.Println(astCoordinate)
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
			}
		}
	}
	return
}
