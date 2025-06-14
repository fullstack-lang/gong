import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'
import { shareReplay } from 'rxjs/operators'

// insertion point sub template for services imports
import { AsSplitAPI } from './assplit-api'
import { AsSplit, CopyAsSplitAPIToAsSplit } from './assplit'
import { AsSplitService } from './assplit.service'

import { AsSplitAreaAPI } from './assplitarea-api'
import { AsSplitArea, CopyAsSplitAreaAPIToAsSplitArea } from './assplitarea'
import { AsSplitAreaService } from './assplitarea.service'

import { ButtonAPI } from './button-api'
import { Button, CopyButtonAPIToButton } from './button'
import { ButtonService } from './button.service'

import { CursorAPI } from './cursor-api'
import { Cursor, CopyCursorAPIToCursor } from './cursor'
import { CursorService } from './cursor.service'

import { DocAPI } from './doc-api'
import { Doc, CopyDocAPIToDoc } from './doc'
import { DocService } from './doc.service'

import { FavIconAPI } from './favicon-api'
import { FavIcon, CopyFavIconAPIToFavIcon } from './favicon'
import { FavIconService } from './favicon.service'

import { FormAPI } from './form-api'
import { Form, CopyFormAPIToForm } from './form'
import { FormService } from './form.service'

import { LoadAPI } from './load-api'
import { Load, CopyLoadAPIToLoad } from './load'
import { LoadService } from './load.service'

import { LogoOnTheLeftAPI } from './logoontheleft-api'
import { LogoOnTheLeft, CopyLogoOnTheLeftAPIToLogoOnTheLeft } from './logoontheleft'
import { LogoOnTheLeftService } from './logoontheleft.service'

import { LogoOnTheRightAPI } from './logoontheright-api'
import { LogoOnTheRight, CopyLogoOnTheRightAPIToLogoOnTheRight } from './logoontheright'
import { LogoOnTheRightService } from './logoontheright.service'

import { SliderAPI } from './slider-api'
import { Slider, CopySliderAPIToSlider } from './slider'
import { SliderService } from './slider.service'

import { SplitAPI } from './split-api'
import { Split, CopySplitAPIToSplit } from './split'
import { SplitService } from './split.service'

import { SvgAPI } from './svg-api'
import { Svg, CopySvgAPIToSvg } from './svg'
import { SvgService } from './svg.service'

import { TableAPI } from './table-api'
import { Table, CopyTableAPIToTable } from './table'
import { TableService } from './table.service'

import { TitleAPI } from './title-api'
import { Title, CopyTitleAPIToTitle } from './title'
import { TitleService } from './title.service'

import { ToneAPI } from './tone-api'
import { Tone, CopyToneAPIToTone } from './tone'
import { ToneService } from './tone.service'

import { TreeAPI } from './tree-api'
import { Tree, CopyTreeAPIToTree } from './tree'
import { TreeService } from './tree.service'

import { ViewAPI } from './view-api'
import { View, CopyViewAPIToView } from './view'
import { ViewService } from './view.service'

