//go:build wasm

package wasmregistry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"syscall/js"

	"github.com/gin-gonic/gin"
)

var ginEngine *gin.Engine

// SetupWasmHooks exposes the HTTP and Socket bridges to the Angular frontend.
func SetupWasmHooks(r *gin.Engine) {
	ginEngine = r

	js.Global().Set("wasmFetch", js.FuncOf(wasmFetch))
	js.Global().Set("openWasmSocket", js.FuncOf(openWasmSocket))
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
	stackType := args[0].String()
	stackPath := args[1].String()
	callback := args[2]

	key := HandleKey{
		StackType: stackType,
		StackPath: stackPath,
	}

	handler, exists := Handlers[key]
	if !exists {
		fmt.Println("ERROR - No WASM handler for:", stackType, stackPath)
		return nil
	}

	go handler(callback)
	return nil
}
