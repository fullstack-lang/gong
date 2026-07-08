package models

import (
	"fmt"
	"time"
)

// enforcePlantNM ensures that N is always less than or equal to M for each Plant
func (stager *Stager) enforcePlantNM() (needCommit bool) {
	for _, plant := range GetGongstrucsSorted[*Plant](stager.stage) {
		if plant.N > plant.M {
			// Swap N and M to ensure N <= M
			plant.N, plant.M = plant.M, plant.N
			needCommit = true

			stager.probeForm.AddNotification(time.Now(),
				fmt.Sprintf("Plant %s: Swapped N and M because N > M (Now N=%d, M=%d)", plant.Name, plant.N, plant.M))
		}
	}
	return
}
