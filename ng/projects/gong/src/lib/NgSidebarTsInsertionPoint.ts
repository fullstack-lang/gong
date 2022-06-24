// generated from ng_file_enum.ts.go
export enum NgSidebarTsInsertionPoint {
	// insertion point	
	NgSidebarTsInsertionPerStruct = 0,
	NgSidebarTsInsertionPerStructImports = 1,
	NgSidebarTsInsertionPerStructServiceDeclaration = 2,
	NgSidebarTsInsertionPerStructObservableForRefresh = 3,
	NgSidebarTsInsertionsNb = 4,
}

export interface NgSidebarTsInsertionPointSelect {
	value: number;
	viewValue: string;
}

export const NgSidebarTsInsertionPointList: NgSidebarTsInsertionPointSelect[] = [ // insertion point	
	{ value: NgSidebarTsInsertionPoint.NgSidebarTsInsertionPerStruct, viewValue: "NgSidebarTsInsertionPerStruct" },
	{ value: NgSidebarTsInsertionPoint.NgSidebarTsInsertionPerStructImports, viewValue: "NgSidebarTsInsertionPerStructImports" },
	{ value: NgSidebarTsInsertionPoint.NgSidebarTsInsertionPerStructServiceDeclaration, viewValue: "NgSidebarTsInsertionPerStructServiceDeclaration" },
	{ value: NgSidebarTsInsertionPoint.NgSidebarTsInsertionPerStructObservableForRefresh, viewValue: "NgSidebarTsInsertionPerStructObservableForRefresh" },
	{ value: NgSidebarTsInsertionPoint.NgSidebarTsInsertionsNb, viewValue: "NgSidebarTsInsertionsNb" },
];
