package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gong/go/orm"
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
	v1 := r.Group("/api/github.com/fullstack-lang/gong/go")
	{ // insertion point for registrations
		v1.GET("/v1/gongbasicfields", GetGongBasicFields)
		v1.GET("/v1/gongbasicfields/:id", GetGongBasicField)
		v1.POST("/v1/gongbasicfields", PostGongBasicField)
		v1.PATCH("/v1/gongbasicfields/:id", UpdateGongBasicField)
		v1.PUT("/v1/gongbasicfields/:id", UpdateGongBasicField)
		v1.DELETE("/v1/gongbasicfields/:id", DeleteGongBasicField)

		v1.GET("/v1/gongenums", GetGongEnums)
		v1.GET("/v1/gongenums/:id", GetGongEnum)
		v1.POST("/v1/gongenums", PostGongEnum)
		v1.PATCH("/v1/gongenums/:id", UpdateGongEnum)
		v1.PUT("/v1/gongenums/:id", UpdateGongEnum)
		v1.DELETE("/v1/gongenums/:id", DeleteGongEnum)

		v1.GET("/v1/gongenumvalues", GetGongEnumValues)
		v1.GET("/v1/gongenumvalues/:id", GetGongEnumValue)
		v1.POST("/v1/gongenumvalues", PostGongEnumValue)
		v1.PATCH("/v1/gongenumvalues/:id", UpdateGongEnumValue)
		v1.PUT("/v1/gongenumvalues/:id", UpdateGongEnumValue)
		v1.DELETE("/v1/gongenumvalues/:id", DeleteGongEnumValue)

		v1.GET("/v1/gongnotes", GetGongNotes)
		v1.GET("/v1/gongnotes/:id", GetGongNote)
		v1.POST("/v1/gongnotes", PostGongNote)
		v1.PATCH("/v1/gongnotes/:id", UpdateGongNote)
		v1.PUT("/v1/gongnotes/:id", UpdateGongNote)
		v1.DELETE("/v1/gongnotes/:id", DeleteGongNote)

		v1.GET("/v1/gongstructs", GetGongStructs)
		v1.GET("/v1/gongstructs/:id", GetGongStruct)
		v1.POST("/v1/gongstructs", PostGongStruct)
		v1.PATCH("/v1/gongstructs/:id", UpdateGongStruct)
		v1.PUT("/v1/gongstructs/:id", UpdateGongStruct)
		v1.DELETE("/v1/gongstructs/:id", DeleteGongStruct)

		v1.GET("/v1/gongtimefields", GetGongTimeFields)
		v1.GET("/v1/gongtimefields/:id", GetGongTimeField)
		v1.POST("/v1/gongtimefields", PostGongTimeField)
		v1.PATCH("/v1/gongtimefields/:id", UpdateGongTimeField)
		v1.PUT("/v1/gongtimefields/:id", UpdateGongTimeField)
		v1.DELETE("/v1/gongtimefields/:id", DeleteGongTimeField)

		v1.GET("/v1/modelpkgs", GetModelPkgs)
		v1.GET("/v1/modelpkgs/:id", GetModelPkg)
		v1.POST("/v1/modelpkgs", PostModelPkg)
		v1.PATCH("/v1/modelpkgs/:id", UpdateModelPkg)
		v1.PUT("/v1/modelpkgs/:id", UpdateModelPkg)
		v1.DELETE("/v1/modelpkgs/:id", DeleteModelPkg)

		v1.GET("/v1/pointertogongstructfields", GetPointerToGongStructFields)
		v1.GET("/v1/pointertogongstructfields/:id", GetPointerToGongStructField)
		v1.POST("/v1/pointertogongstructfields", PostPointerToGongStructField)
		v1.PATCH("/v1/pointertogongstructfields/:id", UpdatePointerToGongStructField)
		v1.PUT("/v1/pointertogongstructfields/:id", UpdatePointerToGongStructField)
		v1.DELETE("/v1/pointertogongstructfields/:id", DeletePointerToGongStructField)

		v1.GET("/v1/sliceofpointertogongstructfields", GetSliceOfPointerToGongStructFields)
		v1.GET("/v1/sliceofpointertogongstructfields/:id", GetSliceOfPointerToGongStructField)
		v1.POST("/v1/sliceofpointertogongstructfields", PostSliceOfPointerToGongStructField)
		v1.PATCH("/v1/sliceofpointertogongstructfields/:id", UpdateSliceOfPointerToGongStructField)
		v1.PUT("/v1/sliceofpointertogongstructfields/:id", UpdateSliceOfPointerToGongStructField)
		v1.DELETE("/v1/sliceofpointertogongstructfields/:id", DeleteSliceOfPointerToGongStructField)

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
