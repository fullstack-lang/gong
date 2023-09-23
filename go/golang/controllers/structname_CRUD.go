package controllers

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fullstack-lang/gong/go/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const controllersTmpl = `// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"{{PkgPathRoot}}/models"
	"{{PkgPathRoot}}/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __{{Structname}}__dummysDeclaration__ models.{{Structname}}
var __{{Structname}}_time__dummyDeclaration time.Duration

var mutex{{Structname}} sync.Mutex

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
// # Get all {{structname}}s
//
// Responses:
// default: genericError
//
//	200: {{structname}}DBResponse
func (controller *Controller) Get{{Structname}}s(c *gin.Context) {

	// source slice
	var {{structname}}DBs []orm.{{Structname}}DB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("Get{{Structname}}s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack {{PkgPathRoot}}/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepo{{Structname}}.GetDB()

	query := db.Find(&{{structname}}DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	{{structname}}APIs := make([]orm.{{Structname}}API, 0)

	// for each {{structname}}, update fields from the database nullable fields
	for idx := range {{structname}}DBs {
		{{structname}}DB := &{{structname}}DBs[idx]
		_ = {{structname}}DB
		var {{structname}}API orm.{{Structname}}API

		// insertion point for updating fields
		{{structname}}API.ID = {{structname}}DB.ID
		{{structname}}DB.CopyBasicFieldsTo{{Structname}}(&{{structname}}API.{{Structname}})
		{{structname}}API.{{Structname}}PointersEnconding = {{structname}}DB.{{Structname}}PointersEnconding
		{{structname}}APIs = append({{structname}}APIs, {{structname}}API)
	}

	c.JSON(http.StatusOK, {{structname}}APIs)
}

// Post{{Structname}}
//
// swagger:route POST /{{structname}}s {{structname}}s post{{Structname}}
//
// Creates a {{structname}}
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) Post{{Structname}}(c *gin.Context) {

	mutex{{Structname}}.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("Post{{Structname}}s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack {{PkgPathRoot}}/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepo{{Structname}}.GetDB()

	// Validate input
	var input orm.{{Structname}}API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create {{structname}}
	{{structname}}DB := orm.{{Structname}}DB{}
	{{structname}}DB.{{Structname}}PointersEnconding = input.{{Structname}}PointersEnconding
	{{structname}}DB.CopyBasicFieldsFrom{{Structname}}(&input.{{Structname}})

	query := db.Create(&{{structname}}DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepo{{Structname}}.CheckoutPhaseOneInstance(&{{structname}}DB)
	{{structname}} := backRepo.BackRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}Ptr[{{structname}}DB.ID]

	if {{structname}} != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), {{structname}})
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, {{structname}}DB)

	mutex{{Structname}}.Unlock()
}

// Get{{Structname}}
//
// swagger:route GET /{{structname}}s/{ID} {{structname}}s get{{Structname}}
//
// Gets the details for a {{structname}}.
//
// Responses:
// default: genericError
//
//	200: {{structname}}DBResponse
func (controller *Controller) Get{{Structname}}(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("Get{{Structname}}", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack {{PkgPathRoot}}/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepo{{Structname}}.GetDB()

	// Get {{structname}}DB in DB
	var {{structname}}DB orm.{{Structname}}DB
	if err := db.First(&{{structname}}DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var {{structname}}API orm.{{Structname}}API
	{{structname}}API.ID = {{structname}}DB.ID
	{{structname}}API.{{Structname}}PointersEnconding = {{structname}}DB.{{Structname}}PointersEnconding
	{{structname}}DB.CopyBasicFieldsTo{{Structname}}(&{{structname}}API.{{Structname}})

	c.JSON(http.StatusOK, {{structname}}API)
}

// Update{{Structname}}
//
// swagger:route PATCH /{{structname}}s/{ID} {{structname}}s update{{Structname}}
//
// # Update a {{structname}}
//
// Responses:
// default: genericError
//
//	200: {{structname}}DBResponse
func (controller *Controller) Update{{Structname}}(c *gin.Context) {

	mutex{{Structname}}.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("Update{{Structname}}", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack {{PkgPathRoot}}/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepo{{Structname}}.GetDB()

	// Validate input
	var input orm.{{Structname}}API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var {{structname}}DB orm.{{Structname}}DB

	// fetch the {{structname}}
	query := db.First(&{{structname}}DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	{{structname}}DB.CopyBasicFieldsFrom{{Structname}}(&input.{{Structname}})
	{{structname}}DB.{{Structname}}PointersEnconding = input.{{Structname}}PointersEnconding

	query = db.Model(&{{structname}}DB).Updates({{structname}}DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	{{structname}}New := new(models.{{Structname}})
	{{structname}}DB.CopyBasicFieldsTo{{Structname}}({{structname}}New)

	// get stage instance from DB instance, and call callback function
	{{structname}}Old := backRepo.BackRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}Ptr[{{structname}}DB.ID]
	if {{structname}}Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), {{structname}}Old, {{structname}}New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the {{structname}}DB
	c.JSON(http.StatusOK, {{structname}}DB)

	mutex{{Structname}}.Unlock()
}

// Delete{{Structname}}
//
// swagger:route DELETE /{{structname}}s/{ID} {{structname}}s delete{{Structname}}
//
// # Delete a {{structname}}
//
// default: genericError
//
//	200: {{structname}}DBResponse
func (controller *Controller) Delete{{Structname}}(c *gin.Context) {

	mutex{{Structname}}.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("Delete{{Structname}}", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack {{PkgPathRoot}}/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepo{{Structname}}.GetDB()

	// Get model if exist
	var {{structname}}DB orm.{{Structname}}DB
	if err := db.First(&{{structname}}DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&{{structname}}DB)

	// get an instance (not staged) from DB instance, and call callback function
	{{structname}}Deleted := new(models.{{Structname}})
	{{structname}}DB.CopyBasicFieldsTo{{Structname}}({{structname}}Deleted)

	// get stage instance from DB instance, and call callback function
	{{structname}}Staged := backRepo.BackRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}Ptr[{{structname}}DB.ID]
	if {{structname}}Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), {{structname}}Staged, {{structname}}Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutex{{Structname}}.Unlock()
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

// Sub Templates
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

	ControllerFileFieldSubTmplGetsTimeField
	ControllerFileFieldSubTmplPostTimeField
	ControllerFileFieldSubTmplGetTimeField
	ControllerFileFieldSubTmplUpdateTimeField

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
	// Time
	//

	ControllerFileFieldSubTmplGetsTimeField: `
		if {{structname}}.{{FieldName}}_Data.Valid {
			{{structname}}.{{FieldName}} = {{structname}}.{{FieldName}}_Data.Time
		}
