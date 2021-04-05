package models

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"

	"golang.org/x/tools/go/packages"
)

// FindComments
func FindComments(pkg *packages.Package, name string) (*ast.CommentGroup, bool) {
	for _, f := range pkg.Syntax {
		for _, d := range f.Decls {
			gd, ok := d.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, s := range gd.Specs {
				if ts, ok := s.(*ast.TypeSpec); ok {
					if ts.Name.Name == name {
						return gd.Doc, true
					}
				}
			}
		}
	}
	return nil, false
}

func FindEnumValues(pkg *packages.Package, enumName string) (list []interface{}, _ bool) {
	for _, f := range pkg.Syntax {
		for _, d := range f.Decls {
			gd, ok := d.(*ast.GenDecl)
			if !ok {
				continue
			}

			if gd.Tok != token.CONST {
				continue
			}

			for _, s := range gd.Specs {
				if vs, ok := s.(*ast.ValueSpec); ok {
					if vsIdent, ok := vs.Type.(*ast.Ident); ok {
						if vsIdent.Name == enumName {
							if len(vs.Values) > 0 {
								if bl, ok := vs.Values[0].(*ast.BasicLit); ok {
									list = append(list, getEnumBasicLitValue(bl))
								}
							}
						}
					}
				}
			}
		}
	}
	return list, true
}

func getEnumBasicLitValue(basicLit *ast.BasicLit) interface{} {
	switch basicLit.Kind.String() {
	case "INT":
		if result, err := strconv.ParseInt(basicLit.Value, 10, 64); err == nil {
			return result
		}
	case "FLOAT":
		if result, err := strconv.ParseFloat(basicLit.Value, 64); err == nil {
			return result
		}
	default:
		return strings.Trim(basicLit.Value, "\"")
	}
	return nil
}
