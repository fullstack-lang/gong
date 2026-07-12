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

	"github.com/gin-gonic/gin"

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
func registerControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gong/lib/threejs/go")
	{ // insertion point for registrations
		v1.GET("/v1/ambiantlights", GetController().GetAmbiantLights)
		v1.GET("/v1/ambiantlights/:id", GetController().GetAmbiantLight)
		v1.POST("/v1/ambiantlights", GetController().PostAmbiantLight)
		v1.PATCH("/v1/ambiantlights/:id", GetController().UpdateAmbiantLight)
		v1.PUT("/v1/ambiantlights/:id", GetController().UpdateAmbiantLight)
		v1.DELETE("/v1/ambiantlights/:id", GetController().DeleteAmbiantLight)

		v1.GET("/v1/boxgeometrys", GetController().GetBoxGeometrys)
		v1.GET("/v1/boxgeometrys/:id", GetController().GetBoxGeometry)
		v1.POST("/v1/boxgeometrys", GetController().PostBoxGeometry)
		v1.PATCH("/v1/boxgeometrys/:id", GetController().UpdateBoxGeometry)
		v1.PUT("/v1/boxgeometrys/:id", GetController().UpdateBoxGeometry)
		v1.DELETE("/v1/boxgeometrys/:id", GetController().DeleteBoxGeometry)

		v1.GET("/v1/cameras", GetController().GetCameras)
		v1.GET("/v1/cameras/:id", GetController().GetCamera)
		v1.POST("/v1/cameras", GetController().PostCamera)
		v1.PATCH("/v1/cameras/:id", GetController().UpdateCamera)
		v1.PUT("/v1/cameras/:id", GetController().UpdateCamera)
		v1.DELETE("/v1/cameras/:id", GetController().DeleteCamera)

		v1.GET("/v1/canvass", GetController().GetCanvass)
		v1.GET("/v1/canvass/:id", GetController().GetCanvas)
		v1.POST("/v1/canvass", GetController().PostCanvas)
		v1.PATCH("/v1/canvass/:id", GetController().UpdateCanvas)
		v1.PUT("/v1/canvass/:id", GetController().UpdateCanvas)
		v1.DELETE("/v1/canvass/:id", GetController().DeleteCanvas)

		v1.GET("/v1/curves", GetController().GetCurves)
		v1.GET("/v1/curves/:id", GetController().GetCurve)
		v1.POST("/v1/curves", GetController().PostCurve)
		v1.PATCH("/v1/curves/:id", GetController().UpdateCurve)
		v1.PUT("/v1/curves/:id", GetController().UpdateCurve)
		v1.DELETE("/v1/curves/:id", GetController().DeleteCurve)

		v1.GET("/v1/cylindergeometrys", GetController().GetCylinderGeometrys)
		v1.GET("/v1/cylindergeometrys/:id", GetController().GetCylinderGeometry)
		v1.POST("/v1/cylindergeometrys", GetController().PostCylinderGeometry)
		v1.PATCH("/v1/cylindergeometrys/:id", GetController().UpdateCylinderGeometry)
		v1.PUT("/v1/cylindergeometrys/:id", GetController().UpdateCylinderGeometry)
		v1.DELETE("/v1/cylindergeometrys/:id", GetController().DeleteCylinderGeometry)

		v1.GET("/v1/directionallights", GetController().GetDirectionalLights)
		v1.GET("/v1/directionallights/:id", GetController().GetDirectionalLight)
		v1.POST("/v1/directionallights", GetController().PostDirectionalLight)
		v1.PATCH("/v1/directionallights/:id", GetController().UpdateDirectionalLight)
		v1.PUT("/v1/directionallights/:id", GetController().UpdateDirectionalLight)
		v1.DELETE("/v1/directionallights/:id", GetController().DeleteDirectionalLight)

		v1.GET("/v1/extrudegeometrys", GetController().GetExtrudeGeometrys)
		v1.GET("/v1/extrudegeometrys/:id", GetController().GetExtrudeGeometry)
		v1.POST("/v1/extrudegeometrys", GetController().PostExtrudeGeometry)
		v1.PATCH("/v1/extrudegeometrys/:id", GetController().UpdateExtrudeGeometry)
		v1.PUT("/v1/extrudegeometrys/:id", GetController().UpdateExtrudeGeometry)
		v1.DELETE("/v1/extrudegeometrys/:id", GetController().DeleteExtrudeGeometry)

		v1.GET("/v1/meshs", GetController().GetMeshs)
		v1.GET("/v1/meshs/:id", GetController().GetMesh)
		v1.POST("/v1/meshs", GetController().PostMesh)
		v1.PATCH("/v1/meshs/:id", GetController().UpdateMesh)
		v1.PUT("/v1/meshs/:id", GetController().UpdateMesh)
		v1.DELETE("/v1/meshs/:id", GetController().DeleteMesh)

		v1.GET("/v1/meshmaterialbasics", GetController().GetMeshMaterialBasics)
		v1.GET("/v1/meshmaterialbasics/:id", GetController().GetMeshMaterialBasic)
		v1.POST("/v1/meshmaterialbasics", GetController().PostMeshMaterialBasic)
		v1.PATCH("/v1/meshmaterialbasics/:id", GetController().UpdateMeshMaterialBasic)
		v1.PUT("/v1/meshmaterialbasics/:id", GetController().UpdateMeshMaterialBasic)
		v1.DELETE("/v1/meshmaterialbasics/:id", GetController().DeleteMeshMaterialBasic)

		v1.GET("/v1/meshphysicalmaterials", GetController().GetMeshPhysicalMaterials)
		v1.GET("/v1/meshphysicalmaterials/:id", GetController().GetMeshPhysicalMaterial)
		v1.POST("/v1/meshphysicalmaterials", GetController().PostMeshPhysicalMaterial)
		v1.PATCH("/v1/meshphysicalmaterials/:id", GetController().UpdateMeshPhysicalMaterial)
		v1.PUT("/v1/meshphysicalmaterials/:id", GetController().UpdateMeshPhysicalMaterial)
		v1.DELETE("/v1/meshphysicalmaterials/:id", GetController().DeleteMeshPhysicalMaterial)

		v1.GET("/v1/planegeometrys", GetController().GetPlaneGeometrys)
		v1.GET("/v1/planegeometrys/:id", GetController().GetPlaneGeometry)
		v1.POST("/v1/planegeometrys", GetController().PostPlaneGeometry)
		v1.PATCH("/v1/planegeometrys/:id", GetController().UpdatePlaneGeometry)
		v1.PUT("/v1/planegeometrys/:id", GetController().UpdatePlaneGeometry)
		v1.DELETE("/v1/planegeometrys/:id", GetController().DeletePlaneGeometry)

		v1.GET("/v1/shapes", GetController().GetShapes)
		v1.GET("/v1/shapes/:id", GetController().GetShape)
		v1.POST("/v1/shapes", GetController().PostShape)
		v1.PATCH("/v1/shapes/:id", GetController().UpdateShape)
		v1.PUT("/v1/shapes/:id", GetController().UpdateShape)
		v1.DELETE("/v1/shapes/:id", GetController().DeleteShape)

		v1.GET("/v1/spheregeometrys", GetController().GetSphereGeometrys)
		v1.GET("/v1/spheregeometrys/:id", GetController().GetSphereGeometry)
		v1.POST("/v1/spheregeometrys", GetController().PostSphereGeometry)
		v1.PATCH("/v1/spheregeometrys/:id", GetController().UpdateSphereGeometry)
		v1.PUT("/v1/spheregeometrys/:id", GetController().UpdateSphereGeometry)
		v1.DELETE("/v1/spheregeometrys/:id", GetController().DeleteSphereGeometry)

		v1.GET("/v1/torusgeometrys", GetController().GetTorusGeometrys)
		v1.GET("/v1/torusgeometrys/:id", GetController().GetTorusGeometry)
		v1.POST("/v1/torusgeometrys", GetController().PostTorusGeometry)
		v1.PATCH("/v1/torusgeometrys/:id", GetController().UpdateTorusGeometry)
		v1.PUT("/v1/torusgeometrys/:id", GetController().UpdateTorusGeometry)
		v1.DELETE("/v1/torusgeometrys/:id", GetController().DeleteTorusGeometry)

		v1.GET("/v1/tubegeometrys", GetController().GetTubeGeometrys)
		v1.GET("/v1/tubegeometrys/:id", GetController().GetTubeGeometry)
		v1.POST("/v1/tubegeometrys", GetController().PostTubeGeometry)
		v1.PATCH("/v1/tubegeometrys/:id", GetController().UpdateTubeGeometry)
		v1.PUT("/v1/tubegeometrys/:id", GetController().UpdateTubeGeometry)
		v1.DELETE("/v1/tubegeometrys/:id", GetController().DeleteTubeGeometry)

		v1.GET("/v1/vector2s", GetController().GetVector2s)
		v1.GET("/v1/vector2s/:id", GetController().GetVector2)
		v1.POST("/v1/vector2s", GetController().PostVector2)
		v1.PATCH("/v1/vector2s/:id", GetController().UpdateVector2)
		v1.PUT("/v1/vector2s/:id", GetController().UpdateVector2)
		v1.DELETE("/v1/vector2s/:id", GetController().DeleteVector2)

		v1.GET("/v1/vector3s", GetController().GetVector3s)
		v1.GET("/v1/vector3s/:id", GetController().GetVector3)
		v1.POST("/v1/vector3s", GetController().PostVector3)
		v1.PATCH("/v1/vector3s/:id", GetController().UpdateVector3)
		v1.PUT("/v1/vector3s/:id", GetController().UpdateVector3)
		v1.DELETE("/v1/vector3s/:id", GetController().DeleteVector3)

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

		// 2. Use a single, formatted log line
		log.Printf(
			"%-12s | %-85s | Idx: %d | Size: %-9s",
			component,
			stackPath,
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

					// 2. Use a single, formatted log line
					log.Printf(
						"%-12s | %-85s | Idx: %d | Size: %-9s",
						component,
						stackPath,
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
		message := "Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

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
		message := "GET Stack github.com/fullstack-lang/gong/lib/threejs/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
