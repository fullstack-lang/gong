package models

import (
	"fmt"
	"sort"
	"time"
)

func (stager *Stager) enforcePartShapeSemantic() (needCommit bool) {
	for _, diagramStructure := range GetGongstrucsSorted[*DiagramStructure](stager.stage) {

		for _, partShape := range diagramStructure.Part_Shapes {
			if partShape.WidthWeight == 0 {
				partShape.WidthWeight = 1.0
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("DiagramStructure %s: part shape %s WidthWeight set to 1.0", diagramStructure.Name, partShape.Name))
				}
			}
		}

		owningSystem := diagramStructure.owningSystem
		if owningSystem == nil {
			continue
		}

		// Map the expected order of parts based on the owning System
		partOrder := make(map[*Part]int)
		for i, part := range owningSystem.Parts {
			partOrder[part] = i
		}

		// Helper function to get order with fallback
		getOrder := func(p *Part) int {
			if order, ok := partOrder[p]; ok {
				return order
			}
			return 999999 // If not found, move it to the end
		}

		// Check if it's already sorted to avoid unnecessary commits
		isSorted := sort.SliceIsSorted(diagramStructure.Part_Shapes, func(i, j int) bool {
			p1 := diagramStructure.Part_Shapes[i].Part
			p2 := diagramStructure.Part_Shapes[j].Part
			return getOrder(p1) < getOrder(p2)
		})

		// Sort and flag the commit if they are out of order
		if !isSorted {
			sort.SliceStable(diagramStructure.Part_Shapes, func(i, j int) bool {
				p1 := diagramStructure.Part_Shapes[i].Part
				p2 := diagramStructure.Part_Shapes[j].Part
				return getOrder(p1) < getOrder(p2)
			})

			needCommit = true
			stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("DiagramStructure %s: part shapes reordered to match System definition", diagramStructure.Name))
		}
	}
	return
}
