// generated from ng_file_enum.ts.go
export enum BackRepoInsertionPoint {
	// insertion point	
	BackRepoBasicFieldsDeclaration = 0,
	BackRepoBasicAndTimeFieldsName = 1,
	BackRepoWOPInitialIndex = 2,
	BackRepoBasicAndTimeFieldsWOPDeclaration = 3,
	BackRepoPointerEncodingFieldsDeclaration = 4,
	BackRepoPointerEncodingFieldsWOPDeclaration = 5,
	BackRepoBasicFieldsCommit = 6,
	BackRepoPointerEncodingFieldsCommit = 7,
	BackRepoBasicFieldsCheckout = 8,
	BackRepoPointerEncodingFieldsCheckout = 9,
	BackRepoPointerEncodingFieldsReindexing = 10,
	BackRepoNbInsertionPoints = 11,
}

export interface BackRepoInsertionPointSelect {
	value: number;
	viewValue: string;
}

export const BackRepoInsertionPointList: BackRepoInsertionPointSelect[] = [ // insertion point	
	{ value: BackRepoInsertionPoint.BackRepoBasicFieldsDeclaration, viewValue: "BackRepoBasicFieldsDeclaration" },
	{ value: BackRepoInsertionPoint.BackRepoBasicAndTimeFieldsName, viewValue: "BackRepoBasicAndTimeFieldsName" },
	{ value: BackRepoInsertionPoint.BackRepoWOPInitialIndex, viewValue: "BackRepoWOPInitialIndex" },
	{ value: BackRepoInsertionPoint.BackRepoBasicAndTimeFieldsWOPDeclaration, viewValue: "BackRepoBasicAndTimeFieldsWOPDeclaration" },
	{ value: BackRepoInsertionPoint.BackRepoPointerEncodingFieldsDeclaration, viewValue: "BackRepoPointerEncodingFieldsDeclaration" },
	{ value: BackRepoInsertionPoint.BackRepoPointerEncodingFieldsWOPDeclaration, viewValue: "BackRepoPointerEncodingFieldsWOPDeclaration" },
	{ value: BackRepoInsertionPoint.BackRepoBasicFieldsCommit, viewValue: "BackRepoBasicFieldsCommit" },
	{ value: BackRepoInsertionPoint.BackRepoPointerEncodingFieldsCommit, viewValue: "BackRepoPointerEncodingFieldsCommit" },
	{ value: BackRepoInsertionPoint.BackRepoBasicFieldsCheckout, viewValue: "BackRepoBasicFieldsCheckout" },
	{ value: BackRepoInsertionPoint.BackRepoPointerEncodingFieldsCheckout, viewValue: "BackRepoPointerEncodingFieldsCheckout" },
	{ value: BackRepoInsertionPoint.BackRepoPointerEncodingFieldsReindexing, viewValue: "BackRepoPointerEncodingFieldsReindexing" },
	{ value: BackRepoInsertionPoint.BackRepoNbInsertionPoints, viewValue: "BackRepoNbInsertionPoints" },
];
