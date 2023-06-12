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
	v1 := r.Group("/api/github.com/fullstack-lang/gongsvg/go")
	{ // insertion point for registrations
		v1.GET("/v1/animates", GetController().GetAnimates)
		v1.GET("/v1/animates/:id", GetController().GetAnimate)
		v1.POST("/v1/animates", GetController().PostAnimate)
		v1.PATCH("/v1/animates/:id", GetController().UpdateAnimate)
		v1.PUT("/v1/animates/:id", GetController().UpdateAnimate)
		v1.DELETE("/v1/animates/:id", GetController().DeleteAnimate)

		v1.GET("/v1/circles", GetController().GetCircles)
		v1.GET("/v1/circles/:id", GetController().GetCircle)
		v1.POST("/v1/circles", GetController().PostCircle)
		v1.PATCH("/v1/circles/:id", GetController().UpdateCircle)
		v1.PUT("/v1/circles/:id", GetController().UpdateCircle)
		v1.DELETE("/v1/circles/:id", GetController().DeleteCircle)

		v1.GET("/v1/ellipses", GetController().GetEllipses)
		v1.GET("/v1/ellipses/:id", GetController().GetEllipse)
		v1.POST("/v1/ellipses", GetController().PostEllipse)
		v1.PATCH("/v1/ellipses/:id", GetController().UpdateEllipse)
		v1.PUT("/v1/ellipses/:id", GetController().UpdateEllipse)
		v1.DELETE("/v1/ellipses/:id", GetController().DeleteEllipse)

		v1.GET("/v1/layers", GetController().GetLayers)
		v1.GET("/v1/layers/:id", GetController().GetLayer)
		v1.POST("/v1/layers", GetController().PostLayer)
		v1.PATCH("/v1/layers/:id", GetController().UpdateLayer)
		v1.PUT("/v1/layers/:id", GetController().UpdateLayer)
		v1.DELETE("/v1/layers/:id", GetController().DeleteLayer)

		v1.GET("/v1/lines", GetController().GetLines)
		v1.GET("/v1/lines/:id", GetController().GetLine)
		v1.POST("/v1/lines", GetController().PostLine)
		v1.PATCH("/v1/lines/:id", GetController().UpdateLine)
		v1.PUT("/v1/lines/:id", GetController().UpdateLine)
		v1.DELETE("/v1/lines/:id", GetController().DeleteLine)

		v1.GET("/v1/links", GetController().GetLinks)
		v1.GET("/v1/links/:id", GetController().GetLink)
		v1.POST("/v1/links", GetController().PostLink)
		v1.PATCH("/v1/links/:id", GetController().UpdateLink)
		v1.PUT("/v1/links/:id", GetController().UpdateLink)
		v1.DELETE("/v1/links/:id", GetController().DeleteLink)

		v1.GET("/v1/linkanchoredtexts", GetController().GetLinkAnchoredTexts)
		v1.GET("/v1/linkanchoredtexts/:id", GetController().GetLinkAnchoredText)
		v1.POST("/v1/linkanchoredtexts", GetController().PostLinkAnchoredText)
		v1.PATCH("/v1/linkanchoredtexts/:id", GetController().UpdateLinkAnchoredText)
		v1.PUT("/v1/linkanchoredtexts/:id", GetController().UpdateLinkAnchoredText)
		v1.DELETE("/v1/linkanchoredtexts/:id", GetController().DeleteLinkAnchoredText)

		v1.GET("/v1/paths", GetController().GetPaths)
		v1.GET("/v1/paths/:id", GetController().GetPath)
		v1.POST("/v1/paths", GetController().PostPath)
		v1.PATCH("/v1/paths/:id", GetController().UpdatePath)
		v1.PUT("/v1/paths/:id", GetController().UpdatePath)
		v1.DELETE("/v1/paths/:id", GetController().DeletePath)

		v1.GET("/v1/points", GetController().GetPoints)
		v1.GET("/v1/points/:id", GetController().GetPoint)
		v1.POST("/v1/points", GetController().PostPoint)
		v1.PATCH("/v1/points/:id", GetController().UpdatePoint)
		v1.PUT("/v1/points/:id", GetController().UpdatePoint)
		v1.DELETE("/v1/points/:id", GetController().DeletePoint)

		v1.GET("/v1/polygones", GetController().GetPolygones)
		v1.GET("/v1/polygones/:id", GetController().GetPolygone)
		v1.POST("/v1/polygones", GetController().PostPolygone)
		v1.PATCH("/v1/polygones/:id", GetController().UpdatePolygone)
		v1.PUT("/v1/polygones/:id", GetController().UpdatePolygone)
		v1.DELETE("/v1/polygones/:id", GetController().DeletePolygone)

		v1.GET("/v1/polylines", GetController().GetPolylines)
		v1.GET("/v1/polylines/:id", GetController().GetPolyline)
		v1.POST("/v1/polylines", GetController().PostPolyline)
		v1.PATCH("/v1/polylines/:id", GetController().UpdatePolyline)
		v1.PUT("/v1/polylines/:id", GetController().UpdatePolyline)
		v1.DELETE("/v1/polylines/:id", GetController().DeletePolyline)

		v1.GET("/v1/rects", GetController().GetRects)
		v1.GET("/v1/rects/:id", GetController().GetRect)
		v1.POST("/v1/rects", GetController().PostRect)
		v1.PATCH("/v1/rects/:id", GetController().UpdateRect)
		v1.PUT("/v1/rects/:id", GetController().UpdateRect)
		v1.DELETE("/v1/rects/:id", GetController().DeleteRect)

		v1.GET("/v1/rectanchoredrects", GetController().GetRectAnchoredRects)
		v1.GET("/v1/rectanchoredrects/:id", GetController().GetRectAnchoredRect)
		v1.POST("/v1/rectanchoredrects", GetController().PostRectAnchoredRect)
		v1.PATCH("/v1/rectanchoredrects/:id", GetController().UpdateRectAnchoredRect)
		v1.PUT("/v1/rectanchoredrects/:id", GetController().UpdateRectAnchoredRect)
		v1.DELETE("/v1/rectanchoredrects/:id", GetController().DeleteRectAnchoredRect)

		v1.GET("/v1/rectanchoredtexts", GetController().GetRectAnchoredTexts)
		v1.GET("/v1/rectanchoredtexts/:id", GetController().GetRectAnchoredText)
		v1.POST("/v1/rectanchoredtexts", GetController().PostRectAnchoredText)
		v1.PATCH("/v1/rectanchoredtexts/:id", GetController().UpdateRectAnchoredText)
		v1.PUT("/v1/rectanchoredtexts/:id", GetController().UpdateRectAnchoredText)
		v1.DELETE("/v1/rectanchoredtexts/:id", GetController().DeleteRectAnchoredText)

		v1.GET("/v1/rectlinklinks", GetController().GetRectLinkLinks)
		v1.GET("/v1/rectlinklinks/:id", GetController().GetRectLinkLink)
		v1.POST("/v1/rectlinklinks", GetController().PostRectLinkLink)
		v1.PATCH("/v1/rectlinklinks/:id", GetController().UpdateRectLinkLink)
		v1.PUT("/v1/rectlinklinks/:id", GetController().UpdateRectLinkLink)
		v1.DELETE("/v1/rectlinklinks/:id", GetController().DeleteRectLinkLink)

		v1.GET("/v1/svgs", GetController().GetSVGs)
		v1.GET("/v1/svgs/:id", GetController().GetSVG)
		v1.POST("/v1/svgs", GetController().PostSVG)
		v1.PATCH("/v1/svgs/:id", GetController().UpdateSVG)
		v1.PUT("/v1/svgs/:id", GetController().UpdateSVG)
		v1.DELETE("/v1/svgs/:id", GetController().DeleteSVG)

		v1.GET("/v1/texts", GetController().GetTexts)
		v1.GET("/v1/texts/:id", GetController().GetText)
		v1.POST("/v1/texts", GetController().PostText)
		v1.PATCH("/v1/texts/:id", GetController().UpdateText)
		v1.PUT("/v1/texts/:id", GetController().UpdateText)
		v1.DELETE("/v1/texts/:id", GetController().DeleteText)

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
