package models

type LinkAnchoredText struct {
	Name string

	//gong:text width:300 height:300
	Content string

	// AutomaticLayout is true will have the front compute the
	// X_Offset / Y_Offset of the anchored text
	AutomaticLayout bool
	LinkAnchorType  LinkAnchorType

	// values if AutomaticLayout is false
	X_Offset float64
	Y_Offset float64

	TextAttributes

	Presentation
	Animates []*Animate

	Impl LinkAnchoredTextImplInterface
}

func (linkAnchoredText *LinkAnchoredText) OnAfterUpdate(stage *Stage, _, frontLinkAnchoredText *LinkAnchoredText) {

	if linkAnchoredText.Impl != nil {
		linkAnchoredText.Impl.AnchoredTextUpdated(frontLinkAnchoredText)
	}
}
