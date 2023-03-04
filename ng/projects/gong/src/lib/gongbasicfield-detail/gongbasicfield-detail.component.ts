// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { GongBasicFieldDB } from '../gongbasicfield-db'
import { GongBasicFieldService } from '../gongbasicfield.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { GongStructDB } from '../gongstruct-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// GongBasicFieldDetailComponent is initilizaed from different routes
// GongBasicFieldDetailComponentState detail different cases 
enum GongBasicFieldDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_GongStruct_GongBasicFields_SET,
}

@Component({
	selector: 'app-gongbasicfield-detail',
	templateUrl: './gongbasicfield-detail.component.html',
	styleUrls: ['./gongbasicfield-detail.component.css'],
})
export class GongBasicFieldDetailComponent implements OnInit {

	// insertion point for declarations
	IsDocLinkFormControl: UntypedFormControl = new UntypedFormControl(false);

	// the GongBasicFieldDB of interest
	gongbasicfield: GongBasicFieldDB = new GongBasicFieldDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: GongBasicFieldDetailComponentState = GongBasicFieldDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private gongbasicfieldService: GongBasicFieldService,
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
			this.state = GongBasicFieldDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = GongBasicFieldDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "GongBasicFields":
						// console.log("GongBasicField" + " is instanciated with back pointer to instance " + this.id + " GongStruct association GongBasicFields")
						this.state = GongBasicFieldDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_GongStruct_GongBasicFields_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getGongBasicField()

		// observable for changes in structs
		this.gongbasicfieldService.GongBasicFieldServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getGongBasicField()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getGongBasicField(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case GongBasicFieldDetailComponentState.CREATE_INSTANCE:
						this.gongbasicfield = new (GongBasicFieldDB)
						break;
					case GongBasicFieldDetailComponentState.UPDATE_INSTANCE:
						let gongbasicfield = frontRepo.GongBasicFields.get(this.id)
						console.assert(gongbasicfield != undefined, "missing gongbasicfield with id:" + this.id)
						this.gongbasicfield = gongbasicfield!
						break;
					// insertion point for init of association field
					case GongBasicFieldDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_GongStruct_GongBasicFields_SET:
						this.gongbasicfield = new (GongBasicFieldDB)
						this.gongbasicfield.GongStruct_GongBasicFields_reverse = frontRepo.GongStructs.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
				this.IsDocLinkFormControl.setValue(this.gongbasicfield.IsDocLink)
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		if (this.gongbasicfield.GongEnumID == undefined) {
			this.gongbasicfield.GongEnumID = new NullInt64
		}
		if (this.gongbasicfield.GongEnum != undefined) {
			this.gongbasicfield.GongEnumID.Int64 = this.gongbasicfield.GongEnum.ID
			this.gongbasicfield.GongEnumID.Valid = true
		} else {
			this.gongbasicfield.GongEnumID.Int64 = 0
			this.gongbasicfield.GongEnumID.Valid = true
		}
		this.gongbasicfield.IsDocLink = this.IsDocLinkFormControl.value

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.gongbasicfield.GongStruct_GongBasicFields_reverse != undefined) {
			if (this.gongbasicfield.GongStruct_GongBasicFieldsDBID == undefined) {
				this.gongbasicfield.GongStruct_GongBasicFieldsDBID = new NullInt64
			}
			this.gongbasicfield.GongStruct_GongBasicFieldsDBID.Int64 = this.gongbasicfield.GongStruct_GongBasicFields_reverse.ID
			this.gongbasicfield.GongStruct_GongBasicFieldsDBID.Valid = true
			if (this.gongbasicfield.GongStruct_GongBasicFieldsDBID_Index == undefined) {
				this.gongbasicfield.GongStruct_GongBasicFieldsDBID_Index = new NullInt64
			}
			this.gongbasicfield.GongStruct_GongBasicFieldsDBID_Index.Valid = true
			this.gongbasicfield.GongStruct_GongBasicFields_reverse = new GongStructDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case GongBasicFieldDetailComponentState.UPDATE_INSTANCE:
				this.gongbasicfieldService.updateGongBasicField(this.gongbasicfield, this.GONG__StackPath)
					.subscribe(gongbasicfield => {
						this.gongbasicfieldService.GongBasicFieldServiceChanged.next("update")
					});
				break;
			default:
				this.gongbasicfieldService.postGongBasicField(this.gongbasicfield, this.GONG__StackPath).subscribe(gongbasicfield => {
					this.gongbasicfieldService.GongBasicFieldServiceChanged.next("post")
					this.gongbasicfield = new (GongBasicFieldDB) // reset fields
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

			dialogData.ID = this.gongbasicfield.ID!
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
			dialogData.ID = this.gongbasicfield.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "GongBasicField"
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
			ID: this.gongbasicfield.ID,
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
		if (this.gongbasicfield.Name == "") {
			this.gongbasicfield.Name = event.value.Name
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
