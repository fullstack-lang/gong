package models

func (stager *Stager) computePartFields() {
	stage := stager.stage

	// reset fields
	for _, part := range stage.GetInstancesSorted[*Part]() {
		part.owningSystem = nil
		part.inDataFlows = nil
		part.outDataFlows = nil
	}

	// compute owningSystem
	for _, system := range stage.GetInstancesSorted[*System]() {
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

	for _, dataFlow := range stage.GetInstancesSorted[*DataFlow]() {
		if startExternalPart := dataFlow.StartExternalPart; startExternalPart != nil {
			startExternalPart.outDataFlows = append(startExternalPart.outDataFlows, dataFlow)
		}
		if endExternalPart := dataFlow.EndExternalPart; endExternalPart != nil {
			endExternalPart.inDataFlows = append(endExternalPart.inDataFlows, dataFlow)
		}
	}
}
