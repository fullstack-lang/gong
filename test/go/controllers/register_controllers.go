package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gong/test/go/orm"
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
	v1 := r.Group("/api/github.com/fullstack-lang/gong/test/go")
	{ // insertion point for registrations
		v1.GET("/v1/aclasss", GetAclasss)
		v1.GET("/v1/aclasss/:id", GetAclass)
		v1.POST("/v1/aclasss", PostAclass)
		v1.PATCH("/v1/aclasss/:id", UpdateAclass)
		v1.PUT("/v1/aclasss/:id", UpdateAclass)
		v1.DELETE("/v1/aclasss/:id", DeleteAclass)

		v1.GET("/v1/aclassbclass2uses", GetAclassBclass2Uses)
		v1.GET("/v1/aclassbclass2uses/:id", GetAclassBclass2Use)
		v1.POST("/v1/aclassbclass2uses", PostAclassBclass2Use)
		v1.PATCH("/v1/aclassbclass2uses/:id", UpdateAclassBclass2Use)
		v1.PUT("/v1/aclassbclass2uses/:id", UpdateAclassBclass2Use)
		v1.DELETE("/v1/aclassbclass2uses/:id", DeleteAclassBclass2Use)

		v1.GET("/v1/aclassbclassuses", GetAclassBclassUses)
		v1.GET("/v1/aclassbclassuses/:id", GetAclassBclassUse)
		v1.POST("/v1/aclassbclassuses", PostAclassBclassUse)
		v1.PATCH("/v1/aclassbclassuses/:id", UpdateAclassBclassUse)
		v1.PUT("/v1/aclassbclassuses/:id", UpdateAclassBclassUse)
		v1.DELETE("/v1/aclassbclassuses/:id", DeleteAclassBclassUse)

		v1.GET("/v1/bclasss", GetBclasss)
		v1.GET("/v1/bclasss/:id", GetBclass)
		v1.POST("/v1/bclasss", PostBclass)
		v1.PATCH("/v1/bclasss/:id", UpdateBclass)
		v1.PUT("/v1/bclasss/:id", UpdateBclass)
		v1.DELETE("/v1/bclasss/:id", DeleteBclass)

		v1.GET("/v1/dclasss", GetDclasss)
		v1.GET("/v1/dclasss/:id", GetDclass)
		v1.POST("/v1/dclasss", PostDclass)
		v1.PATCH("/v1/dclasss/:id", UpdateDclass)
		v1.PUT("/v1/dclasss/:id", UpdateDclass)
		v1.DELETE("/v1/dclasss/:id", DeleteDclass)

		v1.GET("/commitnb", GetLastCommitNb)
		v1.GET("/pushfromfrontnb", GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitnb backrepo GetLastCommitNb
func GetLastCommitNb(c *gin.Context) {
	res := orm.GetLastCommitNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func GetLastPushFromFrontNb(c *gin.Context) {
	res := orm.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
