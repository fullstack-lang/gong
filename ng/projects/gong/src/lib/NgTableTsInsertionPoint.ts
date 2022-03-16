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
	value: number;
	viewValue: string;
}

export const NgTableTsInsertionPointList: NgTableTsInsertionPointSelect[] = [ // insertion point	
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionPerStructColumns, viewValue: "NgTableTsInsertionPerStructColumns" },
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionPerStructColumnsFiltering, viewValue: "NgTableTsInsertionPerStructColumnsFiltering" },
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionPerStructColumnsSorting, viewValue: "NgTableTsInsertionPerStructColumnsSorting" },
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionPerStructRecoveries, viewValue: "NgTableTsInsertionPerStructRecoveries" },
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionsNb, viewValue: "NgTableTsInsertionsNb" },
];
