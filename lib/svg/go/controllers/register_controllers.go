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

	"github.com/fullstack-lang/gong/lib/svg/go/orm"

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
	v1 := r.Group("/api/github.com/fullstack-lang/gong/lib/svg/go")
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

		v1.GET("/v1/svgtexts", GetController().GetSvgTexts)
		v1.GET("/v1/svgtexts/:id", GetController().GetSvgText)
		v1.POST("/v1/svgtexts", GetController().PostSvgText)
		v1.PATCH("/v1/svgtexts/:id", GetController().UpdateSvgText)
		v1.PUT("/v1/svgtexts/:id", GetController().UpdateSvgText)
		v1.DELETE("/v1/svgtexts/:id", GetController().DeleteSvgText)

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

	// log.Println("Stack github.com/fullstack-lang/gong/lib/svg/go, onWebSocketRequestForBackRepoContent")

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
		message := "WebSocket Connect, Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

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
				log.Println("github.com/fullstack-lang/gong/lib/svg/go", stackPath, "WS client disconnected:", err)
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
		log.Println("github.com/fullstack-lang/gong/lib/svg/go:\n",
			"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	} else {
		// 1. Extract the component name from the long path for cleaner logs
		// For example, "github.com/fullstack-lang/gong/lib/table/go" becomes "table"
		parts := strings.Split("github.com/fullstack-lang/gong/lib/svg/go", "/") // Assuming goFilePath holds the path
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
					log.Println("github.com/fullstack-lang/gong/lib/svg/go:\n",
						"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
					fmt.Println(err)
					return
				} else {
					// 1. Extract the component name from the long path for cleaner logs
					// For example, "github.com/fullstack-lang/gong/lib/table/go" becomes "table"
					parts := strings.Split("github.com/fullstack-lang/gong/lib/svg/go", "/") // Assuming goFilePath holds the path
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
		message := "Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

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
		message := "GET Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
