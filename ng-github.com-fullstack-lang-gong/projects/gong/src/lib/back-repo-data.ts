// generated code - do not edit

//insertion point for imports
import { GongBasicFieldAPI } from './gongbasicfield-api'

import { GongEnumAPI } from './gongenum-api'

import { GongEnumValueAPI } from './gongenumvalue-api'

import { GongLinkAPI } from './gonglink-api'

import { GongNoteAPI } from './gongnote-api'

import { GongStructAPI } from './gongstruct-api'

import { GongTimeFieldAPI } from './gongtimefield-api'

import { MetaAPI } from './meta-api'

import { MetaReferenceAPI } from './metareference-api'

import { ModelPkgAPI } from './modelpkg-api'

import { PointerToGongStructFieldAPI } from './pointertogongstructfield-api'

import { SliceOfPointerToGongStructFieldAPI } from './sliceofpointertogongstructfield-api'


export class BackRepoData {
	// insertion point for declarations
	GongBasicFieldAPIs = new Array<GongBasicFieldAPI>()

	GongEnumAPIs = new Array<GongEnumAPI>()

	GongEnumValueAPIs = new Array<GongEnumValueAPI>()

	GongLinkAPIs = new Array<GongLinkAPI>()

	GongNoteAPIs = new Array<GongNoteAPI>()

	GongStructAPIs = new Array<GongStructAPI>()

	GongTimeFieldAPIs = new Array<GongTimeFieldAPI>()

	MetaAPIs = new Array<MetaAPI>()

	MetaReferenceAPIs = new Array<MetaReferenceAPI>()

	ModelPkgAPIs = new Array<ModelPkgAPI>()

	PointerToGongStructFieldAPIs = new Array<PointerToGongStructFieldAPI>()

	SliceOfPointerToGongStructFieldAPIs = new Array<SliceOfPointerToGongStructFieldAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.GongBasicFieldAPIs = data?.GongBasicFieldAPIs || [];

		this.GongEnumAPIs = data?.GongEnumAPIs || [];

		this.GongEnumValueAPIs = data?.GongEnumValueAPIs || [];

		this.GongLinkAPIs = data?.GongLinkAPIs || [];

		this.GongNoteAPIs = data?.GongNoteAPIs || [];

		this.GongStructAPIs = data?.GongStructAPIs || [];

		this.GongTimeFieldAPIs = data?.GongTimeFieldAPIs || [];

		this.MetaAPIs = data?.MetaAPIs || [];

		this.MetaReferenceAPIs = data?.MetaReferenceAPIs || [];

		this.ModelPkgAPIs = data?.ModelPkgAPIs || [];

		this.PointerToGongStructFieldAPIs = data?.PointerToGongStructFieldAPIs || [];

		this.SliceOfPointerToGongStructFieldAPIs = data?.SliceOfPointerToGongStructFieldAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}