// generated code - do not edit
package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gong/lib/split/go/orm"

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
	v1 := r.Group("/api/github.com/fullstack-lang/gong/lib/split/go")
	{ // insertion point for registrations
		v1.GET("/v1/assplits", GetController().GetAsSplits)
		v1.GET("/v1/assplits/:id", GetController().GetAsSplit)
		v1.POST("/v1/assplits", GetController().PostAsSplit)
		v1.PATCH("/v1/assplits/:id", GetController().UpdateAsSplit)
		v1.PUT("/v1/assplits/:id", GetController().UpdateAsSplit)
		v1.DELETE("/v1/assplits/:id", GetController().DeleteAsSplit)

		v1.GET("/v1/assplitareas", GetController().GetAsSplitAreas)
		v1.GET("/v1/assplitareas/:id", GetController().GetAsSplitArea)
		v1.POST("/v1/assplitareas", GetController().PostAsSplitArea)
		v1.PATCH("/v1/assplitareas/:id", GetController().UpdateAsSplitArea)
		v1.PUT("/v1/assplitareas/:id", GetController().UpdateAsSplitArea)
		v1.DELETE("/v1/assplitareas/:id", GetController().DeleteAsSplitArea)

		v1.GET("/v1/buttons", GetController().GetButtons)
		v1.GET("/v1/buttons/:id", GetController().GetButton)
		v1.POST("/v1/buttons", GetController().PostButton)
		v1.PATCH("/v1/buttons/:id", GetController().UpdateButton)
		v1.PUT("/v1/buttons/:id", GetController().UpdateButton)
		v1.DELETE("/v1/buttons/:id", GetController().DeleteButton)

		v1.GET("/v1/cursors", GetController().GetCursors)
		v1.GET("/v1/cursors/:id", GetController().GetCursor)
		v1.POST("/v1/cursors", GetController().PostCursor)
		v1.PATCH("/v1/cursors/:id", GetController().UpdateCursor)
		v1.PUT("/v1/cursors/:id", GetController().UpdateCursor)
		v1.DELETE("/v1/cursors/:id", GetController().DeleteCursor)

		v1.GET("/v1/docs", GetController().GetDocs)
		v1.GET("/v1/docs/:id", GetController().GetDoc)
		v1.POST("/v1/docs", GetController().PostDoc)
		v1.PATCH("/v1/docs/:id", GetController().UpdateDoc)
		v1.PUT("/v1/docs/:id", GetController().UpdateDoc)
		v1.DELETE("/v1/docs/:id", GetController().DeleteDoc)

		v1.GET("/v1/forms", GetController().GetForms)
		v1.GET("/v1/forms/:id", GetController().GetForm)
		v1.POST("/v1/forms", GetController().PostForm)
		v1.PATCH("/v1/forms/:id", GetController().UpdateForm)
		v1.PUT("/v1/forms/:id", GetController().UpdateForm)
		v1.DELETE("/v1/forms/:id", GetController().DeleteForm)

		v1.GET("/v1/loads", GetController().GetLoads)
		v1.GET("/v1/loads/:id", GetController().GetLoad)
		v1.POST("/v1/loads", GetController().PostLoad)
		v1.PATCH("/v1/loads/:id", GetController().UpdateLoad)
		v1.PUT("/v1/loads/:id", GetController().UpdateLoad)
		v1.DELETE("/v1/loads/:id", GetController().DeleteLoad)

		v1.GET("/v1/sliders", GetController().GetSliders)
		v1.GET("/v1/sliders/:id", GetController().GetSlider)
		v1.POST("/v1/sliders", GetController().PostSlider)
		v1.PATCH("/v1/sliders/:id", GetController().UpdateSlider)
		v1.PUT("/v1/sliders/:id", GetController().UpdateSlider)
		v1.DELETE("/v1/sliders/:id", GetController().DeleteSlider)

		v1.GET("/v1/splits", GetController().GetSplits)
		v1.GET("/v1/splits/:id", GetController().GetSplit)
		v1.POST("/v1/splits", GetController().PostSplit)
		v1.PATCH("/v1/splits/:id", GetController().UpdateSplit)
		v1.PUT("/v1/splits/:id", GetController().UpdateSplit)
		v1.DELETE("/v1/splits/:id", GetController().DeleteSplit)

		v1.GET("/v1/svgs", GetController().GetSvgs)
		v1.GET("/v1/svgs/:id", GetController().GetSvg)
		v1.POST("/v1/svgs", GetController().PostSvg)
		v1.PATCH("/v1/svgs/:id", GetController().UpdateSvg)
		v1.PUT("/v1/svgs/:id", GetController().UpdateSvg)
		v1.DELETE("/v1/svgs/:id", GetController().DeleteSvg)

		v1.GET("/v1/tables", GetController().GetTables)
		v1.GET("/v1/tables/:id", GetController().GetTable)
		v1.POST("/v1/tables", GetController().PostTable)
		v1.PATCH("/v1/tables/:id", GetController().UpdateTable)
		v1.PUT("/v1/tables/:id", GetController().UpdateTable)
		v1.DELETE("/v1/tables/:id", GetController().DeleteTable)

		v1.GET("/v1/tones", GetController().GetTones)
		v1.GET("/v1/tones/:id", GetController().GetTone)
		v1.POST("/v1/tones", GetController().PostTone)
		v1.PATCH("/v1/tones/:id", GetController().UpdateTone)
		v1.PUT("/v1/tones/:id", GetController().UpdateTone)
		v1.DELETE("/v1/tones/:id", GetController().DeleteTone)

		v1.GET("/v1/trees", GetController().GetTrees)
		v1.GET("/v1/trees/:id", GetController().GetTree)
		v1.POST("/v1/trees", GetController().PostTree)
		v1.PATCH("/v1/trees/:id", GetController().UpdateTree)
		v1.PUT("/v1/trees/:id", GetController().UpdateTree)
		v1.DELETE("/v1/trees/:id", GetController().DeleteTree)

		v1.GET("/v1/views", GetController().GetViews)
		v1.GET("/v1/views/:id", GetController().GetView)
		v1.POST("/v1/views", GetController().PostView)
		v1.PATCH("/v1/views/:id", GetController().UpdateView)
		v1.PUT("/v1/views/:id", GetController().UpdateView)
		v1.DELETE("/v1/views/:id", GetController().DeleteView)

		v1.GET("/v1/xlsxs", GetController().GetXlsxs)
		v1.GET("/v1/xlsxs/:id", GetController().GetXlsx)
		v1.POST("/v1/xlsxs", GetController().PostXlsx)
		v1.PATCH("/v1/xlsxs/:id", GetController().UpdateXlsx)
		v1.PUT("/v1/xlsxs/:id", GetController().UpdateXlsx)
		v1.DELETE("/v1/xlsxs/:id", GetController().DeleteXlsx)

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

	// log.Println("Stack github.com/fullstack-lang/gong/lib/split/go, onWebSocketRequestForBackRepoContent")

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

	log.Printf("Stack github.com/fullstack-lang/gong/lib/split/go: stack path: '%s', new ws index %d",
		stackPath, controller.listenerIndex,
	)
	index := controller.listenerIndex
	controller.listenerIndex++

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "WebSocket Connect, Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
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
				log.Println("github.com/fullstack-lang/gong/lib/split/go", stackPath, "WS client disconnected:", err)
				cancel() // Cancel the context
				return
			}
		}
	}()

	backRepoData := new(orm.BackRepoData)
	orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)
	backRepoData.GONG__Index = index

	err = wsConnection.WriteJSON(backRepoData)
	if err != nil {
		log.Println("github.com/fullstack-lang/gong/lib/split/go:\n",
			"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	} else {
		log.Println(time.Now().Format("2006-01-02 15:04:05.000000"), "github.com/fullstack-lang/gong/lib/split/go: 1st sent backRepoData of stack:", stackPath, "index", index)
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
				backRepoData.GONG__Index = index

				// Set write deadline to prevent blocking indefinitely
				wsConnection.SetWriteDeadline(time.Now().Add(10 * time.Second))

				// Send backRepo data
				err = wsConnection.WriteJSON(backRepoData)
				if err != nil {
					log.Println("github.com/fullstack-lang/gong/lib/split/go:\n", stackPath,
						"client no longer receiver web socket message,closing websocket handler")
					fmt.Println(err)
					cancel() // Cancel the context
					return
				} else {
					log.Println(time.Now().Format("2006-01-02 15:04:05.000000"), "github.com/fullstack-lang/gong/lib/split/go: sent backRepoData of stack:", stackPath, "index", index)
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
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
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
		message := "GET Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"
		
		message += "Availabe stack names are:\n"
		for k, _ := range controller.Map_BackRepos {
			message += k + "\n"
		}
			
		log.Panic(message)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
