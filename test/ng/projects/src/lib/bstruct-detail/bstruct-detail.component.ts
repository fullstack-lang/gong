// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { BstructDB } from '../bstruct-db'
import { BstructService } from '../bstruct.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { AstructDB } from '../astruct-db'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// BstructDetailComponent is initilizaed from different routes
// BstructDetailComponentState detail different cases 
enum BstructDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Astruct_Anarrayofb_SET,
	CREATE_INSTANCE_WITH_ASSOCIATION_Astruct_Anotherarrayofb_SET,
}

@Component({
	selector: 'app-bstruct-detail',
	templateUrl: './bstruct-detail.component.html',
	styleUrls: ['./bstruct-detail.component.css'],
})
export class BstructDetailComponent implements OnInit {

	// insertion point for declarations

	// the BstructDB of interest
	bstruct: BstructDB = new BstructDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: BstructDetailComponentState = BstructDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	constructor(
		private bstructService: BstructService,
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
			this.state = BstructDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = BstructDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "Anarrayofb":
						// console.log("Bstruct" + " is instanciated with back pointer to instance " + this.id + " Astruct association Anarrayofb")
						this.state = BstructDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Astruct_Anarrayofb_SET
						break;
					case "Anotherarrayofb":
						// console.log("Bstruct" + " is instanciated with back pointer to instance " + this.id + " Astruct association Anotherarrayofb")
						this.state = BstructDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Astruct_Anotherarrayofb_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getBstruct()

		// observable for changes in structs
		this.bstructService.BstructServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getBstruct()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getBstruct(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case BstructDetailComponentState.CREATE_INSTANCE:
						this.bstruct = new (BstructDB)
						break;
					case BstructDetailComponentState.UPDATE_INSTANCE:
						let bstruct = frontRepo.Bstructs.get(this.id)
						console.assert(bstruct != undefined, "missing bstruct with id:" + this.id)
						this.bstruct = bstruct!
						break;
					// insertion point for init of association field
					case BstructDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Astruct_Anarrayofb_SET:
						this.bstruct = new (BstructDB)
						this.bstruct.Astruct_Anarrayofb_reverse = frontRepo.Astructs.get(this.id)!
						break;
					case BstructDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Astruct_Anotherarrayofb_SET:
						this.bstruct = new (BstructDB)
						this.bstruct.Astruct_Anotherarrayofb_reverse = frontRepo.Astructs.get(this.id)!
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

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.bstruct.Astruct_Anarrayofb_reverse != undefined) {
			if (this.bstruct.Astruct_AnarrayofbDBID == undefined) {
				this.bstruct.Astruct_AnarrayofbDBID = new NullInt64
			}
			this.bstruct.Astruct_AnarrayofbDBID.Int64 = this.bstruct.Astruct_Anarrayofb_reverse.ID
			this.bstruct.Astruct_AnarrayofbDBID.Valid = true
			if (this.bstruct.Astruct_AnarrayofbDBID_Index == undefined) {
				this.bstruct.Astruct_AnarrayofbDBID_Index = new NullInt64
			}
			this.bstruct.Astruct_AnarrayofbDBID_Index.Valid = true
			this.bstruct.Astruct_Anarrayofb_reverse = new AstructDB // very important, otherwise, circular JSON
		}
		if (this.bstruct.Astruct_Anotherarrayofb_reverse != undefined) {
			if (this.bstruct.Astruct_AnotherarrayofbDBID == undefined) {
				this.bstruct.Astruct_AnotherarrayofbDBID = new NullInt64
			}
			this.bstruct.Astruct_AnotherarrayofbDBID.Int64 = this.bstruct.Astruct_Anotherarrayofb_reverse.ID
			this.bstruct.Astruct_AnotherarrayofbDBID.Valid = true
			if (this.bstruct.Astruct_AnotherarrayofbDBID_Index == undefined) {
				this.bstruct.Astruct_AnotherarrayofbDBID_Index = new NullInt64
			}
			this.bstruct.Astruct_AnotherarrayofbDBID_Index.Valid = true
			this.bstruct.Astruct_Anotherarrayofb_reverse = new AstructDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case BstructDetailComponentState.UPDATE_INSTANCE:
				this.bstructService.updateBstruct(this.bstruct)
					.subscribe(bstruct => {
						this.bstructService.BstructServiceChanged.next("update")
					});
				break;
			default:
				this.bstructService.postBstruct(this.bstruct).subscribe(bstruct => {
					this.bstructService.BstructServiceChanged.next("post")
					this.bstruct = new (BstructDB) // reset fields
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

			dialogData.ID = this.bstruct.ID!
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
			dialogData.ID = this.bstruct.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "Bstruct"
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
			ID: this.bstruct.ID,
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
		if (this.bstruct.Name == "") {
			this.bstruct.Name = event.value.Name
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
