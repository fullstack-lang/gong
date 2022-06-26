// generated from ng_file_enum.ts.go
export enum EditionMode {
	// insertion point	
	PRODUCTION_MODE = "PRODUCTION_MODE",
	DEVELOPMENT_MODE = "DEVELOPMENT_MODE",
}

export interface EditionModeSelect {
	value: string;
	viewValue: string;
}

export const EditionModeList: EditionModeSelect[] = [ // insertion point	
	{ value: EditionMode.PRODUCTION_MODE, viewValue: "PRODUCTION_MODE" },
	{ value: EditionMode.DEVELOPMENT_MODE, viewValue: "DEVELOPMENT_MODE" },
];
