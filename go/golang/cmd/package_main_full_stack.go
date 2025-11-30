package cmd

const PackageMainFullStack = `package main

import (
	"flag"
	"log"
	"strconv"

	// insertion point for models import{{modelsImportDirective}}
	{{pkgname}}_stack "{{PkgPathRoot}}/stack"
	{{pkgname}}_static "{{PkgPathRoot}}/static"
	{{pkgname}}_models "{{PkgPathRoot}}/models"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("{{pkgname}}: ")
	log.SetFlags(log.Lmicroseconds)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := {{pkgname}}_static.ServeStaticFiles(*logGINFlag)

	// setup model stack with its probe
	stack := {{pkgname}}_stack.NewStack(r, "{{pkgname}}", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()
	stack.Stage.Commit()

	{{pkgname}}_models.NewStager(r, stack.Stage, stack.Probe)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
`
