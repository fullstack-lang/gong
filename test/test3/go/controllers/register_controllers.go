// generated code - do not edit
package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/fullstack-lang/gong/test/test3/go/orm"

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
	v1 := r.Group("/api/github.com/fullstack-lang/gong/test/test3/go")
	{ // insertion point for registrations
		v1.GET("/v1/as", GetController().GetAs)
		v1.GET("/v1/as/:id", GetController().GetA)
		v1.POST("/v1/as", GetController().PostA)
		v1.PATCH("/v1/as/:id", GetController().UpdateA)
		v1.PUT("/v1/as/:id", GetController().UpdateA)
		v1.DELETE("/v1/as/:id", GetController().DeleteA)

		v1.GET("/v1/bs", GetController().GetBs)
		v1.GET("/v1/bs/:id", GetController().GetB)
		v1.POST("/v1/bs", GetController().PostB)
		v1.PATCH("/v1/bs/:id", GetController().UpdateB)
		v1.PUT("/v1/bs/:id", GetController().UpdateB)
		v1.DELETE("/v1/bs/:id", GetController().DeleteB)

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

	// log.Println("Stack github.com/fullstack-lang/gong/test/test3/go, onWebSocketRequestForBackRepoContent")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.

	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			if origin == "" {
				log.Printf("CheckOrigin: Rejected - Origin header is empty. Request from: %s", r.RemoteAddr)
				return false // Or handle as per your security policy
			}

			u, err := url.Parse(origin)
			if err != nil {
				log.Printf("CheckOrigin: Rejected - Invalid Origin URL '%s'. Error: %v", origin, err)
				return false // Invalid URL
			}

			portStr := u.Port()

			if portStr == "" {
				// If no port is specified, it might be using default HTTP/HTTPS ports.
				// For this specific request, we'll assume a port must be present.
				log.Printf("CheckOrigin: Rejected - No port specified in Origin URL '%s'", origin)
				return false
			}

			port, err := strconv.Atoi(portStr)
			if err != nil {
				log.Printf("CheckOrigin: Rejected - Port '%s' in Origin URL '%s' is not a valid number. Error: %v", portStr, origin, err)
				return false // Port is not a valid number
			}

			// Check if the port is 4200 OR in the range 8000-9000
			allowed := port == 4200 || (port >= 8000 && port <= 9000)
			if !allowed {
				log.Printf("CheckOrigin: Rejected - Port %d from Origin '%s' is not in the allowed list (4200 or 8000-9000)", port, origin)
				return false
			}

			log.Printf("CheckOrigin: Accepted - Origin '%s' with port %d", origin, port)
			return true
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
		"%s github.com/fullstack-lang/gong/test/test3/go: Con: '%s', index %d",
		time.Now().Format("2006-01-02 15:04:05.000000"),
		stackPath, index,
	)

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "WebSocket Connect, Stack github.com/fullstack-lang/gong/test/test3/go, Unkown stack: \"" + stackPath + "\"\n"

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
				log.Println("github.com/fullstack-lang/gong/test/test3/go", stackPath, "WS client disconnected:", err)
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
		log.Println("github.com/fullstack-lang/gong/test/test3/go:\n",
			"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	} else {
	log.Printf(
		"%s github.com/fullstack-lang/gong/test/test3/go: %03d: '%s', index %d",
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
					log.Println("github.com/fullstack-lang/gong/test/test3/go:\n", stackPath,
						"client no longer receiver web socket message,closing websocket handler")
					fmt.Println(err)
					cancel() // Cancel the context
					return
				} else {
					log.Printf(
						"%s github.com/fullstack-lang/gong/test/test3/go: %03d: '%s', index %d",
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
		message := "Stack github.com/fullstack-lang/gong/test/test3/go, Unkown stack: \"" + stackPath + "\"\n"

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
		message := "GET Stack github.com/fullstack-lang/gong/test/test3/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
