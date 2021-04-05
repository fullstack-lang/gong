package main

import (
	"flag"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/fullstack-lang/gong/stacks/gongsim/go/controllers"
	"github.com/fullstack-lang/gong/stacks/gongsim/go/models"
	"github.com/fullstack-lang/gong/stacks/gongsim/go/orm"
)

var (
	logDBFlag         = flag.Bool("logDB", false, "log mode for db")
	logGINFlag        = flag.Bool("logGIN", false, "log mode for gin")
	clientControlFlag = flag.Bool("client-control", false, "if true, engine waits for API calls")
)

func main() {

	log.SetPrefix("gongsim: ")
	log.SetFlags(0)

	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatal("surplus arguments")
	}

	// setup GORM
	db := orm.SetupModels(*logDBFlag, ":memory:")
	// mandatory, otherwise, bizarre errors occurs
	db.DB().SetMaxOpenConns(1)

	if *clientControlFlag {
		models.EngineSingloton.ControlMode = models.CLIENT_CONTROL
	} else {
		models.EngineSingloton.ControlMode = models.AUTONOMOUS
	}
	orm.BackRepo.Init(db)

	// seven days of simulation
	models.EngineSingloton.SetStartTime(time.Date(1676, time.January, 1, 0, 0, 0, 0, time.UTC))
	models.EngineSingloton.SetCurrentTime(models.EngineSingloton.GetStartTime())
	models.EngineSingloton.State = models.PAUSED
	models.EngineSingloton.Speed = 0.5 * 24 * 3600.0 // days per second
	log.Printf("Sim start \t\t\t%s\n", models.EngineSingloton.GetStartTime())

	// Three years
	models.EngineSingloton.SetEndTime(time.Date(1680, time.January, 1, 0, 0, 0, 0, time.UTC))
	log.Printf("Sim end  \t\t\t%s\n", models.EngineSingloton.GetEndTime())

	// PLUMBING n°1: callback for treating model specific action. In this case, see specific engine
	var simulation Simulation
	simulation.db = db
	models.EngineSingloton.Simulation = &simulation

	// attach the callback on the action singloton of the gongsim stack
	// legacy to be removed
	// controllers.ActionSinglotonID.ActionBackRepoCallback = singlotons.ActionSinglotonID

	// append a dummy agent to feed the discrete event engine with at least an event
	dummyAgent := new(DummyAgent)
	dummyAgent.Name = "Dummy"

	models.EngineSingloton.AppendAgent(dummyAgent)
	var step models.UpdateState
	step.SetFireTime(models.EngineSingloton.GetStartTime())
	step.Period = 1 * time.Second //
	step.Name = "update of planetary motion"
	dummyAgent.QueueEvent(&step)

	// setup controlers
	r := gin.Default()
	r.Use(cors.Default())

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db) // a gin Context can have a map of variable that is set up at runtime
		c.Next()
	})

	controllers.RegisterControllers(r)

	if *clientControlFlag {
		r.Run()
	} else {
		models.EngineSingloton.Run()
	}

	// r.StaticFS("/static/", http.Dir("/Users/thomaspeugeot/go/src/github.com/fullstack-lang/gong/stacks/gongsim/ng/dist/ng"))

}

type Simulation struct {
	db *gorm.DB
}

func (specificEngine *Simulation) EventFired(engine *models.Engine) {}

func (specificEngine *Simulation) HasAnyStateChanged(engine *models.Engine) bool { return true }

func (specificEngine *Simulation) Reset(engine *models.Engine) {}

func (specificEngine *Simulation) CommitAgents(engine *models.Engine)   {}
func (specificEngine *Simulation) CheckoutAgents(engine *models.Engine) {}
func (specificEngine *Simulation) GetLastCommitNb() uint                { return 0 }

type DummyAgent struct {
	models.Agent

	Name string
}

func (dummyAgent *DummyAgent) FireNextEvent() {
	event, _ := dummyAgent.GetNextEventAndRemoveIt()

	switch event.(type) {
	case *models.UpdateState:
		checkStateEvent := event.(*models.UpdateState)

		// post next event
		checkStateEvent.SetFireTime(checkStateEvent.GetFireTime().Add(checkStateEvent.Period))
		dummyAgent.QueueEvent(checkStateEvent)
	}
}
