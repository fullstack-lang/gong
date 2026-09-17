// generated code - do not edit
import { Injectable, NgZone } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, BehaviorSubject, of } from 'rxjs'
import { shareReplay } from 'rxjs/operators'

// insertion point sub template for services imports
import { AsSplitAPI } from './assplit-api'
import { AsSplit, CopyAsSplitAPIToAsSplit } from './assplit'

import { AsSplitAreaAPI } from './assplitarea-api'
import { AsSplitArea, CopyAsSplitAreaAPIToAsSplitArea } from './assplitarea'

import { ButtonAPI } from './button-api'
import { Button, CopyButtonAPIToButton } from './button'

import { CursorAPI } from './cursor-api'
import { Cursor, CopyCursorAPIToCursor } from './cursor'

import { FavIconAPI } from './favicon-api'
import { FavIcon, CopyFavIconAPIToFavIcon } from './favicon'

import { FormAPI } from './form-api'
import { Form, CopyFormAPIToForm } from './form'

import { LoadAPI } from './load-api'
import { Load, CopyLoadAPIToLoad } from './load'

import { LogoOnTheLeftAPI } from './logoontheleft-api'
import { LogoOnTheLeft, CopyLogoOnTheLeftAPIToLogoOnTheLeft } from './logoontheleft'

import { LogoOnTheRightAPI } from './logoontheright-api'
import { LogoOnTheRight, CopyLogoOnTheRightAPIToLogoOnTheRight } from './logoontheright'

import { MarkdownAPI } from './markdown-api'
import { Markdown, CopyMarkdownAPIToMarkdown } from './markdown'

import { SliderAPI } from './slider-api'
import { Slider, CopySliderAPIToSlider } from './slider'

import { SplitAPI } from './split-api'
import { Split, CopySplitAPIToSplit } from './split'

import { SvgAPI } from './svg-api'
import { Svg, CopySvgAPIToSvg } from './svg'

import { TableAPI } from './table-api'
import { Table, CopyTableAPIToTable } from './table'

import { ThreejsAPI } from './threejs-api'
import { Threejs, CopyThreejsAPIToThreejs } from './threejs'

import { TitleAPI } from './title-api'
import { Title, CopyTitleAPIToTitle } from './title'

import { ToneAPI } from './tone-api'
import { Tone, CopyToneAPIToTone } from './tone'

import { TreeAPI } from './tree-api'
import { Tree, CopyTreeAPIToTree } from './tree'

import { ViewAPI } from './view-api'
import { View, CopyViewAPIToView } from './view'

