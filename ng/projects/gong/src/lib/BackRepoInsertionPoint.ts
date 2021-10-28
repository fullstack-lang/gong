// generated from ng_file_enum.ts.go
export enum BackRepoInsertionPoint {
	// insertion point	
	BackRepoBasicAndTimeFieldsName = 1,
	BackRepoBasicAndTimeFieldsWOPDeclaration = 3,
	BackRepoBasicFieldsCheckout = 8,
	BackRepoBasicFieldsCommit = 6,
	BackRepoBasicFieldsDeclaration = 0,
	BackRepoNbInsertionPoints = 11,
	BackRepoPointerEncodingFieldsCheckout = 9,
	BackRepoPointerEncodingFieldsCommit = 7,
	BackRepoPointerEncodingFieldsDeclaration = 4,
	BackRepoPointerEncodingFieldsReindexing = 10,
	BackRepoPointerEncodingFieldsWOPDeclaration = 5,
	BackRepoWOPInitialIndex = 2,
}

export interface BackRepoInsertionPointSelect {
	value: string;
	viewValue: string;
}

export const BackRepoInsertionPointList: BackRepoInsertionPointSelect[] = [ // insertion point	
	{ value: 'BackRepoBasicAndTimeFieldsName', viewValue: '1' },
	{ value: 'BackRepoBasicAndTimeFieldsWOPDeclaration', viewValue: '3' },
	{ value: 'BackRepoBasicFieldsCheckout', viewValue: '8' },
	{ value: 'BackRepoBasicFieldsCommit', viewValue: '6' },
	{ value: 'BackRepoBasicFieldsDeclaration', viewValue: '0' },
	{ value: 'BackRepoNbInsertionPoints', viewValue: '11' },
	{ value: 'BackRepoPointerEncodingFieldsCheckout', viewValue: '9' },
	{ value: 'BackRepoPointerEncodingFieldsCommit', viewValue: '7' },
	{ value: 'BackRepoPointerEncodingFieldsDeclaration', viewValue: '4' },
	{ value: 'BackRepoPointerEncodingFieldsReindexing', viewValue: '10' },
	{ value: 'BackRepoPointerEncodingFieldsWOPDeclaration', viewValue: '5' },
	{ value: 'BackRepoWOPInitialIndex', viewValue: '2' },
];
