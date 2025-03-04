package doc2svg

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// AnchoredTextImplLinkFieldName
// it meets the interface of AnchoredTextImplInterface
// and updates the link field name position
type AnchoredTextImplLinkFieldName struct {
	link         *gongdoc_models.Link
	gongdocStage *gongdoc_models.StageStruct
}

func NewAnchoredTextImplLinkFieldName(
	anchoredtext *gongdoc_models.Link,
	gongdocStage *gongdoc_models.StageStruct) (anchoredtextImplAnchoredText *AnchoredTextImplLinkFieldName) {

	anchoredtextImplAnchoredText = new(AnchoredTextImplLinkFieldName)
	anchoredtextImplAnchoredText.link = anchoredtext
	anchoredtextImplAnchoredText.gongdocStage = gongdocStage

	return
}

func (anchoredtextImplAnchoredText *AnchoredTextImplLinkFieldName) AnchoredTextUpdated(
	updatedAnchoredText *gongsvg_models.LinkAnchoredText) {

	log.Println("AnchoredTextImplAnchoredText:AnchoredTextUpdated")

	// update the anchoredtext
	anchoredtextImplAnchoredText.link.FieldOffsetX = updatedAnchoredText.X_Offset
	anchoredtextImplAnchoredText.link.FieldOffsetY = updatedAnchoredText.Y_Offset

	anchoredtextImplAnchoredText.gongdocStage.CommitWithSuspendedCallbacks()
}
