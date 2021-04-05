package models

// swagger:enum GongdocCommandType
type GongdocCommandType string

// values for Action Type
const (
	MARSHALL_DIAGRAM       GongdocCommandType = "MARSHALL_ALL_DIAGRAMS"
	PRINT_ALL_DOCUMENTS    GongdocCommandType = "PRINT_ALL_DOCUMENTS"
	DIAGRAM_ELEMENT_DELETE GongdocCommandType = "DIAGRAM_ELEMENT_DELETE"
	DIAGRAM_ELEMENT_CREATE GongdocCommandType = "DIAGRAM_ELEMENT_CREATE"
)
