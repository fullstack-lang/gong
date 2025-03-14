// generated code - do not edit
package static

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	// this package contains ...
	"github.com/fullstack-lang/gong/lib/xlsx"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func ServeStaticFiles(logGINFlag bool) (r *gin.Engine) {

	// setup controlers
	if !logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}
	r = gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080", "http://localhost:4200"} // Allow requests from localhost:8080 and localhost:4200

	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"} // Allow specific HTTP methods

	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"} // Allow specific headers

	r.Use(cors.New(config))

	// insertion point for serving the static file
	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder(xlsx.NgDistNg, "ng-github.com-fullstack-lang-gong-lib-xlsx/dist/ng-github.com-fullstack-lang-gong-lib-xlsx/browser")))
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
