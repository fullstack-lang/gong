package models

import (
	"go/ast"
)

// IsNamedStructWithoutEmbedded checks if any field (including embedded fields) has a field named "Name" of type string.
func IsNamedStructWithoutEmbedded(_type *ast.StructType, mapStructWithNameField map[string]any) bool {
	for _, field := range _type.Fields.List {
		// Check if the field is an embedded struct
		if len(field.Names) == 0 {
			// If the embedded field is itself a struct, check recursively
			if ident, ok := field.Type.(*ast.Ident); ok {

				if mapStructWithNameField != nil {
					if _, ok := mapStructWithNameField[ident.Name]; ok {
						return true
					}
				}
			}
		} else {
			// Check if the field is named "Name" and is of type string
			for _, name := range field.Names {
				if name.Name == "Name" {
					if ident, ok := field.Type.(*ast.Ident); ok && ident.Name == "string" {
						return true
					}
				}
			}
		}
	}
	return false
}
