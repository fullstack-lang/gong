// generated from ng_file_enum.ts.go
export enum ButtonType {
	// insertion point	
	DUPLICATE = 0,
	EDIT_CANCEL = 1,
	EDIT = 2,
	REMOVE = 3,
	RENAME_CANCEL = 4,
	RENAME = 5,
	SAVE = 6,
}

export interface ButtonTypeSelect {
	value: number;
	viewValue: string;
}

export const ButtonTypeList: ButtonTypeSelect[] = [ // insertion point	
	{ value: ButtonType.DUPLICATE, viewValue: "DUPLICATE" },
	{ value: ButtonType.EDIT_CANCEL, viewValue: "EDIT_CANCEL" },
	{ value: ButtonType.EDIT, viewValue: "EDIT" },
	{ value: ButtonType.REMOVE, viewValue: "REMOVE" },
	{ value: ButtonType.RENAME_CANCEL, viewValue: "RENAME_CANCEL" },
	{ value: ButtonType.RENAME, viewValue: "RENAME" },
	{ value: ButtonType.SAVE, viewValue: "SAVE" },
];
