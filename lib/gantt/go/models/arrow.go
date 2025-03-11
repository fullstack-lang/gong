package models

type Arrow struct {
	Name string
	From *Bar
	To   *Bar

	// if not Zero, is taken into account when drawing the bar
	OptionnalColor  string
	OptionnalStroke string
}
