package models

func (stager *Stager) computePortFields() {
	stage := stager.stage

	// reset fields
	for _, port := range GetGongstrucsSorted[*Port](stage) {
		port.inControlFlows = nil
		port.outControlFlows = nil
		port.inDataFlows = nil
		port.outDataFlows = nil
		port.owningPart = nil
	}

	// compute owningPart
	for _, part := range GetGongstrucsSorted[*Part](stage) {
		for _, port := range part.Ports {
			if port != nil {
				port.owningPart = part
			}
		}
	}

	// compute outControlFlows and inControlFlows
	for _, controlFlow := range GetGongstrucsSorted[*ControlFlow](stage) {
		if controlFlow.Start != nil {
			controlFlow.Start.outControlFlows = append(controlFlow.Start.outControlFlows, controlFlow)
		}
		if controlFlow.End != nil {
			controlFlow.End.inControlFlows = append(controlFlow.End.inControlFlows, controlFlow)
		}
	}

	// compute outDataFlows and inDataFlows
	for _, dataFlow := range GetGongstrucsSorted[*DataFlow](stage) {
		if dataFlow.StartPort != nil {
			dataFlow.StartPort.outDataFlows = append(dataFlow.StartPort.outDataFlows, dataFlow)
		}
		if dataFlow.EndPort != nil {
			dataFlow.EndPort.inDataFlows = append(dataFlow.EndPort.inDataFlows, dataFlow)
		}
	}

}
