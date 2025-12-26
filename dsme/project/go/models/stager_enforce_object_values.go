package models

func (stager *Stager) enforceObjectValues() (needCommit bool) {
	const (
		defaultBoxWidth  = 250.0
		defaultBoxHeigth = 100.0
	)

	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
		if diagram.DefaultBoxHeigth == 0 {
			diagram.DefaultBoxHeigth = defaultBoxHeigth
			needCommit = true
		}
		if diagram.DefaultBoxWidth == 0 {
			diagram.DefaultBoxWidth = defaultBoxWidth
			needCommit = true
		}
	}
	return
}
