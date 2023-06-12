package models

type LinkAnchoredText struct {
	Name    string
	Content string

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
