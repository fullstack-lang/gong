package models

import "time"

func (stager *Stager) enforce_semantic_diagrams() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {

		var zeroDate time.Time
		if diagram.StartDate.Equal(zeroDate) {
			diagram.StartDate = time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)
			needCommit = true
		}
		if diagram.EndDate.Equal(zeroDate) {
			diagram.EndDate = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
			needCommit = true
		}

		if diagram.AlignDatesToFiveYears() {
			needCommit = true
		}

		if diagram.EndDate.Before(diagram.StartDate) {
			diagram.EndDate = diagram.StartDate
			needCommit = true
		}

		if diagram.MovementRectAnchorType == "" {
			diagram.MovementRectAnchorType = RECT_RIGHT
			needCommit = true
		}
		if diagram.MovementTextAnchorType == "" {
			diagram.MovementTextAnchorType = TEXT_ANCHOR_END
			needCommit = true
		}
		if diagram.MovementDateRectAnchorType == "" {
			diagram.MovementDateRectAnchorType = RECT_BOTTOM_LEFT
			needCommit = true
		}
		if diagram.MovementDateTextAnchorType == "" {
			diagram.MovementDateTextAnchorType = TEXT_ANCHOR_START
			needCommit = true
		}
		if diagram.MovementPlacesRectAnchorType == "" {
			diagram.MovementPlacesRectAnchorType = RECT_BOTTOM_RIGHT
			needCommit = true
		}
		if diagram.MovementPlacesTextAnchorType == "" {
			diagram.MovementPlacesTextAnchorType = TEXT_ANCHOR_END
			needCommit = true
		}
	}

	return
}
