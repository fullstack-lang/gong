package models

func PostProcessingAnalyzeXSDStructure(stage *Stage) {

	schemas := stage.GetInstancesSorted[*Schema]()
	if len(schemas) == 1 {
		schemas[0].SetParentAndChildren(schemas[0])
	}

	// characterize complex types that are inlined
	for element := range *stage.GetInstancesSet[*Element]() {

		if element.ComplexType != nil {
			element.ComplexType.IsAnonymous = true
			element.ComplexType.OuterElement = element
		}

		for _, group := range element.Groups {
			group.IsAnonymous = true
			group.OuterElement = element
		}
	}

}
