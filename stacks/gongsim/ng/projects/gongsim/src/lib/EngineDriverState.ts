// generated from ng_file_enum.ts.go
export enum EngineDriverState {
	// insertion point	
	CHECKOUT_AGENT_STATES = 1,
	COMMIT_AGENT_STATES = 0,
	FIRE_ONE_EVENT = 2,
	RESET_SIMULATION = 4,
	SLEEP_100_MS = 3,
	UNKOWN = 5,
}

export interface EngineDriverStateSelect {
	value: string;
	viewValue: string;
}

export const EngineDriverStateList: EngineDriverStateSelect[] = [ // insertion point	
	{ value: 'CHECKOUT_AGENT_STATES', viewValue: '1' },
	{ value: 'COMMIT_AGENT_STATES', viewValue: '0' },
	{ value: 'FIRE_ONE_EVENT', viewValue: '2' },
	{ value: 'RESET_SIMULATION', viewValue: '4' },
	{ value: 'SLEEP_100_MS', viewValue: '3' },
	{ value: 'UNKOWN', viewValue: '5' },
];
