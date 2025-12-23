package models

// enforceProducersConsumers computes the producers and consumers for each product
// based on the inputs and outputs of all tasks.
// It first clears existing producer/consumer relationships, then rebuilds them
// by iterating through the tasks in the stage.
func (stager *Stager) enforceProducersConsumers() {
	stage := stager.stage

	// reset producers and consumers
	for _, product := range GetGongstrucsSorted[*Product](stage) {
		product.producers = nil
		product.consumers = nil
	}

	for _, task := range GetGongstrucsSorted[*Task](stage) {
		// consumers are computed from [models.Task.Inputs]
		for _, product := range task.Inputs {
			product.consumers = append(product.consumers, task)
		}

		// producers are computed from [models.Task.Outputs]
		for _, product := range task.Outputs {
			product.producers = append(product.producers, task)
		}
	}
}
