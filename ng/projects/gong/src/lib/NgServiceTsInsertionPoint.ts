// generated from ng_file_enum.ts.go
export enum NgServiceTsInsertionPoint {
	// insertion point	
	NgServiceTsInsertionPointerReset = 0,
	NgServiceTsInsertionPointerRestore = 1,
	NgServiceTsInsertionImports = 2,
	NgServiceTsInsertionsNb = 3,
}

export interface NgServiceTsInsertionPointSelect {
	value: number;
	viewValue: string;
}

export const NgServiceTsInsertionPointList: NgServiceTsInsertionPointSelect[] = [ // insertion point	
	{ value: NgServiceTsInsertionPoint.NgServiceTsInsertionPointerReset, viewValue: "NgServiceTsInsertionPointerReset" },
	{ value: NgServiceTsInsertionPoint.NgServiceTsInsertionPointerRestore, viewValue: "NgServiceTsInsertionPointerRestore" },
	{ value: NgServiceTsInsertionPoint.NgServiceTsInsertionImports, viewValue: "NgServiceTsInsertionImports" },
	{ value: NgServiceTsInsertionPoint.NgServiceTsInsertionsNb, viewValue: "NgServiceTsInsertionsNb" },
];
