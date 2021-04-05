// generated from ng_file_enum.ts.go
export enum SpeedCommandType {
	// insertion point	
	COMMAND_DECREASE_SPEED_50_PERCENTS = "COMMAND_DECREASE_SPEED_50_PERCENTS ",
	COMMAND_INCREASE_SPEED_100_PERCENTS = "INCREASE_SPEED_100_PERCENTS",
	COMMAND_SPEED_STEADY = "COMMAND_SPEED_STEADY",
}

export interface SpeedCommandTypeSelect {
	value: string;
	viewValue: string;
}

export const SpeedCommandTypeList: SpeedCommandTypeSelect[] = [ // insertion point	
	{ value: 'COMMAND_DECREASE_SPEED_50_PERCENTS', viewValue: '"COMMAND_DECREASE_SPEED_50_PERCENTS "' },
	{ value: 'COMMAND_INCREASE_SPEED_100_PERCENTS', viewValue: '"INCREASE_SPEED_100_PERCENTS"' },
	{ value: 'COMMAND_SPEED_STEADY', viewValue: '"COMMAND_SPEED_STEADY"' },
];
