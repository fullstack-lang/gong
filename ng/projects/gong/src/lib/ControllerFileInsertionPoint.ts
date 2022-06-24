// generated from ng_file_enum.ts.go
export enum ControllerFileInsertionPoint {
	// insertion point	
	ControllerFileGetsInsertion = 0,
	ControllerFilePostInsertion = 1,
	ControllerFileGetInsertion = 2,
	ControllerFileUpdateInsertion = 3,
	ControllerFileNbInsertionPoints = 4,
}

export interface ControllerFileInsertionPointSelect {
	value: number;
	viewValue: string;
}

export const ControllerFileInsertionPointList: ControllerFileInsertionPointSelect[] = [ // insertion point	
	{ value: ControllerFileInsertionPoint.ControllerFileGetsInsertion, viewValue: "ControllerFileGetsInsertion" },
	{ value: ControllerFileInsertionPoint.ControllerFilePostInsertion, viewValue: "ControllerFilePostInsertion" },
	{ value: ControllerFileInsertionPoint.ControllerFileGetInsertion, viewValue: "ControllerFileGetInsertion" },
	{ value: ControllerFileInsertionPoint.ControllerFileUpdateInsertion, viewValue: "ControllerFileUpdateInsertion" },
	{ value: ControllerFileInsertionPoint.ControllerFileNbInsertionPoints, viewValue: "ControllerFileNbInsertionPoints" },
];
