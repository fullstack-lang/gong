package cmd

const PackageMainServerFullStack = `//go:build !js

package main

import (
	"log"
	"strconv"

	// insertion point for models import{{modelsImportDirective}}
	{{pkgname}}_stack "{{PkgPathRoot}}/stack"
	{{pkgname}}_static "{{PkgPathRoot}}/static"
	{{pkgname}}_models "{{PkgPathRoot}}/models"
)

func executeServer() {

	// setup the static file server and get the controller
	r := {{pkgname}}_static.ServeStaticFiles(logGINFlag)

	// setup model stack with its probe
	stack := {{pkgname}}_stack.NewStack(r, "{{pkgname}}", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Probe.Refresh()
	stack.Stage.Commit()

	{{pkgname}}_models.NewStager(r, stack.Stage, stack.Probe)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
`
