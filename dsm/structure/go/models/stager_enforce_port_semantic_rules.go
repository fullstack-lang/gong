package models

import (
	"fmt"
	"time"
)

// enforcePortSemanticRules checks
// - port cannot be both start and end
// - per part, cannot have more than one start port
func (stager *Stager) enforcePortSemanticRules() (needCommit bool) {

	for port := range *GetGongstructInstancesSetFromPointerType[*Port](stager.stage) {

		if port.IsStartPort && port.IsEndPort {
			needCommit = true
			port.IsEndPort = false
			stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Port \"%s\" cannot be both start and end",
				port.GetName()))
		}

		if port.Type != nil && !port.IsPortNameNotSystemName && port.GetName() != port.Type.GetName() {
			needCommit = true
			port.SetName(port.Type.GetName())
			stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Port \"%s\" name overridden to \"%s\"",
				port.GetName(), port.GetName()))
		}
	}

	for part := range *GetGongstructInstancesSetFromPointerType[*Part](stager.stage) {
		nbStartPort := 0
		for _, port := range part.Ports {
			if nbStartPort == 1 && port.IsStartPort {
				needCommit = true
				port.IsStartPort = false
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Port \"%s\" is set to not Start since there cant be more than one start in part \"%s\"",
						port.GetName(), part.GetName()))
			}
			if nbStartPort == 0 && port.IsStartPort {
				nbStartPort++
			}
		}
	}

	return
}
