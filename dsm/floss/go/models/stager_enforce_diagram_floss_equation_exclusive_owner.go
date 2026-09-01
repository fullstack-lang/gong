package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceDiagramFlossEquationExclusiveOwner() (needCommit bool) {
	rmSys := GetSliceOfPointersReverseMap[System, DiagramFlossEquation](
		GetAssociationName[System]().DiagramFlossEquations[0].Name, stager.stage)
	rmCA := GetSliceOfPointersReverseMap[CompareAnalysis, DiagramFlossEquation](
		GetAssociationName[CompareAnalysis]().DiagramFlossEquations[0].Name, stager.stage)

	for diagram := range *GetGongstructInstancesSetFromPointerType[*DiagramFlossEquation](stager.stage) {
		owningSystems := rmSys[diagram]
		owningCompareAnalyses := rmCA[diagram]
		totalOwners := len(owningSystems) + len(owningCompareAnalyses)

		if totalOwners == 0 {
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("DiagramFlossEquation \"%s\" has no owning System or CompareAnalysis, unstaging", diagram.Name))
			}
			diagram.SetOwningSystem(nil)
			diagram.SetOwningCompareAnalysis(nil)
			diagram.UnstageVoid(stager.stage)
			needCommit = true
			continue
		}

		if totalOwners > 1 {
			if stager.probeForm != nil {
				if len(owningSystems) > 0 && len(owningCompareAnalyses) > 0 {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("DiagramFlossEquation \"%s\" is attached to both System and CompareAnalysis (mutually exclusive), unstaging", diagram.Name))
				} else if len(owningSystems) > 1 {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("DiagramFlossEquation \"%s\" has more than one owning System, unstaging", diagram.Name))
				} else {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("DiagramFlossEquation \"%s\" has more than one owning CompareAnalysis, unstaging", diagram.Name))
				}
			}
			diagram.SetOwningSystem(nil)
			diagram.SetOwningCompareAnalysis(nil)
			diagram.UnstageVoid(stager.stage)
			needCommit = true
			continue
		}

		// Exactly one owner (mutually exclusive)
		if len(owningSystems) == 1 {
			diagram.SetOwningSystem(owningSystems[0])
			diagram.SetOwningCompareAnalysis(nil)
		} else if len(owningCompareAnalyses) == 1 {
			diagram.SetOwningCompareAnalysis(owningCompareAnalyses[0])
			diagram.SetOwningSystem(nil)
		}
	}
	return
}
