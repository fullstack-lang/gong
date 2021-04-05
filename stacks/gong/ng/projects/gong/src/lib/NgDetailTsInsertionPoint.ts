// generated from ng_file_enum.ts.go
export enum NgDetailTsInsertionPoint {
	// insertion point	
	NgDetailTsInsertionPerStructDeclarations = 1,
	NgDetailTsInsertionPerStructImports = 0,
	NgDetailTsInsertionPerStructInits = 2,
	NgDetailTsInsertionPerStructRecoveries = 3,
	NgDetailTsInsertionPerStructReversePointerSaveWhenCreateFromOwner = 6,
	NgDetailTsInsertionPerStructReversePointerSaveWhenUpdate = 5,
	NgDetailTsInsertionPerStructSaves = 4,
	NgDetailTsInsertionsNb = 7,
}

export interface NgDetailTsInsertionPointSelect {
	value: string;
	viewValue: string;
}

export const NgDetailTsInsertionPointList: NgDetailTsInsertionPointSelect[] = [ // insertion point	
	{ value: 'NgDetailTsInsertionPerStructDeclarations', viewValue: '1' },
	{ value: 'NgDetailTsInsertionPerStructImports', viewValue: '0' },
	{ value: 'NgDetailTsInsertionPerStructInits', viewValue: '2' },
	{ value: 'NgDetailTsInsertionPerStructRecoveries', viewValue: '3' },
	{ value: 'NgDetailTsInsertionPerStructReversePointerSaveWhenCreateFromOwner', viewValue: '6' },
	{ value: 'NgDetailTsInsertionPerStructReversePointerSaveWhenUpdate', viewValue: '5' },
	{ value: 'NgDetailTsInsertionPerStructSaves', viewValue: '4' },
	{ value: 'NgDetailTsInsertionsNb', viewValue: '7' },
];