import { XlsxAPI } from './xlsx-api'
import { Xlsx, CopyXlsxAPIToXlsx } from './xlsx'
import { XlsxService } from './xlsx.service'


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

	array_Docs = new Array<Doc>() // array of front instances
	map_ID_Doc = new Map<number, Doc>() // map of front instances

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

	array_Sliders = new Array<Slider>() // array of front instances
	map_ID_Slider = new Map<number, Slider>() // map of front instances

	array_Splits = new Array<Split>() // array of front instances
	map_ID_Split = new Map<number, Split>() // map of front instances

	array_Svgs = new Array<Svg>() // array of front instances
	map_ID_Svg = new Map<number, Svg>() // map of front instances

	array_Tables = new Array<Table>() // array of front instances
	map_ID_Table = new Map<number, Table>() // map of front instances

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
			case 'Doc':
				return this.array_Docs as unknown as Array<Type>
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
			case 'Slider':
				return this.array_Sliders as unknown as Array<Type>
			case 'Split':
				return this.array_Splits as unknown as Array<Type>
			case 'Svg':
				return this.array_Svgs as unknown as Array<Type>
			case 'Table':
				return this.array_Tables as unknown as Array<Type>
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
			case 'Doc':
				return this.map_ID_Doc as unknown as Map<number, Type>
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
			case 'Slider':
				return this.map_ID_Slider as unknown as Map<number, Type>
			case 'Split':
				return this.map_ID_Split as unknown as Map<number, Type>
			case 'Svg':
				return this.map_ID_Svg as unknown as Map<number, Type>
			case 'Table':
				return this.map_ID_Table as unknown as Map<number, Type>
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
		private http: HttpClient, // insertion point sub template 
		private assplitService: AsSplitService,
		private assplitareaService: AsSplitAreaService,
		private buttonService: ButtonService,
		private cursorService: CursorService,
		private docService: DocService,
		private faviconService: FavIconService,
		private formService: FormService,
		private loadService: LoadService,
		private logoontheleftService: LogoOnTheLeftService,
		private logoontherightService: LogoOnTheRightService,
		private sliderService: SliderService,
		private splitService: SplitService,
		private svgService: SvgService,
		private tableService: TableService,
		private titleService: TitleService,
		private toneService: ToneService,
		private treeService: TreeService,
		private viewService: ViewService,
		private xlsxService: XlsxService,
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
		Observable<AsSplitAPI[]>,
		Observable<AsSplitAreaAPI[]>,
		Observable<ButtonAPI[]>,
		Observable<CursorAPI[]>,
		Observable<DocAPI[]>,
		Observable<FavIconAPI[]>,
		Observable<FormAPI[]>,
		Observable<LoadAPI[]>,
		Observable<LogoOnTheLeftAPI[]>,
		Observable<LogoOnTheRightAPI[]>,
		Observable<SliderAPI[]>,
		Observable<SplitAPI[]>,
		Observable<SvgAPI[]>,
		Observable<TableAPI[]>,
		Observable<TitleAPI[]>,
		Observable<ToneAPI[]>,
		Observable<TreeAPI[]>,
		Observable<ViewAPI[]>,
		Observable<XlsxAPI[]>,
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
			this.assplitService.getAsSplits(this.Name, this.frontRepo),
			this.assplitareaService.getAsSplitAreas(this.Name, this.frontRepo),
			this.buttonService.getButtons(this.Name, this.frontRepo),
			this.cursorService.getCursors(this.Name, this.frontRepo),
			this.docService.getDocs(this.Name, this.frontRepo),
			this.faviconService.getFavIcons(this.Name, this.frontRepo),
			this.formService.getForms(this.Name, this.frontRepo),
			this.loadService.getLoads(this.Name, this.frontRepo),
			this.logoontheleftService.getLogoOnTheLefts(this.Name, this.frontRepo),
			this.logoontherightService.getLogoOnTheRights(this.Name, this.frontRepo),
			this.sliderService.getSliders(this.Name, this.frontRepo),
			this.splitService.getSplits(this.Name, this.frontRepo),
			this.svgService.getSvgs(this.Name, this.frontRepo),
			this.tableService.getTables(this.Name, this.frontRepo),
			this.titleService.getTitles(this.Name, this.frontRepo),
			this.toneService.getTones(this.Name, this.frontRepo),
			this.treeService.getTrees(this.Name, this.frontRepo),
			this.viewService.getViews(this.Name, this.frontRepo),
			this.xlsxService.getXlsxs(this.Name, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						assplits_,
						assplitareas_,
						buttons_,
						cursors_,
						docs_,
						favicons_,
						forms_,
						loads_,
						logoonthelefts_,
						logoontherights_,
						sliders_,
						splits_,
						svgs_,
						tables_,
						titles_,
						tones_,
						trees_,
						views_,
						xlsxs_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var assplits: AsSplitAPI[]
						assplits = assplits_ as AsSplitAPI[]
						var assplitareas: AsSplitAreaAPI[]
						assplitareas = assplitareas_ as AsSplitAreaAPI[]
						var buttons: ButtonAPI[]
						buttons = buttons_ as ButtonAPI[]
						var cursors: CursorAPI[]
						cursors = cursors_ as CursorAPI[]
						var docs: DocAPI[]
						docs = docs_ as DocAPI[]
						var favicons: FavIconAPI[]
						favicons = favicons_ as FavIconAPI[]
						var forms: FormAPI[]
						forms = forms_ as FormAPI[]
						var loads: LoadAPI[]
						loads = loads_ as LoadAPI[]
						var logoonthelefts: LogoOnTheLeftAPI[]
						logoonthelefts = logoonthelefts_ as LogoOnTheLeftAPI[]
						var logoontherights: LogoOnTheRightAPI[]
						logoontherights = logoontherights_ as LogoOnTheRightAPI[]
						var sliders: SliderAPI[]
						sliders = sliders_ as SliderAPI[]
						var splits: SplitAPI[]
						splits = splits_ as SplitAPI[]
						var svgs: SvgAPI[]
						svgs = svgs_ as SvgAPI[]
						var tables: TableAPI[]
						tables = tables_ as TableAPI[]
						var titles: TitleAPI[]
						titles = titles_ as TitleAPI[]
						var tones: ToneAPI[]
						tones = tones_ as ToneAPI[]
						var trees: TreeAPI[]
						trees = trees_ as TreeAPI[]
						var views: ViewAPI[]
						views = views_ as ViewAPI[]
						var xlsxs: XlsxAPI[]
						xlsxs = xlsxs_ as XlsxAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_AsSplits = []
						this.frontRepo.map_ID_AsSplit.clear()

						assplits.forEach(
							assplitAPI => {
								let assplit = new AsSplit
								this.frontRepo.array_AsSplits.push(assplit)
								this.frontRepo.map_ID_AsSplit.set(assplitAPI.ID, assplit)
							}
						)

						// init the arrays
						this.frontRepo.array_AsSplitAreas = []
						this.frontRepo.map_ID_AsSplitArea.clear()

						assplitareas.forEach(
							assplitareaAPI => {
								let assplitarea = new AsSplitArea
								this.frontRepo.array_AsSplitAreas.push(assplitarea)
								this.frontRepo.map_ID_AsSplitArea.set(assplitareaAPI.ID, assplitarea)
							}
						)

						// init the arrays
						this.frontRepo.array_Buttons = []
						this.frontRepo.map_ID_Button.clear()

						buttons.forEach(
							buttonAPI => {
								let button = new Button
								this.frontRepo.array_Buttons.push(button)
								this.frontRepo.map_ID_Button.set(buttonAPI.ID, button)
							}
						)

						// init the arrays
						this.frontRepo.array_Cursors = []
						this.frontRepo.map_ID_Cursor.clear()

						cursors.forEach(
							cursorAPI => {
								let cursor = new Cursor
								this.frontRepo.array_Cursors.push(cursor)
								this.frontRepo.map_ID_Cursor.set(cursorAPI.ID, cursor)
							}
						)

						// init the arrays
						this.frontRepo.array_Docs = []
						this.frontRepo.map_ID_Doc.clear()

						docs.forEach(
							docAPI => {
								let doc = new Doc
								this.frontRepo.array_Docs.push(doc)
								this.frontRepo.map_ID_Doc.set(docAPI.ID, doc)
							}
						)

						// init the arrays
						this.frontRepo.array_FavIcons = []
						this.frontRepo.map_ID_FavIcon.clear()

						favicons.forEach(
							faviconAPI => {
								let favicon = new FavIcon
								this.frontRepo.array_FavIcons.push(favicon)
								this.frontRepo.map_ID_FavIcon.set(faviconAPI.ID, favicon)
							}
						)

						// init the arrays
						this.frontRepo.array_Forms = []
						this.frontRepo.map_ID_Form.clear()

						forms.forEach(
							formAPI => {
								let form = new Form
								this.frontRepo.array_Forms.push(form)
								this.frontRepo.map_ID_Form.set(formAPI.ID, form)
							}
						)

						// init the arrays
						this.frontRepo.array_Loads = []
						this.frontRepo.map_ID_Load.clear()

						loads.forEach(
							loadAPI => {
								let load = new Load
								this.frontRepo.array_Loads.push(load)
								this.frontRepo.map_ID_Load.set(loadAPI.ID, load)
							}
						)

						// init the arrays
						this.frontRepo.array_LogoOnTheLefts = []
						this.frontRepo.map_ID_LogoOnTheLeft.clear()

						logoonthelefts.forEach(
							logoontheleftAPI => {
								let logoontheleft = new LogoOnTheLeft
								this.frontRepo.array_LogoOnTheLefts.push(logoontheleft)
								this.frontRepo.map_ID_LogoOnTheLeft.set(logoontheleftAPI.ID, logoontheleft)
							}
						)

						// init the arrays
						this.frontRepo.array_LogoOnTheRights = []
						this.frontRepo.map_ID_LogoOnTheRight.clear()

						logoontherights.forEach(
							logoontherightAPI => {
								let logoontheright = new LogoOnTheRight
								this.frontRepo.array_LogoOnTheRights.push(logoontheright)
								this.frontRepo.map_ID_LogoOnTheRight.set(logoontherightAPI.ID, logoontheright)
							}
						)

						// init the arrays
						this.frontRepo.array_Sliders = []
						this.frontRepo.map_ID_Slider.clear()

						sliders.forEach(
							sliderAPI => {
								let slider = new Slider
								this.frontRepo.array_Sliders.push(slider)
								this.frontRepo.map_ID_Slider.set(sliderAPI.ID, slider)
							}
						)

						// init the arrays
						this.frontRepo.array_Splits = []
						this.frontRepo.map_ID_Split.clear()

						splits.forEach(
							splitAPI => {
								let split = new Split
								this.frontRepo.array_Splits.push(split)
								this.frontRepo.map_ID_Split.set(splitAPI.ID, split)
							}
						)

						// init the arrays
						this.frontRepo.array_Svgs = []
						this.frontRepo.map_ID_Svg.clear()

						svgs.forEach(
							svgAPI => {
								let svg = new Svg
								this.frontRepo.array_Svgs.push(svg)
								this.frontRepo.map_ID_Svg.set(svgAPI.ID, svg)
							}
						)

						// init the arrays
						this.frontRepo.array_Tables = []
						this.frontRepo.map_ID_Table.clear()

						tables.forEach(
							tableAPI => {
								let table = new Table
								this.frontRepo.array_Tables.push(table)
								this.frontRepo.map_ID_Table.set(tableAPI.ID, table)
							}
						)

						// init the arrays
						this.frontRepo.array_Titles = []
						this.frontRepo.map_ID_Title.clear()

						titles.forEach(
							titleAPI => {
								let title = new Title
								this.frontRepo.array_Titles.push(title)
								this.frontRepo.map_ID_Title.set(titleAPI.ID, title)
							}
						)

						// init the arrays
						this.frontRepo.array_Tones = []
						this.frontRepo.map_ID_Tone.clear()

						tones.forEach(
							toneAPI => {
								let tone = new Tone
								this.frontRepo.array_Tones.push(tone)
								this.frontRepo.map_ID_Tone.set(toneAPI.ID, tone)
							}
						)

						// init the arrays
						this.frontRepo.array_Trees = []
						this.frontRepo.map_ID_Tree.clear()

						trees.forEach(
							treeAPI => {
								let tree = new Tree
								this.frontRepo.array_Trees.push(tree)
								this.frontRepo.map_ID_Tree.set(treeAPI.ID, tree)
							}
						)

						// init the arrays
						this.frontRepo.array_Views = []
						this.frontRepo.map_ID_View.clear()

						views.forEach(
							viewAPI => {
								let view = new View
								this.frontRepo.array_Views.push(view)
								this.frontRepo.map_ID_View.set(viewAPI.ID, view)
							}
						)

						// init the arrays
						this.frontRepo.array_Xlsxs = []
						this.frontRepo.map_ID_Xlsx.clear()

						xlsxs.forEach(
							xlsxAPI => {
								let xlsx = new Xlsx
								this.frontRepo.array_Xlsxs.push(xlsx)
								this.frontRepo.map_ID_Xlsx.set(xlsxAPI.ID, xlsx)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						assplits.forEach(
							assplitAPI => {
								let assplit = this.frontRepo.map_ID_AsSplit.get(assplitAPI.ID)
								CopyAsSplitAPIToAsSplit(assplitAPI, assplit!, this.frontRepo)
							}
						)

						// fill up front objects
						assplitareas.forEach(
							assplitareaAPI => {
								let assplitarea = this.frontRepo.map_ID_AsSplitArea.get(assplitareaAPI.ID)
								CopyAsSplitAreaAPIToAsSplitArea(assplitareaAPI, assplitarea!, this.frontRepo)
							}
						)

						// fill up front objects
						buttons.forEach(
							buttonAPI => {
								let button = this.frontRepo.map_ID_Button.get(buttonAPI.ID)
								CopyButtonAPIToButton(buttonAPI, button!, this.frontRepo)
							}
						)

						// fill up front objects
						cursors.forEach(
							cursorAPI => {
								let cursor = this.frontRepo.map_ID_Cursor.get(cursorAPI.ID)
								CopyCursorAPIToCursor(cursorAPI, cursor!, this.frontRepo)
							}
						)

						// fill up front objects
						docs.forEach(
							docAPI => {
								let doc = this.frontRepo.map_ID_Doc.get(docAPI.ID)
								CopyDocAPIToDoc(docAPI, doc!, this.frontRepo)
							}
						)

						// fill up front objects
						favicons.forEach(
							faviconAPI => {
								let favicon = this.frontRepo.map_ID_FavIcon.get(faviconAPI.ID)
								CopyFavIconAPIToFavIcon(faviconAPI, favicon!, this.frontRepo)
							}
						)

						// fill up front objects
						forms.forEach(
							formAPI => {
								let form = this.frontRepo.map_ID_Form.get(formAPI.ID)
								CopyFormAPIToForm(formAPI, form!, this.frontRepo)
							}
						)

						// fill up front objects
						loads.forEach(
							loadAPI => {
								let load = this.frontRepo.map_ID_Load.get(loadAPI.ID)
								CopyLoadAPIToLoad(loadAPI, load!, this.frontRepo)
							}
						)

						// fill up front objects
						logoonthelefts.forEach(
							logoontheleftAPI => {
								let logoontheleft = this.frontRepo.map_ID_LogoOnTheLeft.get(logoontheleftAPI.ID)
								CopyLogoOnTheLeftAPIToLogoOnTheLeft(logoontheleftAPI, logoontheleft!, this.frontRepo)
							}
						)

						// fill up front objects
						logoontherights.forEach(
							logoontherightAPI => {
								let logoontheright = this.frontRepo.map_ID_LogoOnTheRight.get(logoontherightAPI.ID)
								CopyLogoOnTheRightAPIToLogoOnTheRight(logoontherightAPI, logoontheright!, this.frontRepo)
							}
						)

						// fill up front objects
						sliders.forEach(
							sliderAPI => {
								let slider = this.frontRepo.map_ID_Slider.get(sliderAPI.ID)
								CopySliderAPIToSlider(sliderAPI, slider!, this.frontRepo)
							}
						)

						// fill up front objects
						splits.forEach(
							splitAPI => {
								let split = this.frontRepo.map_ID_Split.get(splitAPI.ID)
								CopySplitAPIToSplit(splitAPI, split!, this.frontRepo)
							}
						)

						// fill up front objects
						svgs.forEach(
							svgAPI => {
								let svg = this.frontRepo.map_ID_Svg.get(svgAPI.ID)
								CopySvgAPIToSvg(svgAPI, svg!, this.frontRepo)
							}
						)

						// fill up front objects
						tables.forEach(
							tableAPI => {
								let table = this.frontRepo.map_ID_Table.get(tableAPI.ID)
								CopyTableAPIToTable(tableAPI, table!, this.frontRepo)
							}
						)

						// fill up front objects
						titles.forEach(
							titleAPI => {
								let title = this.frontRepo.map_ID_Title.get(titleAPI.ID)
								CopyTitleAPIToTitle(titleAPI, title!, this.frontRepo)
							}
						)

						// fill up front objects
						tones.forEach(
							toneAPI => {
								let tone = this.frontRepo.map_ID_Tone.get(toneAPI.ID)
								CopyToneAPIToTone(toneAPI, tone!, this.frontRepo)
							}
						)

						// fill up front objects
						trees.forEach(
							treeAPI => {
								let tree = this.frontRepo.map_ID_Tree.get(treeAPI.ID)
								CopyTreeAPIToTree(treeAPI, tree!, this.frontRepo)
							}
						)

						// fill up front objects
						views.forEach(
							viewAPI => {
								let view = this.frontRepo.map_ID_View.get(viewAPI.ID)
								CopyViewAPIToView(viewAPI, view!, this.frontRepo)
							}
						)

						// fill up front objects
						xlsxs.forEach(
							xlsxAPI => {
								let xlsx = this.frontRepo.map_ID_Xlsx.get(xlsxAPI.ID)
								CopyXlsxAPIToXlsx(xlsxAPI, xlsx!, this.frontRepo)
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
		let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/lib/split/go/v1/ws/stage`

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
				frontRepo.array_Docs = []
				frontRepo.map_ID_Doc.clear()

				backRepoData.DocAPIs.forEach(
					docAPI => {
						let doc = new Doc
						frontRepo.array_Docs.push(doc)
						frontRepo.map_ID_Doc.set(docAPI.ID, doc)
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
				backRepoData.DocAPIs.forEach(
					docAPI => {
						let doc = frontRepo.map_ID_Doc.get(docAPI.ID)
						CopyDocAPIToDoc(docAPI, doc!, frontRepo)
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
export function getDocUniqueID(id: number): number {
	return 47 * id
}
export function getFavIconUniqueID(id: number): number {
	return 53 * id
}
export function getFormUniqueID(id: number): number {
	return 59 * id
}
export function getLoadUniqueID(id: number): number {
	return 61 * id
}
export function getLogoOnTheLeftUniqueID(id: number): number {
	return 67 * id
}
export function getLogoOnTheRightUniqueID(id: number): number {
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
export function getTitleUniqueID(id: number): number {
	return 97 * id
}
export function getToneUniqueID(id: number): number {
	return 101 * id
}
export function getTreeUniqueID(id: number): number {
	return 103 * id
}
export function getViewUniqueID(id: number): number {
	return 107 * id
}
export function getXlsxUniqueID(id: number): number {
	return 109 * id
}
