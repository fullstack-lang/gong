package models

import (
	"log"
	"runtime/pprof"
	"time"
)

var DisplayWatch bool
var CpuProfile bool

// the watcher thread inspects the status of the simulation
func (command *Command) watcher() {

	var lastSimTime = command.Engine.currentTime

	// The period shall not too short for performance sake but not too long because the end user needs a responsive application
	//
	// checkoutSchedulerPeriod is the period of the "checkout scheduler"
	watcherPeriod := 500 * time.Millisecond
	var WatcherSchedulerPeriod = time.NewTicker(watcherPeriod)

	realtimeSimStart := time.Now()
	for {
		select {
		case t := <-WatcherSchedulerPeriod.C:

			_ = t

			// const layout = "Jan 2, 2006 at 15:04:05 (MST)"
			const layout = "15:04:05.999"
			measuredSimSpeed := float64(command.Engine.currentTime.Sub(lastSimTime)) / float64(watcherPeriod)
			if DisplayWatch {
				log.Printf("time %s, next %s, status %s, speed %f, speed request %f, Sim %s, Ho %s",
					time.Now().Format(layout), command.Engine.nextRealtimeHorizon.Format(layout),
					command.Engine.State, measuredSimSpeed, command.Engine.Speed,
					command.Engine.currentTime.Format(layout), command.Engine.nextSimulatedTimeHorizon.Format(layout))
			}
			lastSimTime = command.Engine.currentTime

			if CpuProfile {
				if time.Since(realtimeSimStart) > 20*time.Second {
					pprof.StopCPUProfile()
					log.Println("generated CPU profile")
				}
			}
		}
	}
}
