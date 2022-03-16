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
	value: number;
	viewValue: string;
}

export const BackRepoInsertionPointList: BackRepoInsertionPointSelect[] = [ // insertion point	
	{ value: BackRepoInsertionPoint.BackRepoBasicAndTimeFieldsName, viewValue: "BackRepoBasicAndTimeFieldsName" },
	{ value: BackRepoInsertionPoint.BackRepoBasicAndTimeFieldsWOPDeclaration, viewValue: "BackRepoBasicAndTimeFieldsWOPDeclaration" },
	{ value: BackRepoInsertionPoint.BackRepoBasicFieldsCheckout, viewValue: "BackRepoBasicFieldsCheckout" },
	{ value: BackRepoInsertionPoint.BackRepoBasicFieldsCommit, viewValue: "BackRepoBasicFieldsCommit" },
	{ value: BackRepoInsertionPoint.BackRepoBasicFieldsDeclaration, viewValue: "BackRepoBasicFieldsDeclaration" },
	{ value: BackRepoInsertionPoint.BackRepoNbInsertionPoints, viewValue: "BackRepoNbInsertionPoints" },
	{ value: BackRepoInsertionPoint.BackRepoPointerEncodingFieldsCheckout, viewValue: "BackRepoPointerEncodingFieldsCheckout" },
	{ value: BackRepoInsertionPoint.BackRepoPointerEncodingFieldsCommit, viewValue: "BackRepoPointerEncodingFieldsCommit" },
	{ value: BackRepoInsertionPoint.BackRepoPointerEncodingFieldsDeclaration, viewValue: "BackRepoPointerEncodingFieldsDeclaration" },
	{ value: BackRepoInsertionPoint.BackRepoPointerEncodingFieldsReindexing, viewValue: "BackRepoPointerEncodingFieldsReindexing" },
	{ value: BackRepoInsertionPoint.BackRepoPointerEncodingFieldsWOPDeclaration, viewValue: "BackRepoPointerEncodingFieldsWOPDeclaration" },
	{ value: BackRepoInsertionPoint.BackRepoWOPInitialIndex, viewValue: "BackRepoWOPInitialIndex" },
];
