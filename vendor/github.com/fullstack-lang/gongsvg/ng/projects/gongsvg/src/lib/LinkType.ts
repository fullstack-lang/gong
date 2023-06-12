// generated from ng_file_enum.ts.go
export enum LinkType {
	// insertion point	
	LINK_TYPE_LINE_WITH_CONTROL_POINTS = "LINK_TYPE_LINE_WITH_CONTROL_POINTS",
	LINK_TYPE_FLOATING_ORTHOGONAL = "LINK_TYPE_FLOATING_ORTHOGONAL",
}

export interface LinkTypeSelect {
	value: string;
	viewValue: string;
}

export const LinkTypeList: LinkTypeSelect[] = [ // insertion point	
	{ value: LinkType.LINK_TYPE_LINE_WITH_CONTROL_POINTS, viewValue: "LINK_TYPE_LINE_WITH_CONTROL_POINTS" },
	{ value: LinkType.LINK_TYPE_FLOATING_ORTHOGONAL, viewValue: "LINK_TYPE_FLOATING_ORTHOGONAL" },
];
