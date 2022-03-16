// generated from ng_file_enum.ts.go
export enum NgServiceTsInsertionPoint {
	// insertion point	
	NgServiceTsInsertionImports = 2,
	NgServiceTsInsertionPointerReset = 0,
	NgServiceTsInsertionPointerRestore = 1,
	NgServiceTsInsertionsNb = 3,
}

export interface NgServiceTsInsertionPointSelect {
	value: number;
	viewValue: string;
}

export const NgServiceTsInsertionPointList: NgServiceTsInsertionPointSelect[] = [ // insertion point	
	{ value: NgServiceTsInsertionPoint.NgServiceTsInsertionImports, viewValue: "NgServiceTsInsertionImports" },
	{ value: NgServiceTsInsertionPoint.NgServiceTsInsertionPointerReset, viewValue: "NgServiceTsInsertionPointerReset" },
	{ value: NgServiceTsInsertionPoint.NgServiceTsInsertionPointerRestore, viewValue: "NgServiceTsInsertionPointerRestore" },
	{ value: NgServiceTsInsertionPoint.NgServiceTsInsertionsNb, viewValue: "NgServiceTsInsertionsNb" },
];
