package models

import (
	"go/ast"
	"log"
)

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging[Type Gongstruct](assignStmt *ast.AssignStmt) (instance *Type, identifier string) {

	for _, expr := range assignStmt.Lhs {
		switch expr := expr.(type) {
		case *ast.Ident:
			ident := expr
			log.Println("\t\t\tAST Lhs: ", ident.Name)
			identifier = ident.Name
		}
	}
	for _, expr := range assignStmt.Rhs {
		switch expr := expr.(type) {
		case *ast.CallExpr:
			callExpr := expr
			switch fun := callExpr.Fun.(type) {
			// the is Fun      Expr
			// function expression xxx.Stage()
			case *ast.SelectorExpr:
				selectorExpr := fun
				switch x := selectorExpr.X.(type) {
				case *ast.ParenExpr:
					// A ParenExpr node represents a parenthesized expression.
					// the is the
					//   { Name : "A1"}
					parenExpr := x
					switch x := parenExpr.X.(type) {
					case *ast.UnaryExpr:
						unaryExpr := x
						switch x := unaryExpr.X.(type) {
						case *ast.CompositeLit:
							compositeLit := x
							for _, elt := range compositeLit.Elts {
								switch elt := elt.(type) {
								case *ast.KeyValueExpr:
									// This is expression
									//     Name: "A1"
									keyValueExpr := elt
									switch key := keyValueExpr.Key.(type) {
									case *ast.Ident:
										ident := key
										log.Println("\t\t\t\tAST Rhs Fun Sel X: ", ident.Name)
									}
									switch value := keyValueExpr.Value.(type) {
									case *ast.BasicLit:
										basicLit := value
										log.Println("\t\t\t\tAST Rhs Fun Sel X: ", basicLit.Value)
									}
								}
							}
							switch type_ := compositeLit.Type.(type) {
							case *ast.SelectorExpr:
								slcExpr := type_
								switch X := slcExpr.X.(type) {
								case *ast.Ident:
									ident := X
									log.Println("\t\t\tAST ------- X : ", ident)
								}
								if Sel := slcExpr.Sel; Sel != nil {
									log.Println("\t\t\tAST ------- Sel : ", Sel.Name)
								}
							}
						}
					}
				}
				if sel := selectorExpr.Sel; sel != nil {
					log.Println("\t\t\tAST Rhs Fun Sel Sel: ", sel.Name)
				}
			}
		}
	}
	return
}
