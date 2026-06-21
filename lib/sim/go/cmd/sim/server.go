//go:build !js

package main

import (
	"log"
	"strconv"
	"time"

	m "github.com/fullstack-lang/gong/lib/sim/go/models"
	sim_stack "github.com/fullstack-lang/gong/lib/sim/go/stack"
	sim_static "github.com/fullstack-lang/gong/lib/sim/go/static"
)

func executeServer(args []string) {
	// setup the static file server and get the controller
	r := sim_static.ServeStaticFiles(logGINFlag)

	// setup stack
	stack := sim_stack.NewStack(r, "sim", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Probe.Refresh()
	stage := stack.Stage

	engine := new(m.Engine)
	engine.Name = "Simulation Engine"
	engine.DisplayFormat = "02 Jan 2006"
	engine.Stage(stage)

	// the gongsim command orchestrates the simulation engine regarding to the
	// the rest of the stack. It manages when the stage has to be commited to the
	// back repo or when the back repo has to be checked out to the stage
	command := m.NewCommand(stage, engine)
	_ = command

	// seven days of simulation
	engine.SetStartTime(time.Date(1676, time.January, 1, 0, 0, 0, 0, time.UTC))
	engine.SetCurrentTime(engine.GetStartTime())
	engine.State = m.PAUSED
	engine.Speed = 30 * 0.5 * 24 * 3600.0 // days per second
	log.Printf("Sim start \t\t\t%s\n", engine.GetStartTime())

	// Three years
	engine.SetEndTime(time.Date(1680, time.January, 1, 0, 0, 0, 0, time.UTC))
	log.Printf("Sim end  \t\t\t%s\n", engine.GetEndTime())

	// PLUMBING n°1: callback for treating model specific action. In this case, see specific engine
	var simulation m.Simulation
	engine.Simulation = &simulation

	// append a dummy agent to feed the discrete event engine with at least an event
	dummyAgent := new(m.DummyAgent)
	dummyAgent.Name = "Dummy"

	engine.AppendAgent(dummyAgent)
	var step m.UpdateState
	step.SetFireTime(engine.GetStartTime())
	step.Period = 1 * time.Second //
	step.Name = "update of planetary motion"
	dummyAgent.QueueEvent(&step)

	// start right away
	if play {
		command.Command = m.COMMAND_PLAY
	}
	if displayWatch {
		m.DisplayWatch = true
	}

	// commit simulation stage
	stage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
