package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fullstack-lang/gong/go/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const PackageMain = `package main

import (
	"flag"
	"log"
	"strconv"

	// insertion point for models import{{modelsImportDirective}}
	{{pkgname}}_stack "{{PkgPathRoot}}/stack"
	{{pkgname}}_static "{{PkgPathRoot}}/static"
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
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := {{pkgname}}_static.ServeStaticFiles(*logGINFlag)

	// setup model stack with its probe
	stack := {{pkgname}}_stack.NewStack(r, "{{pkgname}}", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	// insertion point for call to stager{{newStagerCall}}

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
`

const modelsImportDirective = `
	{{pkgname}}_models "{{PkgPathRoot}}/models"`

const newStagerCall = `
	{{pkgname}}_models.NewStager(r, stack.Stage)`

func CodeGeneratorPackageMain(
	modelPkg *models.ModelPkg,
	pkgName string,
	pkgGoPath string,
	mainFilePath string,
	skipStager bool) {
	{
		codeGo := PackageMain

		if !skipStager {
			codeGo = strings.ReplaceAll(codeGo, "{{modelsImportDirective}}", modelsImportDirective)
			codeGo = strings.ReplaceAll(codeGo, "{{newStagerCall}}", newStagerCall)
		} else {
			codeGo = strings.ReplaceAll(codeGo, "{{modelsImportDirective}}", "")
			codeGo = strings.ReplaceAll(codeGo, "{{newStagerCall}}", "")
		}

		caserEnglish := cases.Title(language.English)
		codeGo = models.Replace4(codeGo,
			"{{PkgName}}", pkgName,
			"{{TitlePkgName}}", caserEnglish.String(pkgName),
			"{{pkgname}}", strings.ToLower(pkgName),
			"{{PkgPathRoot}}", strings.ReplaceAll(pkgGoPath, "/models", ""))
		codeGo = strings.ReplaceAll(codeGo, "{{PkgPathAboveRoot}}",
			strings.ReplaceAll(pkgGoPath, "/go/models", ""))

		file, err := os.Create(mainFilePath)
		if err != nil {
			log.Panic(err)
		}
		defer file.Close()
		fmt.Fprint(file, codeGo)
	}
}
