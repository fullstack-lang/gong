package models

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const controllersTmpl = `// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"net/http"

	"{{PkgPathRoot}}/models"
	"{{PkgPathRoot}}/orm"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// declaration in order to justify use of the models import
var __{{Structname}}__dummysDeclaration__ models.{{Structname}}

// An {{Structname}}ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters get{{Structname}} update{{Structname}} delete{{Structname}}
type {{Structname}}ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// {{Structname}}Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters post{{Structname}} update{{Structname}}
type {{Structname}}Input struct {
	// The {{Structname}} to submit or modify
	// in: body
	{{Structname}} *orm.{{Structname}}API
}

// Get{{Structname}}s
//
// swagger:route GET /{{structname}}s {{structname}}s get{{Structname}}s
//
// Get all {{structname}}s
//
// Responses:
//    default: genericError
//        200: {{structname}}DBsResponse
func Get{{Structname}}s(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var {{structname}}s []orm.{{Structname}}DB
	query := db.Find(&{{structname}}s)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// for each {{structname}}, update fields from the database nullable fields
	for idx := range {{structname}}s {
		{{structname}} := &{{structname}}s[idx]
		_ = {{structname}}
		// insertion point for updating fields{{` + string(rune(ControllerFileGetInsertion)) + `}}
	}

	c.JSON(http.StatusOK, {{structname}}s)
}

// Post{{Structname}}
//
// swagger:route POST /{{structname}}s {{structname}}s post{{Structname}}
//
// Creates a {{structname}}
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: {{structname}}DBResponse
func Post{{Structname}}(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input orm.{{Structname}}API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create {{structname}}
	{{structname}}DB := orm.{{Structname}}DB{}
	{{structname}}DB.{{Structname}}API = input
	// insertion point for nullable field set{{` + string(rune(ControllerFilePostInsertion)) + `}}

	query := db.Create(&{{structname}}DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementCommitNb()

	c.JSON(http.StatusOK, {{structname}}DB)
}

// Get{{Structname}}
//
// swagger:route GET /{{structname}}s/{ID} {{structname}}s get{{Structname}}
//
// Gets the details for a {{structname}}.
//
// Responses:
//    default: genericError
//        200: {{structname}}DBResponse
func Get{{Structname}}(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get {{structname}} in DB
	var {{structname}} orm.{{Structname}}DB
	if err := db.First(&{{structname}}, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// insertion point for fields value set from nullable fields{{` + string(rune(ControllerFileGetInsertion)) + `}}

	c.JSON(http.StatusOK, {{structname}})
}

// Update{{Structname}}
//
// swagger:route PATCH /{{structname}}s/{ID} {{structname}}s update{{Structname}}
//
// Update a {{structname}}
//
// Responses:
//    default: genericError
//        200: {{structname}}DBResponse
func Update{{Structname}}(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var {{structname}}DB orm.{{Structname}}DB

	// fetch the {{structname}}
	query := db.First(&{{structname}}DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.{{Structname}}API
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	// insertion point for nullable field set{{` + string(rune(ControllerFileUpdateInsertion)) + `}}

	query = db.Model(&{{structname}}DB).Updates(input)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementCommitNb()

	// return status OK with the marshalling of the the {{structname}}DB
	c.JSON(http.StatusOK, {{structname}}DB)
}

// Delete{{Structname}}
//
// swagger:route DELETE /{{structname}}s/{ID} {{structname}}s delete{{Structname}}
//
// Delete a {{structname}}
//
// Responses:
//    default: genericError
func Delete{{Structname}}(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var {{structname}}DB orm.{{Structname}}DB
	if err := db.First(&{{structname}}DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&{{structname}}DB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementCommitNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
`

// insertion points
type ControllerFileInsertionPoint int

const (
	ControllerFileGetsInsertion ControllerFileInsertionPoint = iota
	ControllerFilePostInsertion
	ControllerFileGetInsertion
	ControllerFileUpdateInsertion
	ControllerFileNbInsertionPoints
)

//
// Sub Templates
//
type ControllerFilPerStructSubTemplate int

