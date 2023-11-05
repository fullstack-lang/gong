package golang

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

	{{pkgname}}_stack "{{PkgPathRoot}}/stack"
	{{pkgname}}_static "{{PkgPathRoot}}/static"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
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

	// setup stack
	stack := {{pkgname}}_stack.NewStack(r, "{{pkgname}}", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
`

const codeForNgStaticService = `
	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder({{pkgname}}.NgDistNg, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})`

func CodeGeneratorPackageMain(
	modelPkg *models.ModelPkg,
	pkgName string,
	pkgGoPath string,
	mainFilePath string,
	skipNg bool) {
	{
		codeGo := PackageMain

		if !skipNg {
			codeGo = strings.ReplaceAll(codeGo, "{{staticCodeServiceCode}}", codeForNgStaticService)
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
