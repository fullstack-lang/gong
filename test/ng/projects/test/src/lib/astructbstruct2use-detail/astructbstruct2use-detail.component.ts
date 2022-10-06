// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { AstructBstruct2UseDB } from '../astructbstruct2use-db'
import { AstructBstruct2UseService } from '../astructbstruct2use.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { AstructDB } from '../astruct-db'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// AstructBstruct2UseDetailComponent is initilizaed from different routes
// AstructBstruct2UseDetailComponentState detail different cases 
enum AstructBstruct2UseDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Astruct_Anarrayofb2Use_SET,
}

@Component({
	selector: 'app-astructbstruct2use-detail',
	templateUrl: './astructbstruct2use-detail.component.html',
	styleUrls: ['./astructbstruct2use-detail.component.css'],
})
export class AstructBstruct2UseDetailComponent implements OnInit {

	// insertion point for declarations

	// the AstructBstruct2UseDB of interest
	astructbstruct2use: AstructBstruct2UseDB = new AstructBstruct2UseDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: AstructBstruct2UseDetailComponentState = AstructBstruct2UseDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	constructor(
		private astructbstruct2useService: AstructBstruct2UseService,
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
			this.state = AstructBstruct2UseDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = AstructBstruct2UseDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "Anarrayofb2Use":
						// console.log("AstructBstruct2Use" + " is instanciated with back pointer to instance " + this.id + " Astruct association Anarrayofb2Use")
						this.state = AstructBstruct2UseDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Astruct_Anarrayofb2Use_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getAstructBstruct2Use()

		// observable for changes in structs
		this.astructbstruct2useService.AstructBstruct2UseServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getAstructBstruct2Use()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getAstructBstruct2Use(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case AstructBstruct2UseDetailComponentState.CREATE_INSTANCE:
						this.astructbstruct2use = new (AstructBstruct2UseDB)
						break;
					case AstructBstruct2UseDetailComponentState.UPDATE_INSTANCE:
						let astructbstruct2use = frontRepo.AstructBstruct2Uses.get(this.id)
						console.assert(astructbstruct2use != undefined, "missing astructbstruct2use with id:" + this.id)
						this.astructbstruct2use = astructbstruct2use!
						break;
					// insertion point for init of association field
					case AstructBstruct2UseDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Astruct_Anarrayofb2Use_SET:
						this.astructbstruct2use = new (AstructBstruct2UseDB)
						this.astructbstruct2use.Astruct_Anarrayofb2Use_reverse = frontRepo.Astructs.get(this.id)!
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
		if (this.astructbstruct2use.Bstrcut2ID == undefined) {
			this.astructbstruct2use.Bstrcut2ID = new NullInt64
		}
		if (this.astructbstruct2use.Bstrcut2 != undefined) {
			this.astructbstruct2use.Bstrcut2ID.Int64 = this.astructbstruct2use.Bstrcut2.ID
			this.astructbstruct2use.Bstrcut2ID.Valid = true
		} else {
			this.astructbstruct2use.Bstrcut2ID.Int64 = 0
			this.astructbstruct2use.Bstrcut2ID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.astructbstruct2use.Astruct_Anarrayofb2Use_reverse != undefined) {
			if (this.astructbstruct2use.Astruct_Anarrayofb2UseDBID == undefined) {
				this.astructbstruct2use.Astruct_Anarrayofb2UseDBID = new NullInt64
			}
			this.astructbstruct2use.Astruct_Anarrayofb2UseDBID.Int64 = this.astructbstruct2use.Astruct_Anarrayofb2Use_reverse.ID
			this.astructbstruct2use.Astruct_Anarrayofb2UseDBID.Valid = true
			if (this.astructbstruct2use.Astruct_Anarrayofb2UseDBID_Index == undefined) {
				this.astructbstruct2use.Astruct_Anarrayofb2UseDBID_Index = new NullInt64
			}
			this.astructbstruct2use.Astruct_Anarrayofb2UseDBID_Index.Valid = true
			this.astructbstruct2use.Astruct_Anarrayofb2Use_reverse = new AstructDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case AstructBstruct2UseDetailComponentState.UPDATE_INSTANCE:
				this.astructbstruct2useService.updateAstructBstruct2Use(this.astructbstruct2use)
					.subscribe(astructbstruct2use => {
						this.astructbstruct2useService.AstructBstruct2UseServiceChanged.next("update")
					});
				break;
			default:
				this.astructbstruct2useService.postAstructBstruct2Use(this.astructbstruct2use).subscribe(astructbstruct2use => {
					this.astructbstruct2useService.AstructBstruct2UseServiceChanged.next("post")
					this.astructbstruct2use = new (AstructBstruct2UseDB) // reset fields
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

			dialogData.ID = this.astructbstruct2use.ID!
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
			dialogData.ID = this.astructbstruct2use.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "AstructBstruct2Use"
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
			ID: this.astructbstruct2use.ID,
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
		if (this.astructbstruct2use.Name == "") {
			this.astructbstruct2use.Name = event.value.Name
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
