//go:build wasm

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http/httptest"
	"syscall/js"
	"time"

	"github.com/fullstack-lang/gong/lib/wasmregistry"
	"github.com/fullstack-lang/gong/test/test4/go/models"
	"github.com/gin-gonic/gin"
)

var ginEngine *gin.Engine

func main() {
	flag.Parse()

	// 1. Initialize the exact same Gong app in memory
	// Capture the Gin router in a temporary variable 'r'
	r, stack := setupApp()

	// Assign it to our global variable so wasmFetch can use it!
	ginEngine = r

	fmt.Println("From the backend. after setupApp")

	// 2. Expose the HTTP and Socket bridges to the Angular frontend
	js.Global().Set("wasmFetch", js.FuncOf(wasmFetch))

	fmt.Println("From the backend. after wasmFetch")

	js.Global().Set("openWasmSocket", js.FuncOf(openWasmSocket))

	fmt.Println("From the backend. after openWasmSocket")

	err := models.ParseAstEmbeddedFile(stack.Stage, stage_content, "data/stage.go")
	if err != nil {
		log.Fatalln(err.Error())
		fmt.Println("Error while parsing")
	} else {
		stack.Stage.Commit()
		fmt.Println("After the first commit")
	}

	// get the first Astruct
	astructs := *models.GetGongstructInstancesSet[models.Astruct](stack.Stage)

	var astruct *models.Astruct
	for astruct = range astructs {
		break
	}

	if astruct != nil {
		idx := 0
		for {
			time.Sleep(time.Second)
			astruct.Name = fmt.Sprintf("A%d", idx%3)
			idx++
			stack.Stage.Commit()

			if idx > 2 {
				break
			}
		}
	}

	// 3. Prevent the WASM module from exiting
	select {}
}

// wasmFetch intercepts HTTP calls and routes them through Gin (Task 4)
func wasmFetch(this js.Value, args []js.Value) any {

	fmt.Println("From the backend. in wasmFetch")

	reqJson := args[0].String()

	var reqData struct {
		Method string `json:"method"`
		URL    string `json:"url"`
		Body   string `json:"body"`
	}
	json.Unmarshal([]byte(reqJson), &reqData)

	fmt.Printf("From the backend, json unmarshall %v\n", reqData)

	var body io.Reader
	if reqData.Body != "" {
		body = bytes.NewBufferString(reqData.Body)
	}

	// Pass the fake request to the Gin Engine
	req := httptest.NewRequest(reqData.Method, reqData.URL, body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	ginEngine.ServeHTTP(w, req)

	res := map[string]any{
		"status": w.Code,
		"body":   w.Body.String(),
	}

	jsRes, _ := json.Marshal(res)
	return string(jsRes)
}

func openWasmSocket(this js.Value, args []js.Value) any {
	stackPath := args[0].String()
	callback := args[1]

	handler, exists := wasmregistry.Handlers[stackPath]
	if !exists {
		fmt.Println("ERROR - No WASM handler for:", stackPath)
		return nil
	}

	go handler(callback)
	return nil
}
