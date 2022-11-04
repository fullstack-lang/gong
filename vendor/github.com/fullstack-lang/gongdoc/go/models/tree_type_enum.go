package models

// swagger:enum GongdocNodeType
type GongdocNodeType string

// values for Action Type
const (
	ROOT_OF_DIAGRAMS       GongdocNodeType = "ROOT_OF_DIAGRAMS"
	ROOT_OF_CLASS_DIAGRAMS GongdocNodeType = "ROOT_OF_CLASS_DIAGRAMS"
	ROOT_OF_STATE_DIAGRAMS GongdocNodeType = "ROOT_OF_STATE_DIAGRAMS"

	CLASS_DIAGRAM GongdocNodeType = "CLASS_DIAGRAM"
	STATE_DIAGRAM GongdocNodeType = "STATE_DIAGRAM"

	ROOT_OF_GONG_STRUCTS GongdocNodeType = "ROOT_OF_GONG_STRUCTS"
	GONG_STRUCT          GongdocNodeType = "GONG_STRUCT"
	GONG_STRUCT_FIELD    GongdocNodeType = "GONG_STRUCT_FIELD"

	ROOT_OF_GONG_ENUMS GongdocNodeType = "ROOT_OF_GONG_ENUMS"
	GONG_ENUM          GongdocNodeType = "GONG_ENUM"
	GONG_ENUM_VALUE    GongdocNodeType = "GONG_ENUM_VALUE"

	ROOT_OF_GONG_NOTES GongdocNodeType = "ROOT_OF_GONG_NOTES"
	GONG_NOTE          GongdocNodeType = "GONG_NOTE"
)
