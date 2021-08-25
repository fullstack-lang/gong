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

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gong/go/controllers"
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")
	apiFlag    = flag.Bool("api", false, "it true, use api controllers instead of default controllers")

	pkgPath = flag.String("pkgPath", "go/models", "path to the models package in order to reveal gong elements in the package")
)

func main() {

	log.SetPrefix("gong: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	if len(flag.Args()) == 1 {
		*pkgPath = flag.Arg(0)
	}

	// setup controlers
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}
	r := gin.Default()
	r.Use(cors.Default())

	// setup GORM
	orm.SetupModels(*logDBFlag, "./test.db")

	// compute absolute path
	absPkgPath, _ := filepath.Abs(*pkgPath)
	*pkgPath = absPkgPath

	// load package to analyse
	modelPkg := &gong_models.ModelPkg{}
	gong_models.Walk(*pkgPath, modelPkg)
	modelPkg.SerializeToStage()
	gong_models.Stage.Commit()

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
