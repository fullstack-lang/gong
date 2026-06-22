//go:build js && wasm

package main

import (
	"log"
	"syscall/js"

	"github.com/fullstack-lang/gong/dsm/project/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/project/go/models"
	"github.com/fullstack-lang/gong/lib/wasmregistry"
)

func main() {
	log.SetOutput(&wasmregistry.ConsoleWriter{})
	log.SetPrefix("project: ")
	log.SetFlags(log.Lmicroseconds)

	log.Println("Initializing project WASM Backend...")

	unmarshallFromCode := ""
	marshallOnCommit := ""
	embeddedDiagrams := true

	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1StackDelta("project", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams, true)
	stack.Stage.SetGongMarshallingMode(models.GongMarshallingAppendCommit)
	stack.Stage.SetIsWithGenesisCommit(true) // the genesis commit is the first commit of the stage, it is the one that contains the initial data. It cannot be rollbacked.

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	// initiates the UX loop
	models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
		"",
	)

	// Expose the HTTP and Socket bridges to the Angular frontend
	wasmregistry.SetupWasmHooks(stack.R)

	// Add event listener for postMessage to handle injected files
	js.Global().Call("eval", `
window.addEventListener('message', (event) => {
    if (event.data && event.data.action === 'PROCESS_INJECTED_FILE') {
        const fileContent = event.data.fileData;
        const fileName = event.data.fileName;
        
        const byteCharacters = atob(fileContent);
        const byteNumbers = new Array(byteCharacters.length);
        for (let i = 0; i < byteCharacters.length; i++) {
            byteNumbers[i] = byteCharacters.charCodeAt(i);
        }
        const byteArray = new Uint8Array(byteNumbers);
        const injectedFile = new File([byteArray], fileName, { type: "text/plain" });

        const fileInput = document.querySelector('input[type="file"]');
        if (fileInput) {
            const dataTransfer = new DataTransfer();
            dataTransfer.items.add(injectedFile);
            fileInput.files = dataTransfer.files;
            const changeEvent = new Event('change', { bubbles: true });
            fileInput.dispatchEvent(changeEvent);
        } else {
            console.error("file input not found!");
        }
    }
});
	`)

	select {} // Keep the WASM instance running indefinitely
}
