package fullstack

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/fullstack-lang/gong"
	"github.com/fullstack-lang/gong/go/controllers"
	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	_ "github.com/fullstack-lang/gong/ng/projects"
)

// NewStackInstance creates a new stack instance from the Stack Model
// and returns the backRepo of the stack instance (you can get the stage from backRepo.GetStage()
//
// - the stackPath is the unique identifier of the stack
// - the optional parameter filenames is for the name of the database filename
// if filenames is omitted, the database is persisted in memory
func NewStackInstance(
	r *gin.Engine,
	stackPath string,
	// filesnames is an optional parameter for the name of the database
	filenames ...string) (
	stage *models.StageStruct) {

	// temporary
	if stackPath == "" {
		stage = models.GetDefaultStage()
	} else {
		stage = models.NewStage()
	}

	if len(filenames) == 0 {
		filenames = append(filenames, ":memory:")
	}

	backRepo := orm.NewBackRepo(stage, filenames[0])

	if stackPath != "" {
		controllers.GetController().AddBackRepo(backRepo, stackPath)
	}

	controllers.Register(r)

	return
}


func ServeStaticFiles(logGINFlag bool) (r *gin.Engine) {

	// setup controlers
	if !logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}
	r = gin.Default()
	r.Use(cors.Default())

	// insertion point for serving the static file
	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder(gong.NgDistNg, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	return
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
