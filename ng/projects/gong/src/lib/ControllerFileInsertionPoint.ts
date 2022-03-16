// generated from ng_file_enum.ts.go
export enum ControllerFileInsertionPoint {
	// insertion point	
	ControllerFileGetInsertion = 2,
	ControllerFileGetsInsertion = 0,
	ControllerFileNbInsertionPoints = 4,
	ControllerFilePostInsertion = 1,
	ControllerFileUpdateInsertion = 3,
}

export interface ControllerFileInsertionPointSelect {
	value: number;
	viewValue: string;
}

export const ControllerFileInsertionPointList: ControllerFileInsertionPointSelect[] = [ // insertion point	
	{ value: ControllerFileInsertionPoint.ControllerFileGetInsertion, viewValue: "ControllerFileGetInsertion" },
	{ value: ControllerFileInsertionPoint.ControllerFileGetsInsertion, viewValue: "ControllerFileGetsInsertion" },
	{ value: ControllerFileInsertionPoint.ControllerFileNbInsertionPoints, viewValue: "ControllerFileNbInsertionPoints" },
	{ value: ControllerFileInsertionPoint.ControllerFilePostInsertion, viewValue: "ControllerFilePostInsertion" },
	{ value: ControllerFileInsertionPoint.ControllerFileUpdateInsertion, viewValue: "ControllerFileUpdateInsertion" },
];
