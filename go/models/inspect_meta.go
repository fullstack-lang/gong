package models

import (
	"go/ast"
	"log"
	"strings"
)

// inspectMeta look for
// expression in the AST as
// var OnTriangleImplementation = Meta{
//
// That is a General Declaration (because it is exported)
// The Type in the Value Specification of the declaration must be "Meta"
func inspectMeta(stage *StageStruct, astPackage *ast.Package) {

	ast.Inspect(astPackage, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			if len(x.Specs) != 1 {
				break
			}
			// look for cases where type is "Meta"
			switch vs := x.Specs[0].(type) {
			case *ast.ValueSpec:
				if len(vs.Values) != 1 {
					break
				}
				switch cl := vs.Values[0].(type) {
				case *ast.CompositeLit:
					switch ident := cl.Type.(type) {
					case *ast.Ident:
						if ident.Name != "Meta" {
							break
						}
						meta := (&Meta{Name: vs.Names[0].Name}).Stage(stage)
						_ = meta
						log.Println("A Meta declaration ", meta.Name)
						// parse both elements
						//		Text: `Astruct is the typical gongstruct and it is references Bstruct in different ways`,
						//	 References: []any{
						if len(cl.Elts) != 2 {
							log.Fatalln("There should be 2 elts in meta declaration")
						}
						switch kv := cl.Elts[0].(type) {
						case *ast.KeyValueExpr:
							switch bl := kv.Value.(type) {
							case *ast.BasicLit:
								meta.Text = strings.TrimPrefix(bl.Value, "`")
								meta.Text = strings.TrimSuffix(meta.Text, "`")
							}
						}
						switch kv := cl.Elts[1].(type) {
						case *ast.KeyValueExpr:
							switch cl := kv.Value.(type) {
							case *ast.CompositeLit:
								for _, elt := range cl.Elts {
									switch x := elt.(type) {
									case *ast.CompositeLit:
										switch i := x.Type.(type) {
										case *ast.Ident:
											meta.MetaReferences = append(meta.MetaReferences,
												(&MetaReference{Name: i.Name}).Stage(stage))
										}
									case *ast.SelectorExpr:
										metaRef := (&MetaReference{Name: "tmp"}).Stage(stage)
										meta.MetaReferences = append(meta.MetaReferences,
											metaRef)
										switch cl := x.X.(type) {
										case *ast.CompositeLit:
											switch i := cl.Type.(type) {
											case *ast.Ident:
												metaRef.Name = i.Name
											}
										}
										metaRef.Name += "." + x.Sel.Name
									}
								}
							}
						}
					}
				}
			}
		}
		return true
	})

}
