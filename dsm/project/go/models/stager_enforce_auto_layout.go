package models

// enforceAutoLayout enforces that all non-time diagrams in auto layout mode
// have their shapes and links laid out according to the abstract hierarchy and order.
func (stager *Stager) enforceAutoLayout() (needCommit bool) {
	for _, diagram := range stager.stage.GetInstancesSorted[*Diagram]() {
		if !diagram.IsInAutoLayoutMode || diagram.IsTimeDiagram {
			continue
		}
		if layoutDiagram(diagram, stager) {
			needCommit = true
		}
	}
	return
}
