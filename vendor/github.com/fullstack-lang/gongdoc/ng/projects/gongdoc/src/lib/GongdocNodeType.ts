// generated from ng_file_enum.ts.go
export enum GongdocNodeType {
	// insertion point	
	ROOT_OF_GONG_STRUCTS = "ROOT_OF_GONG_STRUCTS",
	GONG_STRUCT = "GONG_STRUCT",
	ROOT_OF_BASIC_FIELDS = "ROOT_OF_BASIC_FIELDS",
	BASIC_FIELD = "BASIC_FIELD",
	ROOT_OF_TIME_FIELDS = "ROOT_OF_TIME_FIELDS",
	TIME_FIELD = "TIME_FIELD",
	ROOT_OF_POINTER_TO_STRUCT_FIELDS = "ROOT_OF_POINTER_TO_STRUCT_FIELDS",
	POINTER_TO_STRUCT = "POINTER_TO_STRUCT",
	ROOT_OF_SLICE_OF_POINTER_TO_GONG_STRUCT_FIELDS = "ROOT_OF_SLICE_OF_POINTER_TO_GONG_STRUCT_FIELDS",
	SLICE_OF_POINTER_TO_STRUCT = "SLICE_OF_POINTER_TO_STRUCT",
	ROOT_OF_M_N_ASSOCIATION_FIELDS = "ROOT_OF_M_N_ASSOCIATION_FIELDS",
	M_N_ASSOCIATION_FIELD = "M_N_ASSOCIATION_FIELD",
	ROOT_OF_GONG_NOTES = "ROOT_OF_GONG_NOTES",
	GONG_NOTE = "GONG_NOTE",
}

export interface GongdocNodeTypeSelect {
	value: string;
	viewValue: string;
}

export const GongdocNodeTypeList: GongdocNodeTypeSelect[] = [ // insertion point	
	{ value: GongdocNodeType.ROOT_OF_GONG_STRUCTS, viewValue: "ROOT_OF_GONG_STRUCTS" },
	{ value: GongdocNodeType.GONG_STRUCT, viewValue: "GONG_STRUCT" },
	{ value: GongdocNodeType.ROOT_OF_BASIC_FIELDS, viewValue: "ROOT_OF_BASIC_FIELDS" },
	{ value: GongdocNodeType.BASIC_FIELD, viewValue: "BASIC_FIELD" },
	{ value: GongdocNodeType.ROOT_OF_TIME_FIELDS, viewValue: "ROOT_OF_TIME_FIELDS" },
	{ value: GongdocNodeType.TIME_FIELD, viewValue: "TIME_FIELD" },
	{ value: GongdocNodeType.ROOT_OF_POINTER_TO_STRUCT_FIELDS, viewValue: "ROOT_OF_POINTER_TO_STRUCT_FIELDS" },
	{ value: GongdocNodeType.POINTER_TO_STRUCT, viewValue: "POINTER_TO_STRUCT" },
	{ value: GongdocNodeType.ROOT_OF_SLICE_OF_POINTER_TO_GONG_STRUCT_FIELDS, viewValue: "ROOT_OF_SLICE_OF_POINTER_TO_GONG_STRUCT_FIELDS" },
	{ value: GongdocNodeType.SLICE_OF_POINTER_TO_STRUCT, viewValue: "SLICE_OF_POINTER_TO_STRUCT" },
	{ value: GongdocNodeType.ROOT_OF_M_N_ASSOCIATION_FIELDS, viewValue: "ROOT_OF_M_N_ASSOCIATION_FIELDS" },
	{ value: GongdocNodeType.M_N_ASSOCIATION_FIELD, viewValue: "M_N_ASSOCIATION_FIELD" },
	{ value: GongdocNodeType.ROOT_OF_GONG_NOTES, viewValue: "ROOT_OF_GONG_NOTES" },
	{ value: GongdocNodeType.GONG_NOTE, viewValue: "GONG_NOTE" },
];
