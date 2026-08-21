package models

func (stager *Stager) enforceStagerMaps() {
	stager.map_Element_Diagrams = make(map[AbstractType][]*DiagramFlossEquation)

	for _, diagramEquation := range GetGongstrucsSorted[*DiagramFlossEquation](stager.stage) {
		diagramEquation.map_Note_NoteShape = make(map[*Note]*NoteShape)
		for _, shape := range diagramEquation.Note_Shapes {
			if shape.Note != nil {
				diagramEquation.map_Note_NoteShape[shape.Note] = shape
			}
		}

		diagramEquation.map_Note_NoteComplexityShape = make(map[noteComplexityKey]*NoteComplexityShape)
		for _, shape := range diagramEquation.NoteComplexityShapes {
			key := noteComplexityKey{
				Note:       shape.Note,
				Complexity: shape.Complexity,
			}
			diagramEquation.map_Note_NoteComplexityShape[key] = shape
		}

		diagramEquation.map_Note_NotePerformanceShape = make(map[notePerformanceKey]*NotePerformanceShape)
		for _, shape := range diagramEquation.NotePerformanceShapes {
			key := notePerformanceKey{
				Note:        shape.Note,
				Performance: shape.Performance,
			}
			diagramEquation.map_Note_NotePerformanceShape[key] = shape
		}

		diagramEquation.map_Note_NoteEffortShape = make(map[noteEffortKey]*NoteEffortShape)
		for _, shape := range diagramEquation.NoteEffortShapes {
			key := noteEffortKey{
				Note:   shape.Note,
				Effort: shape.Effort,
			}
			diagramEquation.map_Note_NoteEffortShape[key] = shape
		}
	}
}
