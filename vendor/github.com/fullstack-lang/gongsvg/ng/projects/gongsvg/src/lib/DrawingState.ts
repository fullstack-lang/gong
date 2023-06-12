// generated from ng_file_enum.ts.go
export enum DrawingState {
	// insertion point	
	NOT_DRAWING_LINE = "NOT_DRAWING_LINE",
	DRAWING_LINE = "DRAWING_LINE",
}

export interface DrawingStateSelect {
	value: string;
	viewValue: string;
}

export const DrawingStateList: DrawingStateSelect[] = [ // insertion point	
	{ value: DrawingState.NOT_DRAWING_LINE, viewValue: "NOT_DRAWING_LINE" },
	{ value: DrawingState.DRAWING_LINE, viewValue: "DRAWING_LINE" },
];
