// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { FormFieldSelectDB } from '../formfieldselect-db'
import { FormFieldSelectService } from '../formfieldselect.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// FormFieldSelectDetailComponent is initilizaed from different routes
// FormFieldSelectDetailComponentState detail different cases 
enum FormFieldSelectDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
}

@Component({
	selector: 'app-formfieldselect-detail',
	templateUrl: './formfieldselect-detail.component.html',
	styleUrls: ['./formfieldselect-detail.component.css'],
})
export class FormFieldSelectDetailComponent implements OnInit {

	// insertion point for declarations
	CanBeEmptyFormControl: UntypedFormControl = new UntypedFormControl(false);

	// the FormFieldSelectDB of interest
	formfieldselect: FormFieldSelectDB = new FormFieldSelectDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: FormFieldSelectDetailComponentState = FormFieldSelectDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private formfieldselectService: FormFieldSelectService,
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
			this.state = FormFieldSelectDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = FormFieldSelectDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getFormFieldSelect()

		// observable for changes in structs
		this.formfieldselectService.FormFieldSelectServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getFormFieldSelect()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getFormFieldSelect(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case FormFieldSelectDetailComponentState.CREATE_INSTANCE:
						this.formfieldselect = new (FormFieldSelectDB)
						break;
					case FormFieldSelectDetailComponentState.UPDATE_INSTANCE:
						let formfieldselect = frontRepo.FormFieldSelects.get(this.id)
						console.assert(formfieldselect != undefined, "missing formfieldselect with id:" + this.id)
						this.formfieldselect = formfieldselect!
						break;
					// insertion point for init of association field
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
				this.CanBeEmptyFormControl.setValue(this.formfieldselect.CanBeEmpty)
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		if (this.formfieldselect.ValueID == undefined) {
			this.formfieldselect.ValueID = new NullInt64
		}
		if (this.formfieldselect.Value != undefined) {
			this.formfieldselect.ValueID.Int64 = this.formfieldselect.Value.ID
			this.formfieldselect.ValueID.Valid = true
		} else {
			this.formfieldselect.ValueID.Int64 = 0
			this.formfieldselect.ValueID.Valid = true
		}
		this.formfieldselect.CanBeEmpty = this.CanBeEmptyFormControl.value

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers

		switch (this.state) {
			case FormFieldSelectDetailComponentState.UPDATE_INSTANCE:
				this.formfieldselectService.updateFormFieldSelect(this.formfieldselect, this.GONG__StackPath)
					.subscribe(formfieldselect => {
						this.formfieldselectService.FormFieldSelectServiceChanged.next("update")
					});
				break;
			default:
				this.formfieldselectService.postFormFieldSelect(this.formfieldselect, this.GONG__StackPath).subscribe(formfieldselect => {
					this.formfieldselectService.FormFieldSelectServiceChanged.next("post")
					this.formfieldselect = new (FormFieldSelectDB) // reset fields
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

			dialogData.ID = this.formfieldselect.ID!
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
			dialogData.ID = this.formfieldselect.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "FormFieldSelect"
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
			ID: this.formfieldselect.ID,
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
		if (this.formfieldselect.Name == "") {
			this.formfieldselect.Name = event.value.Name
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
