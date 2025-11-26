// generated from ng_file_enum.ts.go
export enum MatButtonType {
	// insertion point	
	MatButtonTypeBasic = "matButton",
	MatButtonTypeIcon = "matIconButton",
	MatButtonTypeFab = "matFab",
	MatButtonTypeMiniFab = "matMiniFab",
	MatButtonTypeExtendedFab = "matFab extended",
}

export interface MatButtonTypeSelect {
	value: string;
	viewValue: string;
}

export const MatButtonTypeList: MatButtonTypeSelect[] = [ // insertion point	
	{ value: MatButtonType.MatButtonTypeBasic, viewValue: "matButton" },
	{ value: MatButtonType.MatButtonTypeIcon, viewValue: "matIconButton" },
	{ value: MatButtonType.MatButtonTypeFab, viewValue: "matFab" },
	{ value: MatButtonType.MatButtonTypeMiniFab, viewValue: "matMiniFab" },
	{ value: MatButtonType.MatButtonTypeExtendedFab, viewValue: "matFab extended" },
];
