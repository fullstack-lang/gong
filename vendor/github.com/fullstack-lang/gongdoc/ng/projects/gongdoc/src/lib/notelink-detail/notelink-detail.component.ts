// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { NoteLinkDB } from '../notelink-db'
import { NoteLinkService } from '../notelink.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { ReferenceTypeSelect, ReferenceTypeList } from '../ReferenceType'
import { NoteShapeDB } from '../noteshape-db'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// NoteLinkDetailComponent is initilizaed from different routes
// NoteLinkDetailComponentState detail different cases 
enum NoteLinkDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_NoteShape_NoteLinks_SET,
}

@Component({
	selector: 'app-notelink-detail',
	templateUrl: './notelink-detail.component.html',
	styleUrls: ['./notelink-detail.component.css'],
})
export class NoteLinkDetailComponent implements OnInit {

	// insertion point for declarations
	ReferenceTypeList: ReferenceTypeSelect[] = []

	// the NoteLinkDB of interest
	notelink: NoteLinkDB = new NoteLinkDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: NoteLinkDetailComponentState = NoteLinkDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	constructor(
		private notelinkService: NoteLinkService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {

		// compute state
		this.id = +this.route.snapshot.paramMap.get('id')!;
		this.originStruct = this.route.snapshot.paramMap.get('originStruct')!;
		this.originStructFieldName = this.route.snapshot.paramMap.get('originStructFieldName')!;

		const association = this.route.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = NoteLinkDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = NoteLinkDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "NoteLinks":
						// console.log("NoteLink" + " is instanciated with back pointer to instance " + this.id + " NoteShape association NoteLinks")
						this.state = NoteLinkDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_NoteShape_NoteLinks_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getNoteLink()

		// observable for changes in structs
		this.notelinkService.NoteLinkServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getNoteLink()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.ReferenceTypeList = ReferenceTypeList
	}

	getNoteLink(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case NoteLinkDetailComponentState.CREATE_INSTANCE:
						this.notelink = new (NoteLinkDB)
						break;
					case NoteLinkDetailComponentState.UPDATE_INSTANCE:
						let notelink = frontRepo.NoteLinks.get(this.id)
						console.assert(notelink != undefined, "missing notelink with id:" + this.id)
						this.notelink = notelink!
						break;
					// insertion point for init of association field
					case NoteLinkDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_NoteShape_NoteLinks_SET:
						this.notelink = new (NoteLinkDB)
						this.notelink.NoteShape_NoteLinks_reverse = frontRepo.NoteShapes.get(this.id)!
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
		if (this.notelink.ClassshapeID == undefined) {
			this.notelink.ClassshapeID = new NullInt64
		}
		if (this.notelink.Classshape != undefined) {
			this.notelink.ClassshapeID.Int64 = this.notelink.Classshape.ID
			this.notelink.ClassshapeID.Valid = true
		} else {
			this.notelink.ClassshapeID.Int64 = 0
			this.notelink.ClassshapeID.Valid = true
		}
		if (this.notelink.LinkID == undefined) {
			this.notelink.LinkID = new NullInt64
		}
		if (this.notelink.Link != undefined) {
			this.notelink.LinkID.Int64 = this.notelink.Link.ID
			this.notelink.LinkID.Valid = true
		} else {
			this.notelink.LinkID.Int64 = 0
			this.notelink.LinkID.Valid = true
		}
		if (this.notelink.MiddleverticeID == undefined) {
			this.notelink.MiddleverticeID = new NullInt64
		}
		if (this.notelink.Middlevertice != undefined) {
			this.notelink.MiddleverticeID.Int64 = this.notelink.Middlevertice.ID
			this.notelink.MiddleverticeID.Valid = true
		} else {
			this.notelink.MiddleverticeID.Int64 = 0
			this.notelink.MiddleverticeID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.notelink.NoteShape_NoteLinks_reverse != undefined) {
			if (this.notelink.NoteShape_NoteLinksDBID == undefined) {
				this.notelink.NoteShape_NoteLinksDBID = new NullInt64
			}
			this.notelink.NoteShape_NoteLinksDBID.Int64 = this.notelink.NoteShape_NoteLinks_reverse.ID
			this.notelink.NoteShape_NoteLinksDBID.Valid = true
			if (this.notelink.NoteShape_NoteLinksDBID_Index == undefined) {
				this.notelink.NoteShape_NoteLinksDBID_Index = new NullInt64
			}
			this.notelink.NoteShape_NoteLinksDBID_Index.Valid = true
			this.notelink.NoteShape_NoteLinks_reverse = new NoteShapeDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case NoteLinkDetailComponentState.UPDATE_INSTANCE:
				this.notelinkService.updateNoteLink(this.notelink)
					.subscribe(notelink => {
						this.notelinkService.NoteLinkServiceChanged.next("update")
					});
				break;
			default:
				this.notelinkService.postNoteLink(this.notelink).subscribe(notelink => {
					this.notelinkService.NoteLinkServiceChanged.next("post")
					this.notelink = new (NoteLinkDB) // reset fields
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

			dialogData.ID = this.notelink.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

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
			dialogData.ID = this.notelink.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "NoteLink"
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
			ID: this.notelink.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
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
		if (this.notelink.Name == "") {
			this.notelink.Name = event.value.Name
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
