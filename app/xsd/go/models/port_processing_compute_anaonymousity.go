package models

func PostProcessingComputeAnonymousity(stage *Stage) {

	for ct := range *stage.GetInstancesSet[*ComplexType]() {
		ct.IsAnonymous = true
	}
	for _, ct := range SchemaSingloton.ComplexTypes {
		ct.IsAnonymous = false
	}
}
