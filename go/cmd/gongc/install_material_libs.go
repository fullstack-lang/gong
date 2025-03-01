package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fullstack-lang/gong/go/angular"
	gong_models "github.com/fullstack-lang/gong/go/models"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var caserEnglish = cases.Title(language.English)

func installMaterialLibs(modelPkg *gong_models.ModelPkg) {
	{
		start := time.Now()
		//  --skip-confirmation is necessary when executing angular without user interaction
		// otherwise, one meet the error
		// "No terminal detected. '--skip-confirmation' can be used to bypass installation confirmation.
		// Ensure package name is correct prior to '--skip-confirmation' option usage."
		cmd := exec.Command("ng", "add", "@angular/material@19", "--skip-confirmation")
		cmd.Dir = modelPkg.NgWorkspacePath
		log.Printf("Adding angular material\n")

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
		log.Printf("ng add material is over and took %s", time.Since(start))
	}
	{
		start := time.Now()
		//  --skip-confirmation is necessary when executing angular without user interaction
		// otherwise, one meet the error
		// "No terminal detected. '--skip-confirmation' can be used to bypass installation confirmation.
		// Ensure package name is correct prior to '--skip-confirmation' option usage."
		cmd := exec.Command("npm", "install", "@angular/animations@19")
		cmd.Dir = modelPkg.NgWorkspacePath
		log.Printf("Adding angular animation\n")

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
		log.Printf("ng add animation is over and took %s", time.Since(start))
	}

	{
		start := time.Now()
		cmd := exec.Command("npm", "install", "--save",
			"angular-split@19",
			"material-icons")
		cmd.Dir = modelPkg.NgWorkspacePath
		log.Printf("Installing some packages\n")

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
		log.Printf("npm install is over and took %s", time.Since(start))
	}

	// generate default app.component.ts, app.component.html and app.module.ts
	{
		gong_models.VerySimpleCodeGenerator(
			modelPkg,
			filepath.Join(modelPkg.NgWorkspacePath, "src/app/app.config.ts"),
			angular.NgFileAppConfig)

		gong_models.VerySimpleCodeGenerator(
			modelPkg,
			filepath.Join(modelPkg.NgWorkspacePath, "src/index.html"),
			angular.NgFileIndex)

		gong_models.VerySimpleCodeGenerator(
			modelPkg,
			filepath.Join(modelPkg.NgWorkspacePath, "src/app/app.component.ts"),
			angular.NgFileAppComponentTs)

		gong_models.VerySimpleCodeGenerator(
			modelPkg,
			filepath.Join(modelPkg.NgWorkspacePath, "src/app/app.component.html"),
			angular.NgFileAppComponentHtml)

	}

}
