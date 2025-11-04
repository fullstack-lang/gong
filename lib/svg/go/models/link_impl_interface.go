package models

type LinkImplInterface interface {

	// LinkUpdated function is called each time a Link is modified
	LinkUpdated(updatedLink *Link)
}

type LinkImplWithMouseEventInterface interface {

	// LinkUpdated function is called each time a Link is modified
	LinkUpdatedWithMouseEvent(updatedLink *Link, mouseEvent *Gong__MouseEvent)
}
