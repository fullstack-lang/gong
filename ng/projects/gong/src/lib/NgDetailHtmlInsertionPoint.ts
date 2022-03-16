// generated from ng_file_enum.ts.go
export enum NgDetailHtmlInsertionPoint {
	// insertion point	
	NgDetailHtmlInsertionPerStructFields = 0,
	NgDetailHtmlInsertionPerStructFieldsManyMany = 1,
	NgDetailHtmlInsertionsNb = 2,
}

export interface NgDetailHtmlInsertionPointSelect {
	value: number;
	viewValue: string;
}

export const NgDetailHtmlInsertionPointList: NgDetailHtmlInsertionPointSelect[] = [ // insertion point	
	{ value: NgDetailHtmlInsertionPoint.NgDetailHtmlInsertionPerStructFields, viewValue: "NgDetailHtmlInsertionPerStructFields" },
	{ value: NgDetailHtmlInsertionPoint.NgDetailHtmlInsertionPerStructFieldsManyMany, viewValue: "NgDetailHtmlInsertionPerStructFieldsManyMany" },
	{ value: NgDetailHtmlInsertionPoint.NgDetailHtmlInsertionsNb, viewValue: "NgDetailHtmlInsertionsNb" },
];
