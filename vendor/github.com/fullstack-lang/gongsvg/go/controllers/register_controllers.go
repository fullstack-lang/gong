// generated code - do not edit
package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gongsvg/go/orm"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
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

		v1.GET("/v1/rectanchoredpaths", GetController().GetRectAnchoredPaths)
		v1.GET("/v1/rectanchoredpaths/:id", GetController().GetRectAnchoredPath)
		v1.POST("/v1/rectanchoredpaths", GetController().PostRectAnchoredPath)
		v1.PATCH("/v1/rectanchoredpaths/:id", GetController().UpdateRectAnchoredPath)
		v1.PUT("/v1/rectanchoredpaths/:id", GetController().UpdateRectAnchoredPath)
		v1.DELETE("/v1/rectanchoredpaths/:id", GetController().DeleteRectAnchoredPath)

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

		v1.GET("/v1/ws/stage", GetController().onWebSocketRequestForBackRepoContent)

		v1.GET("/v1/stacks", GetController().stacks)
	}
}

func (controller *Controller) stacks(c *gin.Context) {

	var res []string

	for k := range controller.Map_BackRepos {
		res = append(res, k)
	}

	c.JSON(http.StatusOK, res)
}

// onWebSocketRequestForBackRepoContent is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequestForBackRepoContent(c *gin.Context) {

	// log.Println("Stack github.com/fullstack-lang/gongsvg/go, onWebSocketRequestForBackRepoContent")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://localhost:8080" || origin == "http://localhost:4200"
		},
	}

	wsConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wsConnection.Close()

	// Create a context that is canceled when the connection is closed
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}

	log.Printf("Stack github.com/fullstack-lang/gongsvg/go: stack path: '%s', new ws index %d",
		stackPath, controller.listenerIndex,
	)
	controller.listenerIndex++

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go, Unkown stack", stackPath)
	}
	updateCommitBackRepoNbChannel := backRepo.SubscribeToCommitNb(ctx)

	// Start a goroutine to read from the WebSocket to detect disconnection
	go func() {
		for {
			// ReadMessage is used to detect client disconnection
			_, _, err := wsConnection.ReadMessage()
			if err != nil {
				log.Println("github.com/fullstack-lang/gongsvg/go", stackPath, "WS client disconnected:", err)
				cancel() // Cancel the context
				return
			}
		}
	}()

	backRepoData := new(orm.BackRepoData)
	orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)

	err = wsConnection.WriteJSON(backRepoData)
	if err != nil {
		log.Println("github.com/fullstack-lang/gongsvg/go:\n",
			"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	} else {
		log.Println(time.Now().Format(time.RFC3339Nano), "github.com/fullstack-lang/gongsvg/go: 1st sent backRepoData of stack:", stackPath)
	}
	for {
		select {
		case <-ctx.Done():
			// Context canceled, exit the loop
			return
		default:
			for nbCommitBackRepo := range updateCommitBackRepoNbChannel {
				_ = nbCommitBackRepo

				backRepoData := new(orm.BackRepoData)
				orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)

				// Set write deadline to prevent blocking indefinitely
				wsConnection.SetWriteDeadline(time.Now().Add(10 * time.Second))

				// Send backRepo data
				err = wsConnection.WriteJSON(backRepoData)
				if err != nil {
					log.Println("github.com/fullstack-lang/gongsvg/go:\n", stackPath,
						"client no longer receiver web socket message,closing websocket handler")
					fmt.Println(err)
					cancel() // Cancel the context
					return
				} else {
					log.Println(time.Now().Format(time.RFC3339Nano), "github.com/fullstack-lang/gongsvg/go: sent backRepoData of stack:", stackPath)
				}
			}
		}
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
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
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
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
