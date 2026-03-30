package models

func PostProcessingAnalyzeXSDStructure(stage *Stage) {

	// characterize complex types that are inlined
	for element := range *GetGongstructInstancesSet[Element](stage) {

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
