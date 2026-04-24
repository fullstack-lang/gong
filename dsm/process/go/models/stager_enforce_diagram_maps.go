package models

func (stager *Stager) enforceDiagramMaps() {
	stager.map_Element_Diagrams = make(map[AbstractType][]*DiagramProcess)

	for _, diagram := range GetGongstrucsSorted[*DiagramProcess](stager.stage) {
		_ = diagram
	}
}
