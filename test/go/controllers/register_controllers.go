// generated code - do not edit
package controllers

import (
	"fmt"
	"log"
	"net/http"

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
	v1 := r.Group("/api/github.com/fullstack-lang/gong/test/go")
	{ // insertion point for registrations
		v1.GET("/v1/astructs", GetController().GetAstructs)
		v1.GET("/v1/astructs/:id", GetController().GetAstruct)
		v1.POST("/v1/astructs", GetController().PostAstruct)
		v1.PATCH("/v1/astructs/:id", GetController().UpdateAstruct)
		v1.PUT("/v1/astructs/:id", GetController().UpdateAstruct)
		v1.DELETE("/v1/astructs/:id", GetController().DeleteAstruct)

		v1.GET("/v1/astructbstruct2uses", GetController().GetAstructBstruct2Uses)
		v1.GET("/v1/astructbstruct2uses/:id", GetController().GetAstructBstruct2Use)
		v1.POST("/v1/astructbstruct2uses", GetController().PostAstructBstruct2Use)
		v1.PATCH("/v1/astructbstruct2uses/:id", GetController().UpdateAstructBstruct2Use)
		v1.PUT("/v1/astructbstruct2uses/:id", GetController().UpdateAstructBstruct2Use)
		v1.DELETE("/v1/astructbstruct2uses/:id", GetController().DeleteAstructBstruct2Use)

		v1.GET("/v1/astructbstructuses", GetController().GetAstructBstructUses)
		v1.GET("/v1/astructbstructuses/:id", GetController().GetAstructBstructUse)
		v1.POST("/v1/astructbstructuses", GetController().PostAstructBstructUse)
		v1.PATCH("/v1/astructbstructuses/:id", GetController().UpdateAstructBstructUse)
		v1.PUT("/v1/astructbstructuses/:id", GetController().UpdateAstructBstructUse)
		v1.DELETE("/v1/astructbstructuses/:id", GetController().DeleteAstructBstructUse)

		v1.GET("/v1/bstructs", GetController().GetBstructs)
		v1.GET("/v1/bstructs/:id", GetController().GetBstruct)
		v1.POST("/v1/bstructs", GetController().PostBstruct)
		v1.PATCH("/v1/bstructs/:id", GetController().UpdateBstruct)
		v1.PUT("/v1/bstructs/:id", GetController().UpdateBstruct)
		v1.DELETE("/v1/bstructs/:id", GetController().DeleteBstruct)

		v1.GET("/v1/dstructs", GetController().GetDstructs)
		v1.GET("/v1/dstructs/:id", GetController().GetDstruct)
		v1.POST("/v1/dstructs", GetController().PostDstruct)
		v1.PATCH("/v1/dstructs/:id", GetController().UpdateDstruct)
		v1.PUT("/v1/dstructs/:id", GetController().UpdateDstruct)
		v1.DELETE("/v1/dstructs/:id", GetController().DeleteDstruct)

		v1.GET("/v1/fstructs", GetController().GetFstructs)
		v1.GET("/v1/fstructs/:id", GetController().GetFstruct)
		v1.POST("/v1/fstructs", GetController().PostFstruct)
		v1.PATCH("/v1/fstructs/:id", GetController().UpdateFstruct)
		v1.PUT("/v1/fstructs/:id", GetController().UpdateFstruct)
		v1.DELETE("/v1/fstructs/:id", GetController().DeleteFstruct)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)
	}

	r.GET("/api/github.com/fullstack-lang/gong/test/go/v1/ws", GetController().onWebSocketRequest)

}

// onWebSocketRequest is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequest(c *gin.Context) {

	log.Println("onWebSocketRequest")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://localhost:8080"
		},
	}

	wsConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wsConnection.Close()

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
		log.Panic("Stack github.com/fullstack-lang/gong/test/go/models, Unkown stack", stackPath)
	}
	updateCommitBackRepoNbChannel := backRepo.Subscribe()

	for nbCommitBackRepo := range updateCommitBackRepoNbChannel {

		// Send elapsed time as a string over the WebSocket connection
		err = wsConnection.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d", nbCommitBackRepo)))
		if err != nil {
			fmt.Println(err)
			return
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
		log.Panic("Stack github.com/fullstack-lang/gong/test/go/models, Unkown stack", stackPath)
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
		log.Panic("Stack github.com/fullstack-lang/gong/test/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
