package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongdoc/go/orm"
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
	v1 := r.Group("/api/github.com/fullstack-lang/gongdoc/go")
	{ // insertion point for registrations
		v1.GET("/v1/classdiagrams", GetClassdiagrams)
		v1.GET("/v1/classdiagrams/:id", GetClassdiagram)
		v1.POST("/v1/classdiagrams", PostClassdiagram)
		v1.PATCH("/v1/classdiagrams/:id", UpdateClassdiagram)
		v1.PUT("/v1/classdiagrams/:id", UpdateClassdiagram)
		v1.DELETE("/v1/classdiagrams/:id", DeleteClassdiagram)

		v1.GET("/v1/classshapes", GetClassshapes)
		v1.GET("/v1/classshapes/:id", GetClassshape)
		v1.POST("/v1/classshapes", PostClassshape)
		v1.PATCH("/v1/classshapes/:id", UpdateClassshape)
		v1.PUT("/v1/classshapes/:id", UpdateClassshape)
		v1.DELETE("/v1/classshapes/:id", DeleteClassshape)

		v1.GET("/v1/fields", GetFields)
		v1.GET("/v1/fields/:id", GetField)
		v1.POST("/v1/fields", PostField)
		v1.PATCH("/v1/fields/:id", UpdateField)
		v1.PUT("/v1/fields/:id", UpdateField)
		v1.DELETE("/v1/fields/:id", DeleteField)

		v1.GET("/v1/gongstructs", GetGongStructs)
		v1.GET("/v1/gongstructs/:id", GetGongStruct)
		v1.POST("/v1/gongstructs", PostGongStruct)
		v1.PATCH("/v1/gongstructs/:id", UpdateGongStruct)
		v1.PUT("/v1/gongstructs/:id", UpdateGongStruct)
		v1.DELETE("/v1/gongstructs/:id", DeleteGongStruct)

		v1.GET("/v1/gongdoccommands", GetGongdocCommands)
		v1.GET("/v1/gongdoccommands/:id", GetGongdocCommand)
		v1.POST("/v1/gongdoccommands", PostGongdocCommand)
		v1.PATCH("/v1/gongdoccommands/:id", UpdateGongdocCommand)
		v1.PUT("/v1/gongdoccommands/:id", UpdateGongdocCommand)
		v1.DELETE("/v1/gongdoccommands/:id", DeleteGongdocCommand)

		v1.GET("/v1/gongdocstatuss", GetGongdocStatuss)
		v1.GET("/v1/gongdocstatuss/:id", GetGongdocStatus)
		v1.POST("/v1/gongdocstatuss", PostGongdocStatus)
		v1.PATCH("/v1/gongdocstatuss/:id", UpdateGongdocStatus)
		v1.PUT("/v1/gongdocstatuss/:id", UpdateGongdocStatus)
		v1.DELETE("/v1/gongdocstatuss/:id", DeleteGongdocStatus)

		v1.GET("/v1/links", GetLinks)
		v1.GET("/v1/links/:id", GetLink)
		v1.POST("/v1/links", PostLink)
		v1.PATCH("/v1/links/:id", UpdateLink)
		v1.PUT("/v1/links/:id", UpdateLink)
		v1.DELETE("/v1/links/:id", DeleteLink)

		v1.GET("/v1/notes", GetNotes)
		v1.GET("/v1/notes/:id", GetNote)
		v1.POST("/v1/notes", PostNote)
		v1.PATCH("/v1/notes/:id", UpdateNote)
		v1.PUT("/v1/notes/:id", UpdateNote)
		v1.DELETE("/v1/notes/:id", DeleteNote)

		v1.GET("/v1/pkgelts", GetPkgelts)
		v1.GET("/v1/pkgelts/:id", GetPkgelt)
		v1.POST("/v1/pkgelts", PostPkgelt)
		v1.PATCH("/v1/pkgelts/:id", UpdatePkgelt)
		v1.PUT("/v1/pkgelts/:id", UpdatePkgelt)
		v1.DELETE("/v1/pkgelts/:id", DeletePkgelt)

		v1.GET("/v1/positions", GetPositions)
		v1.GET("/v1/positions/:id", GetPosition)
		v1.POST("/v1/positions", PostPosition)
		v1.PATCH("/v1/positions/:id", UpdatePosition)
		v1.PUT("/v1/positions/:id", UpdatePosition)
		v1.DELETE("/v1/positions/:id", DeletePosition)

		v1.GET("/v1/umlstates", GetUmlStates)
		v1.GET("/v1/umlstates/:id", GetUmlState)
		v1.POST("/v1/umlstates", PostUmlState)
		v1.PATCH("/v1/umlstates/:id", UpdateUmlState)
		v1.PUT("/v1/umlstates/:id", UpdateUmlState)
		v1.DELETE("/v1/umlstates/:id", DeleteUmlState)

		v1.GET("/v1/umlscs", GetUmlscs)
		v1.GET("/v1/umlscs/:id", GetUmlsc)
		v1.POST("/v1/umlscs", PostUmlsc)
		v1.PATCH("/v1/umlscs/:id", UpdateUmlsc)
		v1.PUT("/v1/umlscs/:id", UpdateUmlsc)
		v1.DELETE("/v1/umlscs/:id", DeleteUmlsc)

		v1.GET("/v1/vertices", GetVertices)
		v1.GET("/v1/vertices/:id", GetVertice)
		v1.POST("/v1/vertices", PostVertice)
		v1.PATCH("/v1/vertices/:id", UpdateVertice)
		v1.PUT("/v1/vertices/:id", UpdateVertice)
		v1.DELETE("/v1/vertices/:id", DeleteVertice)

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
