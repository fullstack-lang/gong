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

	"github.com/fullstack-lang/gong/lib/split/go/orm"

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
	base := "/api/github.com/fullstack-lang/gong/lib/split/go/v1"

	mux.HandleFunc("GET " + base + "/assplits", GetController().GetAsSplits)
	mux.HandleFunc("GET " + base + "/assplits/{id}", GetController().GetAsSplit)
	mux.HandleFunc("POST " + base + "/assplits", GetController().PostAsSplit)
	mux.HandleFunc("PATCH " + base + "/assplits/{id}", GetController().UpdateAsSplit)
	mux.HandleFunc("PUT " + base + "/assplits/{id}", GetController().UpdateAsSplit)
	mux.HandleFunc("DELETE " + base + "/assplits/{id}", GetController().DeleteAsSplit)

	mux.HandleFunc("GET " + base + "/assplitareas", GetController().GetAsSplitAreas)
	mux.HandleFunc("GET " + base + "/assplitareas/{id}", GetController().GetAsSplitArea)
	mux.HandleFunc("POST " + base + "/assplitareas", GetController().PostAsSplitArea)
	mux.HandleFunc("PATCH " + base + "/assplitareas/{id}", GetController().UpdateAsSplitArea)
	mux.HandleFunc("PUT " + base + "/assplitareas/{id}", GetController().UpdateAsSplitArea)
	mux.HandleFunc("DELETE " + base + "/assplitareas/{id}", GetController().DeleteAsSplitArea)

	mux.HandleFunc("GET " + base + "/buttons", GetController().GetButtons)
	mux.HandleFunc("GET " + base + "/buttons/{id}", GetController().GetButton)
	mux.HandleFunc("POST " + base + "/buttons", GetController().PostButton)
	mux.HandleFunc("PATCH " + base + "/buttons/{id}", GetController().UpdateButton)
	mux.HandleFunc("PUT " + base + "/buttons/{id}", GetController().UpdateButton)
	mux.HandleFunc("DELETE " + base + "/buttons/{id}", GetController().DeleteButton)

	mux.HandleFunc("GET " + base + "/cursors", GetController().GetCursors)
	mux.HandleFunc("GET " + base + "/cursors/{id}", GetController().GetCursor)
	mux.HandleFunc("POST " + base + "/cursors", GetController().PostCursor)
	mux.HandleFunc("PATCH " + base + "/cursors/{id}", GetController().UpdateCursor)
	mux.HandleFunc("PUT " + base + "/cursors/{id}", GetController().UpdateCursor)
	mux.HandleFunc("DELETE " + base + "/cursors/{id}", GetController().DeleteCursor)

	mux.HandleFunc("GET " + base + "/favicons", GetController().GetFavIcons)
	mux.HandleFunc("GET " + base + "/favicons/{id}", GetController().GetFavIcon)
	mux.HandleFunc("POST " + base + "/favicons", GetController().PostFavIcon)
	mux.HandleFunc("PATCH " + base + "/favicons/{id}", GetController().UpdateFavIcon)
	mux.HandleFunc("PUT " + base + "/favicons/{id}", GetController().UpdateFavIcon)
	mux.HandleFunc("DELETE " + base + "/favicons/{id}", GetController().DeleteFavIcon)

	mux.HandleFunc("GET " + base + "/forms", GetController().GetForms)
	mux.HandleFunc("GET " + base + "/forms/{id}", GetController().GetForm)
	mux.HandleFunc("POST " + base + "/forms", GetController().PostForm)
	mux.HandleFunc("PATCH " + base + "/forms/{id}", GetController().UpdateForm)
	mux.HandleFunc("PUT " + base + "/forms/{id}", GetController().UpdateForm)
	mux.HandleFunc("DELETE " + base + "/forms/{id}", GetController().DeleteForm)

	mux.HandleFunc("GET " + base + "/loads", GetController().GetLoads)
	mux.HandleFunc("GET " + base + "/loads/{id}", GetController().GetLoad)
	mux.HandleFunc("POST " + base + "/loads", GetController().PostLoad)
	mux.HandleFunc("PATCH " + base + "/loads/{id}", GetController().UpdateLoad)
	mux.HandleFunc("PUT " + base + "/loads/{id}", GetController().UpdateLoad)
	mux.HandleFunc("DELETE " + base + "/loads/{id}", GetController().DeleteLoad)

	mux.HandleFunc("GET " + base + "/logoonthelefts", GetController().GetLogoOnTheLefts)
	mux.HandleFunc("GET " + base + "/logoonthelefts/{id}", GetController().GetLogoOnTheLeft)
	mux.HandleFunc("POST " + base + "/logoonthelefts", GetController().PostLogoOnTheLeft)
	mux.HandleFunc("PATCH " + base + "/logoonthelefts/{id}", GetController().UpdateLogoOnTheLeft)
	mux.HandleFunc("PUT " + base + "/logoonthelefts/{id}", GetController().UpdateLogoOnTheLeft)
	mux.HandleFunc("DELETE " + base + "/logoonthelefts/{id}", GetController().DeleteLogoOnTheLeft)

	mux.HandleFunc("GET " + base + "/logoontherights", GetController().GetLogoOnTheRights)
	mux.HandleFunc("GET " + base + "/logoontherights/{id}", GetController().GetLogoOnTheRight)
	mux.HandleFunc("POST " + base + "/logoontherights", GetController().PostLogoOnTheRight)
	mux.HandleFunc("PATCH " + base + "/logoontherights/{id}", GetController().UpdateLogoOnTheRight)
	mux.HandleFunc("PUT " + base + "/logoontherights/{id}", GetController().UpdateLogoOnTheRight)
	mux.HandleFunc("DELETE " + base + "/logoontherights/{id}", GetController().DeleteLogoOnTheRight)

	mux.HandleFunc("GET " + base + "/markdowns", GetController().GetMarkdowns)
	mux.HandleFunc("GET " + base + "/markdowns/{id}", GetController().GetMarkdown)
	mux.HandleFunc("POST " + base + "/markdowns", GetController().PostMarkdown)
	mux.HandleFunc("PATCH " + base + "/markdowns/{id}", GetController().UpdateMarkdown)
	mux.HandleFunc("PUT " + base + "/markdowns/{id}", GetController().UpdateMarkdown)
	mux.HandleFunc("DELETE " + base + "/markdowns/{id}", GetController().DeleteMarkdown)

	mux.HandleFunc("GET " + base + "/sliders", GetController().GetSliders)
	mux.HandleFunc("GET " + base + "/sliders/{id}", GetController().GetSlider)
	mux.HandleFunc("POST " + base + "/sliders", GetController().PostSlider)
	mux.HandleFunc("PATCH " + base + "/sliders/{id}", GetController().UpdateSlider)
	mux.HandleFunc("PUT " + base + "/sliders/{id}", GetController().UpdateSlider)
	mux.HandleFunc("DELETE " + base + "/sliders/{id}", GetController().DeleteSlider)

	mux.HandleFunc("GET " + base + "/splits", GetController().GetSplits)
	mux.HandleFunc("GET " + base + "/splits/{id}", GetController().GetSplit)
	mux.HandleFunc("POST " + base + "/splits", GetController().PostSplit)
	mux.HandleFunc("PATCH " + base + "/splits/{id}", GetController().UpdateSplit)
	mux.HandleFunc("PUT " + base + "/splits/{id}", GetController().UpdateSplit)
	mux.HandleFunc("DELETE " + base + "/splits/{id}", GetController().DeleteSplit)

	mux.HandleFunc("GET " + base + "/svgs", GetController().GetSvgs)
	mux.HandleFunc("GET " + base + "/svgs/{id}", GetController().GetSvg)
	mux.HandleFunc("POST " + base + "/svgs", GetController().PostSvg)
	mux.HandleFunc("PATCH " + base + "/svgs/{id}", GetController().UpdateSvg)
	mux.HandleFunc("PUT " + base + "/svgs/{id}", GetController().UpdateSvg)
	mux.HandleFunc("DELETE " + base + "/svgs/{id}", GetController().DeleteSvg)

	mux.HandleFunc("GET " + base + "/tables", GetController().GetTables)
	mux.HandleFunc("GET " + base + "/tables/{id}", GetController().GetTable)
	mux.HandleFunc("POST " + base + "/tables", GetController().PostTable)
	mux.HandleFunc("PATCH " + base + "/tables/{id}", GetController().UpdateTable)
	mux.HandleFunc("PUT " + base + "/tables/{id}", GetController().UpdateTable)
	mux.HandleFunc("DELETE " + base + "/tables/{id}", GetController().DeleteTable)

	mux.HandleFunc("GET " + base + "/threejss", GetController().GetThreejss)
	mux.HandleFunc("GET " + base + "/threejss/{id}", GetController().GetThreejs)
	mux.HandleFunc("POST " + base + "/threejss", GetController().PostThreejs)
	mux.HandleFunc("PATCH " + base + "/threejss/{id}", GetController().UpdateThreejs)
	mux.HandleFunc("PUT " + base + "/threejss/{id}", GetController().UpdateThreejs)
	mux.HandleFunc("DELETE " + base + "/threejss/{id}", GetController().DeleteThreejs)

	mux.HandleFunc("GET " + base + "/titles", GetController().GetTitles)
	mux.HandleFunc("GET " + base + "/titles/{id}", GetController().GetTitle)
	mux.HandleFunc("POST " + base + "/titles", GetController().PostTitle)
	mux.HandleFunc("PATCH " + base + "/titles/{id}", GetController().UpdateTitle)
	mux.HandleFunc("PUT " + base + "/titles/{id}", GetController().UpdateTitle)
	mux.HandleFunc("DELETE " + base + "/titles/{id}", GetController().DeleteTitle)

	mux.HandleFunc("GET " + base + "/tones", GetController().GetTones)
	mux.HandleFunc("GET " + base + "/tones/{id}", GetController().GetTone)
	mux.HandleFunc("POST " + base + "/tones", GetController().PostTone)
	mux.HandleFunc("PATCH " + base + "/tones/{id}", GetController().UpdateTone)
	mux.HandleFunc("PUT " + base + "/tones/{id}", GetController().UpdateTone)
	mux.HandleFunc("DELETE " + base + "/tones/{id}", GetController().DeleteTone)

	mux.HandleFunc("GET " + base + "/trees", GetController().GetTrees)
	mux.HandleFunc("GET " + base + "/trees/{id}", GetController().GetTree)
	mux.HandleFunc("POST " + base + "/trees", GetController().PostTree)
	mux.HandleFunc("PATCH " + base + "/trees/{id}", GetController().UpdateTree)
	mux.HandleFunc("PUT " + base + "/trees/{id}", GetController().UpdateTree)
	mux.HandleFunc("DELETE " + base + "/trees/{id}", GetController().DeleteTree)

	mux.HandleFunc("GET " + base + "/views", GetController().GetViews)
	mux.HandleFunc("GET " + base + "/views/{id}", GetController().GetView)
	mux.HandleFunc("POST " + base + "/views", GetController().PostView)
	mux.HandleFunc("PATCH " + base + "/views/{id}", GetController().UpdateView)
	mux.HandleFunc("PUT " + base + "/views/{id}", GetController().UpdateView)
	mux.HandleFunc("DELETE " + base + "/views/{id}", GetController().DeleteView)

	mux.HandleFunc("GET " + base + "/xlsxs", GetController().GetXlsxs)
	mux.HandleFunc("GET " + base + "/xlsxs/{id}", GetController().GetXlsx)
	mux.HandleFunc("POST " + base + "/xlsxs", GetController().PostXlsx)
	mux.HandleFunc("PATCH " + base + "/xlsxs/{id}", GetController().UpdateXlsx)
	mux.HandleFunc("PUT " + base + "/xlsxs/{id}", GetController().UpdateXlsx)
	mux.HandleFunc("DELETE " + base + "/xlsxs/{id}", GetController().DeleteXlsx)

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

	// log.Println("Stack github.com/fullstack-lang/gong/lib/split/go, onWebSocketRequestForBackRepoContent")

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
	_ = shortHash

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
		message := "Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

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
		message := "GET Stack github.com/fullstack-lang/gong/lib/split/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	res := backRepo.GetLastPushFromFrontNb()

	writeJSON(w, http.StatusOK, res)
}