const (
	ControllerFileFieldSubTmplGetsBasicFieldBool ControllerFilPerStructSubTemplate = iota
	ControllerFileFieldSubTmplPostBasicFieldBool
	ControllerFileFieldSubTmplGetBasicFieldBool
	ControllerFileFieldSubTmplUpdateBasicFieldBool

	ControllerFileFieldSubTmplGetsBasicFieldInt
	ControllerFileFieldSubTmplPostBasicFieldInt
	ControllerFileFieldSubTmplGetBasicFieldInt
	ControllerFileFieldSubTmplUpdateBasicFieldInt

	ControllerFileFieldSubTmplGetsBasicFieldFloat64
	ControllerFileFieldSubTmplPostBasicFieldFloat64
	ControllerFileFieldSubTmplGetBasicFieldFloat64
	ControllerFileFieldSubTmplUpdateBasicFieldFloat64

	ControllerFileFieldSubTmplGetsBasicFieldString
	ControllerFileFieldSubTmplPostBasicFieldString
	ControllerFileFieldSubTmplGetBasicFieldString
	ControllerFileFieldSubTmplUpdateBasicFieldString

	ControllerFileFieldSubTmplGetsBasicFieldStringEnum
	ControllerFileFieldSubTmplPostBasicFieldStringEnum
	ControllerFileFieldSubTmplGetBasicFieldStringEnum
	ControllerFileFieldSubTmplUpdateBasicFieldStringEnum
)

var ControllerFileFieldFieldSubTemplateCode map[ControllerFilPerStructSubTemplate]string = // declaration of the sub templates
map[ControllerFilPerStructSubTemplate]string{

	//
	// Bool
	//

	ControllerFileFieldSubTmplGetsBasicFieldBool: `
	{{structname}}.{{FieldName}} = {{structname}}.{{FieldName}}_Data.Bool
`,

	ControllerFileFieldSubTmplPostBasicFieldBool: `
	{{structname}}DB.{{FieldName}}_Data.Bool = input.{{FieldName}}
	{{structname}}DB.{{FieldName}}_Data.Valid = true
`,

	ControllerFileFieldSubTmplGetBasicFieldBool: `
	{{structname}}.{{FieldName}} = {{structname}}.{{FieldName}}_Data.Bool
`,

	ControllerFileFieldSubTmplUpdateBasicFieldBool: `
	input.{{FieldName}}_Data.Bool = input.{{FieldName}}
	input.{{FieldName}}_Data.Valid = true
`,

	//
	// Int
	//

	ControllerFileFieldSubTmplGetsBasicFieldInt: `
	if {{structname}}.{{FieldName}}_Data.Valid {
		{{structname}}.{{FieldName}} = {{structname}}.{{FieldName}}_Data.Int64
	}
`,

	ControllerFileFieldSubTmplPostBasicFieldInt: `
	{{structname}}DB.{{FieldName}}_Data.Int64 = int64(input.{{FieldName}})
	{{structname}}DB.{{FieldName}}_Data.Valid = true
`,

	ControllerFileFieldSubTmplGetBasicFieldInt: `
	if {{structname}}.{{FieldName}}_Data.Valid {
		{{structname}}.{{FieldName}} = int({{structname}}.{{FieldName}}_Data.Int64)
	}
`,

	ControllerFileFieldSubTmplUpdateBasicFieldInt: `
	input.{{FieldName}}_Data.Int64 = int64(input.{{FieldName}})
	input.{{FieldName}}_Data.Valid = true
`,

	//
	// Float64
	//

	ControllerFileFieldSubTmplGetsBasicFieldFloat64: `
	if {{structname}}.{{FieldName}}_Data.Valid {
		{{structname}}.{{FieldName}} = {{structname}}.{{FieldName}}_Data.Float64
	}
`,

	ControllerFileFieldSubTmplPostBasicFieldFloat64: `
	{{structname}}DB.{{FieldName}}_Data.Float64 = input.{{FieldName}}
	{{structname}}DB.{{FieldName}}_Data.Valid = true
`,

	ControllerFileFieldSubTmplGetBasicFieldFloat64: `
	if {{structname}}.{{FieldName}}_Data.Valid {
		{{structname}}.{{FieldName}} = {{structname}}.{{FieldName}}_Data.Float64
	}
`,

	ControllerFileFieldSubTmplUpdateBasicFieldFloat64: `
	input.{{FieldName}}_Data.Float64 = input.{{FieldName}}
	input.{{FieldName}}_Data.Valid = true
`,

	//
	// String
	//

	ControllerFileFieldSubTmplGetsBasicFieldString: `
	if {{structname}}.{{FieldName}}_Data.Valid {
		{{structname}}.{{FieldName}} = {{structname}}.{{FieldName}}_Data.String
	}
`,

	ControllerFileFieldSubTmplPostBasicFieldString: `
	{{structname}}DB.{{FieldName}}_Data.String = input.{{FieldName}}
	{{structname}}DB.{{FieldName}}_Data.Valid = true
`,

	ControllerFileFieldSubTmplGetBasicFieldString: `
	if {{structname}}.{{FieldName}}_Data.Valid {
		{{structname}}.{{FieldName}} = {{structname}}.{{FieldName}}_Data.String
	}
`,

	ControllerFileFieldSubTmplUpdateBasicFieldString: `
	input.{{FieldName}}_Data.String = input.{{FieldName}}
	input.{{FieldName}}_Data.Valid = true
`,

	//
	// String Enum
	//

	ControllerFileFieldSubTmplGetsBasicFieldStringEnum: `
	if {{structname}}.{{FieldName}}_Data.Valid {
		{{structname}}.{{FieldName}} = models.{{EnumType}}({{structname}}.{{FieldName}}_Data.String)
	}
`,

	ControllerFileFieldSubTmplPostBasicFieldStringEnum: `
	{{structname}}DB.{{FieldName}}_Data.String = string(input.{{FieldName}})
	{{structname}}DB.{{FieldName}}_Data.Valid = true
`,

	ControllerFileFieldSubTmplGetBasicFieldStringEnum: `
	if {{structname}}.{{FieldName}}_Data.Valid {
		{{structname}}.{{FieldName}} = models.{{EnumType}}({{structname}}.{{FieldName}}_Data.String)
	}
`,

	ControllerFileFieldSubTmplUpdateBasicFieldStringEnum: `
	input.{{FieldName}}_Data.String = string(input.{{FieldName}})
	input.{{FieldName}}_Data.Valid = true
`,
}

