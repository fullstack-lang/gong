package main

import (
	"flag"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	gong_controllers "github.com/fullstack-lang/gong/stacks/gong/go/controllers"
	gong_models "github.com/fullstack-lang/gong/stacks/gong/go/models"
	gong_orm "github.com/fullstack-lang/gong/stacks/gong/go/orm"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")
	apiFlag    = flag.Bool("api", false, "it true, use api controllers instead of default controllers")

	pkgPath = flag.String("pkgPath", ".", "path to the models package in order to reveal gong elements in the package")
)

func main() {

	log.SetPrefix("gong: ")
	log.SetFlags(0)

	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatal("surplus arguments")
	}

	// load package to analyse
	modelPkg := &gong_models.ModelPkg{}
	gong_models.Walk(*pkgPath, modelPkg)

	modelPkg.SerializeToStage()
	// setup controlers
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}
	r := gin.Default()
	r.Use(cors.Default())

	// setup GORM
	db := gong_orm.SetupModels(*logDBFlag, ":memory:")
	// mandatory, otherwise, bizarre errors occurs
	db.DB().SetMaxOpenConns(1)

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db) // a gin Context can have a map of variable that is set up at runtime
		c.Next()
	})

	gong_orm.BackRepo.Init(db)
	gong_models.Stage.Commit()

	gong_controllers.RegisterControllers(r)

	r.Run()
}
