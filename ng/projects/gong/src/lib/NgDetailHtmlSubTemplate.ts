// generated from ng_file_enum.ts.go
export enum NgDetailHtmlSubTemplate {
	// insertion point	
	NgDetailHtmlBasicField = 1,
	NgDetailHtmlBasicStringField = 2,
	NgDetailHtmlBool = 4,
	NgDetailHtmlEnum = 0,
	NgDetailHtmlTimeDuration = 5,
	NgDetailHtmlTimeField = 3,
	NgDetailPointerToStructHtmlFormField = 6,
	NgDetailSliceOfPointerToStructHtml = 7,
	NgDetailSliceOfPointerToStructManyManyHtml = 8,
	NgDetailSliceOfPointerToStructReverseHtml = 9,
}

export interface NgDetailHtmlSubTemplateSelect {
	value: number;
	viewValue: string;
}

export const NgDetailHtmlSubTemplateList: NgDetailHtmlSubTemplateSelect[] = [ // insertion point	
	{ value: NgDetailHtmlSubTemplate.NgDetailHtmlBasicField, viewValue: "NgDetailHtmlBasicField" },
	{ value: NgDetailHtmlSubTemplate.NgDetailHtmlBasicStringField, viewValue: "NgDetailHtmlBasicStringField" },
	{ value: NgDetailHtmlSubTemplate.NgDetailHtmlBool, viewValue: "NgDetailHtmlBool" },
	{ value: NgDetailHtmlSubTemplate.NgDetailHtmlEnum, viewValue: "NgDetailHtmlEnum" },
	{ value: NgDetailHtmlSubTemplate.NgDetailHtmlTimeDuration, viewValue: "NgDetailHtmlTimeDuration" },
	{ value: NgDetailHtmlSubTemplate.NgDetailHtmlTimeField, viewValue: "NgDetailHtmlTimeField" },
	{ value: NgDetailHtmlSubTemplate.NgDetailPointerToStructHtmlFormField, viewValue: "NgDetailPointerToStructHtmlFormField" },
	{ value: NgDetailHtmlSubTemplate.NgDetailSliceOfPointerToStructHtml, viewValue: "NgDetailSliceOfPointerToStructHtml" },
	{ value: NgDetailHtmlSubTemplate.NgDetailSliceOfPointerToStructManyManyHtml, viewValue: "NgDetailSliceOfPointerToStructManyManyHtml" },
	{ value: NgDetailHtmlSubTemplate.NgDetailSliceOfPointerToStructReverseHtml, viewValue: "NgDetailSliceOfPointerToStructReverseHtml" },
];