import { XlsxAPI } from './xlsx-api'
import { Xlsx, CopyXlsxAPIToXlsx } from './xlsx'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/lib/split/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_AsSplits = new Array<AsSplit>() // array of front instances
	map_ID_AsSplit = new Map<number, AsSplit>() // map of front instances

	array_AsSplitAreas = new Array<AsSplitArea>() // array of front instances
	map_ID_AsSplitArea = new Map<number, AsSplitArea>() // map of front instances

	array_Buttons = new Array<Button>() // array of front instances
	map_ID_Button = new Map<number, Button>() // map of front instances

	array_Cursors = new Array<Cursor>() // array of front instances
	map_ID_Cursor = new Map<number, Cursor>() // map of front instances

	array_FavIcons = new Array<FavIcon>() // array of front instances
	map_ID_FavIcon = new Map<number, FavIcon>() // map of front instances

	array_Forms = new Array<Form>() // array of front instances
	map_ID_Form = new Map<number, Form>() // map of front instances

	array_Loads = new Array<Load>() // array of front instances
	map_ID_Load = new Map<number, Load>() // map of front instances

	array_LogoOnTheLefts = new Array<LogoOnTheLeft>() // array of front instances
	map_ID_LogoOnTheLeft = new Map<number, LogoOnTheLeft>() // map of front instances

	array_LogoOnTheRights = new Array<LogoOnTheRight>() // array of front instances
	map_ID_LogoOnTheRight = new Map<number, LogoOnTheRight>() // map of front instances

	array_Markdowns = new Array<Markdown>() // array of front instances
	map_ID_Markdown = new Map<number, Markdown>() // map of front instances

	array_Sliders = new Array<Slider>() // array of front instances
	map_ID_Slider = new Map<number, Slider>() // map of front instances

	array_Splits = new Array<Split>() // array of front instances
	map_ID_Split = new Map<number, Split>() // map of front instances

	array_Svgs = new Array<Svg>() // array of front instances
	map_ID_Svg = new Map<number, Svg>() // map of front instances

	array_Tables = new Array<Table>() // array of front instances
	map_ID_Table = new Map<number, Table>() // map of front instances

	array_Threejss = new Array<Threejs>() // array of front instances
	map_ID_Threejs = new Map<number, Threejs>() // map of front instances

	array_Titles = new Array<Title>() // array of front instances
	map_ID_Title = new Map<number, Title>() // map of front instances

	array_Tones = new Array<Tone>() // array of front instances
	map_ID_Tone = new Map<number, Tone>() // map of front instances

	array_Trees = new Array<Tree>() // array of front instances
	map_ID_Tree = new Map<number, Tree>() // map of front instances

	array_Views = new Array<View>() // array of front instances
	map_ID_View = new Map<number, View>() // map of front instances

	array_Xlsxs = new Array<Xlsx>() // array of front instances
	map_ID_Xlsx = new Map<number, Xlsx>() // map of front instances


	public GONG__Index = -1

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'AsSplit':
				return this.array_AsSplits as unknown as Array<Type>
			case 'AsSplitArea':
				return this.array_AsSplitAreas as unknown as Array<Type>
			case 'Button':
				return this.array_Buttons as unknown as Array<Type>
			case 'Cursor':
				return this.array_Cursors as unknown as Array<Type>
			case 'FavIcon':
				return this.array_FavIcons as unknown as Array<Type>
			case 'Form':
				return this.array_Forms as unknown as Array<Type>
			case 'Load':
				return this.array_Loads as unknown as Array<Type>
			case 'LogoOnTheLeft':
				return this.array_LogoOnTheLefts as unknown as Array<Type>
			case 'LogoOnTheRight':
				return this.array_LogoOnTheRights as unknown as Array<Type>
			case 'Markdown':
				return this.array_Markdowns as unknown as Array<Type>
			case 'Slider':
				return this.array_Sliders as unknown as Array<Type>
			case 'Split':
				return this.array_Splits as unknown as Array<Type>
			case 'Svg':
				return this.array_Svgs as unknown as Array<Type>
			case 'Table':
				return this.array_Tables as unknown as Array<Type>
			case 'Threejs':
				return this.array_Threejss as unknown as Array<Type>
			case 'Title':
				return this.array_Titles as unknown as Array<Type>
			case 'Tone':
				return this.array_Tones as unknown as Array<Type>
			case 'Tree':
				return this.array_Trees as unknown as Array<Type>
			case 'View':
				return this.array_Views as unknown as Array<Type>
			case 'Xlsx':
				return this.array_Xlsxs as unknown as Array<Type>
			default:
				throw new Error("Type not recognized")
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'AsSplit':
				return this.map_ID_AsSplit as unknown as Map<number, Type>
			case 'AsSplitArea':
				return this.map_ID_AsSplitArea as unknown as Map<number, Type>
			case 'Button':
				return this.map_ID_Button as unknown as Map<number, Type>
			case 'Cursor':
				return this.map_ID_Cursor as unknown as Map<number, Type>
			case 'FavIcon':
				return this.map_ID_FavIcon as unknown as Map<number, Type>
			case 'Form':
				return this.map_ID_Form as unknown as Map<number, Type>
			case 'Load':
				return this.map_ID_Load as unknown as Map<number, Type>
			case 'LogoOnTheLeft':
				return this.map_ID_LogoOnTheLeft as unknown as Map<number, Type>
			case 'LogoOnTheRight':
				return this.map_ID_LogoOnTheRight as unknown as Map<number, Type>
			case 'Markdown':
				return this.map_ID_Markdown as unknown as Map<number, Type>
			case 'Slider':
				return this.map_ID_Slider as unknown as Map<number, Type>
			case 'Split':
				return this.map_ID_Split as unknown as Map<number, Type>
			case 'Svg':
				return this.map_ID_Svg as unknown as Map<number, Type>
			case 'Table':
				return this.map_ID_Table as unknown as Map<number, Type>
			case 'Threejs':
				return this.map_ID_Threejs as unknown as Map<number, Type>
			case 'Title':
				return this.map_ID_Title as unknown as Map<number, Type>
			case 'Tone':
				return this.map_ID_Tone as unknown as Map<number, Type>
			case 'Tree':
				return this.map_ID_Tree as unknown as Map<number, Type>
			case 'View':
				return this.map_ID_View as unknown as Map<number, Type>
			case 'Xlsx':
				return this.map_ID_Xlsx as unknown as Map<number, Type>
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

		// console.log("github.com/fullstack-lang/gong/lib/split/go; connectToWebSocket: started", Name)

		// Check if a connection for this name already exists
		if (this.webSocketConnections.has(Name)) {
			// console.log("github.com/fullstack-lang/gong/lib/split/go; connectToWebSocket: returning existing connection")
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
		let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/lib/split/go/v1/ws/stage`

		let params = new HttpParams().set("Name", Name)
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`

		const newConnection$ = new Observable<FrontRepo>(observer => {
			// console.log("github.com/fullstack-lang/gong/lib/split/go; connectToWebSocket: new Observable created")

			let socket: WebSocket | undefined

			const isOfflineMode = window.location.protocol === 'file:' || window.document.getElementById('wasm-progress-container') !== null

			const processData = (dataString: string) => {
				// console.log("github.com/fullstack-lang/gong/lib/split/go; connectToWebSocket: processData called")
				const backRepoData = new BackRepoData(JSON.parse(dataString))
				let frontRepo = new (FrontRepo)()
				frontRepo.GONG__Index = backRepoData.GONG__Index

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				frontRepo.array_AsSplits = []
				frontRepo.map_ID_AsSplit.clear()

				backRepoData.AsSplitAPIs.forEach(
					assplitAPI => {
						let assplit = new AsSplit
						frontRepo.array_AsSplits.push(assplit)
						frontRepo.map_ID_AsSplit.set(assplitAPI.ID, assplit)
					}
				)

				// init the arrays
				frontRepo.array_AsSplitAreas = []
				frontRepo.map_ID_AsSplitArea.clear()

				backRepoData.AsSplitAreaAPIs.forEach(
					assplitareaAPI => {
						let assplitarea = new AsSplitArea
						frontRepo.array_AsSplitAreas.push(assplitarea)
						frontRepo.map_ID_AsSplitArea.set(assplitareaAPI.ID, assplitarea)
					}
				)

				// init the arrays
				frontRepo.array_Buttons = []
				frontRepo.map_ID_Button.clear()

				backRepoData.ButtonAPIs.forEach(
					buttonAPI => {
						let button = new Button
						frontRepo.array_Buttons.push(button)
						frontRepo.map_ID_Button.set(buttonAPI.ID, button)
					}
				)

				// init the arrays
				frontRepo.array_Cursors = []
				frontRepo.map_ID_Cursor.clear()

				backRepoData.CursorAPIs.forEach(
					cursorAPI => {
						let cursor = new Cursor
						frontRepo.array_Cursors.push(cursor)
						frontRepo.map_ID_Cursor.set(cursorAPI.ID, cursor)
					}
				)

				// init the arrays
				frontRepo.array_FavIcons = []
				frontRepo.map_ID_FavIcon.clear()

				backRepoData.FavIconAPIs.forEach(
					faviconAPI => {
						let favicon = new FavIcon
						frontRepo.array_FavIcons.push(favicon)
						frontRepo.map_ID_FavIcon.set(faviconAPI.ID, favicon)
					}
				)

				// init the arrays
				frontRepo.array_Forms = []
				frontRepo.map_ID_Form.clear()

				backRepoData.FormAPIs.forEach(
					formAPI => {
						let form = new Form
						frontRepo.array_Forms.push(form)
						frontRepo.map_ID_Form.set(formAPI.ID, form)
					}
				)

				// init the arrays
				frontRepo.array_Loads = []
				frontRepo.map_ID_Load.clear()

				backRepoData.LoadAPIs.forEach(
					loadAPI => {
						let load = new Load
						frontRepo.array_Loads.push(load)
						frontRepo.map_ID_Load.set(loadAPI.ID, load)
					}
				)

				// init the arrays
				frontRepo.array_LogoOnTheLefts = []
				frontRepo.map_ID_LogoOnTheLeft.clear()

				backRepoData.LogoOnTheLeftAPIs.forEach(
					logoontheleftAPI => {
						let logoontheleft = new LogoOnTheLeft
						frontRepo.array_LogoOnTheLefts.push(logoontheleft)
						frontRepo.map_ID_LogoOnTheLeft.set(logoontheleftAPI.ID, logoontheleft)
					}
				)

				// init the arrays
				frontRepo.array_LogoOnTheRights = []
				frontRepo.map_ID_LogoOnTheRight.clear()

				backRepoData.LogoOnTheRightAPIs.forEach(
					logoontherightAPI => {
						let logoontheright = new LogoOnTheRight
						frontRepo.array_LogoOnTheRights.push(logoontheright)
						frontRepo.map_ID_LogoOnTheRight.set(logoontherightAPI.ID, logoontheright)
					}
				)

				// init the arrays
				frontRepo.array_Markdowns = []
				frontRepo.map_ID_Markdown.clear()

				backRepoData.MarkdownAPIs.forEach(
					markdownAPI => {
						let markdown = new Markdown
						frontRepo.array_Markdowns.push(markdown)
						frontRepo.map_ID_Markdown.set(markdownAPI.ID, markdown)
					}
				)

				// init the arrays
				frontRepo.array_Sliders = []
				frontRepo.map_ID_Slider.clear()

				backRepoData.SliderAPIs.forEach(
					sliderAPI => {
						let slider = new Slider
						frontRepo.array_Sliders.push(slider)
						frontRepo.map_ID_Slider.set(sliderAPI.ID, slider)
					}
				)

				// init the arrays
				frontRepo.array_Splits = []
				frontRepo.map_ID_Split.clear()

				backRepoData.SplitAPIs.forEach(
					splitAPI => {
						let split = new Split
						frontRepo.array_Splits.push(split)
						frontRepo.map_ID_Split.set(splitAPI.ID, split)
					}
				)

				// init the arrays
				frontRepo.array_Svgs = []
				frontRepo.map_ID_Svg.clear()

				backRepoData.SvgAPIs.forEach(
					svgAPI => {
						let svg = new Svg
						frontRepo.array_Svgs.push(svg)
						frontRepo.map_ID_Svg.set(svgAPI.ID, svg)
					}
				)

				// init the arrays
				frontRepo.array_Tables = []
				frontRepo.map_ID_Table.clear()

				backRepoData.TableAPIs.forEach(
					tableAPI => {
						let table = new Table
						frontRepo.array_Tables.push(table)
						frontRepo.map_ID_Table.set(tableAPI.ID, table)
					}
				)

				// init the arrays
				frontRepo.array_Threejss = []
				frontRepo.map_ID_Threejs.clear()

				backRepoData.ThreejsAPIs.forEach(
					threejsAPI => {
						let threejs = new Threejs
						frontRepo.array_Threejss.push(threejs)
						frontRepo.map_ID_Threejs.set(threejsAPI.ID, threejs)
					}
				)

				// init the arrays
				frontRepo.array_Titles = []
				frontRepo.map_ID_Title.clear()

				backRepoData.TitleAPIs.forEach(
					titleAPI => {
						let title = new Title
						frontRepo.array_Titles.push(title)
						frontRepo.map_ID_Title.set(titleAPI.ID, title)
					}
				)

				// init the arrays
				frontRepo.array_Tones = []
				frontRepo.map_ID_Tone.clear()

				backRepoData.ToneAPIs.forEach(
					toneAPI => {
						let tone = new Tone
						frontRepo.array_Tones.push(tone)
						frontRepo.map_ID_Tone.set(toneAPI.ID, tone)
					}
				)

				// init the arrays
				frontRepo.array_Trees = []
				frontRepo.map_ID_Tree.clear()

				backRepoData.TreeAPIs.forEach(
					treeAPI => {
						let tree = new Tree
						frontRepo.array_Trees.push(tree)
						frontRepo.map_ID_Tree.set(treeAPI.ID, tree)
					}
				)

				// init the arrays
				frontRepo.array_Views = []
				frontRepo.map_ID_View.clear()

				backRepoData.ViewAPIs.forEach(
					viewAPI => {
						let view = new View
						frontRepo.array_Views.push(view)
						frontRepo.map_ID_View.set(viewAPI.ID, view)
					}
				)

				// init the arrays
				frontRepo.array_Xlsxs = []
				frontRepo.map_ID_Xlsx.clear()

				backRepoData.XlsxAPIs.forEach(
					xlsxAPI => {
						let xlsx = new Xlsx
						frontRepo.array_Xlsxs.push(xlsx)
						frontRepo.map_ID_Xlsx.set(xlsxAPI.ID, xlsx)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.AsSplitAPIs.forEach(
					assplitAPI => {
						let assplit = frontRepo.map_ID_AsSplit.get(assplitAPI.ID)
						CopyAsSplitAPIToAsSplit(assplitAPI, assplit!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.AsSplitAreaAPIs.forEach(
					assplitareaAPI => {
						let assplitarea = frontRepo.map_ID_AsSplitArea.get(assplitareaAPI.ID)
						CopyAsSplitAreaAPIToAsSplitArea(assplitareaAPI, assplitarea!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.ButtonAPIs.forEach(
					buttonAPI => {
						let button = frontRepo.map_ID_Button.get(buttonAPI.ID)
						CopyButtonAPIToButton(buttonAPI, button!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CursorAPIs.forEach(
					cursorAPI => {
						let cursor = frontRepo.map_ID_Cursor.get(cursorAPI.ID)
						CopyCursorAPIToCursor(cursorAPI, cursor!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FavIconAPIs.forEach(
					faviconAPI => {
						let favicon = frontRepo.map_ID_FavIcon.get(faviconAPI.ID)
						CopyFavIconAPIToFavIcon(faviconAPI, favicon!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormAPIs.forEach(
					formAPI => {
						let form = frontRepo.map_ID_Form.get(formAPI.ID)
						CopyFormAPIToForm(formAPI, form!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.LoadAPIs.forEach(
					loadAPI => {
						let load = frontRepo.map_ID_Load.get(loadAPI.ID)
						CopyLoadAPIToLoad(loadAPI, load!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.LogoOnTheLeftAPIs.forEach(
					logoontheleftAPI => {
						let logoontheleft = frontRepo.map_ID_LogoOnTheLeft.get(logoontheleftAPI.ID)
						CopyLogoOnTheLeftAPIToLogoOnTheLeft(logoontheleftAPI, logoontheleft!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.LogoOnTheRightAPIs.forEach(
					logoontherightAPI => {
						let logoontheright = frontRepo.map_ID_LogoOnTheRight.get(logoontherightAPI.ID)
						CopyLogoOnTheRightAPIToLogoOnTheRight(logoontherightAPI, logoontheright!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.MarkdownAPIs.forEach(
					markdownAPI => {
						let markdown = frontRepo.map_ID_Markdown.get(markdownAPI.ID)
						CopyMarkdownAPIToMarkdown(markdownAPI, markdown!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SliderAPIs.forEach(
					sliderAPI => {
						let slider = frontRepo.map_ID_Slider.get(sliderAPI.ID)
						CopySliderAPIToSlider(sliderAPI, slider!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SplitAPIs.forEach(
					splitAPI => {
						let split = frontRepo.map_ID_Split.get(splitAPI.ID)
						CopySplitAPIToSplit(splitAPI, split!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SvgAPIs.forEach(
					svgAPI => {
						let svg = frontRepo.map_ID_Svg.get(svgAPI.ID)
						CopySvgAPIToSvg(svgAPI, svg!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.TableAPIs.forEach(
					tableAPI => {
						let table = frontRepo.map_ID_Table.get(tableAPI.ID)
						CopyTableAPIToTable(tableAPI, table!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.ThreejsAPIs.forEach(
					threejsAPI => {
						let threejs = frontRepo.map_ID_Threejs.get(threejsAPI.ID)
						CopyThreejsAPIToThreejs(threejsAPI, threejs!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.TitleAPIs.forEach(
					titleAPI => {
						let title = frontRepo.map_ID_Title.get(titleAPI.ID)
						CopyTitleAPIToTitle(titleAPI, title!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.ToneAPIs.forEach(
					toneAPI => {
						let tone = frontRepo.map_ID_Tone.get(toneAPI.ID)
						CopyToneAPIToTone(toneAPI, tone!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.TreeAPIs.forEach(
					treeAPI => {
						let tree = frontRepo.map_ID_Tree.get(treeAPI.ID)
						CopyTreeAPIToTree(treeAPI, tree!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.ViewAPIs.forEach(
					viewAPI => {
						let view = frontRepo.map_ID_View.get(viewAPI.ID)
						CopyViewAPIToView(viewAPI, view!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.XlsxAPIs.forEach(
					xlsxAPI => {
						let xlsx = frontRepo.map_ID_Xlsx.get(xlsxAPI.ID)
						CopyXlsxAPIToXlsx(xlsxAPI, xlsx!, frontRepo)
					}
				)


				this.ngZone.run(() => {
					observer.next(frontRepo)
				})
			}

			// 3. Connection Loop
			const attemptConnection = (retries: number): void => {
				// console.log("github.com/fullstack-lang/gong/lib/split/go; attemptConnection: retries =", retries, "isOfflineMode =", isOfflineMode)

				// A. WASM OFFLINE MODE (Check if Go is ready)
				if ((window as any).openWasmSocket) {
					// console.log("github.com/fullstack-lang/gong/lib/split/go; attemptConnection: openWasmSocket exists, calling it");
					(window as any).openWasmSocket("github.com/fullstack-lang/gong/lib/split/go", Name, processData);
					return;
				}

				// B. WAITING FOR WASM
				if (isOfflineMode && retries > 0) {
					// console.log("github.com/fullstack-lang/gong/lib/split/go; attemptConnection: WAITING FOR WASM. Retries left:", retries)
					setTimeout(() => attemptConnection(retries - 1), 100);
					return;
				}

				// C. STANDARD SERVER MODE
				if (!isOfflineMode) {
					// console.log("github.com/fullstack-lang/gong/lib/split/go; attemptConnection: STANDARD SERVER MODE. url =", url)
					socket = new WebSocket(url)
					socket.onopen = (event) => {
						// console.log("github.com/fullstack-lang/gong/lib/split/go; WebSocket: onopen", event)
					}
					socket.onmessage = event => {
						// console.log("github.com/fullstack-lang/gong/lib/split/go; WebSocket: onmessage")
						processData(event.data)
					}
					socket.onerror = event => {
						console.error("github.com/fullstack-lang/gong/lib/split/go WebSocket: onerror", event)
						observer.error(event)
					}
					socket.onclose = (event) => {
						// console.log("github.com/fullstack-lang/gong/lib/split/go; WebSocket: onclose", event)
						observer.complete()
					}
				} else {
					console.error("github.com/fullstack-lang/gong/lib/split/go, attemptConnection: Offline mode detected, but WASM backend failed to load.")
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
export function getAsSplitUniqueID(id: number): number {
	return 31 * id
}
export function getAsSplitAreaUniqueID(id: number): number {
	return 37 * id
}
export function getButtonUniqueID(id: number): number {
	return 41 * id
}
export function getCursorUniqueID(id: number): number {
	return 43 * id
}
export function getFavIconUniqueID(id: number): number {
	return 47 * id
}
export function getFormUniqueID(id: number): number {
	return 53 * id
}
export function getLoadUniqueID(id: number): number {
	return 59 * id
}
export function getLogoOnTheLeftUniqueID(id: number): number {
	return 61 * id
}
export function getLogoOnTheRightUniqueID(id: number): number {
	return 67 * id
}
export function getMarkdownUniqueID(id: number): number {
	return 71 * id
}
export function getSliderUniqueID(id: number): number {
	return 73 * id
}
export function getSplitUniqueID(id: number): number {
	return 79 * id
}
export function getSvgUniqueID(id: number): number {
	return 83 * id
}
export function getTableUniqueID(id: number): number {
	return 89 * id
}
export function getThreejsUniqueID(id: number): number {
	return 97 * id
}
export function getTitleUniqueID(id: number): number {
	return 101 * id
}
export function getToneUniqueID(id: number): number {
	return 103 * id
}
export function getTreeUniqueID(id: number): number {
	return 107 * id
}
export function getViewUniqueID(id: number): number {
	return 109 * id
}
export function getXlsxUniqueID(id: number): number {
	return 113 * id
}
