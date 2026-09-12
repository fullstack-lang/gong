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
	"strings"
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/orm"

	"github.com/gorilla/websocket"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

var _ = genQuery

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
func registerControllers(mux *http.ServeMux) {
	base := "/api/github.com/fullstack-lang/gong/lib/svg/go/v1"

	mux.HandleFunc("GET " + base + "/animates", GetController().GetAnimates)
	mux.HandleFunc("GET " + base + "/animates/{id}", GetController().GetAnimate)
	mux.HandleFunc("POST " + base + "/animates", GetController().PostAnimate)
	mux.HandleFunc("PATCH " + base + "/animates/{id}", GetController().UpdateAnimate)
	mux.HandleFunc("PUT " + base + "/animates/{id}", GetController().UpdateAnimate)
	mux.HandleFunc("DELETE " + base + "/animates/{id}", GetController().DeleteAnimate)

	mux.HandleFunc("GET " + base + "/circles", GetController().GetCircles)
	mux.HandleFunc("GET " + base + "/circles/{id}", GetController().GetCircle)
	mux.HandleFunc("POST " + base + "/circles", GetController().PostCircle)
	mux.HandleFunc("PATCH " + base + "/circles/{id}", GetController().UpdateCircle)
	mux.HandleFunc("PUT " + base + "/circles/{id}", GetController().UpdateCircle)
	mux.HandleFunc("DELETE " + base + "/circles/{id}", GetController().DeleteCircle)

	mux.HandleFunc("GET " + base + "/conditions", GetController().GetConditions)
	mux.HandleFunc("GET " + base + "/conditions/{id}", GetController().GetCondition)
	mux.HandleFunc("POST " + base + "/conditions", GetController().PostCondition)
	mux.HandleFunc("PATCH " + base + "/conditions/{id}", GetController().UpdateCondition)
	mux.HandleFunc("PUT " + base + "/conditions/{id}", GetController().UpdateCondition)
	mux.HandleFunc("DELETE " + base + "/conditions/{id}", GetController().DeleteCondition)

	mux.HandleFunc("GET " + base + "/controlpoints", GetController().GetControlPoints)
	mux.HandleFunc("GET " + base + "/controlpoints/{id}", GetController().GetControlPoint)
	mux.HandleFunc("POST " + base + "/controlpoints", GetController().PostControlPoint)
	mux.HandleFunc("PATCH " + base + "/controlpoints/{id}", GetController().UpdateControlPoint)
	mux.HandleFunc("PUT " + base + "/controlpoints/{id}", GetController().UpdateControlPoint)
	mux.HandleFunc("DELETE " + base + "/controlpoints/{id}", GetController().DeleteControlPoint)

	mux.HandleFunc("GET " + base + "/ellipses", GetController().GetEllipses)
	mux.HandleFunc("GET " + base + "/ellipses/{id}", GetController().GetEllipse)
	mux.HandleFunc("POST " + base + "/ellipses", GetController().PostEllipse)
	mux.HandleFunc("PATCH " + base + "/ellipses/{id}", GetController().UpdateEllipse)
	mux.HandleFunc("PUT " + base + "/ellipses/{id}", GetController().UpdateEllipse)
	mux.HandleFunc("DELETE " + base + "/ellipses/{id}", GetController().DeleteEllipse)

	mux.HandleFunc("GET " + base + "/filetodownloads", GetController().GetFileToDownloads)
	mux.HandleFunc("GET " + base + "/filetodownloads/{id}", GetController().GetFileToDownload)
	mux.HandleFunc("POST " + base + "/filetodownloads", GetController().PostFileToDownload)
	mux.HandleFunc("PATCH " + base + "/filetodownloads/{id}", GetController().UpdateFileToDownload)
	mux.HandleFunc("PUT " + base + "/filetodownloads/{id}", GetController().UpdateFileToDownload)
	mux.HandleFunc("DELETE " + base + "/filetodownloads/{id}", GetController().DeleteFileToDownload)

	mux.HandleFunc("GET " + base + "/layers", GetController().GetLayers)
	mux.HandleFunc("GET " + base + "/layers/{id}", GetController().GetLayer)
	mux.HandleFunc("POST " + base + "/layers", GetController().PostLayer)
	mux.HandleFunc("PATCH " + base + "/layers/{id}", GetController().UpdateLayer)
	mux.HandleFunc("PUT " + base + "/layers/{id}", GetController().UpdateLayer)
	mux.HandleFunc("DELETE " + base + "/layers/{id}", GetController().DeleteLayer)

	mux.HandleFunc("GET " + base + "/lines", GetController().GetLines)
	mux.HandleFunc("GET " + base + "/lines/{id}", GetController().GetLine)
	mux.HandleFunc("POST " + base + "/lines", GetController().PostLine)
	mux.HandleFunc("PATCH " + base + "/lines/{id}", GetController().UpdateLine)
	mux.HandleFunc("PUT " + base + "/lines/{id}", GetController().UpdateLine)
	mux.HandleFunc("DELETE " + base + "/lines/{id}", GetController().DeleteLine)

	mux.HandleFunc("GET " + base + "/links", GetController().GetLinks)
	mux.HandleFunc("GET " + base + "/links/{id}", GetController().GetLink)
	mux.HandleFunc("POST " + base + "/links", GetController().PostLink)
	mux.HandleFunc("PATCH " + base + "/links/{id}", GetController().UpdateLink)
	mux.HandleFunc("PUT " + base + "/links/{id}", GetController().UpdateLink)
	mux.HandleFunc("DELETE " + base + "/links/{id}", GetController().DeleteLink)

	mux.HandleFunc("GET " + base + "/linkanchoredpaths", GetController().GetLinkAnchoredPaths)
	mux.HandleFunc("GET " + base + "/linkanchoredpaths/{id}", GetController().GetLinkAnchoredPath)
	mux.HandleFunc("POST " + base + "/linkanchoredpaths", GetController().PostLinkAnchoredPath)
	mux.HandleFunc("PATCH " + base + "/linkanchoredpaths/{id}", GetController().UpdateLinkAnchoredPath)
	mux.HandleFunc("PUT " + base + "/linkanchoredpaths/{id}", GetController().UpdateLinkAnchoredPath)
	mux.HandleFunc("DELETE " + base + "/linkanchoredpaths/{id}", GetController().DeleteLinkAnchoredPath)

	mux.HandleFunc("GET " + base + "/linkanchoredtexts", GetController().GetLinkAnchoredTexts)
	mux.HandleFunc("GET " + base + "/linkanchoredtexts/{id}", GetController().GetLinkAnchoredText)
	mux.HandleFunc("POST " + base + "/linkanchoredtexts", GetController().PostLinkAnchoredText)
	mux.HandleFunc("PATCH " + base + "/linkanchoredtexts/{id}", GetController().UpdateLinkAnchoredText)
	mux.HandleFunc("PUT " + base + "/linkanchoredtexts/{id}", GetController().UpdateLinkAnchoredText)
	mux.HandleFunc("DELETE " + base + "/linkanchoredtexts/{id}", GetController().DeleteLinkAnchoredText)

	mux.HandleFunc("GET " + base + "/paths", GetController().GetPaths)
	mux.HandleFunc("GET " + base + "/paths/{id}", GetController().GetPath)
	mux.HandleFunc("POST " + base + "/paths", GetController().PostPath)
	mux.HandleFunc("PATCH " + base + "/paths/{id}", GetController().UpdatePath)
	mux.HandleFunc("PUT " + base + "/paths/{id}", GetController().UpdatePath)
	mux.HandleFunc("DELETE " + base + "/paths/{id}", GetController().DeletePath)

	mux.HandleFunc("GET " + base + "/points", GetController().GetPoints)
	mux.HandleFunc("GET " + base + "/points/{id}", GetController().GetPoint)
	mux.HandleFunc("POST " + base + "/points", GetController().PostPoint)
	mux.HandleFunc("PATCH " + base + "/points/{id}", GetController().UpdatePoint)
	mux.HandleFunc("PUT " + base + "/points/{id}", GetController().UpdatePoint)
	mux.HandleFunc("DELETE " + base + "/points/{id}", GetController().DeletePoint)

	mux.HandleFunc("GET " + base + "/polygones", GetController().GetPolygones)
	mux.HandleFunc("GET " + base + "/polygones/{id}", GetController().GetPolygone)
	mux.HandleFunc("POST " + base + "/polygones", GetController().PostPolygone)
	mux.HandleFunc("PATCH " + base + "/polygones/{id}", GetController().UpdatePolygone)
	mux.HandleFunc("PUT " + base + "/polygones/{id}", GetController().UpdatePolygone)
	mux.HandleFunc("DELETE " + base + "/polygones/{id}", GetController().DeletePolygone)

	mux.HandleFunc("GET " + base + "/polylines", GetController().GetPolylines)
	mux.HandleFunc("GET " + base + "/polylines/{id}", GetController().GetPolyline)
	mux.HandleFunc("POST " + base + "/polylines", GetController().PostPolyline)
	mux.HandleFunc("PATCH " + base + "/polylines/{id}", GetController().UpdatePolyline)
	mux.HandleFunc("PUT " + base + "/polylines/{id}", GetController().UpdatePolyline)
	mux.HandleFunc("DELETE " + base + "/polylines/{id}", GetController().DeletePolyline)

	mux.HandleFunc("GET " + base + "/rects", GetController().GetRects)
	mux.HandleFunc("GET " + base + "/rects/{id}", GetController().GetRect)
	mux.HandleFunc("POST " + base + "/rects", GetController().PostRect)
	mux.HandleFunc("PATCH " + base + "/rects/{id}", GetController().UpdateRect)
	mux.HandleFunc("PUT " + base + "/rects/{id}", GetController().UpdateRect)
	mux.HandleFunc("DELETE " + base + "/rects/{id}", GetController().DeleteRect)

	mux.HandleFunc("GET " + base + "/rectanchoredpaths", GetController().GetRectAnchoredPaths)
	mux.HandleFunc("GET " + base + "/rectanchoredpaths/{id}", GetController().GetRectAnchoredPath)
	mux.HandleFunc("POST " + base + "/rectanchoredpaths", GetController().PostRectAnchoredPath)
	mux.HandleFunc("PATCH " + base + "/rectanchoredpaths/{id}", GetController().UpdateRectAnchoredPath)
	mux.HandleFunc("PUT " + base + "/rectanchoredpaths/{id}", GetController().UpdateRectAnchoredPath)
	mux.HandleFunc("DELETE " + base + "/rectanchoredpaths/{id}", GetController().DeleteRectAnchoredPath)

	mux.HandleFunc("GET " + base + "/rectanchoredpngimages", GetController().GetRectAnchoredPngImages)
	mux.HandleFunc("GET " + base + "/rectanchoredpngimages/{id}", GetController().GetRectAnchoredPngImage)
	mux.HandleFunc("POST " + base + "/rectanchoredpngimages", GetController().PostRectAnchoredPngImage)
	mux.HandleFunc("PATCH " + base + "/rectanchoredpngimages/{id}", GetController().UpdateRectAnchoredPngImage)
	mux.HandleFunc("PUT " + base + "/rectanchoredpngimages/{id}", GetController().UpdateRectAnchoredPngImage)
	mux.HandleFunc("DELETE " + base + "/rectanchoredpngimages/{id}", GetController().DeleteRectAnchoredPngImage)

	mux.HandleFunc("GET " + base + "/rectanchoredrects", GetController().GetRectAnchoredRects)
	mux.HandleFunc("GET " + base + "/rectanchoredrects/{id}", GetController().GetRectAnchoredRect)
	mux.HandleFunc("POST " + base + "/rectanchoredrects", GetController().PostRectAnchoredRect)
	mux.HandleFunc("PATCH " + base + "/rectanchoredrects/{id}", GetController().UpdateRectAnchoredRect)
	mux.HandleFunc("PUT " + base + "/rectanchoredrects/{id}", GetController().UpdateRectAnchoredRect)
	mux.HandleFunc("DELETE " + base + "/rectanchoredrects/{id}", GetController().DeleteRectAnchoredRect)

	mux.HandleFunc("GET " + base + "/rectanchoredtexts", GetController().GetRectAnchoredTexts)
	mux.HandleFunc("GET " + base + "/rectanchoredtexts/{id}", GetController().GetRectAnchoredText)
	mux.HandleFunc("POST " + base + "/rectanchoredtexts", GetController().PostRectAnchoredText)
	mux.HandleFunc("PATCH " + base + "/rectanchoredtexts/{id}", GetController().UpdateRectAnchoredText)
	mux.HandleFunc("PUT " + base + "/rectanchoredtexts/{id}", GetController().UpdateRectAnchoredText)
	mux.HandleFunc("DELETE " + base + "/rectanchoredtexts/{id}", GetController().DeleteRectAnchoredText)

	mux.HandleFunc("GET " + base + "/rectlinklinks", GetController().GetRectLinkLinks)
	mux.HandleFunc("GET " + base + "/rectlinklinks/{id}", GetController().GetRectLinkLink)
	mux.HandleFunc("POST " + base + "/rectlinklinks", GetController().PostRectLinkLink)
	mux.HandleFunc("PATCH " + base + "/rectlinklinks/{id}", GetController().UpdateRectLinkLink)
	mux.HandleFunc("PUT " + base + "/rectlinklinks/{id}", GetController().UpdateRectLinkLink)
	mux.HandleFunc("DELETE " + base + "/rectlinklinks/{id}", GetController().DeleteRectLinkLink)

	mux.HandleFunc("GET " + base + "/svgs", GetController().GetSVGs)
	mux.HandleFunc("GET " + base + "/svgs/{id}", GetController().GetSVG)
	mux.HandleFunc("POST " + base + "/svgs", GetController().PostSVG)
	mux.HandleFunc("PATCH " + base + "/svgs/{id}", GetController().UpdateSVG)
	mux.HandleFunc("PUT " + base + "/svgs/{id}", GetController().UpdateSVG)
	mux.HandleFunc("DELETE " + base + "/svgs/{id}", GetController().DeleteSVG)

	mux.HandleFunc("GET " + base + "/svgtexts", GetController().GetSvgTexts)
	mux.HandleFunc("GET " + base + "/svgtexts/{id}", GetController().GetSvgText)
	mux.HandleFunc("POST " + base + "/svgtexts", GetController().PostSvgText)
	mux.HandleFunc("PATCH " + base + "/svgtexts/{id}", GetController().UpdateSvgText)
	mux.HandleFunc("PUT " + base + "/svgtexts/{id}", GetController().UpdateSvgText)
	mux.HandleFunc("DELETE " + base + "/svgtexts/{id}", GetController().DeleteSvgText)

	mux.HandleFunc("GET " + base + "/texts", GetController().GetTexts)
	mux.HandleFunc("GET " + base + "/texts/{id}", GetController().GetText)
	mux.HandleFunc("POST " + base + "/texts", GetController().PostText)
	mux.HandleFunc("PATCH " + base + "/texts/{id}", GetController().UpdateText)
	mux.HandleFunc("PUT " + base + "/texts/{id}", GetController().UpdateText)
	mux.HandleFunc("DELETE " + base + "/texts/{id}", GetController().DeleteText)

	mux.HandleFunc("GET " + base + "/commitfrombacknb", GetController().GetLastCommitFromBackNb)
	mux.HandleFunc("GET " + base + "/pushfromfrontnb", GetController().GetLastPushFromFrontNb)
	mux.HandleFunc("GET " + base + "/ws/stage", GetController().onWebSocketRequestForBackRepoContent)
	mux.HandleFunc("GET " + base + "/stacks", GetController().stacks)
}

