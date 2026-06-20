package models

import (
	"fmt"
	"sort"
	"time"
)

func (stager *Stager) enforceParticipantShapeSemantic() (needCommit bool) {
	for _, diagramProcess := range GetGongstrucsSorted[*DiagramProcess](stager.stage) {

		for _, participantShape := range diagramProcess.Participant_Shapes {
			if participantShape.WidthWeight == 0 {
				participantShape.WidthWeight = 1.0
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("DiagramProcess %s: participant shape %s WidthWeight set to 1.0", diagramProcess.Name, participantShape.Name))
				}
			}
		}

		owningProcess := diagramProcess.owningProcess
		if owningProcess == nil {
			continue
		}

		// Map the expected order of participants based on the owning Process
		participantOrder := make(map[*Participant]int)
		for i, participant := range owningProcess.Participants {
			participantOrder[participant] = i
		}

		// Helper function to get order with fallback
		getOrder := func(p *Participant) int {
			if order, ok := participantOrder[p]; ok {
				return order
			}
			return 999999 // If not found, move it to the end
		}

		// Check if it's already sorted to avoid unnecessary commits
		isSorted := sort.SliceIsSorted(diagramProcess.Participant_Shapes, func(i, j int) bool {
			p1 := diagramProcess.Participant_Shapes[i].Participant
			p2 := diagramProcess.Participant_Shapes[j].Participant
			return getOrder(p1) < getOrder(p2)
		})

		// Sort and flag the commit if they are out of order
		if !isSorted {
			sort.SliceStable(diagramProcess.Participant_Shapes, func(i, j int) bool {
				p1 := diagramProcess.Participant_Shapes[i].Participant
				p2 := diagramProcess.Participant_Shapes[j].Participant
				return getOrder(p1) < getOrder(p2)
			})

			needCommit = true
			stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("DiagramProcess %s: participant shapes reordered to match Process definition", diagramProcess.Name))
		}
	}
	return
}
