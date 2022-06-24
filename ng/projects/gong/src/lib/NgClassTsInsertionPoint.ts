// generated from ng_file_enum.ts.go
export enum NgClassTsInsertionPoint {
	// insertion point	
	NgClassTsInsertionPerStructImports = 0,
	NgClassTsInsertionPerStructBasicFieldsDecl = 1,
	NgClassTsInsertionPerStructOtherDecls = 2,
	NgClassTsInsertionsNb = 3,
}

export interface NgClassTsInsertionPointSelect {
	value: number;
	viewValue: string;
}

export const NgClassTsInsertionPointList: NgClassTsInsertionPointSelect[] = [ // insertion point	
	{ value: NgClassTsInsertionPoint.NgClassTsInsertionPerStructImports, viewValue: "NgClassTsInsertionPerStructImports" },
	{ value: NgClassTsInsertionPoint.NgClassTsInsertionPerStructBasicFieldsDecl, viewValue: "NgClassTsInsertionPerStructBasicFieldsDecl" },
	{ value: NgClassTsInsertionPoint.NgClassTsInsertionPerStructOtherDecls, viewValue: "NgClassTsInsertionPerStructOtherDecls" },
	{ value: NgClassTsInsertionPoint.NgClassTsInsertionsNb, viewValue: "NgClassTsInsertionsNb" },
];
