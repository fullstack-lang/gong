import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'
import { shareReplay } from 'rxjs/operators'

// insertion point sub template for services imports
import { ContentAPI } from './content-api'
import { Content, CopyContentAPIToContent } from './content'
import { ContentService } from './content.service'

import { PngImageAPI } from './pngimage-api'
import { PngImage, CopyPngImageAPIToPngImage } from './pngimage'
import { PngImageService } from './pngimage.service'

import { SvgImageAPI } from './svgimage-api'
import { SvgImage, CopySvgImageAPIToSvgImage } from './svgimage'
import { SvgImageService } from './svgimage.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/lib/markdown/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Contents = new Array<Content>() // array of front instances
	map_ID_Content = new Map<number, Content>() // map of front instances

	array_PngImages = new Array<PngImage>() // array of front instances
	map_ID_PngImage = new Map<number, PngImage>() // map of front instances

	array_SvgImages = new Array<SvgImage>() // array of front instances
	map_ID_SvgImage = new Map<number, SvgImage>() // map of front instances


	public GONG__Index = -1

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Content':
				return this.array_Contents as unknown as Array<Type>
			case 'PngImage':
				return this.array_PngImages as unknown as Array<Type>
			case 'SvgImage':
				return this.array_SvgImages as unknown as Array<Type>
			default:
				throw new Error("Type not recognized")
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Content':
				return this.map_ID_Content as unknown as Map<number, Type>
			case 'PngImage':
				return this.map_ID_PngImage as unknown as Map<number, Type>
			case 'SvgImage':
				return this.map_ID_SvgImage as unknown as Map<number, Type>
			default:
				throw new Error("Type not recognized")
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

	httpOptions = {
		headers: new HttpHeaders({ 'Content-Type': 'application/json' })
	}

	//
	// Store of all instances of the stack
	//
	frontRepo = new (FrontRepo)

	// Manage open WebSocket connections
	private webSocketConnections = new Map<string, Observable<FrontRepo>>()


	constructor(
		private http: HttpClient, // insertion point sub template 
		private contentService: ContentService,
		private pngimageService: PngImageService,
		private svgimageService: SvgImageService,
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
		)
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
		)
	}

	// typing of observable can be messy in typescript. Therefore, one force the type
	observableFrontRepo!: [
		Observable<null>, // see below for the of(null) observable
		// insertion point sub template 
		Observable<ContentAPI[]>,
		Observable<PngImageAPI[]>,
		Observable<SvgImageAPI[]>,
	]

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
			this.contentService.getContents(this.Name, this.frontRepo),
			this.pngimageService.getPngImages(this.Name, this.frontRepo),
			this.svgimageService.getSvgImages(this.Name, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						contents_,
						pngimages_,
						svgimages_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var contents: ContentAPI[]
						contents = contents_ as ContentAPI[]
						var pngimages: PngImageAPI[]
						pngimages = pngimages_ as PngImageAPI[]
						var svgimages: SvgImageAPI[]
						svgimages = svgimages_ as SvgImageAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_Contents = []
						this.frontRepo.map_ID_Content.clear()

						contents.forEach(
							contentAPI => {
								let content = new Content
								this.frontRepo.array_Contents.push(content)
								this.frontRepo.map_ID_Content.set(contentAPI.ID, content)
							}
						)

						// init the arrays
						this.frontRepo.array_PngImages = []
						this.frontRepo.map_ID_PngImage.clear()

						pngimages.forEach(
							pngimageAPI => {
								let pngimage = new PngImage
								this.frontRepo.array_PngImages.push(pngimage)
								this.frontRepo.map_ID_PngImage.set(pngimageAPI.ID, pngimage)
							}
						)

						// init the arrays
						this.frontRepo.array_SvgImages = []
						this.frontRepo.map_ID_SvgImage.clear()

						svgimages.forEach(
							svgimageAPI => {
								let svgimage = new SvgImage
								this.frontRepo.array_SvgImages.push(svgimage)
								this.frontRepo.map_ID_SvgImage.set(svgimageAPI.ID, svgimage)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						contents.forEach(
							contentAPI => {
								let content = this.frontRepo.map_ID_Content.get(contentAPI.ID)
								CopyContentAPIToContent(contentAPI, content!, this.frontRepo)
							}
						)

						// fill up front objects
						pngimages.forEach(
							pngimageAPI => {
								let pngimage = this.frontRepo.map_ID_PngImage.get(pngimageAPI.ID)
								CopyPngImageAPIToPngImage(pngimageAPI, pngimage!, this.frontRepo)
							}
						)

						// fill up front objects
						svgimages.forEach(
							svgimageAPI => {
								let svgimage = this.frontRepo.map_ID_SvgImage.get(svgimageAPI.ID)
								CopySvgImageAPIToSvgImage(svgimageAPI, svgimage!, this.frontRepo)
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

		// Check if a connection for this name already exists
		if (this.webSocketConnections.has(Name)) {
			return this.webSocketConnections.get(Name)!
		}

		//
		// Create a new connection
		//
		let host = window.location.host
		const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'

		if (host === 'localhost:4200') {
			host = 'localhost:8080'
		}

		// Construct the base path using the dynamic host and protocol
		// The API path remains the same.
		let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/lib/markdown/go/v1/ws/stage`

		let params = new HttpParams().set("Name", Name)
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`

		const newConnection$ = new Observable<FrontRepo>(observer => {
			const socket = new WebSocket(url)

			socket.onmessage = event => {
				const backRepoData = new BackRepoData(JSON.parse(event.data))
				let frontRepo = new (FrontRepo)()
				frontRepo.GONG__Index = backRepoData.GONG__Index

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				// insertion point sub template for init 
				// init the arrays
				frontRepo.array_Contents = []
				frontRepo.map_ID_Content.clear()

				backRepoData.ContentAPIs.forEach(
					contentAPI => {
						let content = new Content
						frontRepo.array_Contents.push(content)
						frontRepo.map_ID_Content.set(contentAPI.ID, content)
					}
				)

				// init the arrays
				frontRepo.array_PngImages = []
				frontRepo.map_ID_PngImage.clear()

				backRepoData.PngImageAPIs.forEach(
					pngimageAPI => {
						let pngimage = new PngImage
						frontRepo.array_PngImages.push(pngimage)
						frontRepo.map_ID_PngImage.set(pngimageAPI.ID, pngimage)
					}
				)

				// init the arrays
				frontRepo.array_SvgImages = []
				frontRepo.map_ID_SvgImage.clear()

				backRepoData.SvgImageAPIs.forEach(
					svgimageAPI => {
						let svgimage = new SvgImage
						frontRepo.array_SvgImages.push(svgimage)
						frontRepo.map_ID_SvgImage.set(svgimageAPI.ID, svgimage)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.ContentAPIs.forEach(
					contentAPI => {
						let content = frontRepo.map_ID_Content.get(contentAPI.ID)
						CopyContentAPIToContent(contentAPI, content!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.PngImageAPIs.forEach(
					pngimageAPI => {
						let pngimage = frontRepo.map_ID_PngImage.get(pngimageAPI.ID)
						CopyPngImageAPIToPngImage(pngimageAPI, pngimage!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SvgImageAPIs.forEach(
					svgimageAPI => {
						let svgimage = frontRepo.map_ID_SvgImage.get(svgimageAPI.ID)
						CopySvgImageAPIToSvgImage(svgimageAPI, svgimage!, frontRepo)
					}
				)


				observer.next(frontRepo)
			}

			socket.onerror = event => observer.error(event)
			socket.onclose = () => observer.complete()

			// Teardown logic: Called when the last subscriber unsubscribes.
			return () => {
				this.webSocketConnections.delete(Name) // Remove from cache
				socket.close()
			}
		}).pipe(
			// This is the key:
			// - shareReplay makes this a "multicast" observable, sharing the single WebSocket among subscribers.
			// - { bufferSize: 1, refCount: true } means:
			//   - bufferSize: 1 => new subscribers get the last emitted value immediately.
			//   - refCount: true => the connection starts with the first subscriber and stops with the last.
			shareReplay({ bufferSize: 1, refCount: true })
		)

		// Store the new connection observable in the map
		this.webSocketConnections.set(Name, newConnection$)
		return newConnection$
	}
}

// insertion point for get unique ID per struct 
export function getContentUniqueID(id: number): number {
	return 31 * id
}
export function getPngImageUniqueID(id: number): number {
	return 37 * id
}
export function getSvgImageUniqueID(id: number): number {
	return 41 * id
}
