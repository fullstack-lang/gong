package models

import (
	"fmt"
	"time"
)

// enforcePortSemanticRules checks
// - port cannot be both start and end
// - per part, cannot have more than one start port
func (stager *Stager) enforcePortSemanticRules() (needCommit bool) {

	for part := range *GetGongstructInstancesSetFromPointerType[*Part](stager.stage) {

		if part.TypeOfPart != nil && !part.IsPartNameNotSystemName && part.GetName() != part.TypeOfPart.GetName() {
			needCommit = true
			part.SetName(part.TypeOfPart.GetName())
			stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Port \"%s\" name overridden to \"%s\"",
				part.GetName(), part.GetName()))
		}
	}

	return
}
