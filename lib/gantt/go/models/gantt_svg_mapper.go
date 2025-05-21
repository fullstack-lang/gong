package models

// type of the singloton for interception of gantt commit in order to generate
// the svg
type GanttSVGMapper struct {
	ganttToRender *Gantt

	// to overwrite the gantt definition file
	GanttOuputFile string
}
