// generated from ng_file_enum.ts.go
export enum NgTableSubTemplate {
	// insertion point	
	NgTableTSBasicFieldSorting = 1,
	NgTableTSNonNumberFieldFiltering = 5,
	NgTableTSNumberFieldFiltering = 6,
	NgTableTSPerStructColumn = 10,
	NgTableTSPerStructTimeDurationRecoveries = 0,
	NgTableTSPointerToStructFiltering = 8,
	NgTableTSPointerToStructSorting = 3,
	NgTableTSSliceOfPointerToStructFiltering = 9,
	NgTableTSSliceOfPointerToStructPerStructColumn = 11,
	NgTableTSSliceOfPointerToStructSorting = 4,
	NgTableTSTimeFieldFiltering = 7,
	NgTableTSTimeFieldSorting = 2,
}

export interface NgTableSubTemplateSelect {
	value: number;
	viewValue: string;
}

export const NgTableSubTemplateList: NgTableSubTemplateSelect[] = [ // insertion point	
	{ value: NgTableSubTemplate.NgTableTSBasicFieldSorting, viewValue: "NgTableTSBasicFieldSorting" },
	{ value: NgTableSubTemplate.NgTableTSNonNumberFieldFiltering, viewValue: "NgTableTSNonNumberFieldFiltering" },
	{ value: NgTableSubTemplate.NgTableTSNumberFieldFiltering, viewValue: "NgTableTSNumberFieldFiltering" },
	{ value: NgTableSubTemplate.NgTableTSPerStructColumn, viewValue: "NgTableTSPerStructColumn" },
	{ value: NgTableSubTemplate.NgTableTSPerStructTimeDurationRecoveries, viewValue: "NgTableTSPerStructTimeDurationRecoveries" },
	{ value: NgTableSubTemplate.NgTableTSPointerToStructFiltering, viewValue: "NgTableTSPointerToStructFiltering" },
	{ value: NgTableSubTemplate.NgTableTSPointerToStructSorting, viewValue: "NgTableTSPointerToStructSorting" },
	{ value: NgTableSubTemplate.NgTableTSSliceOfPointerToStructFiltering, viewValue: "NgTableTSSliceOfPointerToStructFiltering" },
	{ value: NgTableSubTemplate.NgTableTSSliceOfPointerToStructPerStructColumn, viewValue: "NgTableTSSliceOfPointerToStructPerStructColumn" },
	{ value: NgTableSubTemplate.NgTableTSSliceOfPointerToStructSorting, viewValue: "NgTableTSSliceOfPointerToStructSorting" },
	{ value: NgTableSubTemplate.NgTableTSTimeFieldFiltering, viewValue: "NgTableTSTimeFieldFiltering" },
	{ value: NgTableSubTemplate.NgTableTSTimeFieldSorting, viewValue: "NgTableTSTimeFieldSorting" },
];
