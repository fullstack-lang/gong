package models

func PostProcessingComputeAnonymousity(stage *Stage) {

	for ct := range *GetGongstructInstancesSet[ComplexType](stage) {
		ct.IsAnonymous = true
	}
	for _, ct := range SchemaSingloton.ComplexTypes {
		ct.IsAnonymous = false
	}
}
