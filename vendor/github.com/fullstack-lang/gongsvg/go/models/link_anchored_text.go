package models

type LinkAnchoredText struct {
	Name    string
	Content string

	// AutomaticLayout is true will have the front compute the
	// X_Offset / Y_Offset of the anchored text
	AutomaticLayout bool
	LinkAnchorType  LinkAnchorType

	// values if AutomaticLayout is false
	X_Offset float64
	Y_Offset float64

	// "bold", "normal", ...
	FontWeight string

	Presentation
	Animates []*Animate

	Impl LinkAnchoredTextImplInterface
}

func (linkAnchoredText *LinkAnchoredText) OnAfterUpdate(stage *StageStruct, _, frontLinkAnchoredText *LinkAnchoredText) {

	if linkAnchoredText.Impl != nil {
		linkAnchoredText.Impl.AnchoredTextUpdated(frontLinkAnchoredText)
	}
}
