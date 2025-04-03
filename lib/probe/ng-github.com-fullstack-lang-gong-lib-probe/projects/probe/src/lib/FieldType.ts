// generated from ng_file_enum.ts.go
export enum FieldType {
	// insertion point	
	FieldTypeInt = 0,
	FieldTypeFloat64 = 1,
	FieldTypeBool = 2,
	FieldTypeTime = 3,
	FieldTypeDuration = 4,
	FieldTypePointer = 5,
	FieldTypeSliceOfPointer = 6,
}

export interface FieldTypeSelect {
	value: number;
	viewValue: string;
}

export const FieldTypeList: FieldTypeSelect[] = [ // insertion point	
	{ value: FieldType.FieldTypeInt, viewValue: "FieldTypeInt" },
	{ value: FieldType.FieldTypeFloat64, viewValue: "FieldTypeFloat64" },
	{ value: FieldType.FieldTypeBool, viewValue: "FieldTypeBool" },
	{ value: FieldType.FieldTypeTime, viewValue: "FieldTypeTime" },
	{ value: FieldType.FieldTypeDuration, viewValue: "FieldTypeDuration" },
	{ value: FieldType.FieldTypePointer, viewValue: "FieldTypePointer" },
	{ value: FieldType.FieldTypeSliceOfPointer, viewValue: "FieldTypeSliceOfPointer" },
];
