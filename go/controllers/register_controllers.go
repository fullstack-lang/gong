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

	"github.com/fullstack-lang/gong/go/orm"

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
	v1 := r.Group("/api/github.com/fullstack-lang/gong/go")
	{ // insertion point for registrations
		v1.GET("/v1/gongbasicfields", GetController().GetGongBasicFields)
		v1.GET("/v1/gongbasicfields/:id", GetController().GetGongBasicField)
		v1.POST("/v1/gongbasicfields", GetController().PostGongBasicField)
		v1.PATCH("/v1/gongbasicfields/:id", GetController().UpdateGongBasicField)
		v1.PUT("/v1/gongbasicfields/:id", GetController().UpdateGongBasicField)
		v1.DELETE("/v1/gongbasicfields/:id", GetController().DeleteGongBasicField)

		v1.GET("/v1/gongenums", GetController().GetGongEnums)
		v1.GET("/v1/gongenums/:id", GetController().GetGongEnum)
		v1.POST("/v1/gongenums", GetController().PostGongEnum)
		v1.PATCH("/v1/gongenums/:id", GetController().UpdateGongEnum)
		v1.PUT("/v1/gongenums/:id", GetController().UpdateGongEnum)
		v1.DELETE("/v1/gongenums/:id", GetController().DeleteGongEnum)

		v1.GET("/v1/gongenumvalues", GetController().GetGongEnumValues)
		v1.GET("/v1/gongenumvalues/:id", GetController().GetGongEnumValue)
		v1.POST("/v1/gongenumvalues", GetController().PostGongEnumValue)
		v1.PATCH("/v1/gongenumvalues/:id", GetController().UpdateGongEnumValue)
		v1.PUT("/v1/gongenumvalues/:id", GetController().UpdateGongEnumValue)
		v1.DELETE("/v1/gongenumvalues/:id", GetController().DeleteGongEnumValue)

		v1.GET("/v1/gonglinks", GetController().GetGongLinks)
		v1.GET("/v1/gonglinks/:id", GetController().GetGongLink)
		v1.POST("/v1/gonglinks", GetController().PostGongLink)
		v1.PATCH("/v1/gonglinks/:id", GetController().UpdateGongLink)
		v1.PUT("/v1/gonglinks/:id", GetController().UpdateGongLink)
		v1.DELETE("/v1/gonglinks/:id", GetController().DeleteGongLink)

		v1.GET("/v1/gongnotes", GetController().GetGongNotes)
		v1.GET("/v1/gongnotes/:id", GetController().GetGongNote)
		v1.POST("/v1/gongnotes", GetController().PostGongNote)
		v1.PATCH("/v1/gongnotes/:id", GetController().UpdateGongNote)
		v1.PUT("/v1/gongnotes/:id", GetController().UpdateGongNote)
		v1.DELETE("/v1/gongnotes/:id", GetController().DeleteGongNote)

		v1.GET("/v1/gongstructs", GetController().GetGongStructs)
		v1.GET("/v1/gongstructs/:id", GetController().GetGongStruct)
		v1.POST("/v1/gongstructs", GetController().PostGongStruct)
		v1.PATCH("/v1/gongstructs/:id", GetController().UpdateGongStruct)
		v1.PUT("/v1/gongstructs/:id", GetController().UpdateGongStruct)
		v1.DELETE("/v1/gongstructs/:id", GetController().DeleteGongStruct)

		v1.GET("/v1/gongtimefields", GetController().GetGongTimeFields)
		v1.GET("/v1/gongtimefields/:id", GetController().GetGongTimeField)
		v1.POST("/v1/gongtimefields", GetController().PostGongTimeField)
		v1.PATCH("/v1/gongtimefields/:id", GetController().UpdateGongTimeField)
		v1.PUT("/v1/gongtimefields/:id", GetController().UpdateGongTimeField)
		v1.DELETE("/v1/gongtimefields/:id", GetController().DeleteGongTimeField)

		v1.GET("/v1/metas", GetController().GetMetas)
		v1.GET("/v1/metas/:id", GetController().GetMeta)
		v1.POST("/v1/metas", GetController().PostMeta)
		v1.PATCH("/v1/metas/:id", GetController().UpdateMeta)
		v1.PUT("/v1/metas/:id", GetController().UpdateMeta)
		v1.DELETE("/v1/metas/:id", GetController().DeleteMeta)

		v1.GET("/v1/metareferences", GetController().GetMetaReferences)
		v1.GET("/v1/metareferences/:id", GetController().GetMetaReference)
		v1.POST("/v1/metareferences", GetController().PostMetaReference)
		v1.PATCH("/v1/metareferences/:id", GetController().UpdateMetaReference)
		v1.PUT("/v1/metareferences/:id", GetController().UpdateMetaReference)
		v1.DELETE("/v1/metareferences/:id", GetController().DeleteMetaReference)

		v1.GET("/v1/modelpkgs", GetController().GetModelPkgs)
		v1.GET("/v1/modelpkgs/:id", GetController().GetModelPkg)
		v1.POST("/v1/modelpkgs", GetController().PostModelPkg)
		v1.PATCH("/v1/modelpkgs/:id", GetController().UpdateModelPkg)
		v1.PUT("/v1/modelpkgs/:id", GetController().UpdateModelPkg)
		v1.DELETE("/v1/modelpkgs/:id", GetController().DeleteModelPkg)

		v1.GET("/v1/pointertogongstructfields", GetController().GetPointerToGongStructFields)
		v1.GET("/v1/pointertogongstructfields/:id", GetController().GetPointerToGongStructField)
		v1.POST("/v1/pointertogongstructfields", GetController().PostPointerToGongStructField)
		v1.PATCH("/v1/pointertogongstructfields/:id", GetController().UpdatePointerToGongStructField)
		v1.PUT("/v1/pointertogongstructfields/:id", GetController().UpdatePointerToGongStructField)
		v1.DELETE("/v1/pointertogongstructfields/:id", GetController().DeletePointerToGongStructField)

		v1.GET("/v1/sliceofpointertogongstructfields", GetController().GetSliceOfPointerToGongStructFields)
		v1.GET("/v1/sliceofpointertogongstructfields/:id", GetController().GetSliceOfPointerToGongStructField)
		v1.POST("/v1/sliceofpointertogongstructfields", GetController().PostSliceOfPointerToGongStructField)
		v1.PATCH("/v1/sliceofpointertogongstructfields/:id", GetController().UpdateSliceOfPointerToGongStructField)
		v1.PUT("/v1/sliceofpointertogongstructfields/:id", GetController().UpdateSliceOfPointerToGongStructField)
		v1.DELETE("/v1/sliceofpointertogongstructfields/:id", GetController().DeleteSliceOfPointerToGongStructField)

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

	// log.Println("Stack github.com/fullstack-lang/gong/go, onWebSocketRequestForBackRepoContent")

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
		message := "WebSocket Connect, Stack github.com/fullstack-lang/gong/go, Unkown stack: \"" + stackPath + "\"\n"

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
				log.Println("github.com/fullstack-lang/gong/go", stackPath, "WS client disconnected:", err)
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
		log.Println("github.com/fullstack-lang/gong/go:\n",
			"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	} else {
		// 1. Extract the component name from the long path for cleaner logs
		// For example, "github.com/fullstack-lang/gong/lib/table/go" becomes "table"
		parts := strings.Split("github.com/fullstack-lang/gong/go", "/") // Assuming goFilePath holds the path
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
					log.Println("github.com/fullstack-lang/gong/go:\n",
						"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
					fmt.Println(err)
					return
				} else {
					// 1. Extract the component name from the long path for cleaner logs
					// For example, "github.com/fullstack-lang/gong/lib/table/go" becomes "table"
					parts := strings.Split("github.com/fullstack-lang/gong/go", "/") // Assuming goFilePath holds the path
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
		message := "Stack github.com/fullstack-lang/gong/go, Unkown stack: \"" + stackPath + "\"\n"

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
		message := "GET Stack github.com/fullstack-lang/gong/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
