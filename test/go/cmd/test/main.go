package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gong/test/go/controllers"
	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"

	test "github.com/fullstack-lang/gong/test"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")
	marshall   = flag.Bool("marshall", false, "marshall data from models.StageReference")
	unmarshall = flag.Bool("unmarshall", false, "unmarshall data from models.StageReference")
)

func main() {

	log.SetPrefix("test: ")
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

	// reset stage and copy from models.StageReference
	if *marshall {
		file, err := os.Create("./stage.go")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()

		models.Stage.Checkout()
		models.Stage.Marshall(file, "github.com/fullstack-lang/gong/test/go/models", "main")
		os.Exit(0)
	}

	// reset stage and copy from models.StageReference
	if *unmarshall {
		models.Stage.Checkout()
		models.Stage.Reset()
		models.Stage.Commit()
		Unmarshall(&models.Stage)
		models.Stage.Commit()
		os.Exit(0)
	}

	// since the stack can be a multi threaded application. It is important to set up
	// only one open connexion at a time
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	dbDB.SetMaxOpenConns(1)

	controllers.RegisterControllers(r)

	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder(test.NgDistNg, "ng/dist/ng")))
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
