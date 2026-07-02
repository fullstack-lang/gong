// generated from ng_file_enum.ts.go
export enum ControlMode {
	// insertion point	
	AUTONOMOUS = "Autonomous",
	CLIENT_CONTROL = "ClientControl",
}

export interface ControlModeSelect {
	value: string;
	viewValue: string;
}

export const ControlModeList: ControlModeSelect[] = [ // insertion point	
	{ value: ControlMode.AUTONOMOUS, viewValue: "Autonomous" },
	{ value: ControlMode.CLIENT_CONTROL, viewValue: "ClientControl" },
];
