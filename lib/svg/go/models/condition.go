package models

// A Condition is something that is triggered by a something (like hovering a shape)
// and that condition the display of another shape
type Condition struct {
	Name string
}

// ShapeConditions
type ShapeConditions struct {
	HoveringTrigger   []*Condition
	DisplayConditions []*Condition
}
