package models

import (
	"go/ast"
	"go/token"
	"log"
	"strings"
)

func parsestringfield(fieldname string, kve *ast.KeyValueExpr, fset *token.FileSet, target *string) (match bool) {

	var ident *ast.Ident
	var ok bool
	if ident, ok = kve.Key.(*ast.Ident); !ok {
		log.Panic("Expecting an ident " + fset.Position(kve.Pos()).String())
	}

	match = false
	if ident.Name == fieldname {
		var bl *ast.BasicLit
		if bl, ok = kve.Value.(*ast.BasicLit); !ok {
			// TODO deal with recursive Binary Expressions "titi"+"toto"+"tata"
			log.Panic("Expecting a basic lit " + fset.Position(kve.Value.Pos()).String())
		}
		// strangely, on have to remove first and last caracter if they are "
		*target = bl.Value
		*target = strings.TrimSuffix(*target, "\"")
		*target = strings.TrimPrefix(*target, "\"")
		match = true
	}
	return
}
