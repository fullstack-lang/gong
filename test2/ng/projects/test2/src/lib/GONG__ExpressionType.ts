// generated from ng_file_enum.ts.go
export enum GONG__ExpressionType {
	// insertion point	
	GONG__STRUCT_INSTANCE = "STRUCT_INSTANCE",
	GONG__FIELD_OR_CONST_VALUE = "FIELD_OR_CONST_VALUE",
	GONG__FIELD_VALUE = "FIELD_VALUE",
	GONG__ENUM_CAST_INT = "ENUM_CAST_INT",
	GONG__ENUM_CAST_STRING = "ENUM_CAST_STRING",
	GONG__IDENTIFIER_CONST = "IDENTIFIER_CONST",
}

export interface GONG__ExpressionTypeSelect {
	value: string;
	viewValue: string;
}

export const GONG__ExpressionTypeList: GONG__ExpressionTypeSelect[] = [ // insertion point	
	{ value: GONG__ExpressionType.GONG__STRUCT_INSTANCE, viewValue: "STRUCT_INSTANCE" },
	{ value: GONG__ExpressionType.GONG__FIELD_OR_CONST_VALUE, viewValue: "FIELD_OR_CONST_VALUE" },
	{ value: GONG__ExpressionType.GONG__FIELD_VALUE, viewValue: "FIELD_VALUE" },
	{ value: GONG__ExpressionType.GONG__ENUM_CAST_INT, viewValue: "ENUM_CAST_INT" },
	{ value: GONG__ExpressionType.GONG__ENUM_CAST_STRING, viewValue: "ENUM_CAST_STRING" },
	{ value: GONG__ExpressionType.GONG__IDENTIFIER_CONST, viewValue: "IDENTIFIER_CONST" },
];
