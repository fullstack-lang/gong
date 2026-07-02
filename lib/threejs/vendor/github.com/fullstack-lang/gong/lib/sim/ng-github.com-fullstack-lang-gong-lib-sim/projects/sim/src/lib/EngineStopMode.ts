// generated from ng_file_enum.ts.go
export enum EngineStopMode {
	// insertion point	
	TEN_MINUTES = 0,
	STATE_CHANGED = 1,
}

export interface EngineStopModeSelect {
	value: number;
	viewValue: string;
}

export const EngineStopModeList: EngineStopModeSelect[] = [ // insertion point	
	{ value: EngineStopMode.TEN_MINUTES, viewValue: "TEN_MINUTES" },
	{ value: EngineStopMode.STATE_CHANGED, viewValue: "STATE_CHANGED" },
];
