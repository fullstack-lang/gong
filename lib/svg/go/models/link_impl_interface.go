package models

type LinkImplInterface interface {

	// LinkUpdated function is called each time a Link is modified
	LinkUpdated(updatedLink *Link)
}