func (controller *Controller) stacks(w http.ResponseWriter, r *http.Request) {

	var res []string

	for k := range controller.Map_BackRepos {
		res = append(res, k)
	}

	writeJSON(w, http.StatusOK, res)
}

// onWebSocketRequestForBackRepoContent is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequestForBackRepoContent(w http.ResponseWriter, r *http.Request) {

	// log.Println("Stack github.com/fullstack-lang/gong/lib/svg/go, onWebSocketRequestForBackRepoContent")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.

	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			if origin == "" {
				// log.Printf("CheckOrigin: Origin header is empty. Request from: %s", r.RemoteAddr)
			} else {
				// log.Printf("CheckOrigin: Accepted connection from Origin '%s'", origin)
			}

			// Always return true to allow connections from Cloud Run and other environments
			// that do not send explicit ports in their Origin headers.
			return true
		},
	}

	wsConnection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wsConnection.Close()

	// Create a context that is canceled when the connection is closed
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	values := r.URL.Query()
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
	_ = shortHash

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

		displayStackPath := stackPath
		if len(displayStackPath) > 85 {
			displayStackPath = displayStackPath[:82] + "..."
		}

		// 2. Use a single, formatted log line
		log.Printf(
			"%-12s | %-85s | Idx: %d | Size: %-9s",
			component,
			displayStackPath,
			index,
			formatBytes(jsonSize),
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
				_ = shortHash

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

					displayStackPath := stackPath
					if len(displayStackPath) > 85 {
						displayStackPath = displayStackPath[:82] + "..."
					}

					// 2. Use a single, formatted log line
					log.Printf(
						"%-12s | %-85s | Idx: %d | Size: %-9s",
						component,
						displayStackPath,
						index,
						formatBytes(jsonSize),
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
func (controller *Controller) GetLastCommitFromBackNb(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
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

	writeJSON(w, http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func (controller *Controller) GetLastPushFromFrontNb(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
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

	writeJSON(w, http.StatusOK, res)
}
