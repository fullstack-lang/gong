// generated from ng_file_enum.ts.go
export enum GongdocCommandType {
	// insertion point	
	DIAGRAM_ELEMENT_CREATE = "DIAGRAM_ELEMENT_CREATE",
	DIAGRAM_ELEMENT_DELETE = "DIAGRAM_ELEMENT_DELETE",
	MARSHALL_DIAGRAM = "MARSHALL_ALL_DIAGRAMS",
	PRINT_ALL_DOCUMENTS = "PRINT_ALL_DOCUMENTS",
}

export interface GongdocCommandTypeSelect {
	value: string;
	viewValue: string;
}

export const GongdocCommandTypeList: GongdocCommandTypeSelect[] = [ // insertion point	
	{ value: 'DIAGRAM_ELEMENT_CREATE', viewValue: '"DIAGRAM_ELEMENT_CREATE"' },
	{ value: 'DIAGRAM_ELEMENT_DELETE', viewValue: '"DIAGRAM_ELEMENT_DELETE"' },
	{ value: 'MARSHALL_DIAGRAM', viewValue: '"MARSHALL_ALL_DIAGRAMS"' },
	{ value: 'PRINT_ALL_DOCUMENTS', viewValue: '"PRINT_ALL_DOCUMENTS"' },
];
