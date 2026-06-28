package models

func PostProcessing(stage *Stage) {
	PostProcessingComputeAnonymousity(stage)
	PostProcessingNames(stage)
	PostProcessingAnalyzeXSDStructure(stage)
}
