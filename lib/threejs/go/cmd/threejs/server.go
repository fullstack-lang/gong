//go:build !js

package main

import (
	"log"
	"strconv"
	"time"

	"github.com/fullstack-lang/gong/lib/threejs/go/models"
	threejs_stack "github.com/fullstack-lang/gong/lib/threejs/go/stack"
	threejs_static "github.com/fullstack-lang/gong/lib/threejs/go/static"
)

func executeServer() {

	r := threejs_static.ServeStaticFiles(logGINFlag)

	// setup
	// - model stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := threejs_stack.NewStack(r, "threejs", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	if generateStrangeFormsFlag {
		models.GenerateReferenceScene(stack.Stage)
	} else {
		stack.Stage.Commit()
	}

	// initiates the UX loop
	models.NewStager(
		r,
		stack.Stage,
		stack.Probe,
	)

	if testRenderingTimeFlag {
		log.Println("[Backend TRACE] Test rendering time mode enabled: Toggling BoxGeometry height every 5s")

		for canvas := range *models.GetGongstructInstancesSet[models.Canvas](stack.Stage) {
			canvas.IsWithLastRenderingUpdate = true
		}

		go func() {
			var stageCommitTime time.Time
			ticker := time.NewTicker(5 * time.Second)
			delta := 1.0
			for range ticker.C {
				for box := range *models.GetGongstructInstancesSet[models.BoxGeometry](stack.Stage) {
					box.Height += delta
					log.Printf("[Backend TRACE] Toggled BoxGeometry '%s' Height to %.2f", box.Name, box.Height)
				}
				delta = -delta

				for canvas := range *models.GetGongstructInstancesSet[models.Canvas](stack.Stage) {
					canvas.IsWithLastRenderingUpdate = true
				}

				stageCommitTime = time.Now()
				stack.Stage.Commit()
				log.Printf("[Backend TRACE] Stage committed at %s", stageCommitTime.Format("15:04:05.000000"))
			}
		}()
	}

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
