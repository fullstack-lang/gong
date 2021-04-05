package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/examples/bookstore/go/orm"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// RegisterControllers register controllers
func RegisterControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/examples/bookstore/go")
	{ // insertion point for registrations
		v1.GET("/v1/areas", GetAreas)
		v1.GET("/v1/areas/:id", GetArea)
		v1.POST("/v1/areas", PostArea)
		v1.PATCH("/v1/areas/:id", UpdateArea)
		v1.PUT("/v1/areas/:id", UpdateArea)
		v1.DELETE("/v1/areas/:id", DeleteArea)

		v1.GET("/v1/books", GetBooks)
		v1.GET("/v1/books/:id", GetBook)
		v1.POST("/v1/books", PostBook)
		v1.PATCH("/v1/books/:id", UpdateBook)
		v1.PUT("/v1/books/:id", UpdateBook)
		v1.DELETE("/v1/books/:id", DeleteBook)

		v1.GET("/v1/editors", GetEditors)
		v1.GET("/v1/editors/:id", GetEditor)
		v1.POST("/v1/editors", PostEditor)
		v1.PATCH("/v1/editors/:id", UpdateEditor)
		v1.PUT("/v1/editors/:id", UpdateEditor)
		v1.DELETE("/v1/editors/:id", DeleteEditor)

		v1.GET("/commitnb", GetLastCommitNb)
	}
}

// swagger:route GET /commitnb backrepo GetLastCommitNb
func GetLastCommitNb(c *gin.Context) {
	res := orm.GetLastCommitNb()

	c.JSON(http.StatusOK, res)
}
