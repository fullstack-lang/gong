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
	value: string;
	viewValue: string;
}

export const NgTableSubTemplateList: NgTableSubTemplateSelect[] = [ // insertion point	
	{ value: 'NgTableTSBasicFieldSorting', viewValue: '1' },
	{ value: 'NgTableTSNonNumberFieldFiltering', viewValue: '5' },
	{ value: 'NgTableTSNumberFieldFiltering', viewValue: '6' },
	{ value: 'NgTableTSPerStructColumn', viewValue: '10' },
	{ value: 'NgTableTSPerStructTimeDurationRecoveries', viewValue: '0' },
	{ value: 'NgTableTSPointerToStructFiltering', viewValue: '8' },
	{ value: 'NgTableTSPointerToStructSorting', viewValue: '3' },
	{ value: 'NgTableTSSliceOfPointerToStructFiltering', viewValue: '9' },
	{ value: 'NgTableTSSliceOfPointerToStructPerStructColumn', viewValue: '11' },
	{ value: 'NgTableTSSliceOfPointerToStructSorting', viewValue: '4' },
	{ value: 'NgTableTSTimeFieldFiltering', viewValue: '7' },
	{ value: 'NgTableTSTimeFieldSorting', viewValue: '2' },
];
