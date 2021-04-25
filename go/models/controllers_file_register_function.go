package models

const ControllersRegisterTemplate = `
package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"{{PkgPathRoot}}/orm"
)

// genQuery return the name of the column
func genQuery( columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32 ` + "`" + `json:"code"` + "`" + `
		Message string ` + "`" + `json:"message"` + "`" + `
	} ` + "`" + `json:"body"` + "`" + `
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  ` + "`" + `json:"code"` + "`" + `
		Message string ` + "`" + `json:"message"` + "`" + `
		Field   string ` + "`" + `json:"field"` + "`" + `
	} ` + "`" + `json:"body"` + "`" + `
}

// RegisterControllers register controllers
func RegisterControllers(r *gin.Engine) {
	v1 := r.Group("/api/{{PkgPathRoot}}")
	{// insertion point for registrations{{` + string(rune(ControllersDeclaration)) + `}}

		v1.GET("/commitnb", GetLastCommitNb)
	}
}

// swagger:route GET /commitnb backrepo GetLastCommitNb
func GetLastCommitNb(c *gin.Context) {
	res := orm.GetLastCommitNb()

	c.JSON(http.StatusOK, res)
}

`

type ControllersRegistrationsSubTemplateInsertions int

const (
	ControllersDeclaration ControllersRegistrationsSubTemplateInsertions = iota
)

var ControllersRegistrationsSubTemplate map[string]string = // new line
map[string]string{

	string(rune(ControllersDeclaration)): `
		v1.GET("/v1/{{structname}}s", Get{{Structname}}s)
		v1.GET("/v1/{{structname}}s/:id", Get{{Structname}})
		v1.POST("/v1/{{structname}}s", Post{{Structname}})
		v1.PATCH("/v1/{{structname}}s/:id", Update{{Structname}})
		v1.PUT("/v1/{{structname}}s/:id", Update{{Structname}})
		v1.DELETE("/v1/{{structname}}s/:id", Delete{{Structname}})
`,
}
