// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { AclassBclass2UseDB } from '../aclassbclass2use-db'
import { AclassBclass2UseService } from '../aclassbclass2use.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

// AclassBclass2UseDetailComponent is initilizaed from different routes
// AclassBclass2UseDetailComponentState detail different cases 
enum AclassBclass2UseDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Aclass_Anarrayofb2Use_SET,
}

@Component({
	selector: 'app-aclassbclass2use-detail',
	templateUrl: './aclassbclass2use-detail.component.html',
	styleUrls: ['./aclassbclass2use-detail.component.css'],
})
export class AclassBclass2UseDetailComponent implements OnInit {

	// insertion point for declarations

	// the AclassBclass2UseDB of interest
	aclassbclass2use: AclassBclass2UseDB;

	// front repo
	frontRepo: FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: AclassBclass2UseDetailComponentState

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string
	originStructFieldName: string

	constructor(
		private aclassbclass2useService: AclassBclass2UseService,
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
			this.state = AclassBclass2UseDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = AclassBclass2UseDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "Anarrayofb2Use":
						console.log("AclassBclass2Use" + " is instanciated with back pointer to instance " + this.id + " Aclass association Anarrayofb2Use")
						this.state = AclassBclass2UseDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Aclass_Anarrayofb2Use_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getAclassBclass2Use()

		// observable for changes in structs
		this.aclassbclass2useService.AclassBclass2UseServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getAclassBclass2Use()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getAclassBclass2Use(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case AclassBclass2UseDetailComponentState.CREATE_INSTANCE:
						this.aclassbclass2use = new (AclassBclass2UseDB)
						break;
					case AclassBclass2UseDetailComponentState.UPDATE_INSTANCE:
						this.aclassbclass2use = frontRepo.AclassBclass2Uses.get(this.id)
						break;
					// insertion point for init of association field
					case AclassBclass2UseDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Aclass_Anarrayofb2Use_SET:
						this.aclassbclass2use = new (AclassBclass2UseDB)
						this.aclassbclass2use.Aclass_Anarrayofb2Use_reverse = frontRepo.Aclasss.get(this.id)
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
		if (this.aclassbclass2use.Bclass2ID == undefined) {
			this.aclassbclass2use.Bclass2ID = new NullInt64
		}
		if (this.aclassbclass2use.Bclass2 != undefined) {
			this.aclassbclass2use.Bclass2ID.Int64 = this.aclassbclass2use.Bclass2.ID
			this.aclassbclass2use.Bclass2ID.Valid = true
		} else {
			this.aclassbclass2use.Bclass2ID.Int64 = 0
			this.aclassbclass2use.Bclass2ID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.aclassbclass2use.Aclass_Anarrayofb2Use_reverse != undefined) {
			if (this.aclassbclass2use.Aclass_Anarrayofb2UseDBID == undefined) {
				this.aclassbclass2use.Aclass_Anarrayofb2UseDBID = new NullInt64
			}
			this.aclassbclass2use.Aclass_Anarrayofb2UseDBID.Int64 = this.aclassbclass2use.Aclass_Anarrayofb2Use_reverse.ID
			this.aclassbclass2use.Aclass_Anarrayofb2UseDBID.Valid = true
			if (this.aclassbclass2use.Aclass_Anarrayofb2UseDBID_Index == undefined) {
				this.aclassbclass2use.Aclass_Anarrayofb2UseDBID_Index = new NullInt64
			}
			this.aclassbclass2use.Aclass_Anarrayofb2UseDBID_Index.Valid = true
			this.aclassbclass2use.Aclass_Anarrayofb2Use_reverse = undefined // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case AclassBclass2UseDetailComponentState.UPDATE_INSTANCE:
				this.aclassbclass2useService.updateAclassBclass2Use(this.aclassbclass2use)
					.subscribe(aclassbclass2use => {
						this.aclassbclass2useService.AclassBclass2UseServiceChanged.next("update")
					});
				break;
			default:
				this.aclassbclass2useService.postAclassBclass2Use(this.aclassbclass2use).subscribe(aclassbclass2use => {
					this.aclassbclass2useService.AclassBclass2UseServiceChanged.next("post")
					this.aclassbclass2use = {} // reset fields
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

			dialogData.ID = this.aclassbclass2use.ID
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
			dialogData.ID = this.aclassbclass2use.ID
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "AclassBclass2Use"
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
			ID: this.aclassbclass2use.ID,
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
		if (this.aclassbclass2use.Name == undefined) {
			this.aclassbclass2use.Name = event.value.Name
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
