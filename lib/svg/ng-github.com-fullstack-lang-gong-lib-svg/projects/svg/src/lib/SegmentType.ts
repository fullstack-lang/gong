// generated from ng_file_enum.ts.go
export enum SegmentType {
	// insertion point	
	StartSegment = "Start Segment",
	MiddleSegment = "Middle Segment",
	EndSegment = "End Segment",
}

export interface SegmentTypeSelect {
	value: string;
	viewValue: string;
}

export const SegmentTypeList: SegmentTypeSelect[] = [ // insertion point	
	{ value: SegmentType.StartSegment, viewValue: "Start Segment" },
	{ value: SegmentType.MiddleSegment, viewValue: "Middle Segment" },
	{ value: SegmentType.EndSegment, viewValue: "End Segment" },
];
