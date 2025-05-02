package models

import (
	"log"

	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// AnchoredTextImplLinkSourceMultiplicity
// it meets the interface of AnchoredTextImplInterface
// and updates the link field name position
type AnchoredTextImplLinkSourceMultiplicity struct {
	link         *LinkShape
	gongdocStage *Stage
}

func NewAnchoredTextImplLinkSourceMultiplicity(
	anchoredtext *LinkShape,
	gongdocStage *Stage) (anchoredtextImplAnchoredText *AnchoredTextImplLinkSourceMultiplicity) {

	anchoredtextImplAnchoredText = new(AnchoredTextImplLinkSourceMultiplicity)
	anchoredtextImplAnchoredText.link = anchoredtext
	anchoredtextImplAnchoredText.gongdocStage = gongdocStage

	return
}

func (anchoredtextImplAnchoredText *AnchoredTextImplLinkSourceMultiplicity) AnchoredTextUpdated(
	updatedAnchoredText *gongsvg_models.LinkAnchoredText) {

	log.Println("AnchoredTextImplSourceMultiplicity:AnchoredTextUpdated")

	// update the anchoredtext
	anchoredtextImplAnchoredText.link.SourceMultiplicityOffsetX = updatedAnchoredText.X_Offset
	anchoredtextImplAnchoredText.link.SourceMultiplicityOffsetY = updatedAnchoredText.Y_Offset

	anchoredtextImplAnchoredText.gongdocStage.CommitWithSuspendedCallbacks()
}
