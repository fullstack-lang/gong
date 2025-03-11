// generated from ng_file_enum.ts.go
export enum CommandType {
	// insertion point	
	COMMAND_PLAY = "PLAY",
	COMMAND_PAUSE = "PAUSE",
	COMMAND_FIRE_NEXT_EVENT = "FIRE_NEXT_EVENT",
	COMMAND_FIRE_EVENT_TILL_STATES_CHANGE = "FIRE_EVENT_TILL_STATES_CHANGE",
	COMMAND_RESET = "RESET",
	COMMAND_ADVANCE_10_MIN = "ADVANCE_10_MIN",
	INCREASE_SPEED_100_PERCENTS = "INCREASE_SPEED_100_PERCENTS",
	DECREASE_SPEED_50_PERCENTS = "DECREASE_SPEED_50_PERCENTS",
}

export interface CommandTypeSelect {
	value: string;
	viewValue: string;
}

export const CommandTypeList: CommandTypeSelect[] = [ // insertion point	
	{ value: CommandType.COMMAND_PLAY, viewValue: "PLAY" },
	{ value: CommandType.COMMAND_PAUSE, viewValue: "PAUSE" },
	{ value: CommandType.COMMAND_FIRE_NEXT_EVENT, viewValue: "FIRE_NEXT_EVENT" },
	{ value: CommandType.COMMAND_FIRE_EVENT_TILL_STATES_CHANGE, viewValue: "FIRE_EVENT_TILL_STATES_CHANGE" },
	{ value: CommandType.COMMAND_RESET, viewValue: "RESET" },
	{ value: CommandType.COMMAND_ADVANCE_10_MIN, viewValue: "ADVANCE_10_MIN" },
	{ value: CommandType.INCREASE_SPEED_100_PERCENTS, viewValue: "INCREASE_SPEED_100_PERCENTS" },
	{ value: CommandType.DECREASE_SPEED_50_PERCENTS, viewValue: "DECREASE_SPEED_50_PERCENTS" },
];
