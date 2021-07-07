// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { BclassDB } from '../bclass-db'
import { BclassService } from '../bclass.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

// BclassDetailComponent is initilizaed from different routes
// BclassDetailComponentState detail different cases 
enum BclassDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Aclass_Anarrayofb_SET,
	CREATE_INSTANCE_WITH_ASSOCIATION_Aclass_Anotherarrayofb_SET,
}

@Component({
	selector: 'app-bclass-detail',
	templateUrl: './bclass-detail.component.html',
	styleUrls: ['./bclass-detail.component.css'],
})
export class BclassDetailComponent implements OnInit {

	// insertion point for declarations

	// the BclassDB of interest
	bclass: BclassDB;

	// front repo
	frontRepo: FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: BclassDetailComponentState

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string
	originStructFieldName: string

	constructor(
		private bclassService: BclassService,
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
			this.state = BclassDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = BclassDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "Anarrayofb":
						console.log("Bclass" + " is instanciated with back pointer to instance " + this.id + " Aclass association Anarrayofb")
						this.state = BclassDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Aclass_Anarrayofb_SET
						break;
					case "Anotherarrayofb":
						console.log("Bclass" + " is instanciated with back pointer to instance " + this.id + " Aclass association Anotherarrayofb")
						this.state = BclassDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Aclass_Anotherarrayofb_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getBclass()

		// observable for changes in structs
		this.bclassService.BclassServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getBclass()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getBclass(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case BclassDetailComponentState.CREATE_INSTANCE:
						this.bclass = new (BclassDB)
						break;
					case BclassDetailComponentState.UPDATE_INSTANCE:
						this.bclass = frontRepo.Bclasss.get(this.id)
						break;
					// insertion point for init of association field
					case BclassDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Aclass_Anarrayofb_SET:
						this.bclass = new (BclassDB)
						this.bclass.Aclass_Anarrayofb_reverse = frontRepo.Aclasss.get(this.id)
						break;
					case BclassDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Aclass_Anotherarrayofb_SET:
						this.bclass = new (BclassDB)
						this.bclass.Aclass_Anotherarrayofb_reverse = frontRepo.Aclasss.get(this.id)
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
		if (this.bclass.Aclass_Anarrayofb_reverse != undefined) {
			if (this.bclass.Aclass_AnarrayofbDBID == undefined) {
				this.bclass.Aclass_AnarrayofbDBID = new NullInt64
			}
			this.bclass.Aclass_AnarrayofbDBID.Int64 = this.bclass.Aclass_Anarrayofb_reverse.ID
			this.bclass.Aclass_AnarrayofbDBID.Valid = true
			if (this.bclass.Aclass_AnarrayofbDBID_Index == undefined) {
				this.bclass.Aclass_AnarrayofbDBID_Index = new NullInt64
			}
			this.bclass.Aclass_AnarrayofbDBID_Index.Valid = true
			this.bclass.Aclass_Anarrayofb_reverse = undefined // very important, otherwise, circular JSON
		}
		if (this.bclass.Aclass_Anotherarrayofb_reverse != undefined) {
			if (this.bclass.Aclass_AnotherarrayofbDBID == undefined) {
				this.bclass.Aclass_AnotherarrayofbDBID = new NullInt64
			}
			this.bclass.Aclass_AnotherarrayofbDBID.Int64 = this.bclass.Aclass_Anotherarrayofb_reverse.ID
			this.bclass.Aclass_AnotherarrayofbDBID.Valid = true
			if (this.bclass.Aclass_AnotherarrayofbDBID_Index == undefined) {
				this.bclass.Aclass_AnotherarrayofbDBID_Index = new NullInt64
			}
			this.bclass.Aclass_AnotherarrayofbDBID_Index.Valid = true
			this.bclass.Aclass_Anotherarrayofb_reverse = undefined // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case BclassDetailComponentState.UPDATE_INSTANCE:
				this.bclassService.updateBclass(this.bclass)
					.subscribe(bclass => {
						this.bclassService.BclassServiceChanged.next("update")
					});
				break;
			default:
				this.bclassService.postBclass(this.bclass).subscribe(bclass => {
					this.bclassService.BclassServiceChanged.next("post")
					this.bclass = {} // reset fields
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

			dialogData.ID = this.bclass.ID
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
			dialogData.ID = this.bclass.ID
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "Bclass"
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
			ID: this.bclass.ID,
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
		if (this.bclass.Name == undefined) {
			this.bclass.Name = event.value.Name
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
