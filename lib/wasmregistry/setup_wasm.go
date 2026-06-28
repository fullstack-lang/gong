//go:build wasm

package wasmregistry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http/httptest"
	"syscall/js"

	"github.com/gin-gonic/gin"
)

var ginEngine *gin.Engine

// ConsoleWriter redirects writes to the browser's console.log.
type ConsoleWriter struct{}

func (cw *ConsoleWriter) Write(p []byte) (n int, err error) {
	js.Global().Get("console").Call("log", string(bytes.TrimSuffix(p, []byte("\n"))))
	return len(p), nil
}

// SetupWasmHooks exposes the HTTP and Socket bridges to the Angular frontend.
func SetupWasmHooks(r *gin.Engine) {
	ginEngine = r

	// Redirect standard log package output to browser console
	log.SetOutput(&ConsoleWriter{})

	// Redirect Gin logging to browser console
	gin.DefaultWriter = &ConsoleWriter{}
	gin.DefaultErrorWriter = &ConsoleWriter{}

	js.Global().Set("wasmFetch", js.FuncOf(wasmFetch))
	js.Global().Set("openWasmSocket", js.FuncOf(openWasmSocket))

	// Add event listener for postMessage to handle injected files
	js.Global().Call("eval", `
(function() {
    let fileProcessed = false;
    window.addEventListener('message', (event) => {
        console.log("wasmregistry/setup_wasm.go: received message event:", event.data ? event.data.action : "unknown");
        if (event.data && event.data.action === 'PROCESS_INJECTED_FILE') {
            if (fileProcessed) {
                // Already processed, just send ACK again in case it was missed
                if (event.source) {
                    event.source.postMessage({ action: 'FILE_PROCESSED' }, event.origin);
                }
                return;
            }
            fileProcessed = true; // Prevent duplicate processing

            const fileContent = event.data.fileData;
            const fileName = event.data.fileName;
            
            const byteCharacters = atob(fileContent);
            const byteNumbers = new Array(byteCharacters.length);
            for (let i = 0; i < byteCharacters.length; i++) {
                byteNumbers[i] = byteCharacters.charCodeAt(i);
            }
            const byteArray = new Uint8Array(byteNumbers);
            const injectedFile = new File([byteArray], fileName, { type: "text/plain" });

            const tryDrop = () => {
                const fileInput = document.querySelector('input[type="file"]');
                if (fileInput) {
                    const dataTransfer = new DataTransfer();
                    dataTransfer.items.add(injectedFile);
                    fileInput.files = dataTransfer.files;
                    const changeEvent = new Event('change', { bubbles: true });
                    fileInput.dispatchEvent(changeEvent);
                    console.log("File successfully injected into the input element!");
                    if (event.source) {
                        event.source.postMessage({ action: 'FILE_PROCESSED' }, event.origin);
                    }
                } else {
                    console.log("file input not found, retrying in 500ms...");
                    setTimeout(tryDrop, 500);
                }
            };
            tryDrop();
        }
    });
})();
	`)
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

	// fmt.Printf("From the backend, json unmarshall %v\n", reqData)

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

	index := 0
	wrappedCallback := js.FuncOf(func(this js.Value, cbArgs []js.Value) any {
		if len(cbArgs) > 0 {
			payload := cbArgs[0].String()
			LogWasmPush(stackType, stackPath, index, len(payload))
			index++
		}
		return callback.Invoke(cbArgs[0])
	})

	go handler(wrappedCallback.Value)
	return nil
}
