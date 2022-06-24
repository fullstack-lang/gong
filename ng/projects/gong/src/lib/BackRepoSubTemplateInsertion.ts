// generated from ng_file_enum.ts.go
export enum BackRepoSubTemplateInsertion {
	// insertion point	
	BackRepoPerStructDeclarations = 0,
	BackRepoPerStructInits = 1,
	BackRepoPerStructPhaseOneCommits = 2,
	BackRepoPerStructPhaseTwoCommits = 3,
	BackRepoPerStructPhaseOneCheckouts = 4,
	BackRepoPerStructPhaseTwoCheckouts = 5,
	BackRepoInitAndCommit = 6,
	BackRepoInitAndCheckout = 7,
	BackRepoCommit = 8,
	BackRepoCheckout = 9,
	BackRepoBackup = 10,
	BackRepoBackupXL = 11,
	BackRepoRestorePhaseOne = 12,
	BackRepoRestoreXLPhaseOne = 13,
	BackRepoRestorePhaseTwo = 14,
}

export interface BackRepoSubTemplateInsertionSelect {
	value: number;
	viewValue: string;
}

export const BackRepoSubTemplateInsertionList: BackRepoSubTemplateInsertionSelect[] = [ // insertion point	
	{ value: BackRepoSubTemplateInsertion.BackRepoPerStructDeclarations, viewValue: "BackRepoPerStructDeclarations" },
	{ value: BackRepoSubTemplateInsertion.BackRepoPerStructInits, viewValue: "BackRepoPerStructInits" },
	{ value: BackRepoSubTemplateInsertion.BackRepoPerStructPhaseOneCommits, viewValue: "BackRepoPerStructPhaseOneCommits" },
	{ value: BackRepoSubTemplateInsertion.BackRepoPerStructPhaseTwoCommits, viewValue: "BackRepoPerStructPhaseTwoCommits" },
	{ value: BackRepoSubTemplateInsertion.BackRepoPerStructPhaseOneCheckouts, viewValue: "BackRepoPerStructPhaseOneCheckouts" },
	{ value: BackRepoSubTemplateInsertion.BackRepoPerStructPhaseTwoCheckouts, viewValue: "BackRepoPerStructPhaseTwoCheckouts" },
	{ value: BackRepoSubTemplateInsertion.BackRepoInitAndCommit, viewValue: "BackRepoInitAndCommit" },
	{ value: BackRepoSubTemplateInsertion.BackRepoInitAndCheckout, viewValue: "BackRepoInitAndCheckout" },
	{ value: BackRepoSubTemplateInsertion.BackRepoCommit, viewValue: "BackRepoCommit" },
	{ value: BackRepoSubTemplateInsertion.BackRepoCheckout, viewValue: "BackRepoCheckout" },
	{ value: BackRepoSubTemplateInsertion.BackRepoBackup, viewValue: "BackRepoBackup" },
	{ value: BackRepoSubTemplateInsertion.BackRepoBackupXL, viewValue: "BackRepoBackupXL" },
	{ value: BackRepoSubTemplateInsertion.BackRepoRestorePhaseOne, viewValue: "BackRepoRestorePhaseOne" },
	{ value: BackRepoSubTemplateInsertion.BackRepoRestoreXLPhaseOne, viewValue: "BackRepoRestoreXLPhaseOne" },
	{ value: BackRepoSubTemplateInsertion.BackRepoRestorePhaseTwo, viewValue: "BackRepoRestorePhaseTwo" },
];
