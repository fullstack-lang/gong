package main

import (
	"flag"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	target_controllers "github.com/fullstack-lang/gong/examples/laundromat/go/controllers"
	target_models "github.com/fullstack-lang/gong/examples/laundromat/go/models"
	target_orm "github.com/fullstack-lang/gong/examples/laundromat/go/orm"

	gongsim_controllers "github.com/fullstack-lang/gong/stacks/gongsim/go/controllers"
	gongsim_models "github.com/fullstack-lang/gong/stacks/gongsim/go/models"
	gongsim_orm "github.com/fullstack-lang/gong/stacks/gongsim/go/orm"

	gongdoc_controllers "github.com/fullstack-lang/gong/stacks/gongdoc/go/controllers"
	gongdoc_models "github.com/fullstack-lang/gong/stacks/gongdoc/go/models"
	gongdoc_orm "github.com/fullstack-lang/gong/stacks/gongdoc/go/orm"
)

var (
	logDBFlag         = flag.Bool("logDB", false, "log mode for db")
	logGINFlag        = flag.Bool("logGIN", false, "log mode for gin")
	clientControlFlag = flag.Bool("client-control", false, "if true, engine waits for API calls")
)

var db *gorm.DB

//
// generic code
//
// specific code is in target_engine
func main() {

	log.SetPrefix("laundromat: ")
	log.SetFlags(0)

	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatal("surplus arguments")
	}

	// Gongsim
	if *clientControlFlag {
		gongsim_models.EngineSingloton.ControlMode = gongsim_models.CLIENT_CONTROL
	} else {
		gongsim_models.EngineSingloton.ControlMode = gongsim_models.AUTONOMOUS
	}

	// setup GORM
	db = target_orm.SetupModels(*logDBFlag, ":memory:")
	db.DB().SetMaxOpenConns(1)

	// add gongdocatabase
	gongdoc_orm.AutoMigrate(db)

	// add gongsim database
	gongsim_orm.AutoMigrate(db)

	// attach specific engine callback to the laundromat model
	simulation := target_models.NewSimulation()
	simulation.Stage()
	gongsim_models.EngineSingloton.Simulation = simulation

	// attach the callback on the action singloton of the gongsim stack
	// legacy to be removed
	// gongsim_controllers.ActionSinglotonID.ActionBackRepoCallback = gongsim_singlotons.ActionSinglotonID

	//
	//  setup controlers
	//
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}
	r := gin.Default()
	r.Use(cors.Default())

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db) // a gin Context can have a map of variable that is set up at runtime
		c.Next()
	})

	target_controllers.RegisterControllers(r)
	gongsim_controllers.RegisterControllers(r)
	gongdoc_controllers.RegisterControllers(r)

	//
	// init all back repositories
	//
	gongsim_orm.BackRepo.Init(db)
	gongdoc_orm.BackRepo.Init(db)
	target_orm.BackRepo.Init(db)

	// put all to database
	gongsim_models.Stage.Commit()
	gongdoc_models.Stage.Commit()
	target_models.Stage.Commit()

	log.Printf("simulation ready to run")
	r.Run()
	os.Exit(0)
}
