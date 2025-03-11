// generated from ng_file_enum.ts.go
export enum EngineState {
	// insertion point	
	RUNNING = "RUNNING",
	PAUSED = "PAUSED",
	OVER = "OVER",
}

export interface EngineStateSelect {
	value: string;
	viewValue: string;
}

export const EngineStateList: EngineStateSelect[] = [ // insertion point	
	{ value: EngineState.RUNNING, viewValue: "RUNNING" },
	{ value: EngineState.PAUSED, viewValue: "PAUSED" },
	{ value: EngineState.OVER, viewValue: "OVER" },
];
