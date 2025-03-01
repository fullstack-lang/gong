package models

// swagger:enum BEnumType
type BEnumType string

// values for EnumType
const (
	BENUM_VAL1 BEnumType = "BENUM_VAL1_NOT_THE_SAME"
	BENUM_VAL2 BEnumType = "BENUM_VAL2"
)

// BstructMap is a Map of all Bstruct via their Name
type BstructMap map[string]*Bstruct

var BstructStore BstructMap = make(map[string]*Bstruct, 0)
