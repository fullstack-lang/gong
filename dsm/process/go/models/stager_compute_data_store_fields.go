package models

func (stager *Stager) computeDataStoreFields() {
	stage := stager.stage

	for _, dataStore := range stage.GetInstancesSorted[*DataStore]() {
		dataStore.inDataFlows = nil
		dataStore.outDataFlows = nil
	}

	for _, library := range stage.GetInstancesSorted[*Library]() {
		for _, dataStore := range library.RootDataStores {
			if dataStore != nil {
				dataStore.SetOwningLibrary(library)
			}
		}
	}

	for _, dataFlow := range stage.GetInstancesSorted[*DataFlow]() {
		if dataFlow.StartDataStore != nil {
			dataFlow.StartDataStore.outDataFlows = append(dataFlow.StartDataStore.outDataFlows, dataFlow)
		}
		if dataFlow.EndDataStore != nil {
			dataFlow.EndDataStore.inDataFlows = append(dataFlow.EndDataStore.inDataFlows, dataFlow)
		}
	}
}
