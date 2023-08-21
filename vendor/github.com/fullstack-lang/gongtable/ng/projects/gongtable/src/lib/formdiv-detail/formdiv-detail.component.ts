// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { FormDivDB } from '../formdiv-db'
import { FormDivService } from '../formdiv.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { FormGroupDB } from '../formgroup-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// FormDivDetailComponent is initilizaed from different routes
// FormDivDetailComponentState detail different cases 
enum FormDivDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_FormGroup_FormDivs_SET,
}

@Component({
	selector: 'app-formdiv-detail',
	templateUrl: './formdiv-detail.component.html',
	styleUrls: ['./formdiv-detail.component.css'],
})
export class FormDivDetailComponent implements OnInit {

	// insertion point for declarations

	// the FormDivDB of interest
	formdiv: FormDivDB = new FormDivDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: FormDivDetailComponentState = FormDivDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private formdivService: FormDivService,
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
			this.state = FormDivDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = FormDivDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "FormDivs":
						// console.log("FormDiv" + " is instanciated with back pointer to instance " + this.id + " FormGroup association FormDivs")
						this.state = FormDivDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_FormGroup_FormDivs_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getFormDiv()

		// observable for changes in structs
		this.formdivService.FormDivServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getFormDiv()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getFormDiv(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case FormDivDetailComponentState.CREATE_INSTANCE:
						this.formdiv = new (FormDivDB)
						break;
					case FormDivDetailComponentState.UPDATE_INSTANCE:
						let formdiv = frontRepo.FormDivs.get(this.id)
						console.assert(formdiv != undefined, "missing formdiv with id:" + this.id)
						this.formdiv = formdiv!
						break;
					// insertion point for init of association field
					case FormDivDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_FormGroup_FormDivs_SET:
						this.formdiv = new (FormDivDB)
						this.formdiv.FormGroup_FormDivs_reverse = frontRepo.FormGroups.get(this.id)!
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
		if (this.formdiv.FormEditAssocButtonID == undefined) {
			this.formdiv.FormEditAssocButtonID = new NullInt64
		}
		if (this.formdiv.FormEditAssocButton != undefined) {
			this.formdiv.FormEditAssocButtonID.Int64 = this.formdiv.FormEditAssocButton.ID
			this.formdiv.FormEditAssocButtonID.Valid = true
		} else {
			this.formdiv.FormEditAssocButtonID.Int64 = 0
			this.formdiv.FormEditAssocButtonID.Valid = true
		}
		if (this.formdiv.FormSortAssocButtonID == undefined) {
			this.formdiv.FormSortAssocButtonID = new NullInt64
		}
		if (this.formdiv.FormSortAssocButton != undefined) {
			this.formdiv.FormSortAssocButtonID.Int64 = this.formdiv.FormSortAssocButton.ID
			this.formdiv.FormSortAssocButtonID.Valid = true
		} else {
			this.formdiv.FormSortAssocButtonID.Int64 = 0
			this.formdiv.FormSortAssocButtonID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.formdiv.FormGroup_FormDivs_reverse != undefined) {
			if (this.formdiv.FormGroup_FormDivsDBID == undefined) {
				this.formdiv.FormGroup_FormDivsDBID = new NullInt64
			}
			this.formdiv.FormGroup_FormDivsDBID.Int64 = this.formdiv.FormGroup_FormDivs_reverse.ID
			this.formdiv.FormGroup_FormDivsDBID.Valid = true
			if (this.formdiv.FormGroup_FormDivsDBID_Index == undefined) {
				this.formdiv.FormGroup_FormDivsDBID_Index = new NullInt64
			}
			this.formdiv.FormGroup_FormDivsDBID_Index.Valid = true
			this.formdiv.FormGroup_FormDivs_reverse = new FormGroupDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case FormDivDetailComponentState.UPDATE_INSTANCE:
				this.formdivService.updateFormDiv(this.formdiv, this.GONG__StackPath)
					.subscribe(formdiv => {
						this.formdivService.FormDivServiceChanged.next("update")
					});
				break;
			default:
				this.formdivService.postFormDiv(this.formdiv, this.GONG__StackPath).subscribe(formdiv => {
					this.formdivService.FormDivServiceChanged.next("post")
					this.formdiv = new (FormDivDB) // reset fields
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

			dialogData.ID = this.formdiv.ID!
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
			dialogData.ID = this.formdiv.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "FormDiv"
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
			ID: this.formdiv.ID,
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
		if (this.formdiv.Name == "") {
			this.formdiv.Name = event.value.Name
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
