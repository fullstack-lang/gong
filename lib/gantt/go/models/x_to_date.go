package models

import "time"

// convert X coordinate to date
func XtoDate(leftX, x, rightX float64, ganttStartDate, ganttEndDate time.Time) (date time.Time) {

	// compute the geometric ratio of X compared to the the distance between RightX and LeftX
	ratio := (x - leftX) / (rightX - leftX)

	ganttDuration := ganttEndDate.Sub(ganttStartDate)

	durationSinceGanttStart := time.Duration(ratio * float64(ganttDuration))

	date = ganttStartDate.Add(durationSinceGanttStart)

	return
}
