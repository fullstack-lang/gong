// generated from ng_file_enum.ts.go
export enum NgDetailTsInsertionPoint {
	// insertion point	
	NgDetailTsInsertionPerStructCaseInitFieldDeclarations = 3,
	NgDetailTsInsertionPerStructCaseSetField = 4,
	NgDetailTsInsertionPerStructDeclarations = 2,
	NgDetailTsInsertionPerStructEnumFieldDeclarations = 1,
	NgDetailTsInsertionPerStructImports = 0,
	NgDetailTsInsertionPerStructInits = 5,
	NgDetailTsInsertionPerStructRecoveries = 7,
	NgDetailTsInsertionPerStructReversePointerSaveWhenCreateFromOwner = 10,
	NgDetailTsInsertionPerStructReversePointerSaveWhenUpdate = 9,
	NgDetailTsInsertionPerStructSaves = 8,
	NgDetailTsInsertionPerStructSorting = 6,
	NgDetailTsInsertionsNb = 11,
}

export interface NgDetailTsInsertionPointSelect {
	value: number;
	viewValue: string;
}

export const NgDetailTsInsertionPointList: NgDetailTsInsertionPointSelect[] = [ // insertion point	
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructCaseInitFieldDeclarations, viewValue: "NgDetailTsInsertionPerStructCaseInitFieldDeclarations" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructCaseSetField, viewValue: "NgDetailTsInsertionPerStructCaseSetField" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructDeclarations, viewValue: "NgDetailTsInsertionPerStructDeclarations" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructEnumFieldDeclarations, viewValue: "NgDetailTsInsertionPerStructEnumFieldDeclarations" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructImports, viewValue: "NgDetailTsInsertionPerStructImports" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructInits, viewValue: "NgDetailTsInsertionPerStructInits" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructRecoveries, viewValue: "NgDetailTsInsertionPerStructRecoveries" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructReversePointerSaveWhenCreateFromOwner, viewValue: "NgDetailTsInsertionPerStructReversePointerSaveWhenCreateFromOwner" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructReversePointerSaveWhenUpdate, viewValue: "NgDetailTsInsertionPerStructReversePointerSaveWhenUpdate" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructSaves, viewValue: "NgDetailTsInsertionPerStructSaves" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionPerStructSorting, viewValue: "NgDetailTsInsertionPerStructSorting" },
	{ value: NgDetailTsInsertionPoint.NgDetailTsInsertionsNb, viewValue: "NgDetailTsInsertionsNb" },
];
