package models

// GanttStacksNames - enumeration of possible 'type' values for an HTML <input> element
// swagger:enum GanttStacksNames
type GanttStacksNames string

// values for TableExtraNameEnum
const (
	SvgStackName   GanttStacksNames = "svg"   // the gantt application launch a svg stack for visualisation. this is its name
	GanttStackName GanttStacksNames = "gantt" // the default gantt stack has this name
)
