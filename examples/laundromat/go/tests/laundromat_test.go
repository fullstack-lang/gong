package tests

import (
	"log"
	"testing"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gong/examples/laundromat/go/models"
	target_orm "github.com/fullstack-lang/gong/examples/laundromat/go/orm"

	// lib sim

	gongsim_models "github.com/fullstack-lang/gong/stacks/gongsim/go/models"
	gongsim_orm "github.com/fullstack-lang/gong/stacks/gongsim/go/orm"
)

func TestLaundromat(t *testing.T) {

	gongsim_models.EngineSingloton.ControlMode = gongsim_models.AUTONOMOUS

	// setup controlers
	r := gin.Default()
	r.Use(cors.Default())

	// setup GORM with the
	db := target_orm.SetupModels(false, ":memory:")

	// add gongsim database
	gongsim_orm.AutoMigrate(db)

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db) // a gin Context can have a map of variable that is set up at runtime
		c.Next()
	})

	// attach specific engine callback to the model
	simulation := models.NewSimulation()
	gongsim_models.EngineSingloton.Simulation = simulation

	gongsim_orm.BackRepo.Init(db)
	target_orm.BackRepo.Init(db)

	log.Printf("sleep 1/10 second")
	time.Sleep(time.Duration(100 * time.Millisecond))

	log.Printf("put PLAY")
	// setup command of GongsimCommand
	gongsim_models.GongsimCommandSingloton.Command = gongsim_models.COMMAND_PLAY
	gongsim_models.GongsimCommandSingloton.Commit()

	log.Printf("sleep 1/10 second")
	time.Sleep(time.Duration(100 * time.Millisecond))

	gongsim_models.EngineSingloton.ListEvents()

	log.Printf("Washer has %f kg of dirty laundry", simulation.Washer.DirtyLaundryWeight)
	log.Printf("Washer has %f kg of cleaned laundry", simulation.Washer.CleanedLaundryWeight)
	gongsim_models.EngineSingloton.Run()
	log.Printf("Simulation ended at %s", gongsim_models.EngineSingloton.GetCurrentTime().String())
	log.Printf("Washer has %f kg of dirty laundry", simulation.Washer.DirtyLaundryWeight)
	log.Printf("Washer has %f kg of cleaned laundry", simulation.Washer.CleanedLaundryWeight)
}
