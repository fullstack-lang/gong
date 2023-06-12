package models

type RectImplInterface interface {

	// RectUpdated function is called each time a Rect is modified
	RectUpdated(updatedRect *Rect)
}
