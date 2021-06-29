// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { AclassBclassUseDB } from '../aclassbclassuse-db'
import { AclassBclassUseService } from '../aclassbclassuse.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

// AclassBclassUseDetailComponent is initilizaed from different routes
// AclassBclassUseDetailComponentState detail different cases 
enum AclassBclassUseDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Aclass_AnarrayofbUse_SET,
}

@Component({
	selector: 'app-aclassbclassuse-detail',
	templateUrl: './aclassbclassuse-detail.component.html',
	styleUrls: ['./aclassbclassuse-detail.component.css'],
})
export class AclassBclassUseDetailComponent implements OnInit {

	// insertion point for declarations

	// the AclassBclassUseDB of interest
	aclassbclassuse: AclassBclassUseDB;

	// front repo
	frontRepo: FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: AclassBclassUseDetailComponentState

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string
	originStructFieldName: string

	constructor(
		private aclassbclassuseService: AclassBclassUseService,
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
			this.state = AclassBclassUseDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = AclassBclassUseDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "AnarrayofbUse":
						console.log("AclassBclassUse" + " is instanciated with back pointer to instance " + this.id + " Aclass association AnarrayofbUse")
						this.state = AclassBclassUseDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Aclass_AnarrayofbUse_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getAclassBclassUse()

		// observable for changes in structs
		this.aclassbclassuseService.AclassBclassUseServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getAclassBclassUse()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getAclassBclassUse(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case AclassBclassUseDetailComponentState.CREATE_INSTANCE:
						this.aclassbclassuse = new (AclassBclassUseDB)
						break;
					case AclassBclassUseDetailComponentState.UPDATE_INSTANCE:
						this.aclassbclassuse = frontRepo.AclassBclassUses.get(this.id)
						break;
					// insertion point for init of association field
					case AclassBclassUseDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Aclass_AnarrayofbUse_SET:
						this.aclassbclassuse = new (AclassBclassUseDB)
						this.aclassbclassuse.Aclass_AnarrayofbUse_reverse = frontRepo.Aclasss.get(this.id)
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
		if (this.aclassbclassuse.Bclass2ID == undefined) {
			this.aclassbclassuse.Bclass2ID = new NullInt64
		}
		if (this.aclassbclassuse.Bclass2 != undefined) {
			this.aclassbclassuse.Bclass2ID.Int64 = this.aclassbclassuse.Bclass2.ID
			this.aclassbclassuse.Bclass2ID.Valid = true
		} else {
			this.aclassbclassuse.Bclass2ID.Int64 = 0
			this.aclassbclassuse.Bclass2ID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.aclassbclassuse.Aclass_AnarrayofbUse_reverse != undefined) {
			if (this.aclassbclassuse.Aclass_AnarrayofbUseDBID == undefined) {
				this.aclassbclassuse.Aclass_AnarrayofbUseDBID = new NullInt64
			}
			this.aclassbclassuse.Aclass_AnarrayofbUseDBID.Int64 = this.aclassbclassuse.Aclass_AnarrayofbUse_reverse.ID
			this.aclassbclassuse.Aclass_AnarrayofbUseDBID.Valid = true
			if (this.aclassbclassuse.Aclass_AnarrayofbUseDBID_Index == undefined) {
				this.aclassbclassuse.Aclass_AnarrayofbUseDBID_Index = new NullInt64
			}
			this.aclassbclassuse.Aclass_AnarrayofbUseDBID_Index.Valid = true
			this.aclassbclassuse.Aclass_AnarrayofbUse_reverse = undefined // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case AclassBclassUseDetailComponentState.UPDATE_INSTANCE:
				this.aclassbclassuseService.updateAclassBclassUse(this.aclassbclassuse)
					.subscribe(aclassbclassuse => {
						this.aclassbclassuseService.AclassBclassUseServiceChanged.next("update")
					});
				break;
			default:
				this.aclassbclassuseService.postAclassBclassUse(this.aclassbclassuse).subscribe(aclassbclassuse => {
					this.aclassbclassuseService.AclassBclassUseServiceChanged.next("post")
					this.aclassbclassuse = {} // reset fields
				});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string, selectionMode: SelectionMode,
		sourceField: string, intermediateStructField: string, nextAssociatedStruct: string ) {

		console.log("mode " + selectionMode)

		const dialogConfig = new MatDialogConfig();

		let dialogData = new DialogData();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		if (selectionMode == SelectionMode.ONE_MANY_ASSOCIATION_MODE) {

			dialogData.ID = this.aclassbclassuse.ID
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
			dialogData.ID = this.aclassbclassuse.ID
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "AclassBclassUse"
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
			ID: this.aclassbclassuse.ID,
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
		if (this.aclassbclassuse.Name == undefined) {
			this.aclassbclassuse.Name = event.value.Name
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
