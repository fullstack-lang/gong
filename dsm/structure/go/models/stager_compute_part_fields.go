package models

func (stager *Stager) computePartFields() {
	stage := stager.stage

	// reset fields
	for _, part := range GetGongstrucsSorted[*Part](stage) {
		part.owningSystem = nil
		part.inDataFlows = nil
		part.outDataFlows = nil
	}

	// compute owningSystem
	for _, system := range GetGongstrucsSorted[*System](stage) {
		for _, part := range system.Parts {
			if part != nil {
				part.owningSystem = system
			}
		}
		for _, externalPart := range system.ExternalParts {
			if externalPart != nil {
				externalPart.owningSystem = system
			}
		}
	}

	for _, dataFlow := range GetGongstrucsSorted[*DataFlow](stage) {
		if startExternalPart := dataFlow.StartExternalPart; startExternalPart != nil {
			startExternalPart.outDataFlows = append(startExternalPart.outDataFlows, dataFlow)
		}
		if endExternalPart := dataFlow.EndExternalPart; endExternalPart != nil {
			endExternalPart.inDataFlows = append(endExternalPart.inDataFlows, dataFlow)
		}
	}
}
