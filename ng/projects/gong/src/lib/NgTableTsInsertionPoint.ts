// generated from ng_file_enum.ts.go
export enum NgTableTsInsertionPoint {
	// insertion point	
	NgTableTsInsertionPerStructColumns = 3,
	NgTableTsInsertionPerStructColumnsFiltering = 5,
	NgTableTsInsertionPerStructColumnsSorting = 4,
	NgTableTsInsertionPerStructEnumIntRecoveries = 2,
	NgTableTsInsertionPerStructImports = 0,
	NgTableTsInsertionPerStructTimeDurationRecoveries = 1,
	NgTableTsInsertionsNb = 6,
}

export interface NgTableTsInsertionPointSelect {
	value: number;
	viewValue: string;
}

export const NgTableTsInsertionPointList: NgTableTsInsertionPointSelect[] = [ // insertion point	
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionPerStructColumns, viewValue: "NgTableTsInsertionPerStructColumns" },
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionPerStructColumnsFiltering, viewValue: "NgTableTsInsertionPerStructColumnsFiltering" },
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionPerStructColumnsSorting, viewValue: "NgTableTsInsertionPerStructColumnsSorting" },
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionPerStructEnumIntRecoveries, viewValue: "NgTableTsInsertionPerStructEnumIntRecoveries" },
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionPerStructImports, viewValue: "NgTableTsInsertionPerStructImports" },
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionPerStructTimeDurationRecoveries, viewValue: "NgTableTsInsertionPerStructTimeDurationRecoveries" },
	{ value: NgTableTsInsertionPoint.NgTableTsInsertionsNb, viewValue: "NgTableTsInsertionsNb" },
];
