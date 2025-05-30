package doc2svg

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// AnchoredTextImplLinkSourceMultiplicity
// it meets the interface of AnchoredTextImplInterface
// and updates the link field name position
type AnchoredTextImplLinkSourceMultiplicity struct {
	link         *gongdoc_models.Link
	gongdocStage *gongdoc_models.Stage
}

func NewAnchoredTextImplLinkSourceMultiplicity(
	anchoredtext *gongdoc_models.Link,
	gongdocStage *gongdoc_models.Stage) (anchoredtextImplAnchoredText *AnchoredTextImplLinkSourceMultiplicity) {

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
