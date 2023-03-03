// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { NoteShapeDB } from '../noteshape-db'
import { NoteShapeService } from '../noteshape.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { ClassdiagramDB } from '../classdiagram-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// NoteShapeDetailComponent is initilizaed from different routes
// NoteShapeDetailComponentState detail different cases 
enum NoteShapeDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Classdiagram_NoteShapes_SET,
}

@Component({
	selector: 'app-noteshape-detail',
	templateUrl: './noteshape-detail.component.html',
	styleUrls: ['./noteshape-detail.component.css'],
})
export class NoteShapeDetailComponent implements OnInit {

	// insertion point for declarations
	MatchedFormControl: UntypedFormControl = new UntypedFormControl(false);

	// the NoteShapeDB of interest
	noteshape: NoteShapeDB = new NoteShapeDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: NoteShapeDetailComponentState = NoteShapeDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private noteshapeService: NoteShapeService,
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
			this.state = NoteShapeDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = NoteShapeDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "NoteShapes":
						// console.log("NoteShape" + " is instanciated with back pointer to instance " + this.id + " Classdiagram association NoteShapes")
						this.state = NoteShapeDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Classdiagram_NoteShapes_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getNoteShape()

		// observable for changes in structs
		this.noteshapeService.NoteShapeServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getNoteShape()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getNoteShape(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case NoteShapeDetailComponentState.CREATE_INSTANCE:
						this.noteshape = new (NoteShapeDB)
						break;
					case NoteShapeDetailComponentState.UPDATE_INSTANCE:
						let noteshape = frontRepo.NoteShapes.get(this.id)
						console.assert(noteshape != undefined, "missing noteshape with id:" + this.id)
						this.noteshape = noteshape!
						break;
					// insertion point for init of association field
					case NoteShapeDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Classdiagram_NoteShapes_SET:
						this.noteshape = new (NoteShapeDB)
						this.noteshape.Classdiagram_NoteShapes_reverse = frontRepo.Classdiagrams.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
				this.MatchedFormControl.setValue(this.noteshape.Matched)
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		this.noteshape.Matched = this.MatchedFormControl.value

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.noteshape.Classdiagram_NoteShapes_reverse != undefined) {
			if (this.noteshape.Classdiagram_NoteShapesDBID == undefined) {
				this.noteshape.Classdiagram_NoteShapesDBID = new NullInt64
			}
			this.noteshape.Classdiagram_NoteShapesDBID.Int64 = this.noteshape.Classdiagram_NoteShapes_reverse.ID
			this.noteshape.Classdiagram_NoteShapesDBID.Valid = true
			if (this.noteshape.Classdiagram_NoteShapesDBID_Index == undefined) {
				this.noteshape.Classdiagram_NoteShapesDBID_Index = new NullInt64
			}
			this.noteshape.Classdiagram_NoteShapesDBID_Index.Valid = true
			this.noteshape.Classdiagram_NoteShapes_reverse = new ClassdiagramDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case NoteShapeDetailComponentState.UPDATE_INSTANCE:
				this.noteshapeService.updateNoteShape(this.noteshape, this.GONG__StackPath)
					.subscribe(noteshape => {
						this.noteshapeService.NoteShapeServiceChanged.next("update")
					});
				break;
			default:
				this.noteshapeService.postNoteShape(this.noteshape, this.GONG__StackPath).subscribe(noteshape => {
					this.noteshapeService.NoteShapeServiceChanged.next("post")
					this.noteshape = new (NoteShapeDB) // reset fields
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

			dialogData.ID = this.noteshape.ID!
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
			dialogData.ID = this.noteshape.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "NoteShape"
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
			ID: this.noteshape.ID,
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
		if (this.noteshape.Name == "") {
			this.noteshape.Name = event.value.Name
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
