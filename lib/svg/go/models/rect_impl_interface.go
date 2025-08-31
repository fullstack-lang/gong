package models

type RectImplInterface interface {

	// RectUpdated function is called each time a Rect is modified
	RectUpdated(updatedRect *Rect)
}

type RectImplWithMouseEventInterface interface {

	// RectUpdated function is called each time a Rect is modified
	RectUpdatedWithMouseEvent(updatedRect *Rect, mouseEvent *Gong__MouseEvent)
}
