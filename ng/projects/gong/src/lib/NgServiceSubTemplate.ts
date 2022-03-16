// generated from ng_file_enum.ts.go
export enum NgServiceSubTemplate {
	// insertion point	
	NgServiceTSPointerToGongStructImports = 0,
	NgServiceTSPointerToGongStructReset = 1,
	NgServiceTSReversePointerToSliceOfGongStructImports = 5,
	NgServiceTSSliceOfPointerToGongStructReset = 2,
	NgServiceTSSliceOfPointerToGongStructReversePointerReset = 3,
	NgServiceTSSliceOfPointerToGongStructReversePointerRestore = 4,
}

export interface NgServiceSubTemplateSelect {
	value: number;
	viewValue: string;
}

export const NgServiceSubTemplateList: NgServiceSubTemplateSelect[] = [ // insertion point	
	{ value: NgServiceSubTemplate.NgServiceTSPointerToGongStructImports, viewValue: "NgServiceTSPointerToGongStructImports" },
	{ value: NgServiceSubTemplate.NgServiceTSPointerToGongStructReset, viewValue: "NgServiceTSPointerToGongStructReset" },
	{ value: NgServiceSubTemplate.NgServiceTSReversePointerToSliceOfGongStructImports, viewValue: "NgServiceTSReversePointerToSliceOfGongStructImports" },
	{ value: NgServiceSubTemplate.NgServiceTSSliceOfPointerToGongStructReset, viewValue: "NgServiceTSSliceOfPointerToGongStructReset" },
	{ value: NgServiceSubTemplate.NgServiceTSSliceOfPointerToGongStructReversePointerReset, viewValue: "NgServiceTSSliceOfPointerToGongStructReversePointerReset" },
	{ value: NgServiceSubTemplate.NgServiceTSSliceOfPointerToGongStructReversePointerRestore, viewValue: "NgServiceTSSliceOfPointerToGongStructReversePointerRestore" },
];