// MultiCodeGeneratorControllers parses mdlPkg and generates the code for the
// back repository code
func MultiCodeGeneratorControllers(
	mdlPkg *ModelPkg,
	pkgName string,
	pkgGoPath string,
	dirPath string) {

	// have alphabetical order generation
	structList := []*GongStruct{}
	for _, _struct := range mdlPkg.GongStructs {
		structList = append(structList, _struct)
	}
	sort.Slice(structList[:], func(i, j int) bool {
		return structList[i].Name < structList[j].Name
	})

	for _, _struct := range structList {
		_ = _struct

		if !_struct.HasNameField() {
			continue
		}

		insertions := make(map[ControllerFileInsertionPoint]string)
		for insertion := ControllerFileInsertionPoint(0); insertion < ControllerFileNbInsertionPoints; insertion++ {
			insertions[insertion] = ""
		}

		codeGO := controllersTmpl

		for _, field := range _struct.Fields {
			switch field.(type) {
			case *GongBasicField:
				modelBasicField := field.(*GongBasicField)

				switch modelBasicField.basicKind {
				case types.Bool:
					insertions[ControllerFileGetsInsertion] += Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetsBasicFieldBool],
						"{{FieldName}}", modelBasicField.Name)

					insertions[ControllerFilePostInsertion] += Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplPostBasicFieldBool],
						"{{FieldName}}", modelBasicField.Name)

					insertions[ControllerFileGetInsertion] += Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetBasicFieldBool],
						"{{FieldName}}", modelBasicField.Name)

					insertions[ControllerFileUpdateInsertion] += Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplUpdateBasicFieldBool],
						"{{FieldName}}", modelBasicField.Name)

				case types.Int:
					insertions[ControllerFileGetsInsertion] += Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetsBasicFieldInt],
						"{{FieldName}}", modelBasicField.Name)

					insertions[ControllerFilePostInsertion] += Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplPostBasicFieldInt],
						"{{FieldName}}", modelBasicField.Name)

					insertions[ControllerFileGetInsertion] += Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetBasicFieldInt],
						"{{FieldName}}", modelBasicField.Name)

					insertions[ControllerFileUpdateInsertion] += Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplUpdateBasicFieldInt],
						"{{FieldName}}", modelBasicField.Name)

				case types.Float64:
					insertions[ControllerFileGetsInsertion] += Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetsBasicFieldFloat64],
						"{{FieldName}}", modelBasicField.Name)

					insertions[ControllerFilePostInsertion] += Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplPostBasicFieldFloat64],
						"{{FieldName}}", modelBasicField.Name)

					insertions[ControllerFileGetInsertion] += Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetBasicFieldFloat64],
						"{{FieldName}}", modelBasicField.Name)

					insertions[ControllerFileUpdateInsertion] += Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplUpdateBasicFieldFloat64],
						"{{FieldName}}", modelBasicField.Name)

				case types.String:
					if modelBasicField.GongEnum != nil {
						insertions[ControllerFileGetsInsertion] += Replace2(
							ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetsBasicFieldStringEnum],
							"{{FieldName}}", modelBasicField.Name,
							"{{EnumType}}", modelBasicField.GongEnum.Name)

						insertions[ControllerFilePostInsertion] += Replace2(
							ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplPostBasicFieldStringEnum],
							"{{FieldName}}", modelBasicField.Name,
							"{{EnumType}}", modelBasicField.GongEnum.Name)

						insertions[ControllerFileGetInsertion] += Replace2(
							ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetBasicFieldStringEnum],
							"{{FieldName}}", modelBasicField.Name,
							"{{EnumType}}", modelBasicField.GongEnum.Name)

						insertions[ControllerFileUpdateInsertion] += Replace2(
							ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplUpdateBasicFieldStringEnum],
							"{{FieldName}}", modelBasicField.Name,
							"{{EnumType}}", modelBasicField.GongEnum.Name)
					} else {
						insertions[ControllerFileGetsInsertion] += Replace1(
							ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetsBasicFieldString],
							"{{FieldName}}", modelBasicField.Name)

						insertions[ControllerFilePostInsertion] += Replace1(
							ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplPostBasicFieldString],
							"{{FieldName}}", modelBasicField.Name)

						insertions[ControllerFileGetInsertion] += Replace1(
							ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetBasicFieldString],
							"{{FieldName}}", modelBasicField.Name)

						insertions[ControllerFileUpdateInsertion] += Replace1(
							ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplUpdateBasicFieldString],
							"{{FieldName}}", modelBasicField.Name)
					}
				}
			default:
			}
		}

		// substitutes {{<<insertion points>>}} stuff with generated code
		for insertion := ControllerFileInsertionPoint(0); insertion < ControllerFileNbInsertionPoints; insertion++ {
			toReplace := "{{" + string(rune(insertion)) + "}}"
			codeGO = strings.ReplaceAll(codeGO, toReplace, insertions[insertion])
		}

		codeGO = Replace6(codeGO,
			"{{PkgName}}", pkgName,
			"{{TitlePkgName}}", strings.Title(pkgName),
			"{{pkgname}}", strings.ToLower(pkgName),
			"{{PkgPathRoot}}", strings.ReplaceAll(pkgGoPath, "/models", ""),
			"{{Structname}}", _struct.Name,
			"{{structname}}", strings.ToLower(_struct.Name))

		file, err := os.Create(filepath.Join(dirPath, _struct.Name+"_CRUDs.go"))
		if err != nil {
			log.Panic(err)
		}
		defer file.Close()
		fmt.Fprint(file, codeGO)
	}
}