`,

	ControllerFileFieldSubTmplPostTimeField: `
	{{structname}}DB.{{FieldName}}_Data.Time = input.{{FieldName}}
	{{structname}}DB.{{FieldName}}_Data.Valid = true
`,

	ControllerFileFieldSubTmplGetTimeField: `
	if {{structname}}.{{FieldName}}_Data.Valid {
		{{structname}}.{{FieldName}} = {{structname}}.{{FieldName}}_Data.Time
	}
`,

	ControllerFileFieldSubTmplUpdateTimeField: `
	input.{{FieldName}}_Data.Time = input.{{FieldName}}
	input.{{FieldName}}_Data.Valid = true
`,

	//
	// Int
	//

	ControllerFileFieldSubTmplGetsBasicFieldInt: `
		if {{structname}}.{{FieldName}}_Data.Valid {
			{{structname}}.{{FieldName}} = {{FieldType}}({{structname}}.{{FieldName}}_Data.Int64)
		}
`,

	ControllerFileFieldSubTmplPostBasicFieldInt: `
	{{structname}}DB.{{FieldName}}_Data.Int64 = int64(input.{{FieldName}})
	{{structname}}DB.{{FieldName}}_Data.Valid = true
`,

	ControllerFileFieldSubTmplGetBasicFieldInt: `
	if {{structname}}.{{FieldName}}_Data.Valid {
		{{structname}}.{{FieldName}} = {{FieldType}}({{structname}}.{{FieldName}}_Data.Int64)
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
	mdlPkg *models.ModelPkg,
	pkgName string,
	pkgGoPath string,
	dirPath string) {

	// have alphabetical order generation
	structList := []*models.GongStruct{}
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
			switch field := field.(type) {
			case *models.GongBasicField:

				switch field.GetBasicKind() {
				case types.Bool:
					insertions[ControllerFileGetsInsertion] += models.Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetsBasicFieldBool],
						"{{FieldName}}", field.Name)

					insertions[ControllerFilePostInsertion] += models.Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplPostBasicFieldBool],
						"{{FieldName}}", field.Name)

					insertions[ControllerFileGetInsertion] += models.Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetBasicFieldBool],
						"{{FieldName}}", field.Name)

					insertions[ControllerFileUpdateInsertion] += models.Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplUpdateBasicFieldBool],
						"{{FieldName}}", field.Name)

				case types.Int, types.Int64:
					insertions[ControllerFileGetsInsertion] += models.Replace2(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetsBasicFieldInt],
						"{{FieldName}}", field.Name,
						"{{FieldType}}", field.DeclaredType)

					insertions[ControllerFilePostInsertion] += models.Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplPostBasicFieldInt],
						"{{FieldName}}", field.Name)

					insertions[ControllerFileGetInsertion] += models.Replace2(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetBasicFieldInt],
						"{{FieldName}}", field.Name,
						"{{FieldType}}", field.DeclaredType)

					insertions[ControllerFileUpdateInsertion] += models.Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplUpdateBasicFieldInt],
						"{{FieldName}}", field.Name)

				case types.Float64:
					insertions[ControllerFileGetsInsertion] += models.Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetsBasicFieldFloat64],
						"{{FieldName}}", field.Name)

					insertions[ControllerFilePostInsertion] += models.Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplPostBasicFieldFloat64],
						"{{FieldName}}", field.Name)

					insertions[ControllerFileGetInsertion] += models.Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetBasicFieldFloat64],
						"{{FieldName}}", field.Name)

					insertions[ControllerFileUpdateInsertion] += models.Replace1(
						ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplUpdateBasicFieldFloat64],
						"{{FieldName}}", field.Name)

				case types.String:
					if field.GongEnum != nil {
						insertions[ControllerFileGetsInsertion] += models.Replace2(
							ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetsBasicFieldStringEnum],
							"{{FieldName}}", field.Name,
							"{{EnumType}}", field.GongEnum.Name)

						insertions[ControllerFilePostInsertion] += models.Replace2(
							ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplPostBasicFieldStringEnum],
							"{{FieldName}}", field.Name,
							"{{EnumType}}", field.GongEnum.Name)

						insertions[ControllerFileGetInsertion] += models.Replace2(
							ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetBasicFieldStringEnum],
							"{{FieldName}}", field.Name,
							"{{EnumType}}", field.GongEnum.Name)

						insertions[ControllerFileUpdateInsertion] += models.Replace2(
							ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplUpdateBasicFieldStringEnum],
							"{{FieldName}}", field.Name,
							"{{EnumType}}", field.GongEnum.Name)
					} else {
						insertions[ControllerFileGetsInsertion] += models.Replace1(
							ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetsBasicFieldString],
							"{{FieldName}}", field.Name)

						insertions[ControllerFilePostInsertion] += models.Replace1(
							ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplPostBasicFieldString],
							"{{FieldName}}", field.Name)

						insertions[ControllerFileGetInsertion] += models.Replace1(
							ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetBasicFieldString],
							"{{FieldName}}", field.Name)

						insertions[ControllerFileUpdateInsertion] += models.Replace1(
							ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplUpdateBasicFieldString],
							"{{FieldName}}", field.Name)
					}
				}
			case *models.GongTimeField:

				insertions[ControllerFileGetsInsertion] += models.Replace1(
					ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetsTimeField],
					"{{FieldName}}", field.Name)

				insertions[ControllerFilePostInsertion] += models.Replace1(
					ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplPostTimeField],
					"{{FieldName}}", field.Name)

				insertions[ControllerFileGetInsertion] += models.Replace1(
					ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplGetTimeField],
					"{{FieldName}}", field.Name)

				insertions[ControllerFileUpdateInsertion] += models.Replace1(
					ControllerFileFieldFieldSubTemplateCode[ControllerFileFieldSubTmplUpdateTimeField],
					"{{FieldName}}", field.Name)
			default:
			}
		}

		// substitutes {{<<insertion points>>}} stuff with generated code
		for insertion := ControllerFileInsertionPoint(0); insertion < ControllerFileNbInsertionPoints; insertion++ {
			toReplace := "{{" + string(rune(insertion)) + "}}"
			codeGO = strings.ReplaceAll(codeGO, toReplace, insertions[insertion])
		}

		caserEnglish := cases.Title(language.English)
		codeGO = models.Replace6(codeGO,
			"{{PkgName}}", pkgName,
			"{{TitlePkgName}}", caserEnglish.String(pkgName),
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
