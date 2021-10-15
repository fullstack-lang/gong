// generated from ng_file_enum.ts.go
export enum BackRepoInsertionPoint {
	// insertion point	
	BackRepoBasicAndTimeFieldsName = 1,
	BackRepoBasicAndTimeFieldsWOPDeclaration = 2,
	BackRepoBasicFieldsCheckout = 7,
	BackRepoBasicFieldsCommit = 5,
	BackRepoBasicFieldsDeclaration = 0,
	BackRepoNbInsertionPoints = 10,
	BackRepoPointerEncodingFieldsCheckout = 8,
	BackRepoPointerEncodingFieldsCommit = 6,
	BackRepoPointerEncodingFieldsDeclaration = 3,
	BackRepoPointerEncodingFieldsReindexing = 9,
	BackRepoPointerEncodingFieldsWOPDeclaration = 4,
}

export interface BackRepoInsertionPointSelect {
	value: string;
	viewValue: string;
}

export const BackRepoInsertionPointList: BackRepoInsertionPointSelect[] = [ // insertion point	
	{ value: 'BackRepoBasicAndTimeFieldsName', viewValue: '1' },
	{ value: 'BackRepoBasicAndTimeFieldsWOPDeclaration', viewValue: '2' },
	{ value: 'BackRepoBasicFieldsCheckout', viewValue: '7' },
	{ value: 'BackRepoBasicFieldsCommit', viewValue: '5' },
	{ value: 'BackRepoBasicFieldsDeclaration', viewValue: '0' },
	{ value: 'BackRepoNbInsertionPoints', viewValue: '10' },
	{ value: 'BackRepoPointerEncodingFieldsCheckout', viewValue: '8' },
	{ value: 'BackRepoPointerEncodingFieldsCommit', viewValue: '6' },
	{ value: 'BackRepoPointerEncodingFieldsDeclaration', viewValue: '3' },
	{ value: 'BackRepoPointerEncodingFieldsReindexing', viewValue: '9' },
	{ value: 'BackRepoPointerEncodingFieldsWOPDeclaration', viewValue: '4' },
];
