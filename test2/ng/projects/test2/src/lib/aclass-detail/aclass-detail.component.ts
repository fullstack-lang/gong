// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { AclassDB } from '../aclass-db'
import { AclassService } from '../aclass.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

// AclassDetailComponent is initilizaed from different routes
// AclassDetailComponentState detail different cases 
enum AclassDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Aclass_Anarrayofa_SET,
}

@Component({
	selector: 'app-aclass-detail',
	templateUrl: './aclass-detail.component.html',
	styleUrls: ['./aclass-detail.component.css'],
})
export class AclassDetailComponent implements OnInit {

	// insertion point for declarations
	BooleanfieldFormControl = new FormControl(false);
	AnotherbooleanfieldFormControl = new FormControl(false);
	Duration1_Hours: number
	Duration1_Minutes: number
	Duration1_Seconds: number

	// the AclassDB of interest
	aclass: AclassDB;

	// front repo
	frontRepo: FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: AclassDetailComponentState

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string
	originStructFieldName: string

	constructor(
		private aclassService: AclassService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {

		// compute state
		this.id = +this.route.snapshot.paramMap.get('id');
		this.originStruct = this.route.snapshot.paramMap.get('originStruct');
		this.originStructFieldName = this.route.snapshot.paramMap.get('originStructFieldName');

		const association = this.route.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = AclassDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = AclassDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "Anarrayofa":
						console.log("Aclass" + " is instanciated with back pointer to instance " + this.id + " Aclass association Anarrayofa")
						this.state = AclassDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Aclass_Anarrayofa_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getAclass()

		// observable for changes in structs
		this.aclassService.AclassServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getAclass()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getAclass(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case AclassDetailComponentState.CREATE_INSTANCE:
						this.aclass = new (AclassDB)
						break;
					case AclassDetailComponentState.UPDATE_INSTANCE:
						this.aclass = frontRepo.Aclasss.get(this.id)
						break;
					// insertion point for init of association field
					case AclassDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Aclass_Anarrayofa_SET:
						this.aclass = new (AclassDB)
						this.aclass.Aclass_Anarrayofa_reverse = frontRepo.Aclasss.get(this.id)
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
				this.BooleanfieldFormControl.setValue(this.aclass.Booleanfield)
				this.AnotherbooleanfieldFormControl.setValue(this.aclass.Anotherbooleanfield)
				// computation of Hours, Minutes, Seconds for Duration1
				this.Duration1_Hours = Math.floor(this.aclass.Duration1 / (3600 * 1000 * 1000 * 1000))
				this.Duration1_Minutes = Math.floor(this.aclass.Duration1 % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.Duration1_Seconds = this.aclass.Duration1 % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		this.aclass.Booleanfield = this.BooleanfieldFormControl.value
		this.aclass.Anotherbooleanfield = this.AnotherbooleanfieldFormControl.value
		this.aclass.Duration1 =
			this.Duration1_Hours * (3600 * 1000 * 1000 * 1000) +
			this.Duration1_Minutes * (60 * 1000 * 1000 * 1000) +
			this.Duration1_Seconds * (1000 * 1000 * 1000)

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.aclass.Aclass_Anarrayofa_reverse != undefined) {
			if (this.aclass.Aclass_AnarrayofaDBID == undefined) {
				this.aclass.Aclass_AnarrayofaDBID = new NullInt64
			}
			this.aclass.Aclass_AnarrayofaDBID.Int64 = this.aclass.Aclass_Anarrayofa_reverse.ID
			this.aclass.Aclass_AnarrayofaDBID.Valid = true
			if (this.aclass.Aclass_AnarrayofaDBID_Index == undefined) {
				this.aclass.Aclass_AnarrayofaDBID_Index = new NullInt64
			}
			this.aclass.Aclass_AnarrayofaDBID_Index.Valid = true
			this.aclass.Aclass_Anarrayofa_reverse = undefined // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case AclassDetailComponentState.UPDATE_INSTANCE:
				this.aclassService.updateAclass(this.aclass)
					.subscribe(aclass => {
						this.aclassService.AclassServiceChanged.next("update")
					});
				break;
			default:
				this.aclassService.postAclass(this.aclass).subscribe(aclass => {
					this.aclassService.AclassServiceChanged.next("post")
					this.aclass = {} // reset fields
				});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string, selectionMode: string,
		sourceField: string, intermediateStructField: string, nextAssociatedStruct: string ) {

		console.log("mode " + selectionMode)

		const dialogConfig = new MatDialogConfig();

		let dialogData = new DialogData();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		if (selectionMode == SelectionMode.ONE_MANY_ASSOCIATION_MODE) {

			dialogData.ID = this.aclass.ID
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
			dialogData.ID = this.aclass.ID
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "Aclass"
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
			ID: this.aclass.ID,
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

	fillUpNameIfEmpty(event) {
		if (this.aclass.Name == undefined) {
			this.aclass.Name = event.value.Name
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
			return this.mapFields_displayAsTextArea.get(fieldName)
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
