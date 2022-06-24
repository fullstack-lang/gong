// generated from ng_file_enum.ts.go
export enum NgDetailTsInsertionPoint {
	// insertion point	
	NgDetailTsInsertionPerStructImports = 0,
	NgDetailTsInsertionPerStructEnumFieldDeclarations = 1,
	NgDetailTsInsertionPerStructDeclarations = 2,
	NgDetailTsInsertionPerStructCaseInitFieldDeclarations = 3,
	NgDetailTsInsertionPerStructCaseSetField = 4,
	NgDetailTsInsertionPerStructInits = 5,
	NgDetailTsInsertionPerStructSorting = 6,
	NgDetailTsInsertionPerStructRecoveries = 7,
	NgDetailTsInsertionPerStructSaves = 8,
	NgDetailTsInsertionPerStructReversePointerSaveWhenUpdate = 9,
	NgDetailTsInsertionPerStructReversePointerSaveWhenCreateFromOwner = 10,
	NgDetailTsInsertionsNb = 11,
}

export interface NgDetailTsInsertionPointSelect {
	value: number;
	viewValue: string;
}

export const NgDetailTsInsertionPointList: NgDetailTsInsertionPointSelect[] = [ // insertion point	
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructImports, viewValue: "NgDetailTsInsertionPerStructImports" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructEnumFieldDeclarations, viewValue: "NgDetailTsInsertionPerStructEnumFieldDeclarations" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructDeclarations, viewValue: "NgDetailTsInsertionPerStructDeclarations" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructCaseInitFieldDeclarations, viewValue: "NgDetailTsInsertionPerStructCaseInitFieldDeclarations" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructCaseSetField, viewValue: "NgDetailTsInsertionPerStructCaseSetField" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructInits, viewValue: "NgDetailTsInsertionPerStructInits" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructSorting, viewValue: "NgDetailTsInsertionPerStructSorting" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructRecoveries, viewValue: "NgDetailTsInsertionPerStructRecoveries" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructSaves, viewValue: "NgDetailTsInsertionPerStructSaves" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructReversePointerSaveWhenUpdate, viewValue: "NgDetailTsInsertionPerStructReversePointerSaveWhenUpdate" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructReversePointerSaveWhenCreateFromOwner, viewValue: "NgDetailTsInsertionPerStructReversePointerSaveWhenCreateFromOwner" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionsNb, viewValue: "NgDetailTsInsertionsNb" },
];
