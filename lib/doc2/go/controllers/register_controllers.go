// generated code - do not edit
package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/orm"

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
	v1 := r.Group("/api/github.com/fullstack-lang/gong/lib/doc2/go")
	{ // insertion point for registrations
		v1.GET("/v1/attributeshapes", GetController().GetAttributeShapes)
		v1.GET("/v1/attributeshapes/:id", GetController().GetAttributeShape)
		v1.POST("/v1/attributeshapes", GetController().PostAttributeShape)
		v1.PATCH("/v1/attributeshapes/:id", GetController().UpdateAttributeShape)
		v1.PUT("/v1/attributeshapes/:id", GetController().UpdateAttributeShape)
		v1.DELETE("/v1/attributeshapes/:id", GetController().DeleteAttributeShape)

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

		v1.GET("/v1/gongenumshapes", GetController().GetGongEnumShapes)
		v1.GET("/v1/gongenumshapes/:id", GetController().GetGongEnumShape)
		v1.POST("/v1/gongenumshapes", GetController().PostGongEnumShape)
		v1.PATCH("/v1/gongenumshapes/:id", GetController().UpdateGongEnumShape)
		v1.PUT("/v1/gongenumshapes/:id", GetController().UpdateGongEnumShape)
		v1.DELETE("/v1/gongenumshapes/:id", GetController().DeleteGongEnumShape)

		v1.GET("/v1/gongenumvalueshapes", GetController().GetGongEnumValueShapes)
		v1.GET("/v1/gongenumvalueshapes/:id", GetController().GetGongEnumValueShape)
		v1.POST("/v1/gongenumvalueshapes", GetController().PostGongEnumValueShape)
		v1.PATCH("/v1/gongenumvalueshapes/:id", GetController().UpdateGongEnumValueShape)
		v1.PUT("/v1/gongenumvalueshapes/:id", GetController().UpdateGongEnumValueShape)
		v1.DELETE("/v1/gongenumvalueshapes/:id", GetController().DeleteGongEnumValueShape)

		v1.GET("/v1/gongnotelinkshapes", GetController().GetGongNoteLinkShapes)
		v1.GET("/v1/gongnotelinkshapes/:id", GetController().GetGongNoteLinkShape)
		v1.POST("/v1/gongnotelinkshapes", GetController().PostGongNoteLinkShape)
		v1.PATCH("/v1/gongnotelinkshapes/:id", GetController().UpdateGongNoteLinkShape)
		v1.PUT("/v1/gongnotelinkshapes/:id", GetController().UpdateGongNoteLinkShape)
		v1.DELETE("/v1/gongnotelinkshapes/:id", GetController().DeleteGongNoteLinkShape)

		v1.GET("/v1/gongnoteshapes", GetController().GetGongNoteShapes)
		v1.GET("/v1/gongnoteshapes/:id", GetController().GetGongNoteShape)
		v1.POST("/v1/gongnoteshapes", GetController().PostGongNoteShape)
		v1.PATCH("/v1/gongnoteshapes/:id", GetController().UpdateGongNoteShape)
		v1.PUT("/v1/gongnoteshapes/:id", GetController().UpdateGongNoteShape)
		v1.DELETE("/v1/gongnoteshapes/:id", GetController().DeleteGongNoteShape)

		v1.GET("/v1/gongstructshapes", GetController().GetGongStructShapes)
		v1.GET("/v1/gongstructshapes/:id", GetController().GetGongStructShape)
		v1.POST("/v1/gongstructshapes", GetController().PostGongStructShape)
		v1.PATCH("/v1/gongstructshapes/:id", GetController().UpdateGongStructShape)
		v1.PUT("/v1/gongstructshapes/:id", GetController().UpdateGongStructShape)
		v1.DELETE("/v1/gongstructshapes/:id", GetController().DeleteGongStructShape)

		v1.GET("/v1/linkshapes", GetController().GetLinkShapes)
		v1.GET("/v1/linkshapes/:id", GetController().GetLinkShape)
		v1.POST("/v1/linkshapes", GetController().PostLinkShape)
		v1.PATCH("/v1/linkshapes/:id", GetController().UpdateLinkShape)
		v1.PUT("/v1/linkshapes/:id", GetController().UpdateLinkShape)
		v1.DELETE("/v1/linkshapes/:id", GetController().DeleteLinkShape)

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

	// log.Println("Stack github.com/fullstack-lang/gong/lib/doc2/go, onWebSocketRequestForBackRepoContent")

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
		value := values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "Name", stackPath)
		}
	}

	index := controller.listenerIndex
	controller.listenerIndex++
	log.Printf(
		"%s github.com/fullstack-lang/gong/lib/doc2/go: Con: '%s', index %d",
		time.Now().Format("2006-01-02 15:04:05.000000"),
		stackPath, index,
	)

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "WebSocket Connect, Stack github.com/fullstack-lang/gong/lib/doc2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	updateCommitBackRepoNbChannel := backRepo.SubscribeToCommitNb(ctx)

	// Start a goroutine to read from the WebSocket to detect disconnection
	go func() {
		for {
			// ReadMessage is used to detect client disconnection
			_, _, err := wsConnection.ReadMessage()
			if err != nil {
				log.Println("github.com/fullstack-lang/gong/lib/doc2/go", stackPath, "WS client disconnected:", err)
				cancel() // Cancel the context
				return
			}
		}
	}()

	backRepoData := new(orm.BackRepoData)
	orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)
	backRepoData.GONG__Index = index
	
	refresh := 0
	err = wsConnection.WriteJSON(backRepoData)
	if err != nil {
		log.Println("github.com/fullstack-lang/gong/lib/doc2/go:\n",
			"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	} else {
	log.Printf(
		"%s github.com/fullstack-lang/gong/lib/doc2/go: %03d: '%s', index %d",
		time.Now().Format("2006-01-02 15:04:05.000000"),
		refresh,
		stackPath, index,
	)
	}
	for {
		select {
		case <-ctx.Done():
			// Context canceled, exit the loop
			return
		default:
			for nbCommitBackRepo := range updateCommitBackRepoNbChannel {
				_ = nbCommitBackRepo
				refresh += 1

				backRepoData := new(orm.BackRepoData)
				orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)
				backRepoData.GONG__Index = index

				// Set write deadline to prevent blocking indefinitely
				wsConnection.SetWriteDeadline(time.Now().Add(10 * time.Second))

				// Send backRepo data
				err = wsConnection.WriteJSON(backRepoData)
				if err != nil {
					log.Println("github.com/fullstack-lang/gong/lib/doc2/go:\n", stackPath,
						"client no longer receiver web socket message,closing websocket handler")
					fmt.Println(err)
					cancel() // Cancel the context
					return
				} else {
					log.Printf(
						"%s github.com/fullstack-lang/gong/lib/doc2/go: %03d: '%s', index %d",
						time.Now().Format("2006-01-02 15:04:05.000000"),
						refresh,
						stackPath, index,
					)
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
		value := values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/doc2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func (controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastPushFromFrontNb", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/doc2/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
