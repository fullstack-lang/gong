import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { FileToDownloadAPI } from './filetodownload-api'
import { FileToDownload, CopyFileToDownloadAPIToFileToDownload } from './filetodownload'
import { FileToDownloadService } from './filetodownload.service'

import { FileToUploadAPI } from './filetoupload-api'
import { FileToUpload, CopyFileToUploadAPIToFileToUpload } from './filetoupload'
import { FileToUploadService } from './filetoupload.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/lib/load/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_FileToDownloads = new Array<FileToDownload>() // array of front instances
	map_ID_FileToDownload = new Map<number, FileToDownload>() // map of front instances

	array_FileToUploads = new Array<FileToUpload>() // array of front instances
	map_ID_FileToUpload = new Map<number, FileToUpload>() // map of front instances


	public GONG__Index = -1

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'FileToDownload':
				return this.array_FileToDownloads as unknown as Array<Type>
			case 'FileToUpload':
				return this.array_FileToUploads as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'FileToDownload':
				return this.map_ID_FileToDownload as unknown as Map<number, Type>
			case 'FileToUpload':
				return this.map_ID_FileToUpload as unknown as Map<number, Type>
			default:
				throw new Error("Type not recognized");
		}
	}
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
	ID: number = 0 // ID of the calling instance

	// the reverse pointer is the name of the generated field on the destination
	// struct of the ONE-MANY association
	ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
	OrderingMode: boolean = false // if true, this is for ordering items

	// there are different selection mode : ONE_MANY or MANY_MANY
	SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

	// used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
	//
	// In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
	// 
	// in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
	// at the end of the ONE-MANY association
	SourceStruct: string = ""	// The "Aclass"
	SourceField: string = "" // the "AnarrayofbUse"
	IntermediateStruct: string = "" // the "AclassBclassUse" 
	IntermediateStructField: string = "" // the "Bclass" as field
	NextAssociationStruct: string = "" // the "Bclass"

	Name: string = ""
}

export enum SelectionMode {
	ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
	MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
	providedIn: 'root'
})
export class FrontRepoService {

	Name: string = ""
	private socket: WebSocket | undefined

	httpOptions = {
		headers: new HttpHeaders({ 'Content-Type': 'application/json' })
	};

	//
	// Store of all instances of the stack
	//
	frontRepo = new (FrontRepo)

	constructor(
		private http: HttpClient, // insertion point sub template 
		private filetodownloadService: FileToDownloadService,
		private filetouploadService: FileToUploadService,
	) { }

	// postService provides a post function for each struct name
	postService(structName: string, instanceToBePosted: any) {
		let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
		let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

		servicePostFunction(instanceToBePosted).subscribe(
			instance => {
				let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("post")
			}
		);
	}

	// deleteService provides a delete function for each struct name
	deleteService(structName: string, instanceToBeDeleted: any) {
		let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
		let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

		serviceDeleteFunction(instanceToBeDeleted).subscribe(
			instance => {
				let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("delete")
			}
		);
	}

	// typing of observable can be messy in typescript. Therefore, one force the type
	observableFrontRepo!: [
		Observable<null>, // see below for the of(null) observable
		// insertion point sub template 
		Observable<FileToDownloadAPI[]>,
		Observable<FileToUploadAPI[]>,
	];

