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
		v1.GET("/v1/astructs", GetAstructs)
		v1.GET("/v1/astructs/:id", GetAstruct)
		v1.POST("/v1/astructs", PostAstruct)
		v1.PATCH("/v1/astructs/:id", UpdateAstruct)
		v1.PUT("/v1/astructs/:id", UpdateAstruct)
		v1.DELETE("/v1/astructs/:id", DeleteAstruct)

		v1.GET("/v1/astructbstruct2uses", GetAstructBstruct2Uses)
		v1.GET("/v1/astructbstruct2uses/:id", GetAstructBstruct2Use)
		v1.POST("/v1/astructbstruct2uses", PostAstructBstruct2Use)
		v1.PATCH("/v1/astructbstruct2uses/:id", UpdateAstructBstruct2Use)
		v1.PUT("/v1/astructbstruct2uses/:id", UpdateAstructBstruct2Use)
		v1.DELETE("/v1/astructbstruct2uses/:id", DeleteAstructBstruct2Use)

		v1.GET("/v1/astructbstructuses", GetAstructBstructUses)
		v1.GET("/v1/astructbstructuses/:id", GetAstructBstructUse)
		v1.POST("/v1/astructbstructuses", PostAstructBstructUse)
		v1.PATCH("/v1/astructbstructuses/:id", UpdateAstructBstructUse)
		v1.PUT("/v1/astructbstructuses/:id", UpdateAstructBstructUse)
		v1.DELETE("/v1/astructbstructuses/:id", DeleteAstructBstructUse)

		v1.GET("/v1/bstructs", GetBstructs)
		v1.GET("/v1/bstructs/:id", GetBstruct)
		v1.POST("/v1/bstructs", PostBstruct)
		v1.PATCH("/v1/bstructs/:id", UpdateBstruct)
		v1.PUT("/v1/bstructs/:id", UpdateBstruct)
		v1.DELETE("/v1/bstructs/:id", DeleteBstruct)

		v1.GET("/v1/dstructs", GetDstructs)
		v1.GET("/v1/dstructs/:id", GetDstruct)
		v1.POST("/v1/dstructs", PostDstruct)
		v1.PATCH("/v1/dstructs/:id", UpdateDstruct)
		v1.PUT("/v1/dstructs/:id", UpdateDstruct)
		v1.DELETE("/v1/dstructs/:id", DeleteDstruct)

		v1.GET("/commitfrombacknb", GetLastCommitFromBackNb)
		v1.GET("/pushfromfrontnb", GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func GetLastCommitFromBackNb(c *gin.Context) {
	res := orm.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func GetLastPushFromFrontNb(c *gin.Context) {
	res := orm.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
