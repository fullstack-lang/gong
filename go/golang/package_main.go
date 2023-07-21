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
	"fmt"
	"log"
	"os"
	"strconv"

	{{pkgname}}_go "{{PkgPathRoot}}"
	{{pkgname}}_data "{{PkgPathRoot}}/data"
	{{pkgname}}_fullstack "{{PkgPathRoot}}/fullstack"
	{{pkgname}}_models "{{PkgPathRoot}}/models"
	{{pkgname}}_static "{{PkgPathRoot}}/static"

	gongdoc_load "github.com/fullstack-lang/gongdoc/go/load"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

// InjectionGateway is the singloton that stores all functions
// that can set the objects the stage
// InjectionGateway stores function as a map of names
var InjectionGateway = make(map[string](func()))

// hook marhalling to stage
type BeforeCommitImplementation struct {
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *{{pkgname}}_models.StageStruct) {
	file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnCommit))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stage.Checkout()
	stage.Marshall(file, "{{PkgPathRoot}}/models", "main")
}

func main() {

	log.SetPrefix("{{pkgname}}: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := {{pkgname}}_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	var stage *{{pkgname}}_models.StageStruct
	if *marshallOnCommit != "" {
		// persistence in a SQLite file on disk in memory
		stage = {{pkgname}}_fullstack.NewStackInstance(r, "{{pkgname}}")
	} else {
		// persistence in a SQLite file on disk
		stage = {{pkgname}}_fullstack.NewStackInstance(r, "{{pkgname}}", "./test.db")
	}

	{{pkgname}}_data.Load(r, {{pkgname}}_go.GoModelsDir, "{{pkgname}}")

	if *unmarshallFromCode != "" {
		stage.Checkout()
		stage.Reset()
		stage.Commit()
		err := {{pkgname}}_models.ParseAstFile(stage, *unmarshallFromCode)

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		stage.Commit()
	} else {
		// in case the database is used, checkout the content to the stage
		stage.Checkout()
	}

	// hook automatic marshall to go code at every commit
	if *marshallOnCommit != "" {
		hook := new(BeforeCommitImplementation)
		stage.OnInitCommitFromFrontCallback = hook
	}

	gongdoc_load.Load(
		"{{pkgname}}",
		"{{PkgPathRoot}}/models",
		{{pkgname}}_go.GoModelsDir,
		{{pkgname}}_go.GoDiagramsDir,
		r,
		*embeddedDiagrams,
		&stage.Map_GongStructName_InstancesNb)

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
	mdlPkg *models.ModelPkg,
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
