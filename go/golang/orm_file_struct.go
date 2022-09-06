package golang

const OrmFileStructTemplate = `// insertion point for per struct generated file {{` + string(OrmFileStructDeclarations) + `}}`

type OrmFileStructSubTemplate string

const (
	OrmFileStructDeclarations OrmFileStructSubTemplate = "OrmFileStructDeclarations"
)

var OrmFileStructSubTemplateCode map[string]string = // new line
map[string]string{
	string(OrmFileStructDeclarations): `
package orm

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"{{PkgPathRoot}}/models"
)

// {{Structname}}API is the input in POST API
// 
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model {{structname}}API
type {{Structname}}API struct {
	models.{{Structname}}

	// insertion point for implementation of pointer to struct {{` + string(OrmFileStructPtrToStructDecls) + `}}

	// insertion point for implementation of reverse pointer to struct {{` + string(OrmFileStructReversePtrToStructDecls) + `}}
}

// {{Structname}}DB describes a {{structname}} in the database
//
// It incorporates all fields : from the model, from the generated field for the API and the GORM ID
//
// swagger:model {{structname}}DB
type {{Structname}}DB struct {
	gorm.Model

	{{Structname}}API
}

// {{Structname}}DBs arrays {{structname}}DBs
// swagger:response {{structname}}DBsResponse
type {{Structname}}DBs []{{Structname}}DB

// {{Structname}}DBResponse provides response
// swagger:response {{structname}}DBResponse
type {{Structname}}DBResponse struct {
	{{Structname}}DB
}
`,
}

type OrmFileStructSubSubTemplate string

const (
	OrmFileStructPtrToStructDecls        OrmFileStructSubSubTemplate = "OrmFileStructPtrToStructDecls"
	OrmFileStructReversePtrToStructDecls OrmFileStructSubSubTemplate = "OrmFileStructReversePtrToStructDecls"
)

var OrmFileStructSubSubTemplateCode map[string]string = // new line
map[string]string{
	string(OrmFileStructPtrToStructDecls): `
	// field {{FieldName}} is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable a GORM "HAS ONE" association
	{{FieldName}}ID *uint
	`,

	string(OrmFileStructReversePtrToStructDecls): `
	// ID generated for the implementation of the field {{Structname}}{}.{{FieldName}} []*{{AssocStructName}}
	{{Structname}}_{{FieldName}}ID uint`,
}

// for each sub sub template, what sub template it relates to
var OrmFileStructSubSubToSubMap map[string]string = //
map[string]string{
	string(OrmFileStructPtrToStructDecls):        string(OrmFileStructDeclarations),
	string(OrmFileStructReversePtrToStructDecls): string(OrmFileStructDeclarations),
}
