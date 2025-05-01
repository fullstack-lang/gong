package models

import (
	"log"

	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// LinkImplLink
// it meets the interface of LinkImplInterface
type LinkImplLink struct {
	link         *Link
	gongdocStage *Stage
}

func NewLinkImplLink(
	link *Link,
	gongdocStage *Stage) (linkImplLink *LinkImplLink) {

	linkImplLink = new(LinkImplLink)
	linkImplLink.link = link
	linkImplLink.gongdocStage = gongdocStage

	return
}

func (linkImplLink *LinkImplLink) LinkUpdated(updatedLink *gongsvg_models.Link) {

	log.Println("LinkImplLink:LinkUpdated")

	// update the link
	linkImplLink.link.StartOrientation = OrientationType(updatedLink.StartOrientation)
	linkImplLink.link.StartRatio = updatedLink.StartRatio
	linkImplLink.link.EndOrientation = OrientationType(updatedLink.EndOrientation)
	linkImplLink.link.EndRatio = updatedLink.EndRatio
	linkImplLink.link.CornerOffsetRatio = updatedLink.CornerOffsetRatio

	linkImplLink.gongdocStage.CommitWithSuspendedCallbacks()
}
