// generated from ng_file_enum.ts.go
export enum NgTableTsInsertionPoint {
	// insertion point	
	NgTableTsInsertionPerStructImports = 0,
	NgTableTsInsertionPerStructTimeDurationRecoveries = 1,
	NgTableTsInsertionPerStructEnumIntRecoveries = 2,
	NgTableTsInsertionPerStructColumns = 3,
	NgTableTsInsertionPerStructColumnsSorting = 4,
	NgTableTsInsertionPerStructColumnsFiltering = 5,
	NgTableTsInsertionsNb = 6,
}

export interface NgTableTsInsertionPointSelect {
	value: number;
	viewValue: string;
}

export const NgTableTsInsertionPointList: NgTableTsInsertionPointSelect[] = [ // insertion point	
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionPerStructImports, viewValue: "NgTableTsInsertionPerStructImports" },
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionPerStructTimeDurationRecoveries, viewValue: "NgTableTsInsertionPerStructTimeDurationRecoveries" },
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionPerStructEnumIntRecoveries, viewValue: "NgTableTsInsertionPerStructEnumIntRecoveries" },
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionPerStructColumns, viewValue: "NgTableTsInsertionPerStructColumns" },
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionPerStructColumnsSorting, viewValue: "NgTableTsInsertionPerStructColumnsSorting" },
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionPerStructColumnsFiltering, viewValue: "NgTableTsInsertionPerStructColumnsFiltering" },
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionsNb, viewValue: "NgTableTsInsertionsNb" },
];
