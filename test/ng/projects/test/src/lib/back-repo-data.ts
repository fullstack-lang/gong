
import { AstructDB } from './astruct-db'
import { Astruct, CopyAstructDBToAstruct } from './astruct'
import { AstructService } from './astruct.service'

import { AstructBstruct2UseDB } from './astructbstruct2use-db'
import { AstructBstruct2Use, CopyAstructBstruct2UseDBToAstructBstruct2Use } from './astructbstruct2use'
import { AstructBstruct2UseService } from './astructbstruct2use.service'

import { AstructBstructUseDB } from './astructbstructuse-db'
import { AstructBstructUse, CopyAstructBstructUseDBToAstructBstructUse } from './astructbstructuse'
import { AstructBstructUseService } from './astructbstructuse.service'

import { BstructDB } from './bstruct-db'
import { Bstruct, CopyBstructDBToBstruct } from './bstruct'
import { BstructService } from './bstruct.service'

import { DstructDB } from './dstruct-db'
import { Dstruct, CopyDstructDBToDstruct } from './dstruct'
import { DstructService } from './dstruct.service'

export class BackRepoData {
	AstructDBs = new Array<AstructDB>() // array of front instances

	AstructBstruct2UseDBs = new Array<AstructBstruct2UseDB>() // array of front instances

	AstructBstructUseDBs = new Array<AstructBstructUseDB>() // array of front instances

	BstructDBs = new Array<BstructDB>() // array of front instances

	DstructDBs = new Array<DstructDB>() // array of front instances

	constructor(data?: Partial<BackRepoData>) {
		this.AstructDBs = data?.AstructDBs || [];
		this.AstructBstruct2UseDBs = data?.AstructBstruct2UseDBs || [];
		this.AstructBstructUseDBs = data?.AstructBstructUseDBs || [];
		this.BstructDBs = data?.BstructDBs || [];
		this.DstructDBs = data?.DstructDBs || [];
	}

}