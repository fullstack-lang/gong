// generated code - do not edit
package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gongdoc/go/orm"

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

	// log.Println("Stack github.com/fullstack-lang/gongdoc/go, onWebSocketRequestForBackRepoContent")

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

	log.Printf("Stack github.com/fullstack-lang/gongdoc/go: stack path: '%s', new ws index %d",
		stackPath, controller.listenerIndex,
	)
	controller.listenerIndex++

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go, Unkown stack", stackPath)
	}
	updateCommitBackRepoNbChannel := backRepo.SubscribeToCommitNb(ctx)

	// Start a goroutine to read from the WebSocket to detect disconnection
	go func() {
		for {
			// ReadMessage is used to detect client disconnection
			_, _, err := wsConnection.ReadMessage()
			if err != nil {
				log.Println("github.com/fullstack-lang/gongdoc/go", stackPath, "WS client disconnected:", err)
				cancel() // Cancel the context
				return
			}
		}
	}()

	backRepoData := new(orm.BackRepoData)
	orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)

	err = wsConnection.WriteJSON(backRepoData)
	if err != nil {
		log.Println("github.com/fullstack-lang/gongdoc/go:\n",
			"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	} else {
		log.Println(time.Now().Format(time.RFC3339Nano), "github.com/fullstack-lang/gongdoc/go: 1st sent backRepoData of stack:", stackPath)
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
					log.Println("github.com/fullstack-lang/gongdoc/go:\n", stackPath,
						"client no longer receiver web socket message,closing websocket handler")
					fmt.Println(err)
					cancel() // Cancel the context
					return
				} else {
					log.Println(time.Now().Format(time.RFC3339Nano), "github.com/fullstack-lang/gongdoc/go: sent backRepoData of stack:", stackPath)
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
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
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
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
