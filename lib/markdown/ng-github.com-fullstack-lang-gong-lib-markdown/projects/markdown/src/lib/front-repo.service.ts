// generated code - do not edit
import { Injectable, NgZone } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, BehaviorSubject, of } from 'rxjs'
import { shareReplay } from 'rxjs/operators'

// insertion point sub template for services imports
import { ContentAPI } from './content-api'
import { Content, CopyContentAPIToContent } from './content'

import { JpgImageAPI } from './jpgimage-api'
import { JpgImage, CopyJpgImageAPIToJpgImage } from './jpgimage'

import { PngImageAPI } from './pngimage-api'
import { PngImage, CopyPngImageAPIToPngImage } from './pngimage'

import { SvgImageAPI } from './svgimage-api'
import { SvgImage, CopySvgImageAPIToSvgImage } from './svgimage'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/lib/markdown/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Contents = new Array<Content>() // array of front instances
	map_ID_Content = new Map<number, Content>() // map of front instances

	array_JpgImages = new Array<JpgImage>() // array of front instances
	map_ID_JpgImage = new Map<number, JpgImage>() // map of front instances

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
			case 'JpgImage':
				return this.array_JpgImages as unknown as Array<Type>
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
			case 'JpgImage':
				return this.map_ID_JpgImage as unknown as Map<number, Type>
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
		private http: HttpClient,
		private ngZone: NgZone,
	) { }

	//
	// pull returns the FrontRepo Observable
	//
	pull(Name: string = ""): Observable<FrontRepo> {
		this.Name = Name
		return of(this.frontRepo)
	}

	public connectToWebSocket(Name: string): Observable<FrontRepo> {

		// console.log("github.com/fullstack-lang/gong/lib/markdown/go; connectToWebSocket: started", Name)

		// Check if a connection for this name already exists
		if (this.webSocketConnections.has(Name)) {
			// console.log("github.com/fullstack-lang/gong/lib/markdown/go; connectToWebSocket: returning existing connection")
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
			// console.log("github.com/fullstack-lang/gong/lib/markdown/go; connectToWebSocket: new Observable created")

			let socket: WebSocket | undefined

			const isOfflineMode = window.location.protocol === 'file:' || window.document.getElementById('wasm-progress-container') !== null

			const processData = (dataString: string) => {
				// console.log("github.com/fullstack-lang/gong/lib/markdown/go; connectToWebSocket: processData called")
				const backRepoData = new BackRepoData(JSON.parse(dataString))
				let frontRepo = new (FrontRepo)()
				frontRepo.GONG__Index = backRepoData.GONG__Index

				// 
				// First Step: init map of instances
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
				frontRepo.array_JpgImages = []
				frontRepo.map_ID_JpgImage.clear()

				backRepoData.JpgImageAPIs.forEach(
					jpgimageAPI => {
						let jpgimage = new JpgImage
						frontRepo.array_JpgImages.push(jpgimage)
						frontRepo.map_ID_JpgImage.set(jpgimageAPI.ID, jpgimage)
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
				backRepoData.ContentAPIs.forEach(
					contentAPI => {
						let content = frontRepo.map_ID_Content.get(contentAPI.ID)
						CopyContentAPIToContent(contentAPI, content!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.JpgImageAPIs.forEach(
					jpgimageAPI => {
						let jpgimage = frontRepo.map_ID_JpgImage.get(jpgimageAPI.ID)
						CopyJpgImageAPIToJpgImage(jpgimageAPI, jpgimage!, frontRepo)
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


				this.ngZone.run(() => {
					observer.next(frontRepo)
				})
			}

			// 3. Connection Loop
			const attemptConnection = (retries: number): void => {
				// console.log("github.com/fullstack-lang/gong/lib/markdown/go; attemptConnection: retries =", retries, "isOfflineMode =", isOfflineMode)

				// A. WASM OFFLINE MODE (Check if Go is ready)
				if ((window as any).openWasmSocket) {
					// console.log("github.com/fullstack-lang/gong/lib/markdown/go; attemptConnection: openWasmSocket exists, calling it");
					(window as any).openWasmSocket("github.com/fullstack-lang/gong/lib/markdown/go", Name, processData);
					return;
				}

				// B. WAITING FOR WASM
				if (isOfflineMode && retries > 0) {
					// console.log("github.com/fullstack-lang/gong/lib/markdown/go; attemptConnection: WAITING FOR WASM. Retries left:", retries)
					setTimeout(() => attemptConnection(retries - 1), 100);
					return;
				}

				// C. STANDARD SERVER MODE
				if (!isOfflineMode) {
					// console.log("github.com/fullstack-lang/gong/lib/markdown/go; attemptConnection: STANDARD SERVER MODE. url =", url)
					socket = new WebSocket(url)
					socket.onopen = (event) => {
						// console.log("github.com/fullstack-lang/gong/lib/markdown/go; WebSocket: onopen", event)
					}
					socket.onmessage = event => {
						// console.log("github.com/fullstack-lang/gong/lib/markdown/go; WebSocket: onmessage")
						processData(event.data)
					}
					socket.onerror = event => {
						console.error("github.com/fullstack-lang/gong/lib/markdown/go WebSocket: onerror", event)
						observer.error(event)
					}
					socket.onclose = (event) => {
						// console.log("github.com/fullstack-lang/gong/lib/markdown/go; WebSocket: onclose", event)
						observer.complete()
					}
				} else {
					console.error("github.com/fullstack-lang/gong/lib/markdown/go, attemptConnection: Offline mode detected, but WASM backend failed to load.")
					observer.error("Offline mode detected, but WASM backend failed to load.");
				}
			};

			attemptConnection(50);

			// Teardown logic: Called when the last subscriber unsubscribes.
			return () => {
				this.webSocketConnections.delete(Name) // Remove from cache
				if (socket) {
					socket.close()
				}
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
export function getJpgImageUniqueID(id: number): number {
	return 37 * id
}
export function getPngImageUniqueID(id: number): number {
	return 41 * id
}
export function getSvgImageUniqueID(id: number): number {
	return 43 * id
}
