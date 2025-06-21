// generated code - do not edit
package controllers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

		v1.GET("/v1/favicons", GetController().GetFavIcons)
		v1.GET("/v1/favicons/:id", GetController().GetFavIcon)
		v1.POST("/v1/favicons", GetController().PostFavIcon)
		v1.PATCH("/v1/favicons/:id", GetController().UpdateFavIcon)
		v1.PUT("/v1/favicons/:id", GetController().UpdateFavIcon)
		v1.DELETE("/v1/favicons/:id", GetController().DeleteFavIcon)

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

		v1.GET("/v1/logoonthelefts", GetController().GetLogoOnTheLefts)
		v1.GET("/v1/logoonthelefts/:id", GetController().GetLogoOnTheLeft)
		v1.POST("/v1/logoonthelefts", GetController().PostLogoOnTheLeft)
		v1.PATCH("/v1/logoonthelefts/:id", GetController().UpdateLogoOnTheLeft)
		v1.PUT("/v1/logoonthelefts/:id", GetController().UpdateLogoOnTheLeft)
		v1.DELETE("/v1/logoonthelefts/:id", GetController().DeleteLogoOnTheLeft)

		v1.GET("/v1/logoontherights", GetController().GetLogoOnTheRights)
		v1.GET("/v1/logoontherights/:id", GetController().GetLogoOnTheRight)
		v1.POST("/v1/logoontherights", GetController().PostLogoOnTheRight)
		v1.PATCH("/v1/logoontherights/:id", GetController().UpdateLogoOnTheRight)
		v1.PUT("/v1/logoontherights/:id", GetController().UpdateLogoOnTheRight)
		v1.DELETE("/v1/logoontherights/:id", GetController().DeleteLogoOnTheRight)

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

		v1.GET("/v1/titles", GetController().GetTitles)
		v1.GET("/v1/titles/:id", GetController().GetTitle)
		v1.POST("/v1/titles", GetController().PostTitle)
		v1.PATCH("/v1/titles/:id", GetController().UpdateTitle)
		v1.PUT("/v1/titles/:id", GetController().UpdateTitle)
		v1.DELETE("/v1/titles/:id", GetController().DeleteTitle)

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

			// log.Printf("CheckOrigin: Accepted - Origin '%s' with port %d", origin, port)
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

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "WebSocket Connect, Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

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
				log.Println("github.com/fullstack-lang/gong/lib/split/go", stackPath, "WS client disconnected:", err)
				cancel() // Cancel the context
				return
			}
		}
	}()

	backRepoData := new(orm.BackRepoData)
	orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)
	backRepoData.GONG__Index = index

	refresh := 0
	// Marshal the data to JSON first to be able to get its size
	jsonData, err := json.Marshal(backRepoData)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return
	}

	// Get the size of the JSON data in bytes
	jsonSize := len(jsonData)

    // Calculate the full SHA-256 hash
    fullHash := sha256.Sum256(jsonData)

    // Use the first 12 characters for a shorter, yet highly unique, signature
    shortHash := hex.EncodeToString(fullHash[:])[0:12]

	// Use WriteMessage to send the pre-marshaled JSON data.
	// websocket.TextMessage is typically what WriteJSON uses.
	err = wsConnection.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		log.Println("github.com/fullstack-lang/gong/lib/split/go:\n",
			"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	} else {
		// 1. Extract the component name from the long path for cleaner logs
		// For example, "github.com/fullstack-lang/gong/lib/table/go" becomes "table"
		parts := strings.Split("github.com/fullstack-lang/gong/lib/split/go", "/") // Assuming goFilePath holds the path
		component := "unknown"
		if len(parts) > 2 {
			component = parts[len(parts)-2]
		}

		// 2. Use a single, formatted log line
		log.Printf(
			"%-12s | %-85s | Idx: %d | Size: %-9s | Hash: %s",
			component,
			stackPath,
			index,
			formatBytes(jsonSize),
			shortHash,
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
				// Marshal the data to JSON first to be able to get its size
				jsonData, err := json.Marshal(backRepoData)
				if err != nil {
					log.Printf("Error marshaling JSON: %v", err)
					return
				}

				// Get the size of the JSON data in bytes
				jsonSize := len(jsonData)

				// Calculate the full SHA-256 hash
				fullHash := sha256.Sum256(jsonData)

				// Use the first 12 characters for a shorter, yet highly unique, signature
				shortHash := hex.EncodeToString(fullHash[:])[0:12]

				// Use WriteMessage to send the pre-marshaled JSON data.
				// websocket.TextMessage is typically what WriteJSON uses.
				err = wsConnection.WriteMessage(websocket.TextMessage, jsonData)
				if err != nil {
					log.Println("github.com/fullstack-lang/gong/lib/split/go:\n",
						"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
					fmt.Println(err)
					return
				} else {
					// 1. Extract the component name from the long path for cleaner logs
					// For example, "github.com/fullstack-lang/gong/lib/table/go" becomes "table"
					parts := strings.Split("github.com/fullstack-lang/gong/lib/split/go", "/") // Assuming goFilePath holds the path
					component := "unknown"
					if len(parts) > 2 {
						component = parts[len(parts)-2]
					}

					// 2. Use a single, formatted log line
					log.Printf(
						"%-12s | %-85s | Idx: %d | Size: %-9s | Hash: %s",
						component,
						stackPath,
						index,
						formatBytes(jsonSize),
						shortHash,
					)
				}
			}
		}
	}
}

// formatBytes converts a size in bytes to a human-readable string (KB, MB, GB).
func formatBytes(size int) string {
    if size < 1024 {
        return fmt.Sprintf("%d B", size)
    }
    sizeInKB := float64(size) / 1024.0
    if sizeInKB < 1024.0 {
        // For KB, show one decimal place if it's not a whole number
        if math.Mod(sizeInKB, 1.0) == 0 {
            return fmt.Sprintf("%.0f KB", sizeInKB)
        }
        return fmt.Sprintf("%.1f KB", sizeInKB)
    }
    sizeInMB := sizeInKB / 1024.0
    return fmt.Sprintf("%.2f MB", sizeInMB)
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
		message := "GET Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
