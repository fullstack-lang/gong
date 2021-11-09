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

	test2 "github.com/fullstack-lang/gong/test2"
	test2_controllers "github.com/fullstack-lang/gong/test2/go/controllers"
	test2_models "github.com/fullstack-lang/gong/test2/go/models"
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
	} else {
		test_orm.AutoMigrate(inMemoryDB)
	}

	astruct := new(test2_models.Astruct).Stage()
	astruct.Name = "Test2 Astruct instance #1"
	// astruct.Date = time.Date(2020, time.January, 1, 10, 11, 12, 0, time.UTC)

	// bstruct1 := new(models.Bstruct).Stage()
	// bstruct1.Name = "Test2 Bstruct instance #1"
	// bstruct2 := new(models.Bstruct).Stage()
	// bstruct2.Name = "Test2 Bstruct instance #2"
	// astruct.Associationtob = bstruct1
	// astruct.Anarrayofb = append(astruct.Anarrayofb, bstruct1)
	// astruct.Anarrayofb = append(astruct.Anarrayofb, bstruct2)

	test2_models.Stage.Commit()

	test2_controllers.RegisterControllers(r)
	test_controllers.RegisterControllers(r)

	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder(test2.NgDistNg, "ng/dist/ng")))
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