	//
	// pull performs a GET on all struct of the stack and redeem association pointers 
	//
	// This is an observable. Therefore, the control flow forks with
	// - pull() return immediatly the observable
	// - the observable observer, if it subscribe, is called when all GET calls are performs
	pull(Name: string = ""): Observable<FrontRepo> {

		this.Name = Name

		this.observableFrontRepo = [
			of(null), // see above for justification
			// insertion point sub template
			this.filetodownloadService.getFileToDownloads(this.Name, this.frontRepo),
			this.filetouploadService.getFileToUploads(this.Name, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						filetodownloads_,
						filetouploads_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var filetodownloads: FileToDownloadAPI[]
						filetodownloads = filetodownloads_ as FileToDownloadAPI[]
						var filetouploads: FileToUploadAPI[]
						filetouploads = filetouploads_ as FileToUploadAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_FileToDownloads = []
						this.frontRepo.map_ID_FileToDownload.clear()

						filetodownloads.forEach(
							filetodownloadAPI => {
								let filetodownload = new FileToDownload
								this.frontRepo.array_FileToDownloads.push(filetodownload)
								this.frontRepo.map_ID_FileToDownload.set(filetodownloadAPI.ID, filetodownload)
							}
						)

						// init the arrays
						this.frontRepo.array_FileToUploads = []
						this.frontRepo.map_ID_FileToUpload.clear()

						filetouploads.forEach(
							filetouploadAPI => {
								let filetoupload = new FileToUpload
								this.frontRepo.array_FileToUploads.push(filetoupload)
								this.frontRepo.map_ID_FileToUpload.set(filetouploadAPI.ID, filetoupload)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						filetodownloads.forEach(
							filetodownloadAPI => {
								let filetodownload = this.frontRepo.map_ID_FileToDownload.get(filetodownloadAPI.ID)
								CopyFileToDownloadAPIToFileToDownload(filetodownloadAPI, filetodownload!, this.frontRepo)
							}
						)

						// fill up front objects
						filetouploads.forEach(
							filetouploadAPI => {
								let filetoupload = this.frontRepo.map_ID_FileToUpload.get(filetouploadAPI.ID)
								CopyFileToUploadAPIToFileToUpload(filetouploadAPI, filetoupload!, this.frontRepo)
							}
						)


						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
	}

	public connectToWebSocket(Name: string): Observable<FrontRepo> {

		this.Name = Name


		// Determine the base URL for the WebSocket connection dynamically
		// window.location.host includes hostname and port (e.g., "localhost:8080" or "yourdomain.com:8090")
		// If running on standard ports (80 for http, 443 for https), the port might not be explicitly in window.location.host
		// but WebSocket constructor handles 'ws://' and 'wss://' correctly with host.
		let host = window.location.host; // e.g., localhost:4200 or myapp.com
		const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'; // Use wss for https, ws for http

		// Check if the host is localhost:4200 and change it to localhost:8080 (when using ng serve)
		if (host === 'localhost:4200') {
			host = 'localhost:8080';
		}

		// Construct the base path using the dynamic host and protocol
		// The API path remains the same.
		let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/lib/load/go/v1/ws/stage`

		let params = new HttpParams().set("Name", this.Name)
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`
		this.socket = new WebSocket(url)

		return new Observable(observer => {
			this.socket!.onmessage = event => {


				const backRepoData = new BackRepoData(JSON.parse(event.data))

				let frontRepo = new (FrontRepo)
				frontRepo.GONG__Index = backRepoData.GONG__Index

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				// insertion point sub template for init 
				// init the arrays
				frontRepo.array_FileToDownloads = []
				frontRepo.map_ID_FileToDownload.clear()

				backRepoData.FileToDownloadAPIs.forEach(
					filetodownloadAPI => {
						let filetodownload = new FileToDownload
						frontRepo.array_FileToDownloads.push(filetodownload)
						frontRepo.map_ID_FileToDownload.set(filetodownloadAPI.ID, filetodownload)
					}
				)

				// init the arrays
				frontRepo.array_FileToUploads = []
				frontRepo.map_ID_FileToUpload.clear()

				backRepoData.FileToUploadAPIs.forEach(
					filetouploadAPI => {
						let filetoupload = new FileToUpload
						frontRepo.array_FileToUploads.push(filetoupload)
						frontRepo.map_ID_FileToUpload.set(filetouploadAPI.ID, filetoupload)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.FileToDownloadAPIs.forEach(
					filetodownloadAPI => {
						let filetodownload = frontRepo.map_ID_FileToDownload.get(filetodownloadAPI.ID)
						CopyFileToDownloadAPIToFileToDownload(filetodownloadAPI, filetodownload!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FileToUploadAPIs.forEach(
					filetouploadAPI => {
						let filetoupload = frontRepo.map_ID_FileToUpload.get(filetouploadAPI.ID)
						CopyFileToUploadAPIToFileToUpload(filetouploadAPI, filetoupload!, frontRepo)
					}
				)



				observer.next(frontRepo)
			}
			this.socket!.onerror = event => {
				observer.error(event)
			}
			this.socket!.onclose = event => {
				observer.complete()
			}

			return () => {
				this.socket!.close()
			}
		})
	}
}

// insertion point for get unique ID per struct 
export function getFileToDownloadUniqueID(id: number): number {
	return 31 * id
}
export function getFileToUploadUniqueID(id: number): number {
	return 37 * id
}
