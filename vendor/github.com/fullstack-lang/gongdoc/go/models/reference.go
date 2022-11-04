package models

// Reference store usefull informations about the gongstruct or the gongenum
// refernced by a shape present in the diagrams
type Reference struct {
	Name        string // name of the gong struct
	NbInstances int
	Type        ReferenceType
}
