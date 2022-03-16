// generated from ng_file_enum.ts.go
export enum NgTableSubTemplate {
	// insertion point	
	NgTableTSBasicFieldSorting = 3,
	NgTableTSEnumIntFiltering = 9,
	NgTableTSNonNumberFieldFiltering = 7,
	NgTableTSNumberFieldFiltering = 8,
	NgTableTSPerStructColumn = 13,
	NgTableTSPerStructEnumIntRecoveries = 1,
	NgTableTSPerStructTimeDurationRecoveries = 2,
	NgTableTSPointerToStructFiltering = 11,
	NgTableTSPointerToStructSorting = 5,
	NgTableTSSliceOfPointerToStructFiltering = 12,
	NgTableTSSliceOfPointerToStructPerStructColumn = 14,
	NgTableTSSliceOfPointerToStructSorting = 6,
	NgTableTSTimeFieldFiltering = 10,
	NgTableTSTimeFieldSorting = 4,
	NgTableTsInsertionPerStructImportsTpl = 0,
}

export interface NgTableSubTemplateSelect {
	value: number;
	viewValue: string;
}

export const NgTableSubTemplateList: NgTableSubTemplateSelect[] = [ // insertion point	
	{ value: NgTableSubTemplate.NgTableTSBasicFieldSorting, viewValue: "NgTableTSBasicFieldSorting" },
	{ value: NgTableSubTemplate.NgTableTSEnumIntFiltering, viewValue: "NgTableTSEnumIntFiltering" },
	{ value: NgTableSubTemplate.NgTableTSNonNumberFieldFiltering, viewValue: "NgTableTSNonNumberFieldFiltering" },
	{ value: NgTableSubTemplate.NgTableTSNumberFieldFiltering, viewValue: "NgTableTSNumberFieldFiltering" },
	{ value: NgTableSubTemplate.NgTableTSPerStructColumn, viewValue: "NgTableTSPerStructColumn" },
	{ value: NgTableSubTemplate.NgTableTSPerStructEnumIntRecoveries, viewValue: "NgTableTSPerStructEnumIntRecoveries" },
	{ value: NgTableSubTemplate.NgTableTSPerStructTimeDurationRecoveries, viewValue: "NgTableTSPerStructTimeDurationRecoveries" },
	{ value: NgTableSubTemplate.NgTableTSPointerToStructFiltering, viewValue: "NgTableTSPointerToStructFiltering" },
	{ value: NgTableSubTemplate.NgTableTSPointerToStructSorting, viewValue: "NgTableTSPointerToStructSorting" },
	{ value: NgTableSubTemplate.NgTableTSSliceOfPointerToStructFiltering, viewValue: "NgTableTSSliceOfPointerToStructFiltering" },
	{ value: NgTableSubTemplate.NgTableTSSliceOfPointerToStructPerStructColumn, viewValue: "NgTableTSSliceOfPointerToStructPerStructColumn" },
	{ value: NgTableSubTemplate.NgTableTSSliceOfPointerToStructSorting, viewValue: "NgTableTSSliceOfPointerToStructSorting" },
	{ value: NgTableSubTemplate.NgTableTSTimeFieldFiltering, viewValue: "NgTableTSTimeFieldFiltering" },
	{ value: NgTableSubTemplate.NgTableTSTimeFieldSorting, viewValue: "NgTableTSTimeFieldSorting" },
	{ value: NgTableSubTemplate.NgTableTsInsertionPerStructImportsTpl, viewValue: "NgTableTsInsertionPerStructImportsTpl" },
];
