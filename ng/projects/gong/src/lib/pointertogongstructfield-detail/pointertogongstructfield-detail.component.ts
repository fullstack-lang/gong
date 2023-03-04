// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { PointerToGongStructFieldDB } from '../pointertogongstructfield-db'
import { PointerToGongStructFieldService } from '../pointertogongstructfield.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { GongStructDB } from '../gongstruct-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// PointerToGongStructFieldDetailComponent is initilizaed from different routes
// PointerToGongStructFieldDetailComponentState detail different cases 
enum PointerToGongStructFieldDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_GongStruct_PointerToGongStructFields_SET,
}

@Component({
	selector: 'app-pointertogongstructfield-detail',
	templateUrl: './pointertogongstructfield-detail.component.html',
	styleUrls: ['./pointertogongstructfield-detail.component.css'],
})
export class PointerToGongStructFieldDetailComponent implements OnInit {

	// insertion point for declarations

	// the PointerToGongStructFieldDB of interest
	pointertogongstructfield: PointerToGongStructFieldDB = new PointerToGongStructFieldDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: PointerToGongStructFieldDetailComponentState = PointerToGongStructFieldDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private pointertogongstructfieldService: PointerToGongStructFieldService,
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
			this.state = PointerToGongStructFieldDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = PointerToGongStructFieldDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "PointerToGongStructFields":
						// console.log("PointerToGongStructField" + " is instanciated with back pointer to instance " + this.id + " GongStruct association PointerToGongStructFields")
						this.state = PointerToGongStructFieldDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_GongStruct_PointerToGongStructFields_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getPointerToGongStructField()

		// observable for changes in structs
		this.pointertogongstructfieldService.PointerToGongStructFieldServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getPointerToGongStructField()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getPointerToGongStructField(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case PointerToGongStructFieldDetailComponentState.CREATE_INSTANCE:
						this.pointertogongstructfield = new (PointerToGongStructFieldDB)
						break;
					case PointerToGongStructFieldDetailComponentState.UPDATE_INSTANCE:
						let pointertogongstructfield = frontRepo.PointerToGongStructFields.get(this.id)
						console.assert(pointertogongstructfield != undefined, "missing pointertogongstructfield with id:" + this.id)
						this.pointertogongstructfield = pointertogongstructfield!
						break;
					// insertion point for init of association field
					case PointerToGongStructFieldDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_GongStruct_PointerToGongStructFields_SET:
						this.pointertogongstructfield = new (PointerToGongStructFieldDB)
						this.pointertogongstructfield.GongStruct_PointerToGongStructFields_reverse = frontRepo.GongStructs.get(this.id)!
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
		if (this.pointertogongstructfield.GongStructID == undefined) {
			this.pointertogongstructfield.GongStructID = new NullInt64
		}
		if (this.pointertogongstructfield.GongStruct != undefined) {
			this.pointertogongstructfield.GongStructID.Int64 = this.pointertogongstructfield.GongStruct.ID
			this.pointertogongstructfield.GongStructID.Valid = true
		} else {
			this.pointertogongstructfield.GongStructID.Int64 = 0
			this.pointertogongstructfield.GongStructID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.pointertogongstructfield.GongStruct_PointerToGongStructFields_reverse != undefined) {
			if (this.pointertogongstructfield.GongStruct_PointerToGongStructFieldsDBID == undefined) {
				this.pointertogongstructfield.GongStruct_PointerToGongStructFieldsDBID = new NullInt64
			}
			this.pointertogongstructfield.GongStruct_PointerToGongStructFieldsDBID.Int64 = this.pointertogongstructfield.GongStruct_PointerToGongStructFields_reverse.ID
			this.pointertogongstructfield.GongStruct_PointerToGongStructFieldsDBID.Valid = true
			if (this.pointertogongstructfield.GongStruct_PointerToGongStructFieldsDBID_Index == undefined) {
				this.pointertogongstructfield.GongStruct_PointerToGongStructFieldsDBID_Index = new NullInt64
			}
			this.pointertogongstructfield.GongStruct_PointerToGongStructFieldsDBID_Index.Valid = true
			this.pointertogongstructfield.GongStruct_PointerToGongStructFields_reverse = new GongStructDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case PointerToGongStructFieldDetailComponentState.UPDATE_INSTANCE:
				this.pointertogongstructfieldService.updatePointerToGongStructField(this.pointertogongstructfield, this.GONG__StackPath)
					.subscribe(pointertogongstructfield => {
						this.pointertogongstructfieldService.PointerToGongStructFieldServiceChanged.next("update")
					});
				break;
			default:
				this.pointertogongstructfieldService.postPointerToGongStructField(this.pointertogongstructfield, this.GONG__StackPath).subscribe(pointertogongstructfield => {
					this.pointertogongstructfieldService.PointerToGongStructFieldServiceChanged.next("post")
					this.pointertogongstructfield = new (PointerToGongStructFieldDB) // reset fields
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

			dialogData.ID = this.pointertogongstructfield.ID!
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
			dialogData.ID = this.pointertogongstructfield.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "PointerToGongStructField"
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
			ID: this.pointertogongstructfield.ID,
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
		if (this.pointertogongstructfield.Name == "") {
			this.pointertogongstructfield.Name = event.value.Name
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
