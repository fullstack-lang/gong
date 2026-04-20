package models

func (stager *Stager) enforceDiagramMaps() {
	stager.map_Element_Diagrams = make(map[AbstractType][]*Diagram)

	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
		_ = diagram
	}
}
