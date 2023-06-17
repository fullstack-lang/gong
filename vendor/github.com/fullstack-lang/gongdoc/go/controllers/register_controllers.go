package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

// registerControllers register controllers
func registerControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gongdoc/go")
	{ // insertion point for registrations
		v1.GET("/v1/classdiagrams", GetController().GetClassdiagrams)
		v1.GET("/v1/classdiagrams/:id", GetController().GetClassdiagram)
		v1.POST("/v1/classdiagrams", GetController().PostClassdiagram)
		v1.PATCH("/v1/classdiagrams/:id", GetController().UpdateClassdiagram)
		v1.PUT("/v1/classdiagrams/:id", GetController().UpdateClassdiagram)
		v1.DELETE("/v1/classdiagrams/:id", GetController().DeleteClassdiagram)

		v1.GET("/v1/diagrampackages", GetController().GetDiagramPackages)
		v1.GET("/v1/diagrampackages/:id", GetController().GetDiagramPackage)
		v1.POST("/v1/diagrampackages", GetController().PostDiagramPackage)
		v1.PATCH("/v1/diagrampackages/:id", GetController().UpdateDiagramPackage)
		v1.PUT("/v1/diagrampackages/:id", GetController().UpdateDiagramPackage)
		v1.DELETE("/v1/diagrampackages/:id", GetController().DeleteDiagramPackage)

		v1.GET("/v1/fields", GetController().GetFields)
		v1.GET("/v1/fields/:id", GetController().GetField)
		v1.POST("/v1/fields", GetController().PostField)
		v1.PATCH("/v1/fields/:id", GetController().UpdateField)
		v1.PUT("/v1/fields/:id", GetController().UpdateField)
		v1.DELETE("/v1/fields/:id", GetController().DeleteField)

		v1.GET("/v1/gongenumshapes", GetController().GetGongEnumShapes)
		v1.GET("/v1/gongenumshapes/:id", GetController().GetGongEnumShape)
		v1.POST("/v1/gongenumshapes", GetController().PostGongEnumShape)
		v1.PATCH("/v1/gongenumshapes/:id", GetController().UpdateGongEnumShape)
		v1.PUT("/v1/gongenumshapes/:id", GetController().UpdateGongEnumShape)
		v1.DELETE("/v1/gongenumshapes/:id", GetController().DeleteGongEnumShape)

		v1.GET("/v1/gongenumvalueentrys", GetController().GetGongEnumValueEntrys)
		v1.GET("/v1/gongenumvalueentrys/:id", GetController().GetGongEnumValueEntry)
		v1.POST("/v1/gongenumvalueentrys", GetController().PostGongEnumValueEntry)
		v1.PATCH("/v1/gongenumvalueentrys/:id", GetController().UpdateGongEnumValueEntry)
		v1.PUT("/v1/gongenumvalueentrys/:id", GetController().UpdateGongEnumValueEntry)
		v1.DELETE("/v1/gongenumvalueentrys/:id", GetController().DeleteGongEnumValueEntry)

		v1.GET("/v1/gongstructshapes", GetController().GetGongStructShapes)
		v1.GET("/v1/gongstructshapes/:id", GetController().GetGongStructShape)
		v1.POST("/v1/gongstructshapes", GetController().PostGongStructShape)
		v1.PATCH("/v1/gongstructshapes/:id", GetController().UpdateGongStructShape)
		v1.PUT("/v1/gongstructshapes/:id", GetController().UpdateGongStructShape)
		v1.DELETE("/v1/gongstructshapes/:id", GetController().DeleteGongStructShape)

		v1.GET("/v1/links", GetController().GetLinks)
		v1.GET("/v1/links/:id", GetController().GetLink)
		v1.POST("/v1/links", GetController().PostLink)
		v1.PATCH("/v1/links/:id", GetController().UpdateLink)
		v1.PUT("/v1/links/:id", GetController().UpdateLink)
		v1.DELETE("/v1/links/:id", GetController().DeleteLink)

		v1.GET("/v1/noteshapes", GetController().GetNoteShapes)
		v1.GET("/v1/noteshapes/:id", GetController().GetNoteShape)
		v1.POST("/v1/noteshapes", GetController().PostNoteShape)
		v1.PATCH("/v1/noteshapes/:id", GetController().UpdateNoteShape)
		v1.PUT("/v1/noteshapes/:id", GetController().UpdateNoteShape)
		v1.DELETE("/v1/noteshapes/:id", GetController().DeleteNoteShape)

		v1.GET("/v1/noteshapelinks", GetController().GetNoteShapeLinks)
		v1.GET("/v1/noteshapelinks/:id", GetController().GetNoteShapeLink)
		v1.POST("/v1/noteshapelinks", GetController().PostNoteShapeLink)
		v1.PATCH("/v1/noteshapelinks/:id", GetController().UpdateNoteShapeLink)
		v1.PUT("/v1/noteshapelinks/:id", GetController().UpdateNoteShapeLink)
		v1.DELETE("/v1/noteshapelinks/:id", GetController().DeleteNoteShapeLink)

		v1.GET("/v1/positions", GetController().GetPositions)
		v1.GET("/v1/positions/:id", GetController().GetPosition)
		v1.POST("/v1/positions", GetController().PostPosition)
		v1.PATCH("/v1/positions/:id", GetController().UpdatePosition)
		v1.PUT("/v1/positions/:id", GetController().UpdatePosition)
		v1.DELETE("/v1/positions/:id", GetController().DeletePosition)

		v1.GET("/v1/umlstates", GetController().GetUmlStates)
		v1.GET("/v1/umlstates/:id", GetController().GetUmlState)
		v1.POST("/v1/umlstates", GetController().PostUmlState)
		v1.PATCH("/v1/umlstates/:id", GetController().UpdateUmlState)
		v1.PUT("/v1/umlstates/:id", GetController().UpdateUmlState)
		v1.DELETE("/v1/umlstates/:id", GetController().DeleteUmlState)

		v1.GET("/v1/umlscs", GetController().GetUmlscs)
		v1.GET("/v1/umlscs/:id", GetController().GetUmlsc)
		v1.POST("/v1/umlscs", GetController().PostUmlsc)
		v1.PATCH("/v1/umlscs/:id", GetController().UpdateUmlsc)
		v1.PUT("/v1/umlscs/:id", GetController().UpdateUmlsc)
		v1.DELETE("/v1/umlscs/:id", GetController().DeleteUmlsc)

		v1.GET("/v1/vertices", GetController().GetVertices)
		v1.GET("/v1/vertices/:id", GetController().GetVertice)
		v1.POST("/v1/vertices", GetController().PostVertice)
		v1.PATCH("/v1/vertices/:id", GetController().UpdateVertice)
		v1.PUT("/v1/vertices/:id", GetController().UpdateVertice)
		v1.DELETE("/v1/vertices/:id", GetController().DeleteVertice)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func (controller *Controller) GetLastCommitFromBackNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func (controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastPushFromFrontNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
