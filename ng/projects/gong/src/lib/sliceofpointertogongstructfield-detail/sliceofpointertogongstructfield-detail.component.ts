// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { SliceOfPointerToGongStructFieldDB } from '../sliceofpointertogongstructfield-db'
import { SliceOfPointerToGongStructFieldService } from '../sliceofpointertogongstructfield.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { GongStructDB } from '../gongstruct-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// SliceOfPointerToGongStructFieldDetailComponent is initilizaed from different routes
// SliceOfPointerToGongStructFieldDetailComponentState detail different cases 
enum SliceOfPointerToGongStructFieldDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_GongStruct_SliceOfPointerToGongStructFields_SET,
}

@Component({
	selector: 'app-sliceofpointertogongstructfield-detail',
	templateUrl: './sliceofpointertogongstructfield-detail.component.html',
	styleUrls: ['./sliceofpointertogongstructfield-detail.component.css'],
})
export class SliceOfPointerToGongStructFieldDetailComponent implements OnInit {

	// insertion point for declarations

	// the SliceOfPointerToGongStructFieldDB of interest
	sliceofpointertogongstructfield: SliceOfPointerToGongStructFieldDB = new SliceOfPointerToGongStructFieldDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: SliceOfPointerToGongStructFieldDetailComponentState = SliceOfPointerToGongStructFieldDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private sliceofpointertogongstructfieldService: SliceOfPointerToGongStructFieldService,
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
			this.state = SliceOfPointerToGongStructFieldDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = SliceOfPointerToGongStructFieldDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "SliceOfPointerToGongStructFields":
						// console.log("SliceOfPointerToGongStructField" + " is instanciated with back pointer to instance " + this.id + " GongStruct association SliceOfPointerToGongStructFields")
						this.state = SliceOfPointerToGongStructFieldDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_GongStruct_SliceOfPointerToGongStructFields_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getSliceOfPointerToGongStructField()

		// observable for changes in structs
		this.sliceofpointertogongstructfieldService.SliceOfPointerToGongStructFieldServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getSliceOfPointerToGongStructField()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getSliceOfPointerToGongStructField(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case SliceOfPointerToGongStructFieldDetailComponentState.CREATE_INSTANCE:
						this.sliceofpointertogongstructfield = new (SliceOfPointerToGongStructFieldDB)
						break;
					case SliceOfPointerToGongStructFieldDetailComponentState.UPDATE_INSTANCE:
						let sliceofpointertogongstructfield = frontRepo.SliceOfPointerToGongStructFields.get(this.id)
						console.assert(sliceofpointertogongstructfield != undefined, "missing sliceofpointertogongstructfield with id:" + this.id)
						this.sliceofpointertogongstructfield = sliceofpointertogongstructfield!
						break;
					// insertion point for init of association field
					case SliceOfPointerToGongStructFieldDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_GongStruct_SliceOfPointerToGongStructFields_SET:
						this.sliceofpointertogongstructfield = new (SliceOfPointerToGongStructFieldDB)
						this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFields_reverse = frontRepo.GongStructs.get(this.id)!
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
		if (this.sliceofpointertogongstructfield.GongStructID == undefined) {
			this.sliceofpointertogongstructfield.GongStructID = new NullInt64
		}
		if (this.sliceofpointertogongstructfield.GongStruct != undefined) {
			this.sliceofpointertogongstructfield.GongStructID.Int64 = this.sliceofpointertogongstructfield.GongStruct.ID
			this.sliceofpointertogongstructfield.GongStructID.Valid = true
		} else {
			this.sliceofpointertogongstructfield.GongStructID.Int64 = 0
			this.sliceofpointertogongstructfield.GongStructID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFields_reverse != undefined) {
			if (this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFieldsDBID == undefined) {
				this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFieldsDBID = new NullInt64
			}
			this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFieldsDBID.Int64 = this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFields_reverse.ID
			this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFieldsDBID.Valid = true
			if (this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFieldsDBID_Index == undefined) {
				this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFieldsDBID_Index = new NullInt64
			}
			this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFieldsDBID_Index.Valid = true
			this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFields_reverse = new GongStructDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case SliceOfPointerToGongStructFieldDetailComponentState.UPDATE_INSTANCE:
				this.sliceofpointertogongstructfieldService.updateSliceOfPointerToGongStructField(this.sliceofpointertogongstructfield, this.GONG__StackPath)
					.subscribe(sliceofpointertogongstructfield => {
						this.sliceofpointertogongstructfieldService.SliceOfPointerToGongStructFieldServiceChanged.next("update")
					});
				break;
			default:
				this.sliceofpointertogongstructfieldService.postSliceOfPointerToGongStructField(this.sliceofpointertogongstructfield, this.GONG__StackPath).subscribe(sliceofpointertogongstructfield => {
					this.sliceofpointertogongstructfieldService.SliceOfPointerToGongStructFieldServiceChanged.next("post")
					this.sliceofpointertogongstructfield = new (SliceOfPointerToGongStructFieldDB) // reset fields
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

			dialogData.ID = this.sliceofpointertogongstructfield.ID!
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
			dialogData.ID = this.sliceofpointertogongstructfield.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "SliceOfPointerToGongStructField"
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
			ID: this.sliceofpointertogongstructfield.ID,
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
		if (this.sliceofpointertogongstructfield.Name == "") {
			this.sliceofpointertogongstructfield.Name = event.value.Name
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
