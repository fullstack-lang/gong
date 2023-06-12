package doc2svg

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

// AnchoredTextImplLinkSourceMultiplicity
// it meets the interface of AnchoredTextImplInterface
// and updates the link field name position
type AnchoredTextImplLinkSourceMultiplicity struct {
	link         *gongdoc_models.Link
	gongdocStage *gongdoc_models.StageStruct
}

func NewAnchoredTextImplLinkSourceMultiplicity(
	anchoredtext *gongdoc_models.Link,
	gongdocStage *gongdoc_models.StageStruct) (anchoredtextImplAnchoredText *AnchoredTextImplLinkSourceMultiplicity) {

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
