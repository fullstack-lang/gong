// generated code - do not edit

//insertion point for imports
import { FileToDownloadAPI } from './filetodownload-api'

import { FileToUploadAPI } from './filetoupload-api'


export class BackRepoData {
	// insertion point for declarations
	FileToDownloadAPIs = new Array<FileToDownloadAPI>()

	FileToUploadAPIs = new Array<FileToUploadAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.FileToDownloadAPIs = data?.FileToDownloadAPIs || [];

		this.FileToUploadAPIs = data?.FileToUploadAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}