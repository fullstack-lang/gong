package models

// swagger:enum Etape
type Etape string

// values for Action Type
const (
	AVANT_CALCUL Etape = "AVANT_CALCUL" // iota // Parse the spinosa model (temp)
	APRES_CALCUL Etape = "APRES_CALCUL"
)