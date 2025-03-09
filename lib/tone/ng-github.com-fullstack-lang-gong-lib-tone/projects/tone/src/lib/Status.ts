// generated from ng_file_enum.ts.go
export enum Status {
	// insertion point	
	PLAYING = "PLAYING",
	PAUSED = "PAUSED",
}

export interface StatusSelect {
	value: string;
	viewValue: string;
}

export const StatusList: StatusSelect[] = [ // insertion point	
	{ value: Status.PLAYING, viewValue: "PLAYING" },
	{ value: Status.PAUSED, viewValue: "PAUSED" },
];
