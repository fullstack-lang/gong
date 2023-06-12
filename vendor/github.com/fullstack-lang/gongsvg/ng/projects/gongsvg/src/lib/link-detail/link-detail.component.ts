// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { LinkDB } from '../link-db'
import { LinkService } from '../link.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { LinkTypeSelect, LinkTypeList } from '../LinkType'
import { AnchorTypeSelect, AnchorTypeList } from '../AnchorType'
import { OrientationTypeSelect, OrientationTypeList } from '../OrientationType'
import { LayerDB } from '../layer-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// LinkDetailComponent is initilizaed from different routes
// LinkDetailComponentState detail different cases 
enum LinkDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Layer_Links_SET,
}

@Component({
	selector: 'app-link-detail',
	templateUrl: './link-detail.component.html',
	styleUrls: ['./link-detail.component.css'],
})
export class LinkDetailComponent implements OnInit {

	// insertion point for declarations
	LinkTypeList: LinkTypeSelect[] = []
	AnchorTypeList: AnchorTypeSelect[] = []
	OrientationTypeList: OrientationTypeSelect[] = []
	HasEndArrowFormControl: UntypedFormControl = new UntypedFormControl(false);

	// the LinkDB of interest
	link: LinkDB = new LinkDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: LinkDetailComponentState = LinkDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private linkService: LinkService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private activatedRoute: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.GONG__StackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')!;

		this.activatedRoute.params.subscribe(params => {
			this.onChangedActivatedRoute()
		});
	}
	onChangedActivatedRoute(): void {

		// compute state
		this.id = +this.activatedRoute.snapshot.paramMap.get('id')!;
		this.originStruct = this.activatedRoute.snapshot.paramMap.get('originStruct')!;
		this.originStructFieldName = this.activatedRoute.snapshot.paramMap.get('originStructFieldName')!;

		this.GONG__StackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')!;

		const association = this.activatedRoute.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = LinkDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = LinkDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "Links":
						// console.log("Link" + " is instanciated with back pointer to instance " + this.id + " Layer association Links")
						this.state = LinkDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Layer_Links_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getLink()

		// observable for changes in structs
		this.linkService.LinkServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getLink()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.LinkTypeList = LinkTypeList
		this.AnchorTypeList = AnchorTypeList
		this.OrientationTypeList = OrientationTypeList
	}

	getLink(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case LinkDetailComponentState.CREATE_INSTANCE:
						this.link = new (LinkDB)
						break;
					case LinkDetailComponentState.UPDATE_INSTANCE:
						let link = frontRepo.Links.get(this.id)
						console.assert(link != undefined, "missing link with id:" + this.id)
						this.link = link!
						break;
					// insertion point for init of association field
					case LinkDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Layer_Links_SET:
						this.link = new (LinkDB)
						this.link.Layer_Links_reverse = frontRepo.Layers.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
				this.HasEndArrowFormControl.setValue(this.link.HasEndArrow)
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		if (this.link.StartID == undefined) {
			this.link.StartID = new NullInt64
		}
		if (this.link.Start != undefined) {
			this.link.StartID.Int64 = this.link.Start.ID
			this.link.StartID.Valid = true
		} else {
			this.link.StartID.Int64 = 0
			this.link.StartID.Valid = true
		}
		if (this.link.EndID == undefined) {
			this.link.EndID = new NullInt64
		}
		if (this.link.End != undefined) {
			this.link.EndID.Int64 = this.link.End.ID
			this.link.EndID.Valid = true
		} else {
			this.link.EndID.Int64 = 0
			this.link.EndID.Valid = true
		}
		this.link.HasEndArrow = this.HasEndArrowFormControl.value

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.link.Layer_Links_reverse != undefined) {
			if (this.link.Layer_LinksDBID == undefined) {
				this.link.Layer_LinksDBID = new NullInt64
			}
			this.link.Layer_LinksDBID.Int64 = this.link.Layer_Links_reverse.ID
			this.link.Layer_LinksDBID.Valid = true
			if (this.link.Layer_LinksDBID_Index == undefined) {
				this.link.Layer_LinksDBID_Index = new NullInt64
			}
			this.link.Layer_LinksDBID_Index.Valid = true
			this.link.Layer_Links_reverse = new LayerDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case LinkDetailComponentState.UPDATE_INSTANCE:
				this.linkService.updateLink(this.link, this.GONG__StackPath)
					.subscribe(link => {
						this.linkService.LinkServiceChanged.next("update")
					});
				break;
			default:
				this.linkService.postLink(this.link, this.GONG__StackPath).subscribe(link => {
					this.linkService.LinkServiceChanged.next("post")
					this.link = new (LinkDB) // reset fields
				});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string, selectionMode: string,
		sourceField: string, intermediateStructField: string, nextAssociatedStruct: string) {

		console.log("mode " + selectionMode)

		const dialogConfig = new MatDialogConfig();

		let dialogData = new DialogData();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		if (selectionMode == SelectionMode.ONE_MANY_ASSOCIATION_MODE) {

			dialogData.ID = this.link.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(AssociatedStruct).get(
					AssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}
		if (selectionMode == SelectionMode.MANY_MANY_ASSOCIATION_MODE) {
			dialogData.ID = this.link.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "Link"
			dialogData.SourceField = sourceField

			// set up the intermediate struct
			dialogData.IntermediateStruct = AssociatedStruct
			dialogData.IntermediateStructField = intermediateStructField

			// set up the end struct
			dialogData.NextAssociationStruct = nextAssociatedStruct

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(nextAssociatedStruct).get(
					nextAssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}

	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.link.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
			GONG__StackPath: this.GONG__StackPath,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfSortingComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'SortingComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
	}

	fillUpNameIfEmpty(event: { value: { Name: string; }; }) {
		if (this.link.Name == "") {
			this.link.Name = event.value.Name
		}
	}

	toggleTextArea(fieldName: string) {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			let displayAsTextArea = this.mapFields_displayAsTextArea.get(fieldName)
			this.mapFields_displayAsTextArea.set(fieldName, !displayAsTextArea)
		} else {
			this.mapFields_displayAsTextArea.set(fieldName, true)
		}
	}

	isATextArea(fieldName: string): boolean {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			return this.mapFields_displayAsTextArea.get(fieldName)!
		} else {
			return false
		}
	}

	compareObjects(o1: any, o2: any) {
		if (o1?.ID == o2?.ID) {
			return true;
		}
		else {
			return false
		}
	}
}
