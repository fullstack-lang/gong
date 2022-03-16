// generated from ng_file_enum.ts.go
export enum NgClassTsInsertionPoint {
	// insertion point	
	NgClassTsInsertionPerStructBasicFieldsDecl = 1,
	NgClassTsInsertionPerStructImports = 0,
	NgClassTsInsertionPerStructOtherDecls = 2,
	NgClassTsInsertionsNb = 3,
}

export interface NgClassTsInsertionPointSelect {
	value: number;
	viewValue: string;
}

export const NgClassTsInsertionPointList: NgClassTsInsertionPointSelect[] = [ // insertion point	
	{ value: NgClassTsInsertionPoint.NgClassTsInsertionPerStructBasicFieldsDecl, viewValue: "NgClassTsInsertionPerStructBasicFieldsDecl" },
	{ value: NgClassTsInsertionPoint.NgClassTsInsertionPerStructImports, viewValue: "NgClassTsInsertionPerStructImports" },
	{ value: NgClassTsInsertionPoint.NgClassTsInsertionPerStructOtherDecls, viewValue: "NgClassTsInsertionPerStructOtherDecls" },
	{ value: NgClassTsInsertionPoint.NgClassTsInsertionsNb, viewValue: "NgClassTsInsertionsNb" },
];
