// generated from ng_file_enum.ts.go
export enum LinkTargetType {
	// insertion point	
	LINK_TARGET_BLANK = "_blank",
	LINK_TARGET_SELF = "_self",
	LINK_TARGET_PARENT = "_parent",
	LINK_TARGET_TOP = "_top",
}

export interface LinkTargetTypeSelect {
	value: string;
	viewValue: string;
}

export const LinkTargetTypeList: LinkTargetTypeSelect[] = [ // insertion point	
	{ value: LinkTargetType.LINK_TARGET_BLANK, viewValue: "_blank" },
	{ value: LinkTargetType.LINK_TARGET_SELF, viewValue: "_self" },
	{ value: LinkTargetType.LINK_TARGET_PARENT, viewValue: "_parent" },
	{ value: LinkTargetType.LINK_TARGET_TOP, viewValue: "_top" },
];
