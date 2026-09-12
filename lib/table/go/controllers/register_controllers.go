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

	"github.com/fullstack-lang/gong/lib/table/go/orm"

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
	base := "/api/github.com/fullstack-lang/gong/lib/table/go/v1"

	mux.HandleFunc("GET " + base + "/buttons", GetController().GetButtons)
	mux.HandleFunc("GET " + base + "/buttons/{id}", GetController().GetButton)
	mux.HandleFunc("POST " + base + "/buttons", GetController().PostButton)
	mux.HandleFunc("PATCH " + base + "/buttons/{id}", GetController().UpdateButton)
	mux.HandleFunc("PUT " + base + "/buttons/{id}", GetController().UpdateButton)
	mux.HandleFunc("DELETE " + base + "/buttons/{id}", GetController().DeleteButton)

	mux.HandleFunc("GET " + base + "/cells", GetController().GetCells)
	mux.HandleFunc("GET " + base + "/cells/{id}", GetController().GetCell)
	mux.HandleFunc("POST " + base + "/cells", GetController().PostCell)
	mux.HandleFunc("PATCH " + base + "/cells/{id}", GetController().UpdateCell)
	mux.HandleFunc("PUT " + base + "/cells/{id}", GetController().UpdateCell)
	mux.HandleFunc("DELETE " + base + "/cells/{id}", GetController().DeleteCell)

	mux.HandleFunc("GET " + base + "/cellbooleans", GetController().GetCellBooleans)
	mux.HandleFunc("GET " + base + "/cellbooleans/{id}", GetController().GetCellBoolean)
	mux.HandleFunc("POST " + base + "/cellbooleans", GetController().PostCellBoolean)
	mux.HandleFunc("PATCH " + base + "/cellbooleans/{id}", GetController().UpdateCellBoolean)
	mux.HandleFunc("PUT " + base + "/cellbooleans/{id}", GetController().UpdateCellBoolean)
	mux.HandleFunc("DELETE " + base + "/cellbooleans/{id}", GetController().DeleteCellBoolean)

	mux.HandleFunc("GET " + base + "/cellfloat64s", GetController().GetCellFloat64s)
	mux.HandleFunc("GET " + base + "/cellfloat64s/{id}", GetController().GetCellFloat64)
	mux.HandleFunc("POST " + base + "/cellfloat64s", GetController().PostCellFloat64)
	mux.HandleFunc("PATCH " + base + "/cellfloat64s/{id}", GetController().UpdateCellFloat64)
	mux.HandleFunc("PUT " + base + "/cellfloat64s/{id}", GetController().UpdateCellFloat64)
	mux.HandleFunc("DELETE " + base + "/cellfloat64s/{id}", GetController().DeleteCellFloat64)

	mux.HandleFunc("GET " + base + "/cellicons", GetController().GetCellIcons)
	mux.HandleFunc("GET " + base + "/cellicons/{id}", GetController().GetCellIcon)
	mux.HandleFunc("POST " + base + "/cellicons", GetController().PostCellIcon)
	mux.HandleFunc("PATCH " + base + "/cellicons/{id}", GetController().UpdateCellIcon)
	mux.HandleFunc("PUT " + base + "/cellicons/{id}", GetController().UpdateCellIcon)
	mux.HandleFunc("DELETE " + base + "/cellicons/{id}", GetController().DeleteCellIcon)

	mux.HandleFunc("GET " + base + "/cellints", GetController().GetCellInts)
	mux.HandleFunc("GET " + base + "/cellints/{id}", GetController().GetCellInt)
	mux.HandleFunc("POST " + base + "/cellints", GetController().PostCellInt)
	mux.HandleFunc("PATCH " + base + "/cellints/{id}", GetController().UpdateCellInt)
	mux.HandleFunc("PUT " + base + "/cellints/{id}", GetController().UpdateCellInt)
	mux.HandleFunc("DELETE " + base + "/cellints/{id}", GetController().DeleteCellInt)

	mux.HandleFunc("GET " + base + "/cellstrings", GetController().GetCellStrings)
	mux.HandleFunc("GET " + base + "/cellstrings/{id}", GetController().GetCellString)
	mux.HandleFunc("POST " + base + "/cellstrings", GetController().PostCellString)
	mux.HandleFunc("PATCH " + base + "/cellstrings/{id}", GetController().UpdateCellString)
	mux.HandleFunc("PUT " + base + "/cellstrings/{id}", GetController().UpdateCellString)
	mux.HandleFunc("DELETE " + base + "/cellstrings/{id}", GetController().DeleteCellString)

	mux.HandleFunc("GET " + base + "/displayedcolumns", GetController().GetDisplayedColumns)
	mux.HandleFunc("GET " + base + "/displayedcolumns/{id}", GetController().GetDisplayedColumn)
	mux.HandleFunc("POST " + base + "/displayedcolumns", GetController().PostDisplayedColumn)
	mux.HandleFunc("PATCH " + base + "/displayedcolumns/{id}", GetController().UpdateDisplayedColumn)
	mux.HandleFunc("PUT " + base + "/displayedcolumns/{id}", GetController().UpdateDisplayedColumn)
	mux.HandleFunc("DELETE " + base + "/displayedcolumns/{id}", GetController().DeleteDisplayedColumn)

	mux.HandleFunc("GET " + base + "/rows", GetController().GetRows)
	mux.HandleFunc("GET " + base + "/rows/{id}", GetController().GetRow)
	mux.HandleFunc("POST " + base + "/rows", GetController().PostRow)
	mux.HandleFunc("PATCH " + base + "/rows/{id}", GetController().UpdateRow)
	mux.HandleFunc("PUT " + base + "/rows/{id}", GetController().UpdateRow)
	mux.HandleFunc("DELETE " + base + "/rows/{id}", GetController().DeleteRow)

	mux.HandleFunc("GET " + base + "/svgicons", GetController().GetSVGIcons)
	mux.HandleFunc("GET " + base + "/svgicons/{id}", GetController().GetSVGIcon)
	mux.HandleFunc("POST " + base + "/svgicons", GetController().PostSVGIcon)
	mux.HandleFunc("PATCH " + base + "/svgicons/{id}", GetController().UpdateSVGIcon)
	mux.HandleFunc("PUT " + base + "/svgicons/{id}", GetController().UpdateSVGIcon)
	mux.HandleFunc("DELETE " + base + "/svgicons/{id}", GetController().DeleteSVGIcon)

	mux.HandleFunc("GET " + base + "/tables", GetController().GetTables)
	mux.HandleFunc("GET " + base + "/tables/{id}", GetController().GetTable)
	mux.HandleFunc("POST " + base + "/tables", GetController().PostTable)
	mux.HandleFunc("PATCH " + base + "/tables/{id}", GetController().UpdateTable)
	mux.HandleFunc("PUT " + base + "/tables/{id}", GetController().UpdateTable)
	mux.HandleFunc("DELETE " + base + "/tables/{id}", GetController().DeleteTable)

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

	// log.Println("Stack github.com/fullstack-lang/gong/lib/table/go, onWebSocketRequestForBackRepoContent")

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
		message := "WebSocket Connect, Stack github.com/fullstack-lang/gong/lib/table/go, Unkown stack: \"" + stackPath + "\"\n"

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
				log.Println("github.com/fullstack-lang/gong/lib/table/go", stackPath, "WS client disconnected:", err)
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
		log.Println("github.com/fullstack-lang/gong/lib/table/go:\n",
			"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	} else {
		// 1. Extract the component name from the long path for cleaner logs
		// For example, "github.com/fullstack-lang/gong/lib/table/go" becomes "table"
		parts := strings.Split("github.com/fullstack-lang/gong/lib/table/go", "/") // Assuming goFilePath holds the path
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
					log.Println("github.com/fullstack-lang/gong/lib/table/go:\n",
						"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
					fmt.Println(err)
					return
				} else {
					// 1. Extract the component name from the long path for cleaner logs
					// For example, "github.com/fullstack-lang/gong/lib/table/go" becomes "table"
					parts := strings.Split("github.com/fullstack-lang/gong/lib/table/go", "/") // Assuming goFilePath holds the path
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
		message := "Stack github.com/fullstack-lang/gong/lib/table/go, Unkown stack: \"" + stackPath + "\"\n"

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
		message := "GET Stack github.com/fullstack-lang/gong/lib/table/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	res := backRepo.GetLastPushFromFrontNb()

	writeJSON(w, http.StatusOK, res)
}
