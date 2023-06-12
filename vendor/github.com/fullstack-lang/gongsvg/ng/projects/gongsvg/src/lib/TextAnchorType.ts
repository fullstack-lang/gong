// generated from ng_file_enum.ts.go
export enum TextAnchorType {
	// insertion point	
	TEXT_ANCHOR_START = "start",
	TEXT_ANCHOR_CENTER = "middle",
	TEXT_ANCHOR_END = "end",
}

export interface TextAnchorTypeSelect {
	value: string;
	viewValue: string;
}

export const TextAnchorTypeList: TextAnchorTypeSelect[] = [ // insertion point	
	{ value: TextAnchorType.TEXT_ANCHOR_START, viewValue: "start" },
	{ value: TextAnchorType.TEXT_ANCHOR_CENTER, viewValue: "middle" },
	{ value: TextAnchorType.TEXT_ANCHOR_END, viewValue: "end" },
];
