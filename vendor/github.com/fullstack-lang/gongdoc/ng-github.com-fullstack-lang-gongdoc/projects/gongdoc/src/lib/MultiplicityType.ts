// generated from ng_file_enum.ts.go
export enum MultiplicityType {
	// insertion point	
	ZERO_ONE = "0..1",
	ONE = "1",
	MANY = "*",
}

export interface MultiplicityTypeSelect {
	value: string;
	viewValue: string;
}

export const MultiplicityTypeList: MultiplicityTypeSelect[] = [ // insertion point	
	{ value: MultiplicityType.ZERO_ONE, viewValue: "0..1" },
	{ value: MultiplicityType.ONE, viewValue: "1" },
	{ value: MultiplicityType.MANY, viewValue: "*" },
];
