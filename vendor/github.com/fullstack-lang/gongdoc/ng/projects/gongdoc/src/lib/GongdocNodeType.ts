// generated from ng_file_enum.ts.go
export enum GongdocNodeType {
	// insertion point	
	ROOT_OF_DIAGRAMS = "ROOT_OF_DIAGRAMS",
	ROOT_OF_CLASS_DIAGRAMS = "ROOT_OF_CLASS_DIAGRAMS",
	ROOT_OF_STATE_DIAGRAMS = "ROOT_OF_STATE_DIAGRAMS",
	CLASS_DIAGRAM = "CLASS_DIAGRAM",
	STATE_DIAGRAM = "STATE_DIAGRAM",
	ROOT_OF_GONG_STRUCTS = "ROOT_OF_GONG_STRUCTS",
	GONG_STRUCT = "GONG_STRUCT",
	GONG_STRUCT_FIELD = "GONG_STRUCT_FIELD",
	ROOT_OF_GONG_ENUMS = "ROOT_OF_GONG_ENUMS",
	GONG_ENUM = "GONG_ENUM",
	GONG_ENUM_VALUE = "GONG_ENUM_VALUE",
	ROOT_OF_GONG_NOTES = "ROOT_OF_GONG_NOTES",
	GONG_NOTE = "GONG_NOTE",
	GONG_NOTE_LINK = "GONG_NOTE_LINK",
}

export interface GongdocNodeTypeSelect {
	value: string;
	viewValue: string;
}

export const GongdocNodeTypeList: GongdocNodeTypeSelect[] = [ // insertion point	
	{ value: GongdocNodeType.ROOT_OF_DIAGRAMS, viewValue: "ROOT_OF_DIAGRAMS" },
	{ value: GongdocNodeType.ROOT_OF_CLASS_DIAGRAMS, viewValue: "ROOT_OF_CLASS_DIAGRAMS" },
	{ value: GongdocNodeType.ROOT_OF_STATE_DIAGRAMS, viewValue: "ROOT_OF_STATE_DIAGRAMS" },
	{ value: GongdocNodeType.CLASS_DIAGRAM, viewValue: "CLASS_DIAGRAM" },
	{ value: GongdocNodeType.STATE_DIAGRAM, viewValue: "STATE_DIAGRAM" },
	{ value: GongdocNodeType.ROOT_OF_GONG_STRUCTS, viewValue: "ROOT_OF_GONG_STRUCTS" },
	{ value: GongdocNodeType.GONG_STRUCT, viewValue: "GONG_STRUCT" },
	{ value: GongdocNodeType.GONG_STRUCT_FIELD, viewValue: "GONG_STRUCT_FIELD" },
	{ value: GongdocNodeType.ROOT_OF_GONG_ENUMS, viewValue: "ROOT_OF_GONG_ENUMS" },
	{ value: GongdocNodeType.GONG_ENUM, viewValue: "GONG_ENUM" },
	{ value: GongdocNodeType.GONG_ENUM_VALUE, viewValue: "GONG_ENUM_VALUE" },
	{ value: GongdocNodeType.ROOT_OF_GONG_NOTES, viewValue: "ROOT_OF_GONG_NOTES" },
	{ value: GongdocNodeType.GONG_NOTE, viewValue: "GONG_NOTE" },
	{ value: GongdocNodeType.GONG_NOTE_LINK, viewValue: "GONG_NOTE_LINK" },
];
