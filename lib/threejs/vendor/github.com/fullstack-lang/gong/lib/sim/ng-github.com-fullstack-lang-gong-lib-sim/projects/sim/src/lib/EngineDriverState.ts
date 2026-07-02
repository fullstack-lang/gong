// generated from ng_file_enum.ts.go
export enum EngineDriverState {
	// insertion point	
	COMMIT_AGENT_STATES = 0,
	CHECKOUT_AGENT_STATES = 1,
	FIRE_ONE_EVENT = 2,
	SLEEP_100_MS = 3,
	RESET_SIMULATION = 4,
	UNKOWN = 5,
}

export interface EngineDriverStateSelect {
	value: number;
	viewValue: string;
}

export const EngineDriverStateList: EngineDriverStateSelect[] = [ // insertion point	
	{ value: EngineDriverState.COMMIT_AGENT_STATES, viewValue: "COMMIT_AGENT_STATES" },
	{ value: EngineDriverState.CHECKOUT_AGENT_STATES, viewValue: "CHECKOUT_AGENT_STATES" },
	{ value: EngineDriverState.FIRE_ONE_EVENT, viewValue: "FIRE_ONE_EVENT" },
	{ value: EngineDriverState.SLEEP_100_MS, viewValue: "SLEEP_100_MS" },
	{ value: EngineDriverState.RESET_SIMULATION, viewValue: "RESET_SIMULATION" },
	{ value: EngineDriverState.UNKOWN, viewValue: "UNKOWN" },
];
