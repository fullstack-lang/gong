// generated from ng_file_enum.ts.go
export enum NgTableTsInsertionPoint {
	// insertion point	
	NgTableTsInsertionPerStructColumns = 1,
	NgTableTsInsertionPerStructColumnsFiltering = 3,
	NgTableTsInsertionPerStructColumnsSorting = 2,
	NgTableTsInsertionPerStructRecoveries = 0,
	NgTableTsInsertionsNb = 4,
}

export interface NgTableTsInsertionPointSelect {
	value: string;
	viewValue: string;
}

export const NgTableTsInsertionPointList: NgTableTsInsertionPointSelect[] = [ // insertion point	
	{ value: 'NgTableTsInsertionPerStructColumns', viewValue: '1' },
	{ value: 'NgTableTsInsertionPerStructColumnsFiltering', viewValue: '3' },
	{ value: 'NgTableTsInsertionPerStructColumnsSorting', viewValue: '2' },
	{ value: 'NgTableTsInsertionPerStructRecoveries', viewValue: '0' },
	{ value: 'NgTableTsInsertionsNb', viewValue: '4' },
];
