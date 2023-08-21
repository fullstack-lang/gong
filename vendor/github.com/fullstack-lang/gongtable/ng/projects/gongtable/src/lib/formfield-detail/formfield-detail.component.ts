// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { FormFieldDB } from '../formfield-db'
import { FormFieldService } from '../formfield.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { InputTypeEnumSelect, InputTypeEnumList } from '../InputTypeEnum'
import { FormDivDB } from '../formdiv-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// FormFieldDetailComponent is initilizaed from different routes
// FormFieldDetailComponentState detail different cases 
enum FormFieldDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_FormDiv_FormFields_SET,
}

@Component({
	selector: 'app-formfield-detail',
	templateUrl: './formfield-detail.component.html',
	styleUrls: ['./formfield-detail.component.css'],
})
export class FormFieldDetailComponent implements OnInit {

	// insertion point for declarations
	InputTypeEnumList: InputTypeEnumSelect[] = []

	// the FormFieldDB of interest
	formfield: FormFieldDB = new FormFieldDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: FormFieldDetailComponentState = FormFieldDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private formfieldService: FormFieldService,
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
			this.state = FormFieldDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = FormFieldDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "FormFields":
						// console.log("FormField" + " is instanciated with back pointer to instance " + this.id + " FormDiv association FormFields")
						this.state = FormFieldDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_FormDiv_FormFields_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getFormField()

		// observable for changes in structs
		this.formfieldService.FormFieldServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getFormField()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.InputTypeEnumList = InputTypeEnumList
	}

	getFormField(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case FormFieldDetailComponentState.CREATE_INSTANCE:
						this.formfield = new (FormFieldDB)
						break;
					case FormFieldDetailComponentState.UPDATE_INSTANCE:
						let formfield = frontRepo.FormFields.get(this.id)
						console.assert(formfield != undefined, "missing formfield with id:" + this.id)
						this.formfield = formfield!
						break;
					// insertion point for init of association field
					case FormFieldDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_FormDiv_FormFields_SET:
						this.formfield = new (FormFieldDB)
						this.formfield.FormDiv_FormFields_reverse = frontRepo.FormDivs.get(this.id)!
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
		if (this.formfield.FormFieldStringID == undefined) {
			this.formfield.FormFieldStringID = new NullInt64
		}
		if (this.formfield.FormFieldString != undefined) {
			this.formfield.FormFieldStringID.Int64 = this.formfield.FormFieldString.ID
			this.formfield.FormFieldStringID.Valid = true
		} else {
			this.formfield.FormFieldStringID.Int64 = 0
			this.formfield.FormFieldStringID.Valid = true
		}
		if (this.formfield.FormFieldFloat64ID == undefined) {
			this.formfield.FormFieldFloat64ID = new NullInt64
		}
		if (this.formfield.FormFieldFloat64 != undefined) {
			this.formfield.FormFieldFloat64ID.Int64 = this.formfield.FormFieldFloat64.ID
			this.formfield.FormFieldFloat64ID.Valid = true
		} else {
			this.formfield.FormFieldFloat64ID.Int64 = 0
			this.formfield.FormFieldFloat64ID.Valid = true
		}
		if (this.formfield.FormFieldIntID == undefined) {
			this.formfield.FormFieldIntID = new NullInt64
		}
		if (this.formfield.FormFieldInt != undefined) {
			this.formfield.FormFieldIntID.Int64 = this.formfield.FormFieldInt.ID
			this.formfield.FormFieldIntID.Valid = true
		} else {
			this.formfield.FormFieldIntID.Int64 = 0
			this.formfield.FormFieldIntID.Valid = true
		}
		if (this.formfield.FormFieldDateID == undefined) {
			this.formfield.FormFieldDateID = new NullInt64
		}
		if (this.formfield.FormFieldDate != undefined) {
			this.formfield.FormFieldDateID.Int64 = this.formfield.FormFieldDate.ID
			this.formfield.FormFieldDateID.Valid = true
		} else {
			this.formfield.FormFieldDateID.Int64 = 0
			this.formfield.FormFieldDateID.Valid = true
		}
		if (this.formfield.FormFieldTimeID == undefined) {
			this.formfield.FormFieldTimeID = new NullInt64
		}
		if (this.formfield.FormFieldTime != undefined) {
			this.formfield.FormFieldTimeID.Int64 = this.formfield.FormFieldTime.ID
			this.formfield.FormFieldTimeID.Valid = true
		} else {
			this.formfield.FormFieldTimeID.Int64 = 0
			this.formfield.FormFieldTimeID.Valid = true
		}
		if (this.formfield.FormFieldDateTimeID == undefined) {
			this.formfield.FormFieldDateTimeID = new NullInt64
		}
		if (this.formfield.FormFieldDateTime != undefined) {
			this.formfield.FormFieldDateTimeID.Int64 = this.formfield.FormFieldDateTime.ID
			this.formfield.FormFieldDateTimeID.Valid = true
		} else {
			this.formfield.FormFieldDateTimeID.Int64 = 0
			this.formfield.FormFieldDateTimeID.Valid = true
		}
		if (this.formfield.FormFieldSelectID == undefined) {
			this.formfield.FormFieldSelectID = new NullInt64
		}
		if (this.formfield.FormFieldSelect != undefined) {
			this.formfield.FormFieldSelectID.Int64 = this.formfield.FormFieldSelect.ID
			this.formfield.FormFieldSelectID.Valid = true
		} else {
			this.formfield.FormFieldSelectID.Int64 = 0
			this.formfield.FormFieldSelectID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.formfield.FormDiv_FormFields_reverse != undefined) {
			if (this.formfield.FormDiv_FormFieldsDBID == undefined) {
				this.formfield.FormDiv_FormFieldsDBID = new NullInt64
			}
			this.formfield.FormDiv_FormFieldsDBID.Int64 = this.formfield.FormDiv_FormFields_reverse.ID
			this.formfield.FormDiv_FormFieldsDBID.Valid = true
			if (this.formfield.FormDiv_FormFieldsDBID_Index == undefined) {
				this.formfield.FormDiv_FormFieldsDBID_Index = new NullInt64
			}
			this.formfield.FormDiv_FormFieldsDBID_Index.Valid = true
			this.formfield.FormDiv_FormFields_reverse = new FormDivDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case FormFieldDetailComponentState.UPDATE_INSTANCE:
				this.formfieldService.updateFormField(this.formfield, this.GONG__StackPath)
					.subscribe(formfield => {
						this.formfieldService.FormFieldServiceChanged.next("update")
					});
				break;
			default:
				this.formfieldService.postFormField(this.formfield, this.GONG__StackPath).subscribe(formfield => {
					this.formfieldService.FormFieldServiceChanged.next("post")
					this.formfield = new (FormFieldDB) // reset fields
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

			dialogData.ID = this.formfield.ID!
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
			dialogData.ID = this.formfield.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "FormField"
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
			ID: this.formfield.ID,
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
		if (this.formfield.Name == "") {
			this.formfield.Name = event.value.Name
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
