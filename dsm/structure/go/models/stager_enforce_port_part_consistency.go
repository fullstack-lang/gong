package models

import (
	"fmt"
	"slices"
	"time"
)

// enforcePortPartConsistency unstages Ports that do not belong to any part
// or that belong to more than one part.
func (stager *Stager) enforcePortPartConsistency() (needCommit bool) {
	stage := stager.stage

	// 1. Build a reverse map from Port to its owning Parts
	portToParts := make(map[*Port][]*Part)
	for _, part := range GetGongstrucsSorted[*Part](stage) {
		for _, port := range part.Ports {
			portToParts[port] = append(portToParts[port], part)
		}
	}

	// 2. Iterate over all ports and check their ownership
	for _, port := range GetGongstrucsSorted[*Port](stage) {
		owners := portToParts[port]

		// Unstage ports with no owner
		if len(owners) == 0 {
			port.UnstageVoid(stage)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Unstaged orphan port \"%s\" (no part)", port.GetName()))
			}
			continue // continue to next port
		}

		// Unstage ports with more than one owner
		if len(owners) > 1 {
			// To make the operation idempotent and clear, we remove the port from all its owners' lists
			// before unstaging it.
			for _, owner := range owners {
				if idx := slices.Index(owner.Ports, port); idx != -1 {
					owner.Ports = slices.Delete(owner.Ports, idx, idx+1)
				}
			}
			port.UnstageVoid(stage)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Unstaged port \"%s\" with multiple parts", port.GetName()))
			}
		}
	}

	return
}
