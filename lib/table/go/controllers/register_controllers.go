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

	"github.com/fullstack-lang/gong/lib/table/go/orm"

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
	v1 := r.Group("/api/github.com/fullstack-lang/gong/lib/table/go")
	{ // insertion point for registrations
		v1.GET("/v1/cells", GetController().GetCells)
		v1.GET("/v1/cells/:id", GetController().GetCell)
		v1.POST("/v1/cells", GetController().PostCell)
		v1.PATCH("/v1/cells/:id", GetController().UpdateCell)
		v1.PUT("/v1/cells/:id", GetController().UpdateCell)
		v1.DELETE("/v1/cells/:id", GetController().DeleteCell)

		v1.GET("/v1/cellbooleans", GetController().GetCellBooleans)
		v1.GET("/v1/cellbooleans/:id", GetController().GetCellBoolean)
		v1.POST("/v1/cellbooleans", GetController().PostCellBoolean)
		v1.PATCH("/v1/cellbooleans/:id", GetController().UpdateCellBoolean)
		v1.PUT("/v1/cellbooleans/:id", GetController().UpdateCellBoolean)
		v1.DELETE("/v1/cellbooleans/:id", GetController().DeleteCellBoolean)

		v1.GET("/v1/cellfloat64s", GetController().GetCellFloat64s)
		v1.GET("/v1/cellfloat64s/:id", GetController().GetCellFloat64)
		v1.POST("/v1/cellfloat64s", GetController().PostCellFloat64)
		v1.PATCH("/v1/cellfloat64s/:id", GetController().UpdateCellFloat64)
		v1.PUT("/v1/cellfloat64s/:id", GetController().UpdateCellFloat64)
		v1.DELETE("/v1/cellfloat64s/:id", GetController().DeleteCellFloat64)

		v1.GET("/v1/cellicons", GetController().GetCellIcons)
		v1.GET("/v1/cellicons/:id", GetController().GetCellIcon)
		v1.POST("/v1/cellicons", GetController().PostCellIcon)
		v1.PATCH("/v1/cellicons/:id", GetController().UpdateCellIcon)
		v1.PUT("/v1/cellicons/:id", GetController().UpdateCellIcon)
		v1.DELETE("/v1/cellicons/:id", GetController().DeleteCellIcon)

		v1.GET("/v1/cellints", GetController().GetCellInts)
		v1.GET("/v1/cellints/:id", GetController().GetCellInt)
		v1.POST("/v1/cellints", GetController().PostCellInt)
		v1.PATCH("/v1/cellints/:id", GetController().UpdateCellInt)
		v1.PUT("/v1/cellints/:id", GetController().UpdateCellInt)
		v1.DELETE("/v1/cellints/:id", GetController().DeleteCellInt)

		v1.GET("/v1/cellstrings", GetController().GetCellStrings)
		v1.GET("/v1/cellstrings/:id", GetController().GetCellString)
		v1.POST("/v1/cellstrings", GetController().PostCellString)
		v1.PATCH("/v1/cellstrings/:id", GetController().UpdateCellString)
		v1.PUT("/v1/cellstrings/:id", GetController().UpdateCellString)
		v1.DELETE("/v1/cellstrings/:id", GetController().DeleteCellString)

		v1.GET("/v1/checkboxs", GetController().GetCheckBoxs)
		v1.GET("/v1/checkboxs/:id", GetController().GetCheckBox)
		v1.POST("/v1/checkboxs", GetController().PostCheckBox)
		v1.PATCH("/v1/checkboxs/:id", GetController().UpdateCheckBox)
		v1.PUT("/v1/checkboxs/:id", GetController().UpdateCheckBox)
		v1.DELETE("/v1/checkboxs/:id", GetController().DeleteCheckBox)

		v1.GET("/v1/displayedcolumns", GetController().GetDisplayedColumns)
		v1.GET("/v1/displayedcolumns/:id", GetController().GetDisplayedColumn)
		v1.POST("/v1/displayedcolumns", GetController().PostDisplayedColumn)
		v1.PATCH("/v1/displayedcolumns/:id", GetController().UpdateDisplayedColumn)
		v1.PUT("/v1/displayedcolumns/:id", GetController().UpdateDisplayedColumn)
		v1.DELETE("/v1/displayedcolumns/:id", GetController().DeleteDisplayedColumn)

		v1.GET("/v1/formdivs", GetController().GetFormDivs)
		v1.GET("/v1/formdivs/:id", GetController().GetFormDiv)
		v1.POST("/v1/formdivs", GetController().PostFormDiv)
		v1.PATCH("/v1/formdivs/:id", GetController().UpdateFormDiv)
		v1.PUT("/v1/formdivs/:id", GetController().UpdateFormDiv)
		v1.DELETE("/v1/formdivs/:id", GetController().DeleteFormDiv)

		v1.GET("/v1/formeditassocbuttons", GetController().GetFormEditAssocButtons)
		v1.GET("/v1/formeditassocbuttons/:id", GetController().GetFormEditAssocButton)
		v1.POST("/v1/formeditassocbuttons", GetController().PostFormEditAssocButton)
		v1.PATCH("/v1/formeditassocbuttons/:id", GetController().UpdateFormEditAssocButton)
		v1.PUT("/v1/formeditassocbuttons/:id", GetController().UpdateFormEditAssocButton)
		v1.DELETE("/v1/formeditassocbuttons/:id", GetController().DeleteFormEditAssocButton)

		v1.GET("/v1/formfields", GetController().GetFormFields)
		v1.GET("/v1/formfields/:id", GetController().GetFormField)
		v1.POST("/v1/formfields", GetController().PostFormField)
		v1.PATCH("/v1/formfields/:id", GetController().UpdateFormField)
		v1.PUT("/v1/formfields/:id", GetController().UpdateFormField)
		v1.DELETE("/v1/formfields/:id", GetController().DeleteFormField)

		v1.GET("/v1/formfielddates", GetController().GetFormFieldDates)
		v1.GET("/v1/formfielddates/:id", GetController().GetFormFieldDate)
		v1.POST("/v1/formfielddates", GetController().PostFormFieldDate)
		v1.PATCH("/v1/formfielddates/:id", GetController().UpdateFormFieldDate)
		v1.PUT("/v1/formfielddates/:id", GetController().UpdateFormFieldDate)
		v1.DELETE("/v1/formfielddates/:id", GetController().DeleteFormFieldDate)

		v1.GET("/v1/formfielddatetimes", GetController().GetFormFieldDateTimes)
		v1.GET("/v1/formfielddatetimes/:id", GetController().GetFormFieldDateTime)
		v1.POST("/v1/formfielddatetimes", GetController().PostFormFieldDateTime)
		v1.PATCH("/v1/formfielddatetimes/:id", GetController().UpdateFormFieldDateTime)
		v1.PUT("/v1/formfielddatetimes/:id", GetController().UpdateFormFieldDateTime)
		v1.DELETE("/v1/formfielddatetimes/:id", GetController().DeleteFormFieldDateTime)

		v1.GET("/v1/formfieldfloat64s", GetController().GetFormFieldFloat64s)
		v1.GET("/v1/formfieldfloat64s/:id", GetController().GetFormFieldFloat64)
		v1.POST("/v1/formfieldfloat64s", GetController().PostFormFieldFloat64)
		v1.PATCH("/v1/formfieldfloat64s/:id", GetController().UpdateFormFieldFloat64)
		v1.PUT("/v1/formfieldfloat64s/:id", GetController().UpdateFormFieldFloat64)
		v1.DELETE("/v1/formfieldfloat64s/:id", GetController().DeleteFormFieldFloat64)

		v1.GET("/v1/formfieldints", GetController().GetFormFieldInts)
		v1.GET("/v1/formfieldints/:id", GetController().GetFormFieldInt)
		v1.POST("/v1/formfieldints", GetController().PostFormFieldInt)
		v1.PATCH("/v1/formfieldints/:id", GetController().UpdateFormFieldInt)
		v1.PUT("/v1/formfieldints/:id", GetController().UpdateFormFieldInt)
		v1.DELETE("/v1/formfieldints/:id", GetController().DeleteFormFieldInt)

		v1.GET("/v1/formfieldselects", GetController().GetFormFieldSelects)
		v1.GET("/v1/formfieldselects/:id", GetController().GetFormFieldSelect)
		v1.POST("/v1/formfieldselects", GetController().PostFormFieldSelect)
		v1.PATCH("/v1/formfieldselects/:id", GetController().UpdateFormFieldSelect)
		v1.PUT("/v1/formfieldselects/:id", GetController().UpdateFormFieldSelect)
		v1.DELETE("/v1/formfieldselects/:id", GetController().DeleteFormFieldSelect)

		v1.GET("/v1/formfieldstrings", GetController().GetFormFieldStrings)
		v1.GET("/v1/formfieldstrings/:id", GetController().GetFormFieldString)
		v1.POST("/v1/formfieldstrings", GetController().PostFormFieldString)
		v1.PATCH("/v1/formfieldstrings/:id", GetController().UpdateFormFieldString)
		v1.PUT("/v1/formfieldstrings/:id", GetController().UpdateFormFieldString)
		v1.DELETE("/v1/formfieldstrings/:id", GetController().DeleteFormFieldString)

		v1.GET("/v1/formfieldtimes", GetController().GetFormFieldTimes)
		v1.GET("/v1/formfieldtimes/:id", GetController().GetFormFieldTime)
		v1.POST("/v1/formfieldtimes", GetController().PostFormFieldTime)
		v1.PATCH("/v1/formfieldtimes/:id", GetController().UpdateFormFieldTime)
		v1.PUT("/v1/formfieldtimes/:id", GetController().UpdateFormFieldTime)
		v1.DELETE("/v1/formfieldtimes/:id", GetController().DeleteFormFieldTime)

		v1.GET("/v1/formgroups", GetController().GetFormGroups)
		v1.GET("/v1/formgroups/:id", GetController().GetFormGroup)
		v1.POST("/v1/formgroups", GetController().PostFormGroup)
		v1.PATCH("/v1/formgroups/:id", GetController().UpdateFormGroup)
		v1.PUT("/v1/formgroups/:id", GetController().UpdateFormGroup)
		v1.DELETE("/v1/formgroups/:id", GetController().DeleteFormGroup)

		v1.GET("/v1/formsortassocbuttons", GetController().GetFormSortAssocButtons)
		v1.GET("/v1/formsortassocbuttons/:id", GetController().GetFormSortAssocButton)
		v1.POST("/v1/formsortassocbuttons", GetController().PostFormSortAssocButton)
		v1.PATCH("/v1/formsortassocbuttons/:id", GetController().UpdateFormSortAssocButton)
		v1.PUT("/v1/formsortassocbuttons/:id", GetController().UpdateFormSortAssocButton)
		v1.DELETE("/v1/formsortassocbuttons/:id", GetController().DeleteFormSortAssocButton)

		v1.GET("/v1/options", GetController().GetOptions)
		v1.GET("/v1/options/:id", GetController().GetOption)
		v1.POST("/v1/options", GetController().PostOption)
		v1.PATCH("/v1/options/:id", GetController().UpdateOption)
		v1.PUT("/v1/options/:id", GetController().UpdateOption)
		v1.DELETE("/v1/options/:id", GetController().DeleteOption)

		v1.GET("/v1/rows", GetController().GetRows)
		v1.GET("/v1/rows/:id", GetController().GetRow)
		v1.POST("/v1/rows", GetController().PostRow)
		v1.PATCH("/v1/rows/:id", GetController().UpdateRow)
		v1.PUT("/v1/rows/:id", GetController().UpdateRow)
		v1.DELETE("/v1/rows/:id", GetController().DeleteRow)

		v1.GET("/v1/tables", GetController().GetTables)
		v1.GET("/v1/tables/:id", GetController().GetTable)
		v1.POST("/v1/tables", GetController().PostTable)
		v1.PATCH("/v1/tables/:id", GetController().UpdateTable)
		v1.PUT("/v1/tables/:id", GetController().UpdateTable)
		v1.DELETE("/v1/tables/:id", GetController().DeleteTable)

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

	// log.Println("Stack github.com/fullstack-lang/gong/lib/table/go, onWebSocketRequestForBackRepoContent")

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
		message := "Stack github.com/fullstack-lang/gong/lib/table/go, Unkown stack: \"" + stackPath + "\"\n"

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
		message := "GET Stack github.com/fullstack-lang/gong/lib/table/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
