// generated from ng_file_enum.ts.go
export enum BackRepoSubTemplateInsertion {
	// insertion point	
	BackRepoBackup = 10,
	BackRepoBackupXL = 11,
	BackRepoCheckout = 9,
	BackRepoCommit = 8,
	BackRepoInitAndCheckout = 7,
	BackRepoInitAndCommit = 6,
	BackRepoPerStructDeclarations = 0,
	BackRepoPerStructInits = 1,
	BackRepoPerStructPhaseOneCheckouts = 4,
	BackRepoPerStructPhaseOneCommits = 2,
	BackRepoPerStructPhaseTwoCheckouts = 5,
	BackRepoPerStructPhaseTwoCommits = 3,
	BackRepoRestorePhaseOne = 12,
	BackRepoRestorePhaseTwo = 14,
	BackRepoRestoreXLPhaseOne = 13,
}

export interface BackRepoSubTemplateInsertionSelect {
	value: number;
	viewValue: string;
}

export const BackRepoSubTemplateInsertionList: BackRepoSubTemplateInsertionSelect[] = [ // insertion point	
	{ value: BackRepoSubTemplateInsertion.BackRepoBackup, viewValue: "BackRepoBackup" },
	{ value: BackRepoSubTemplateInsertion.BackRepoBackupXL, viewValue: "BackRepoBackupXL" },
	{ value: BackRepoSubTemplateInsertion.BackRepoCheckout, viewValue: "BackRepoCheckout" },
	{ value: BackRepoSubTemplateInsertion.BackRepoCommit, viewValue: "BackRepoCommit" },
	{ value: BackRepoSubTemplateInsertion.BackRepoInitAndCheckout, viewValue: "BackRepoInitAndCheckout" },
	{ value: BackRepoSubTemplateInsertion.BackRepoInitAndCommit, viewValue: "BackRepoInitAndCommit" },
	{ value: BackRepoSubTemplateInsertion.BackRepoPerStructDeclarations, viewValue: "BackRepoPerStructDeclarations" },
	{ value: BackRepoSubTemplateInsertion.BackRepoPerStructInits, viewValue: "BackRepoPerStructInits" },
	{ value: BackRepoSubTemplateInsertion.BackRepoPerStructPhaseOneCheckouts, viewValue: "BackRepoPerStructPhaseOneCheckouts" },
	{ value: BackRepoSubTemplateInsertion.BackRepoPerStructPhaseOneCommits, viewValue: "BackRepoPerStructPhaseOneCommits" },
	{ value: BackRepoSubTemplateInsertion.BackRepoPerStructPhaseTwoCheckouts, viewValue: "BackRepoPerStructPhaseTwoCheckouts" },
	{ value: BackRepoSubTemplateInsertion.BackRepoPerStructPhaseTwoCommits, viewValue: "BackRepoPerStructPhaseTwoCommits" },
	{ value: BackRepoSubTemplateInsertion.BackRepoRestorePhaseOne, viewValue: "BackRepoRestorePhaseOne" },
	{ value: BackRepoSubTemplateInsertion.BackRepoRestorePhaseTwo, viewValue: "BackRepoRestorePhaseTwo" },
	{ value: BackRepoSubTemplateInsertion.BackRepoRestoreXLPhaseOne, viewValue: "BackRepoRestoreXLPhaseOne" },
];
