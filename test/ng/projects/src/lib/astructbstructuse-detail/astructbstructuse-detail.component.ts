// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { AstructBstructUseDB } from '../astructbstructuse-db'
import { AstructBstructUseService } from '../astructbstructuse.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { AstructDB } from '../astruct-db'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// AstructBstructUseDetailComponent is initilizaed from different routes
// AstructBstructUseDetailComponentState detail different cases 
enum AstructBstructUseDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Astruct_AnarrayofbUse_SET,
}

@Component({
	selector: 'app-astructbstructuse-detail',
	templateUrl: './astructbstructuse-detail.component.html',
	styleUrls: ['./astructbstructuse-detail.component.css'],
})
export class AstructBstructUseDetailComponent implements OnInit {

	// insertion point for declarations

	// the AstructBstructUseDB of interest
	astructbstructuse: AstructBstructUseDB = new AstructBstructUseDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: AstructBstructUseDetailComponentState = AstructBstructUseDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	constructor(
		private astructbstructuseService: AstructBstructUseService,
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
			this.state = AstructBstructUseDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = AstructBstructUseDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "AnarrayofbUse":
						// console.log("AstructBstructUse" + " is instanciated with back pointer to instance " + this.id + " Astruct association AnarrayofbUse")
						this.state = AstructBstructUseDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Astruct_AnarrayofbUse_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getAstructBstructUse()

		// observable for changes in structs
		this.astructbstructuseService.AstructBstructUseServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getAstructBstructUse()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getAstructBstructUse(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case AstructBstructUseDetailComponentState.CREATE_INSTANCE:
						this.astructbstructuse = new (AstructBstructUseDB)
						break;
					case AstructBstructUseDetailComponentState.UPDATE_INSTANCE:
						let astructbstructuse = frontRepo.AstructBstructUses.get(this.id)
						console.assert(astructbstructuse != undefined, "missing astructbstructuse with id:" + this.id)
						this.astructbstructuse = astructbstructuse!
						break;
					// insertion point for init of association field
					case AstructBstructUseDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Astruct_AnarrayofbUse_SET:
						this.astructbstructuse = new (AstructBstructUseDB)
						this.astructbstructuse.Astruct_AnarrayofbUse_reverse = frontRepo.Astructs.get(this.id)!
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
		if (this.astructbstructuse.Bstruct2ID == undefined) {
			this.astructbstructuse.Bstruct2ID = new NullInt64
		}
		if (this.astructbstructuse.Bstruct2 != undefined) {
			this.astructbstructuse.Bstruct2ID.Int64 = this.astructbstructuse.Bstruct2.ID
			this.astructbstructuse.Bstruct2ID.Valid = true
		} else {
			this.astructbstructuse.Bstruct2ID.Int64 = 0
			this.astructbstructuse.Bstruct2ID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.astructbstructuse.Astruct_AnarrayofbUse_reverse != undefined) {
			if (this.astructbstructuse.Astruct_AnarrayofbUseDBID == undefined) {
				this.astructbstructuse.Astruct_AnarrayofbUseDBID = new NullInt64
			}
			this.astructbstructuse.Astruct_AnarrayofbUseDBID.Int64 = this.astructbstructuse.Astruct_AnarrayofbUse_reverse.ID
			this.astructbstructuse.Astruct_AnarrayofbUseDBID.Valid = true
			if (this.astructbstructuse.Astruct_AnarrayofbUseDBID_Index == undefined) {
				this.astructbstructuse.Astruct_AnarrayofbUseDBID_Index = new NullInt64
			}
			this.astructbstructuse.Astruct_AnarrayofbUseDBID_Index.Valid = true
			this.astructbstructuse.Astruct_AnarrayofbUse_reverse = new AstructDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case AstructBstructUseDetailComponentState.UPDATE_INSTANCE:
				this.astructbstructuseService.updateAstructBstructUse(this.astructbstructuse)
					.subscribe(astructbstructuse => {
						this.astructbstructuseService.AstructBstructUseServiceChanged.next("update")
					});
				break;
			default:
				this.astructbstructuseService.postAstructBstructUse(this.astructbstructuse).subscribe(astructbstructuse => {
					this.astructbstructuseService.AstructBstructUseServiceChanged.next("post")
					this.astructbstructuse = new (AstructBstructUseDB) // reset fields
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

			dialogData.ID = this.astructbstructuse.ID!
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
			dialogData.ID = this.astructbstructuse.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "AstructBstructUse"
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
			ID: this.astructbstructuse.ID,
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
		if (this.astructbstructuse.Name == "") {
			this.astructbstructuse.Name = event.value.Name
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
