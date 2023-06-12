// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { RectLinkLinkDB } from '../rectlinklink-db'
import { RectLinkLinkService } from '../rectlinklink.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { LayerDB } from '../layer-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// RectLinkLinkDetailComponent is initilizaed from different routes
// RectLinkLinkDetailComponentState detail different cases 
enum RectLinkLinkDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Layer_RectLinkLinks_SET,
}

@Component({
	selector: 'app-rectlinklink-detail',
	templateUrl: './rectlinklink-detail.component.html',
	styleUrls: ['./rectlinklink-detail.component.css'],
})
export class RectLinkLinkDetailComponent implements OnInit {

	// insertion point for declarations

	// the RectLinkLinkDB of interest
	rectlinklink: RectLinkLinkDB = new RectLinkLinkDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: RectLinkLinkDetailComponentState = RectLinkLinkDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private rectlinklinkService: RectLinkLinkService,
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
			this.state = RectLinkLinkDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = RectLinkLinkDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "RectLinkLinks":
						// console.log("RectLinkLink" + " is instanciated with back pointer to instance " + this.id + " Layer association RectLinkLinks")
						this.state = RectLinkLinkDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Layer_RectLinkLinks_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getRectLinkLink()

		// observable for changes in structs
		this.rectlinklinkService.RectLinkLinkServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getRectLinkLink()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getRectLinkLink(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case RectLinkLinkDetailComponentState.CREATE_INSTANCE:
						this.rectlinklink = new (RectLinkLinkDB)
						break;
					case RectLinkLinkDetailComponentState.UPDATE_INSTANCE:
						let rectlinklink = frontRepo.RectLinkLinks.get(this.id)
						console.assert(rectlinklink != undefined, "missing rectlinklink with id:" + this.id)
						this.rectlinklink = rectlinklink!
						break;
					// insertion point for init of association field
					case RectLinkLinkDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Layer_RectLinkLinks_SET:
						this.rectlinklink = new (RectLinkLinkDB)
						this.rectlinklink.Layer_RectLinkLinks_reverse = frontRepo.Layers.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		if (this.rectlinklink.StartID == undefined) {
			this.rectlinklink.StartID = new NullInt64
		}
		if (this.rectlinklink.Start != undefined) {
			this.rectlinklink.StartID.Int64 = this.rectlinklink.Start.ID
			this.rectlinklink.StartID.Valid = true
		} else {
			this.rectlinklink.StartID.Int64 = 0
			this.rectlinklink.StartID.Valid = true
		}
		if (this.rectlinklink.EndID == undefined) {
			this.rectlinklink.EndID = new NullInt64
		}
		if (this.rectlinklink.End != undefined) {
			this.rectlinklink.EndID.Int64 = this.rectlinklink.End.ID
			this.rectlinklink.EndID.Valid = true
		} else {
			this.rectlinklink.EndID.Int64 = 0
			this.rectlinklink.EndID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.rectlinklink.Layer_RectLinkLinks_reverse != undefined) {
			if (this.rectlinklink.Layer_RectLinkLinksDBID == undefined) {
				this.rectlinklink.Layer_RectLinkLinksDBID = new NullInt64
			}
			this.rectlinklink.Layer_RectLinkLinksDBID.Int64 = this.rectlinklink.Layer_RectLinkLinks_reverse.ID
			this.rectlinklink.Layer_RectLinkLinksDBID.Valid = true
			if (this.rectlinklink.Layer_RectLinkLinksDBID_Index == undefined) {
				this.rectlinklink.Layer_RectLinkLinksDBID_Index = new NullInt64
			}
			this.rectlinklink.Layer_RectLinkLinksDBID_Index.Valid = true
			this.rectlinklink.Layer_RectLinkLinks_reverse = new LayerDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case RectLinkLinkDetailComponentState.UPDATE_INSTANCE:
				this.rectlinklinkService.updateRectLinkLink(this.rectlinklink, this.GONG__StackPath)
					.subscribe(rectlinklink => {
						this.rectlinklinkService.RectLinkLinkServiceChanged.next("update")
					});
				break;
			default:
				this.rectlinklinkService.postRectLinkLink(this.rectlinklink, this.GONG__StackPath).subscribe(rectlinklink => {
					this.rectlinklinkService.RectLinkLinkServiceChanged.next("post")
					this.rectlinklink = new (RectLinkLinkDB) // reset fields
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

			dialogData.ID = this.rectlinklink.ID!
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
			dialogData.ID = this.rectlinklink.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "RectLinkLink"
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
			ID: this.rectlinklink.ID,
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
		if (this.rectlinklink.Name == "") {
			this.rectlinklink.Name = event.value.Name
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
