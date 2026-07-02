package models

type LinkAnchoredTextImplInterface interface {

	// AnchoredTextUpdated function is called each time a AnchoredText is modified
	AnchoredTextUpdated(updatedAnchoredText *LinkAnchoredText)
}
