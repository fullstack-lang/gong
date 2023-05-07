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
	cmd := exec.Command("ng", "generate", "library", modelPkg.Name+"datamodel", "--defaults=true", "--skip-install=true")
	cmd.Dir = gong_models.NgWorkspacePath
	log.Printf("Creating a specific library %s in the angular workspace\n", modelPkg.Name)

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
	log.Printf("ng generate library is over and took %s", time.Since(start))

	directory, _ := filepath.Abs(
		filepath.Join(gong_models.MaterialLibDatamodelTargetPath, "data-model-panel"))
	_, err := os.Stat(directory)
	if os.IsNotExist(err) {
		err := os.MkdirAll(directory, os.ModePerm)

		if err != nil {
			log.Fatalln("Unable to create dir", directory, err.Error())
		}

		log.Println("creating directory : " + gong_models.OrmPkgGenPath)

		// if data-model-panel does not exists, it is an indication that
		// the module file has to be overwritten
		gong_models.VerySimpleCodeGenerator(
			modelPkg,
			modelPkg.Name,
			modelPkg.PkgPath,
			filepath.Join(gong_models.MaterialLibDatamodelTargetPath,
				fmt.Sprintf("%sspecific.module.ts", modelPkg.Name)),
			angular.NgFileModuleSpecific)

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
	if os.IsExist(err) {
		log.Println("directory " + gong_models.OrmPkgGenPath + " allready exists")
	}

}
