package models

// enforceProducersConsumers computes the producers and consumers for each deliverable
// based on the inputs and outputs of all tasks.
// It first clears existing producer/consumer relationships, then rebuilds them
// by iterating through the tasks in the stage.
func (stager *Stager) enforceProducersConsumers() {
	stage := stager.stage

	// reset producers and consumers
	for _, deliverable := range GetGongstrucsSorted[*Deliverable](stage) {
		deliverable.producers = nil
		deliverable.consumers = nil
	}

	for _, task := range GetGongstrucsSorted[*Concern](stage) {
		// consumers are computed from [models.Task.Inputs]
		for _, deliverable := range task.Inputs {
			deliverable.consumers = append(deliverable.consumers, task)
		}

		// producers are computed from [models.Task.Outputs]
		for _, deliverable := range task.Outputs {
			deliverable.producers = append(deliverable.producers, task)
		}
	}
}
