package models

// swagger:enum GongdocNodeType
type GongdocNodeType string

// values for Action Type
const (
	ROOT_OF_GONG_STRUCTS GongdocNodeType = "ROOT_OF_GONG_STRUCTS"
	GONG_STRUCT          GongdocNodeType = "GONG_STRUCT"

	ROOT_OF_BASIC_FIELDS GongdocNodeType = "ROOT_OF_BASIC_FIELDS"
	BASIC_FIELD          GongdocNodeType = "BASIC_FIELD"

	ROOT_OF_POINTER_TO_STRUCT_FIELDS GongdocNodeType = "ROOT_OF_POINTER_TO_STRUCT_FIELDS"
	POINTER_TO_STRUCT                GongdocNodeType = "POINTER_TO_STRUCT"

	ROOT_OF_SLICE_OF_POINTER_TO_GONG_STRUCT_FIELDS GongdocNodeType = "ROOT_OF_SLICE_OF_POINTER_TO_GONG_STRUCT_FIELDS"
	SLICE_OF_POINTER_TO_STRUCT                     GongdocNodeType = "SLICE_OF_POINTER_TO_STRUCT"
)
