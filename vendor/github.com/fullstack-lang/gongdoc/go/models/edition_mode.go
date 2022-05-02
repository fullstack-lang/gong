package models

// swagger:enum EditionMode
type EditionMode string

// values for Action Type
const (
	PRODUCTION_MODE  EditionMode = "PRODUCTION_MODE"
	DEVELOPMENT_MODE EditionMode = "DEVELOPMENT_MODE"
)
