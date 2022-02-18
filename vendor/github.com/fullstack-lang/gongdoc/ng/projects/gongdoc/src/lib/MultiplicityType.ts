// generated from ng_file_enum.ts.go
export enum MultiplicityType {
	// insertion point	
	MANY = "*",
	ONE = "1",
	ZERO_ONE = "0..1",
}

export interface MultiplicityTypeSelect {
	value: string;
	viewValue: string;
}

export const MultiplicityTypeList: MultiplicityTypeSelect[] = [ // insertion point	
	{ value: 'MANY', viewValue: '"*"' },
	{ value: 'ONE', viewValue: '"1"' },
	{ value: 'ZERO_ONE', viewValue: '"0..1"' },
];
