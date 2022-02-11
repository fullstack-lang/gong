package models

const PackageMain = `package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"{{PkgPathRoot}}/controllers"
	"{{PkgPathRoot}}/models"
	"{{PkgPathRoot}}/orm"

	{{pkgname}} "{{PkgPathAboveRoot}}"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")
	marshall   = flag.String("marshall", "", "marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")
	unmarshall = flag.String("unmarshall", "", "unmarshall data from marshall name and '.go' (must be lowercased without spaces), If unmarshall arg is '', no unmarshalling")
)

// InjectionGateway is the singloton that stores all functions
// that can set the objects the stage
// InjectionGateway stores function as a map of names
var InjectionGateway = make(map[string](func()))

// hook marhalling to stage
type BeforeCommitImplementation struct {
}

// func (impl *BeforeCommitImplementation) BeforeCommit(stage *models.StageStruct) {
// 	file, err := os.Create("./stage.go")
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	defer file.Close()

// 	models.Stage.Checkout()
// 	models.Stage.Marshall(file, "{{PkgPathRoot}}/models", "main")
// }

func main() {

	log.SetPrefix("{{pkgname}}: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup controlers
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}
	r := gin.Default()
	r.Use(cors.Default())

	// setup GORM
	db := orm.SetupModels(*logDBFlag, "./test.db")
	dbDB, err := db.DB()

	// generate injection code from the stage
	if *marshall != "" {

		if strings.Contains(*marshall, " ") {
			log.Fatalln(*marshall + " must not contains blank spaces")
		}
		if strings.ToLower(*marshall) != *marshall {
			log.Fatalln(*marshall + " must be lowercases")
		}

		file, err := os.Create(fmt.Sprintf("./%s.go", *marshall))
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()

		models.Stage.Checkout()
		models.Stage.Marshall(file, "{{PkgPathAboveRoot}}/go/models", "main")
		os.Exit(0)
	}

	// setup the stage by injecting the code from code database
	if *unmarshall == "" {
		models.Stage.Checkout()
		models.Stage.Reset()
		models.Stage.Commit()
		InjectionGateway[*unmarshall]()
		models.Stage.Commit()
	}

	// hook automatic marshall to go code at every commit
	// hook := new(BeforeCommitImplementation)
	// models.Stage.OnInitCommitCallback = hook

	// since the stack can be a multi threaded application. It is important to set up
	// only one open connexion at a time
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	dbDB.SetMaxOpenConns(1)

	controllers.RegisterControllers(r)

	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder({{pkgname}}.NgDistNg, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}

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
`
