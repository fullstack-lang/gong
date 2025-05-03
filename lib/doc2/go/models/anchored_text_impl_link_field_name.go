package models

import (
	"log"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// AnchoredTextImplLinkFieldName
// it meets the interface of AnchoredTextImplInterface
// and updates the link field name position
type AnchoredTextImplLinkFieldName struct {
	link         *LinkShape
	gongdocStage *Stage
}

func NewAnchoredTextImplLinkFieldName(
	anchoredtext *LinkShape,
	gongdocStage *Stage) (anchoredtextImplAnchoredText *AnchoredTextImplLinkFieldName) {

	anchoredtextImplAnchoredText = new(AnchoredTextImplLinkFieldName)
	anchoredtextImplAnchoredText.link = anchoredtext
	anchoredtextImplAnchoredText.gongdocStage = gongdocStage

	return
}

func (anchoredtextImplAnchoredText *AnchoredTextImplLinkFieldName) AnchoredTextUpdated(
	updatedAnchoredText *svg.LinkAnchoredText) {

	log.Println("AnchoredTextImplAnchoredText:AnchoredTextUpdated")

	// update the anchoredtext
	anchoredtextImplAnchoredText.link.FieldOffsetX = updatedAnchoredText.X_Offset
	anchoredtextImplAnchoredText.link.FieldOffsetY = updatedAnchoredText.Y_Offset

	anchoredtextImplAnchoredText.gongdocStage.CommitWithSuspendedCallbacks()
}
