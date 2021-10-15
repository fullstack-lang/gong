// generated from ng_file_enum.ts.go
export enum NgSidebarTsInsertionPoint {
	// insertion point	
	NgSidebarTsInsertionPerStruct = 0,
	NgSidebarTsInsertionPerStructImports = 1,
	NgSidebarTsInsertionPerStructObservableForRefresh = 3,
	NgSidebarTsInsertionPerStructServiceDeclaration = 2,
	NgSidebarTsInsertionsNb = 4,
}

export interface NgSidebarTsInsertionPointSelect {
	value: string;
	viewValue: string;
}

export const NgSidebarTsInsertionPointList: NgSidebarTsInsertionPointSelect[] = [ // insertion point	
	{ value: 'NgSidebarTsInsertionPerStruct', viewValue: '0' },
	{ value: 'NgSidebarTsInsertionPerStructImports', viewValue: '1' },
	{ value: 'NgSidebarTsInsertionPerStructObservableForRefresh', viewValue: '3' },
	{ value: 'NgSidebarTsInsertionPerStructServiceDeclaration', viewValue: '2' },
	{ value: 'NgSidebarTsInsertionsNb', viewValue: '4' },
];
