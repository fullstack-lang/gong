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

	{{pkgname}}_models "{{PkgPathRoot}}/models"
	{{pkgname}}_stack "{{PkgPathRoot}}/stack"
	{{pkgname}}_static "{{PkgPathRoot}}/static"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
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

	// setup root split stack and insert the probe at the root
	splitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage

	(&split.View{
		Name: "Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stack.Stage.GetName() + {{pkgname}}_models.ProbeSplitSuffix,
				}).Stage(splitStage),
			}).Stage(splitStage),
		},
	}).Stage(splitStage)

	splitStage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
`

const codeForNgStaticService = `
	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder({{pkgname}}.NgDistNg, "{{NgWorkspaceName}}/dist/{{NgWorkspaceName}}")))
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
