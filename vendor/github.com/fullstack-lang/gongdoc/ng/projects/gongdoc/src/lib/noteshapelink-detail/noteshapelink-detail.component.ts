// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { NoteShapeLinkDB } from '../noteshapelink-db'
import { NoteShapeLinkService } from '../noteshapelink.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { NoteShapeLinkTypeSelect, NoteShapeLinkTypeList } from '../NoteShapeLinkType'
import { NoteShapeDB } from '../noteshape-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// NoteShapeLinkDetailComponent is initilizaed from different routes
// NoteShapeLinkDetailComponentState detail different cases 
enum NoteShapeLinkDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_NoteShape_NoteShapeLinks_SET,
}

@Component({
	selector: 'app-noteshapelink-detail',
	templateUrl: './noteshapelink-detail.component.html',
	styleUrls: ['./noteshapelink-detail.component.css'],
})
export class NoteShapeLinkDetailComponent implements OnInit {

	// insertion point for declarations
	NoteShapeLinkTypeList: NoteShapeLinkTypeSelect[] = []

	// the NoteShapeLinkDB of interest
	noteshapelink: NoteShapeLinkDB = new NoteShapeLinkDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: NoteShapeLinkDetailComponentState = NoteShapeLinkDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private noteshapelinkService: NoteShapeLinkService,
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
			this.state = NoteShapeLinkDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = NoteShapeLinkDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "NoteShapeLinks":
						// console.log("NoteShapeLink" + " is instanciated with back pointer to instance " + this.id + " NoteShape association NoteShapeLinks")
						this.state = NoteShapeLinkDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_NoteShape_NoteShapeLinks_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getNoteShapeLink()

		// observable for changes in structs
		this.noteshapelinkService.NoteShapeLinkServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getNoteShapeLink()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.NoteShapeLinkTypeList = NoteShapeLinkTypeList
	}

	getNoteShapeLink(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case NoteShapeLinkDetailComponentState.CREATE_INSTANCE:
						this.noteshapelink = new (NoteShapeLinkDB)
						break;
					case NoteShapeLinkDetailComponentState.UPDATE_INSTANCE:
						let noteshapelink = frontRepo.NoteShapeLinks.get(this.id)
						console.assert(noteshapelink != undefined, "missing noteshapelink with id:" + this.id)
						this.noteshapelink = noteshapelink!
						break;
					// insertion point for init of association field
					case NoteShapeLinkDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_NoteShape_NoteShapeLinks_SET:
						this.noteshapelink = new (NoteShapeLinkDB)
						this.noteshapelink.NoteShape_NoteShapeLinks_reverse = frontRepo.NoteShapes.get(this.id)!
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

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.noteshapelink.NoteShape_NoteShapeLinks_reverse != undefined) {
			if (this.noteshapelink.NoteShape_NoteShapeLinksDBID == undefined) {
				this.noteshapelink.NoteShape_NoteShapeLinksDBID = new NullInt64
			}
			this.noteshapelink.NoteShape_NoteShapeLinksDBID.Int64 = this.noteshapelink.NoteShape_NoteShapeLinks_reverse.ID
			this.noteshapelink.NoteShape_NoteShapeLinksDBID.Valid = true
			if (this.noteshapelink.NoteShape_NoteShapeLinksDBID_Index == undefined) {
				this.noteshapelink.NoteShape_NoteShapeLinksDBID_Index = new NullInt64
			}
			this.noteshapelink.NoteShape_NoteShapeLinksDBID_Index.Valid = true
			this.noteshapelink.NoteShape_NoteShapeLinks_reverse = new NoteShapeDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case NoteShapeLinkDetailComponentState.UPDATE_INSTANCE:
				this.noteshapelinkService.updateNoteShapeLink(this.noteshapelink, this.GONG__StackPath)
					.subscribe(noteshapelink => {
						this.noteshapelinkService.NoteShapeLinkServiceChanged.next("update")
					});
				break;
			default:
				this.noteshapelinkService.postNoteShapeLink(this.noteshapelink, this.GONG__StackPath).subscribe(noteshapelink => {
					this.noteshapelinkService.NoteShapeLinkServiceChanged.next("post")
					this.noteshapelink = new (NoteShapeLinkDB) // reset fields
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

			dialogData.ID = this.noteshapelink.ID!
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
			dialogData.ID = this.noteshapelink.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "NoteShapeLink"
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
			ID: this.noteshapelink.ID,
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
		if (this.noteshapelink.Name == "") {
			this.noteshapelink.Name = event.value.Name
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
