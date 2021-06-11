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

	test2_controllers "github.com/fullstack-lang/gong/test2/go/controllers"
	test2_orm "github.com/fullstack-lang/gong/test2/go/orm"

	test_controllers "github.com/fullstack-lang/gong/test/go/controllers"
	test_orm "github.com/fullstack-lang/gong/test/go/orm"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")
	apiFlag    = flag.Bool("api", false, "it true, use api controllers instead of default controllers")
)

func main() {

	log.SetPrefix("test2: ")
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
	inMemoryDB := test2_orm.SetupModels(*logDBFlag, ":memory:")
	// mandatory, otherwise, bizarre errors occurs
	inMemoryDBSqlDB, err := inMemoryDB.DB()
	if err != nil {
		log.Panic("Cannot access to DB of database")
	}
	inMemoryDBSqlDB.SetMaxOpenConns(1)

	// init test2 first
	test2_orm.BackRepo.Init(inMemoryDB)

	var bothStackShareTheSameDB = true

	if !bothStackShareTheSameDB {
		// setup GORM
		inFileDB := test_orm.SetupModels(*logDBFlag, "./test.db")

		// mandatory, otherwise, bizarre errors occurs
		inFileDBSqlDB, err := inFileDB.DB()
		if err != nil {
			log.Panic("Cannot access to DB of database")
		}
		inFileDBSqlDB.SetMaxOpenConns(1)
		test_orm.BackRepo.Init(inFileDB)
	} else {
		test_orm.AutoMigrate(inMemoryDB)
		test_orm.BackRepo.Init(inMemoryDB)
	}

	test2_controllers.RegisterControllers(r)
	test_controllers.RegisterControllers(r)

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
