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

	"github.com/fullstack-lang/gong/lib/threejs/go/orm"

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
	base := "/api/github.com/fullstack-lang/gong/lib/threejs/go/v1"

	mux.HandleFunc("GET " + base + "/ambiantlights", GetController().GetAmbiantLights)
	mux.HandleFunc("GET " + base + "/ambiantlights/{id}", GetController().GetAmbiantLight)
	mux.HandleFunc("POST " + base + "/ambiantlights", GetController().PostAmbiantLight)
	mux.HandleFunc("PATCH " + base + "/ambiantlights/{id}", GetController().UpdateAmbiantLight)
	mux.HandleFunc("PUT " + base + "/ambiantlights/{id}", GetController().UpdateAmbiantLight)
	mux.HandleFunc("DELETE " + base + "/ambiantlights/{id}", GetController().DeleteAmbiantLight)

	mux.HandleFunc("GET " + base + "/boxgeometrys", GetController().GetBoxGeometrys)
	mux.HandleFunc("GET " + base + "/boxgeometrys/{id}", GetController().GetBoxGeometry)
	mux.HandleFunc("POST " + base + "/boxgeometrys", GetController().PostBoxGeometry)
	mux.HandleFunc("PATCH " + base + "/boxgeometrys/{id}", GetController().UpdateBoxGeometry)
	mux.HandleFunc("PUT " + base + "/boxgeometrys/{id}", GetController().UpdateBoxGeometry)
	mux.HandleFunc("DELETE " + base + "/boxgeometrys/{id}", GetController().DeleteBoxGeometry)

	mux.HandleFunc("GET " + base + "/buffergeometrys", GetController().GetBufferGeometrys)
	mux.HandleFunc("GET " + base + "/buffergeometrys/{id}", GetController().GetBufferGeometry)
	mux.HandleFunc("POST " + base + "/buffergeometrys", GetController().PostBufferGeometry)
	mux.HandleFunc("PATCH " + base + "/buffergeometrys/{id}", GetController().UpdateBufferGeometry)
	mux.HandleFunc("PUT " + base + "/buffergeometrys/{id}", GetController().UpdateBufferGeometry)
	mux.HandleFunc("DELETE " + base + "/buffergeometrys/{id}", GetController().DeleteBufferGeometry)

	mux.HandleFunc("GET " + base + "/cameras", GetController().GetCameras)
	mux.HandleFunc("GET " + base + "/cameras/{id}", GetController().GetCamera)
	mux.HandleFunc("POST " + base + "/cameras", GetController().PostCamera)
	mux.HandleFunc("PATCH " + base + "/cameras/{id}", GetController().UpdateCamera)
	mux.HandleFunc("PUT " + base + "/cameras/{id}", GetController().UpdateCamera)
	mux.HandleFunc("DELETE " + base + "/cameras/{id}", GetController().DeleteCamera)

	mux.HandleFunc("GET " + base + "/canvass", GetController().GetCanvass)
	mux.HandleFunc("GET " + base + "/canvass/{id}", GetController().GetCanvas)
	mux.HandleFunc("POST " + base + "/canvass", GetController().PostCanvas)
	mux.HandleFunc("PATCH " + base + "/canvass/{id}", GetController().UpdateCanvas)
	mux.HandleFunc("PUT " + base + "/canvass/{id}", GetController().UpdateCanvas)
	mux.HandleFunc("DELETE " + base + "/canvass/{id}", GetController().DeleteCanvas)

	mux.HandleFunc("GET " + base + "/curves", GetController().GetCurves)
	mux.HandleFunc("GET " + base + "/curves/{id}", GetController().GetCurve)
	mux.HandleFunc("POST " + base + "/curves", GetController().PostCurve)
	mux.HandleFunc("PATCH " + base + "/curves/{id}", GetController().UpdateCurve)
	mux.HandleFunc("PUT " + base + "/curves/{id}", GetController().UpdateCurve)
	mux.HandleFunc("DELETE " + base + "/curves/{id}", GetController().DeleteCurve)

	mux.HandleFunc("GET " + base + "/cylindergeometrys", GetController().GetCylinderGeometrys)
	mux.HandleFunc("GET " + base + "/cylindergeometrys/{id}", GetController().GetCylinderGeometry)
	mux.HandleFunc("POST " + base + "/cylindergeometrys", GetController().PostCylinderGeometry)
	mux.HandleFunc("PATCH " + base + "/cylindergeometrys/{id}", GetController().UpdateCylinderGeometry)
	mux.HandleFunc("PUT " + base + "/cylindergeometrys/{id}", GetController().UpdateCylinderGeometry)
	mux.HandleFunc("DELETE " + base + "/cylindergeometrys/{id}", GetController().DeleteCylinderGeometry)

	mux.HandleFunc("GET " + base + "/directionallights", GetController().GetDirectionalLights)
	mux.HandleFunc("GET " + base + "/directionallights/{id}", GetController().GetDirectionalLight)
	mux.HandleFunc("POST " + base + "/directionallights", GetController().PostDirectionalLight)
	mux.HandleFunc("PATCH " + base + "/directionallights/{id}", GetController().UpdateDirectionalLight)
	mux.HandleFunc("PUT " + base + "/directionallights/{id}", GetController().UpdateDirectionalLight)
	mux.HandleFunc("DELETE " + base + "/directionallights/{id}", GetController().DeleteDirectionalLight)

	mux.HandleFunc("GET " + base + "/extrudegeometrys", GetController().GetExtrudeGeometrys)
	mux.HandleFunc("GET " + base + "/extrudegeometrys/{id}", GetController().GetExtrudeGeometry)
	mux.HandleFunc("POST " + base + "/extrudegeometrys", GetController().PostExtrudeGeometry)
	mux.HandleFunc("PATCH " + base + "/extrudegeometrys/{id}", GetController().UpdateExtrudeGeometry)
	mux.HandleFunc("PUT " + base + "/extrudegeometrys/{id}", GetController().UpdateExtrudeGeometry)
	mux.HandleFunc("DELETE " + base + "/extrudegeometrys/{id}", GetController().DeleteExtrudeGeometry)

	mux.HandleFunc("GET " + base + "/meshs", GetController().GetMeshs)
	mux.HandleFunc("GET " + base + "/meshs/{id}", GetController().GetMesh)
	mux.HandleFunc("POST " + base + "/meshs", GetController().PostMesh)
	mux.HandleFunc("PATCH " + base + "/meshs/{id}", GetController().UpdateMesh)
	mux.HandleFunc("PUT " + base + "/meshs/{id}", GetController().UpdateMesh)
	mux.HandleFunc("DELETE " + base + "/meshs/{id}", GetController().DeleteMesh)

	mux.HandleFunc("GET " + base + "/meshmaterialbasics", GetController().GetMeshMaterialBasics)
	mux.HandleFunc("GET " + base + "/meshmaterialbasics/{id}", GetController().GetMeshMaterialBasic)
	mux.HandleFunc("POST " + base + "/meshmaterialbasics", GetController().PostMeshMaterialBasic)
	mux.HandleFunc("PATCH " + base + "/meshmaterialbasics/{id}", GetController().UpdateMeshMaterialBasic)
	mux.HandleFunc("PUT " + base + "/meshmaterialbasics/{id}", GetController().UpdateMeshMaterialBasic)
	mux.HandleFunc("DELETE " + base + "/meshmaterialbasics/{id}", GetController().DeleteMeshMaterialBasic)

	mux.HandleFunc("GET " + base + "/meshphysicalmaterials", GetController().GetMeshPhysicalMaterials)
	mux.HandleFunc("GET " + base + "/meshphysicalmaterials/{id}", GetController().GetMeshPhysicalMaterial)
	mux.HandleFunc("POST " + base + "/meshphysicalmaterials", GetController().PostMeshPhysicalMaterial)
	mux.HandleFunc("PATCH " + base + "/meshphysicalmaterials/{id}", GetController().UpdateMeshPhysicalMaterial)
	mux.HandleFunc("PUT " + base + "/meshphysicalmaterials/{id}", GetController().UpdateMeshPhysicalMaterial)
	mux.HandleFunc("DELETE " + base + "/meshphysicalmaterials/{id}", GetController().DeleteMeshPhysicalMaterial)

	mux.HandleFunc("GET " + base + "/planegeometrys", GetController().GetPlaneGeometrys)
	mux.HandleFunc("GET " + base + "/planegeometrys/{id}", GetController().GetPlaneGeometry)
	mux.HandleFunc("POST " + base + "/planegeometrys", GetController().PostPlaneGeometry)
	mux.HandleFunc("PATCH " + base + "/planegeometrys/{id}", GetController().UpdatePlaneGeometry)
	mux.HandleFunc("PUT " + base + "/planegeometrys/{id}", GetController().UpdatePlaneGeometry)
	mux.HandleFunc("DELETE " + base + "/planegeometrys/{id}", GetController().DeletePlaneGeometry)

	mux.HandleFunc("GET " + base + "/shapes", GetController().GetShapes)
	mux.HandleFunc("GET " + base + "/shapes/{id}", GetController().GetShape)
	mux.HandleFunc("POST " + base + "/shapes", GetController().PostShape)
	mux.HandleFunc("PATCH " + base + "/shapes/{id}", GetController().UpdateShape)
	mux.HandleFunc("PUT " + base + "/shapes/{id}", GetController().UpdateShape)
	mux.HandleFunc("DELETE " + base + "/shapes/{id}", GetController().DeleteShape)

	mux.HandleFunc("GET " + base + "/spheregeometrys", GetController().GetSphereGeometrys)
	mux.HandleFunc("GET " + base + "/spheregeometrys/{id}", GetController().GetSphereGeometry)
	mux.HandleFunc("POST " + base + "/spheregeometrys", GetController().PostSphereGeometry)
	mux.HandleFunc("PATCH " + base + "/spheregeometrys/{id}", GetController().UpdateSphereGeometry)
	mux.HandleFunc("PUT " + base + "/spheregeometrys/{id}", GetController().UpdateSphereGeometry)
	mux.HandleFunc("DELETE " + base + "/spheregeometrys/{id}", GetController().DeleteSphereGeometry)

	mux.HandleFunc("GET " + base + "/torusgeometrys", GetController().GetTorusGeometrys)
	mux.HandleFunc("GET " + base + "/torusgeometrys/{id}", GetController().GetTorusGeometry)
	mux.HandleFunc("POST " + base + "/torusgeometrys", GetController().PostTorusGeometry)
	mux.HandleFunc("PATCH " + base + "/torusgeometrys/{id}", GetController().UpdateTorusGeometry)
	mux.HandleFunc("PUT " + base + "/torusgeometrys/{id}", GetController().UpdateTorusGeometry)
	mux.HandleFunc("DELETE " + base + "/torusgeometrys/{id}", GetController().DeleteTorusGeometry)

	mux.HandleFunc("GET " + base + "/triangles", GetController().GetTriangles)
	mux.HandleFunc("GET " + base + "/triangles/{id}", GetController().GetTriangle)
	mux.HandleFunc("POST " + base + "/triangles", GetController().PostTriangle)
	mux.HandleFunc("PATCH " + base + "/triangles/{id}", GetController().UpdateTriangle)
	mux.HandleFunc("PUT " + base + "/triangles/{id}", GetController().UpdateTriangle)
	mux.HandleFunc("DELETE " + base + "/triangles/{id}", GetController().DeleteTriangle)

	mux.HandleFunc("GET " + base + "/tubegeometrys", GetController().GetTubeGeometrys)
	mux.HandleFunc("GET " + base + "/tubegeometrys/{id}", GetController().GetTubeGeometry)
	mux.HandleFunc("POST " + base + "/tubegeometrys", GetController().PostTubeGeometry)
	mux.HandleFunc("PATCH " + base + "/tubegeometrys/{id}", GetController().UpdateTubeGeometry)
	mux.HandleFunc("PUT " + base + "/tubegeometrys/{id}", GetController().UpdateTubeGeometry)
	mux.HandleFunc("DELETE " + base + "/tubegeometrys/{id}", GetController().DeleteTubeGeometry)

	mux.HandleFunc("GET " + base + "/vector2s", GetController().GetVector2s)
	mux.HandleFunc("GET " + base + "/vector2s/{id}", GetController().GetVector2)
	mux.HandleFunc("POST " + base + "/vector2s", GetController().PostVector2)
	mux.HandleFunc("PATCH " + base + "/vector2s/{id}", GetController().UpdateVector2)
	mux.HandleFunc("PUT " + base + "/vector2s/{id}", GetController().UpdateVector2)
	mux.HandleFunc("DELETE " + base + "/vector2s/{id}", GetController().DeleteVector2)

	mux.HandleFunc("GET " + base + "/vector3s", GetController().GetVector3s)
	mux.HandleFunc("GET " + base + "/vector3s/{id}", GetController().GetVector3)
	mux.HandleFunc("POST " + base + "/vector3s", GetController().PostVector3)
	mux.HandleFunc("PATCH " + base + "/vector3s/{id}", GetController().UpdateVector3)
	mux.HandleFunc("PUT " + base + "/vector3s/{id}", GetController().UpdateVector3)
	mux.HandleFunc("DELETE " + base + "/vector3s/{id}", GetController().DeleteVector3)

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

	// log.Println("Stack github.com/fullstack-lang/gong/lib/threejs/go, onWebSocketRequestForBackRepoContent")

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
		message := "WebSocket Connect, Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

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
				log.Println("github.com/fullstack-lang/gong/lib/threejs/go", stackPath, "WS client disconnected:", err)
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
		log.Println("github.com/fullstack-lang/gong/lib/threejs/go:\n",
			"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	} else {
		// 1. Extract the component name from the long path for cleaner logs
		// For example, "github.com/fullstack-lang/gong/lib/table/go" becomes "table"
		parts := strings.Split("github.com/fullstack-lang/gong/lib/threejs/go", "/") // Assuming goFilePath holds the path
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
					log.Println("github.com/fullstack-lang/gong/lib/threejs/go:\n",
						"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
					fmt.Println(err)
					return
				} else {
					// 1. Extract the component name from the long path for cleaner logs
					// For example, "github.com/fullstack-lang/gong/lib/table/go" becomes "table"
					parts := strings.Split("github.com/fullstack-lang/gong/lib/threejs/go", "/") // Assuming goFilePath holds the path
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
		message := "Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

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
		message := "GET Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	res := backRepo.GetLastPushFromFrontNb()

	writeJSON(w, http.StatusOK, res)
}
