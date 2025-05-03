package models

import (
	"log"

	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// AnchoredTextImplLinkTargetMultiplicity
// it meets the interface of AnchoredTextImplInterface
// and updates the link field name position
type AnchoredTextImplLinkTargetMultiplicity struct {
	link         *LinkShape
	gongdocStage *Stage
}

func NewAnchoredTextImplLinkTargetMultiplicity(
	anchoredtext *LinkShape,
	gongdocStage *Stage) (anchoredtextImplAnchoredText *AnchoredTextImplLinkTargetMultiplicity) {

	anchoredtextImplAnchoredText = new(AnchoredTextImplLinkTargetMultiplicity)
	anchoredtextImplAnchoredText.link = anchoredtext
	anchoredtextImplAnchoredText.gongdocStage = gongdocStage

	return
}

func (anchoredtextImplAnchoredText *AnchoredTextImplLinkTargetMultiplicity) AnchoredTextUpdated(
	updatedAnchoredText *gongsvg_models.LinkAnchoredText) {

	log.Println("AnchoredTextImplTargetMultiplicity:AnchoredTextUpdated")

	// update the anchoredtext
	anchoredtextImplAnchoredText.link.TargetMultiplicityOffsetX = updatedAnchoredText.X_Offset
	anchoredtextImplAnchoredText.link.TargetMultiplicityOffsetY = updatedAnchoredText.Y_Offset

	anchoredtextImplAnchoredText.gongdocStage.CommitWithSuspendedCallbacks()
}
