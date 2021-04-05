// generated from ng_file_enum.ts.go
export enum BackRepoSubTemplateInsertion {
	// insertion point	
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
}

export interface BackRepoSubTemplateInsertionSelect {
	value: string;
	viewValue: string;
}

export const BackRepoSubTemplateInsertionList: BackRepoSubTemplateInsertionSelect[] = [ // insertion point	
	{ value: 'BackRepoCheckout', viewValue: '9' },
	{ value: 'BackRepoCommit', viewValue: '8' },
	{ value: 'BackRepoInitAndCheckout', viewValue: '7' },
	{ value: 'BackRepoInitAndCommit', viewValue: '6' },
	{ value: 'BackRepoPerStructDeclarations', viewValue: '0' },
	{ value: 'BackRepoPerStructInits', viewValue: '1' },
	{ value: 'BackRepoPerStructPhaseOneCheckouts', viewValue: '4' },
	{ value: 'BackRepoPerStructPhaseOneCommits', viewValue: '2' },
	{ value: 'BackRepoPerStructPhaseTwoCheckouts', viewValue: '5' },
	{ value: 'BackRepoPerStructPhaseTwoCommits', viewValue: '3' },
];
