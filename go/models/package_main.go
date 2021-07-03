package models

const PackageMain = `package main

import (
	"flag"
	"log"
	"os"
	"fmt"
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"{{PkgPathRoot}}/controllers"
	"{{PkgPathRoot}}/orm"
)

var (
	logDBFlag = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")
	apiFlag   = flag.Bool("api", false, "it true, use api controllers instead of default controllers")
)

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

	// since gongsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	dbDB.SetMaxOpenConns(1)


	orm.BackRepo.Init(db)

	controllers.RegisterControllers(r)

	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder(ng, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}

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
`
