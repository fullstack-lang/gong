package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fullstack-lang/gong/go/angular"
	gong_models "github.com/fullstack-lang/gong/go/models"
)

func generateDatamodelLib(modelPkg *gong_models.ModelPkg) {
	start := time.Now()
	cmd := exec.Command("ng", "generate", "component", "data-model-panel", "--project", modelPkg.Name+"datamodel")
	cmd.Dir = gong_models.NgWorkspacePath
	log.Printf("Creating a specific component %s in the angular workspace\n", "data-model-panel")

	// https://stackoverflow.com/questions/48253268/print-the-stdout-from-exec-command-in-real-time-in-go
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	log.Println(cmd.String())
	log.Println(stdBuffer.String())

	// Execute the command
	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}
	log.Printf("ng generate component is over and took %s", time.Since(start))

	// if data-model-panel does not exists, it is an indication that
	// the module file has to be overwritten
	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(gong_models.MaterialLibDatamodelTargetPath,
			fmt.Sprintf("%sdatamodel.module.ts", modelPkg.Name)),
		angular.NgFileModuleDatamodel)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(gong_models.MaterialLibDatamodelTargetPath,
			"data-model-panel", "data-model-panel.component.ts"),
		angular.NgFileDataModelPanelTemplateTs)

	gong_models.VerySimpleCodeGenerator(
		modelPkg,
		modelPkg.Name,
		modelPkg.PkgPath,
		filepath.Join(gong_models.MaterialLibDatamodelTargetPath,
			"data-model-panel", "data-model-panel.component.html"),
		angular.NgFileDataModelPanelTemplateHtml)
}
