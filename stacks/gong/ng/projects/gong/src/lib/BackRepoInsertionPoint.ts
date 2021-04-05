// generated from ng_file_enum.ts.go
export enum BackRepoInsertionPoint {
	// insertion point	
	BackRepoFieldsCheckoutNew = 2,
	BackRepoFieldsCommitNew = 1,
	BackRepoFieldsDeclaration = 0,
	BackRepoNbInsertionPoints = 3,
}

export interface BackRepoInsertionPointSelect {
	value: string;
	viewValue: string;
}

export const BackRepoInsertionPointList: BackRepoInsertionPointSelect[] = [ // insertion point	
	{ value: 'BackRepoFieldsCheckoutNew', viewValue: '2' },
	{ value: 'BackRepoFieldsCommitNew', viewValue: '1' },
	{ value: 'BackRepoFieldsDeclaration', viewValue: '0' },
	{ value: 'BackRepoNbInsertionPoints', viewValue: '3' },
];
