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

		v1.GET("/v1/diagrampackages", GetDiagramPackages)
		v1.GET("/v1/diagrampackages/:id", GetDiagramPackage)
		v1.POST("/v1/diagrampackages", PostDiagramPackage)
		v1.PATCH("/v1/diagrampackages/:id", UpdateDiagramPackage)
		v1.PUT("/v1/diagrampackages/:id", UpdateDiagramPackage)
		v1.DELETE("/v1/diagrampackages/:id", DeleteDiagramPackage)

		v1.GET("/v1/fields", GetFields)
		v1.GET("/v1/fields/:id", GetField)
		v1.POST("/v1/fields", PostField)
		v1.PATCH("/v1/fields/:id", UpdateField)
		v1.PUT("/v1/fields/:id", UpdateField)
		v1.DELETE("/v1/fields/:id", DeleteField)

		v1.GET("/v1/links", GetLinks)
		v1.GET("/v1/links/:id", GetLink)
		v1.POST("/v1/links", PostLink)
		v1.PATCH("/v1/links/:id", UpdateLink)
		v1.PUT("/v1/links/:id", UpdateLink)
		v1.DELETE("/v1/links/:id", DeleteLink)

		v1.GET("/v1/nodes", GetNodes)
		v1.GET("/v1/nodes/:id", GetNode)
		v1.POST("/v1/nodes", PostNode)
		v1.PATCH("/v1/nodes/:id", UpdateNode)
		v1.PUT("/v1/nodes/:id", UpdateNode)
		v1.DELETE("/v1/nodes/:id", DeleteNode)

		v1.GET("/v1/notelinks", GetNoteLinks)
		v1.GET("/v1/notelinks/:id", GetNoteLink)
		v1.POST("/v1/notelinks", PostNoteLink)
		v1.PATCH("/v1/notelinks/:id", UpdateNoteLink)
		v1.PUT("/v1/notelinks/:id", UpdateNoteLink)
		v1.DELETE("/v1/notelinks/:id", DeleteNoteLink)

		v1.GET("/v1/noteshapes", GetNoteShapes)
		v1.GET("/v1/noteshapes/:id", GetNoteShape)
		v1.POST("/v1/noteshapes", PostNoteShape)
		v1.PATCH("/v1/noteshapes/:id", UpdateNoteShape)
		v1.PUT("/v1/noteshapes/:id", UpdateNoteShape)
		v1.DELETE("/v1/noteshapes/:id", DeleteNoteShape)

		v1.GET("/v1/positions", GetPositions)
		v1.GET("/v1/positions/:id", GetPosition)
		v1.POST("/v1/positions", PostPosition)
		v1.PATCH("/v1/positions/:id", UpdatePosition)
		v1.PUT("/v1/positions/:id", UpdatePosition)
		v1.DELETE("/v1/positions/:id", DeletePosition)

		v1.GET("/v1/references", GetReferences)
		v1.GET("/v1/references/:id", GetReference)
		v1.POST("/v1/references", PostReference)
		v1.PATCH("/v1/references/:id", UpdateReference)
		v1.PUT("/v1/references/:id", UpdateReference)
		v1.DELETE("/v1/references/:id", DeleteReference)

		v1.GET("/v1/trees", GetTrees)
		v1.GET("/v1/trees/:id", GetTree)
		v1.POST("/v1/trees", PostTree)
		v1.PATCH("/v1/trees/:id", UpdateTree)
		v1.PUT("/v1/trees/:id", UpdateTree)
		v1.DELETE("/v1/trees/:id", DeleteTree)

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
