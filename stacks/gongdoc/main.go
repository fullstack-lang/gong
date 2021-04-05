package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	gongdoc_controllers "github.com/fullstack-lang/gong/stacks/gongdoc/go/controllers"
	gongdoc_models "github.com/fullstack-lang/gong/stacks/gongdoc/go/models"
	gongdoc_orm "github.com/fullstack-lang/gong/stacks/gongdoc/go/orm"

	gong_controllers "github.com/fullstack-lang/gong/stacks/gong/go/controllers"
	gong_models "github.com/fullstack-lang/gong/stacks/gong/go/models"
	gong_orm "github.com/fullstack-lang/gong/stacks/gong/go/orm"
)

var (
	diagramPkgPath = flag.String("diagramPkgPath", "go/diagrams", "path to package for analysis")

	logBBFlag = flag.Bool("logDB", false, "log mode for db")

	genDefaultDiagramFlag = flag.Bool("genDefaultDiagram", false, "generate default diagram")

	svg = flag.Bool("svg", false, "generate svg output and exits")

	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	pkgPath = flag.String("pkgPath", "go/models", "path to the models package in order to reveal gong elements in the package")
)

//go:embed ng/dist/ng
var ng embed.FS

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}

func main() {
	log.SetPrefix("gongdoc: ")
	log.SetFlags(0)
	flag.Parse()
	if len(flag.Args()) > 1 {
		log.Fatal("surplus arguments")
	}
	if len(flag.Args()) == 1 {
		*diagramPkgPath = flag.Arg(0)
	}

	// setup GORM
	db := gongdoc_orm.SetupModels(*logBBFlag, ":memory:")
	db.DB().SetMaxOpenConns(1)

	gong_orm.AutoMigrate(db)

	// setup controlers
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}

	// setup controlers
	r := gin.Default()
	r.Use(cors.Default())

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db) // a gin Context can have a map of variable that is set up at runtime
		c.Next()
	})

	//
	// Init BackRepo
	gongdoc_orm.BackRepo.Init(db)
	gong_orm.BackRepo.Init(db)

	if *genDefaultDiagramFlag {
		log.Printf("Generating default diagram")

		// generates default UML diagrams
		GenGoDefaultDiagram()
		return
	}

	gongdoc_controllers.RegisterControllers(r)
	gong_controllers.RegisterControllers(r)

	r.Use(static.Serve("/", EmbedFolder(ng, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	// load package to analyse
	modelPkg := &gong_models.ModelPkg{}
	gong_models.Walk(*pkgPath, modelPkg)

	modelPkg.SerializeToStage()
	gong_models.Stage.Commit()

	var pkgelt gongdoc_models.Pkgelt
	// parse the diagram package
	pkgelt.Unmarshall(*diagramPkgPath)

	if *svg {
		for _, classDiagram := range pkgelt.Classdiagrams {
			classDiagram.OutputSVG(*diagramPkgPath)
		}
		os.Exit(0)
	}
	pkgelt.SerializeToStage()
	gongdoc_models.Stage.Commit()
	log.Printf("Parse found %d diagrams\n", len(pkgelt.Classdiagrams))

	r.Run(":8080")
}

// GenGoDefaultDiagram generates the Ng Enum
func GenGoDefaultDiagram() {

	// load package into database
	modelPkg := &gong_models.ModelPkg{}
	gong_models.Walk("../models", modelPkg)

	// generate diagrams for documentation
	var pkgelt gongdoc_models.Pkgelt
	// parse the diagram package
	diagramPkgPath := filepath.Join("../models", "../diagrams")

	// check existance of path
	_, err := os.Stat(diagramPkgPath)

	// if directory does not exist, creates it
	if os.IsNotExist(err) {

		errd := os.Mkdir(diagramPkgPath, os.ModePerm)
		if os.IsNotExist(errd) {
			log.Println("creating directory : " + diagramPkgPath)
		}

	}

	// generates default diagram
	{
		var pkgelt_default gongdoc_models.Pkgelt
		pkgelt_default.Name = gong_models.PkgGoPath

		defaultClassDiagramm := new(gongdoc_models.Classdiagram)
		defaultClassDiagramm.Name = "defaultDiagram"
		pkgelt_default.Classdiagrams = append(pkgelt_default.Classdiagrams, defaultClassDiagramm)

		idx := 0.0
		for _, _enum := range modelPkg.GongEnums {
			classshape := new(gongdoc_models.Classshape)
			classshape.ClassshapeTargetType = gongdoc_models.ENUM
			classshape.Name = _enum.Name

			classshape.Position = new(gongdoc_models.Position)
			classshape.Structname = _enum.Name

			classshape.Position.X = 40.0 + 300.0*idx
			idx = idx + 1
			classshape.Position.Y = 500.0

			defaultClassDiagramm.Classshapes = append(defaultClassDiagramm.Classshapes, classshape)

			for _, value := range _enum.GongEnumValues {
				field := new(gongdoc_models.Field)
				field.Fieldname = value.Name
				field.Structname = _enum.Name
				field.Field = value
				classshape.Fields = append(classshape.Fields, field)
			}
		}

		idx = 0.0
		for _, _struct := range modelPkg.GongStructs {
			classshape := new(gongdoc_models.Classshape)
			classshape.ClassshapeTargetType = gongdoc_models.STRUCT

			classshape.Name = _struct.Name

			classshape.Position = new(gongdoc_models.Position)
			classshape.Structname = _struct.Name

			classshape.Position.X = 40.0 + 300.0*idx
			idx = idx + 1.0
			classshape.Position.Y = 40.0

			defaultClassDiagramm.Classshapes = append(defaultClassDiagramm.Classshapes, classshape)

			// get fields

			linkIndex := 0
			for _, _field := range _struct.Fields {

				// only put exported fields
				if _field.GetName() == strings.Title(_field.GetName()) {

					switch _field.(type) {
					case *gong_models.PointerToGongStructField:
						pointerToGongStructField := _field.(*gong_models.PointerToGongStructField)

						link := new(gongdoc_models.Link)
						link.Fieldname = pointerToGongStructField.Name
						link.Structname = _struct.Name
						link.Field = _field
						link.Multiplicity = gongdoc_models.ZERO_ONE

						link.Middlevertice = new(gongdoc_models.Vertice)
						link.Middlevertice.X = 40 + 300.0*float64(idx) + 250
						link.Middlevertice.Y = 200.0 + float64(linkIndex)*50
						classshape.Links = append(classshape.Links, link)

						linkIndex = linkIndex + 1
					case *gong_models.SliceOfPointerToGongStructField:
						sliceOfPointerToGongStructField := _field.(*gong_models.SliceOfPointerToGongStructField)

						link := new(gongdoc_models.Link)
						link.Fieldname = sliceOfPointerToGongStructField.Name
						link.Structname = _struct.Name
						link.Field = _field
						link.Multiplicity = gongdoc_models.MANY

						link.Middlevertice = new(gongdoc_models.Vertice)
						link.Middlevertice.X = 40 + 300.0*float64(idx) + 250
						link.Middlevertice.Y = 200.0 + float64(linkIndex)*50
						classshape.Links = append(classshape.Links, link)

						linkIndex = linkIndex + 1
					case *gong_models.GongBasicField:
						basicField := _field.(*gong_models.GongBasicField)

						field := new(gongdoc_models.Field)
						field.Fieldname = basicField.Name
						field.Structname = _struct.Name
						field.Field = _field
						classshape.Fields = append(classshape.Fields, field)

					default:
					}

				}

			}
		}
		pkgelt_default.Marshall(diagramPkgPath)
	}

	// generates all diagrams
	if !os.IsNotExist(err) {
		pkgelt.Unmarshall(diagramPkgPath)

		for _, classDiagram := range pkgelt.Classdiagrams {
			classDiagram.OutputSVG(diagramPkgPath)
		}
	}

}
