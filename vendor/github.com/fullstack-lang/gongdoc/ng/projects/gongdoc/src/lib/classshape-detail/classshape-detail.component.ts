// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { ClassshapeDB } from '../classshape-db'
import { ClassshapeService } from '../classshape.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { ClassshapeTargetTypeSelect, ClassshapeTargetTypeList } from '../ClassshapeTargetType'
import { ClassdiagramDB } from '../classdiagram-db'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// ClassshapeDetailComponent is initilizaed from different routes
// ClassshapeDetailComponentState detail different cases 
enum ClassshapeDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Classdiagram_Classshapes_SET,
}

@Component({
	selector: 'app-classshape-detail',
	templateUrl: './classshape-detail.component.html',
	styleUrls: ['./classshape-detail.component.css'],
})
export class ClassshapeDetailComponent implements OnInit {

	// insertion point for declarations
	ShowNbInstancesFormControl: UntypedFormControl = new UntypedFormControl(false);
	ClassshapeTargetTypeList: ClassshapeTargetTypeSelect[] = []

	// the ClassshapeDB of interest
	classshape: ClassshapeDB = new ClassshapeDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: ClassshapeDetailComponentState = ClassshapeDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	constructor(
		private classshapeService: ClassshapeService,
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
			this.state = ClassshapeDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = ClassshapeDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "Classshapes":
						// console.log("Classshape" + " is instanciated with back pointer to instance " + this.id + " Classdiagram association Classshapes")
						this.state = ClassshapeDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Classdiagram_Classshapes_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getClassshape()

		// observable for changes in structs
		this.classshapeService.ClassshapeServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getClassshape()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.ClassshapeTargetTypeList = ClassshapeTargetTypeList
	}

	getClassshape(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case ClassshapeDetailComponentState.CREATE_INSTANCE:
						this.classshape = new (ClassshapeDB)
						break;
					case ClassshapeDetailComponentState.UPDATE_INSTANCE:
						let classshape = frontRepo.Classshapes.get(this.id)
						console.assert(classshape != undefined, "missing classshape with id:" + this.id)
						this.classshape = classshape!
						break;
					// insertion point for init of association field
					case ClassshapeDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Classdiagram_Classshapes_SET:
						this.classshape = new (ClassshapeDB)
						this.classshape.Classdiagram_Classshapes_reverse = frontRepo.Classdiagrams.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
				this.ShowNbInstancesFormControl.setValue(this.classshape.ShowNbInstances)
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		if (this.classshape.PositionID == undefined) {
			this.classshape.PositionID = new NullInt64
		}
		if (this.classshape.Position != undefined) {
			this.classshape.PositionID.Int64 = this.classshape.Position.ID
			this.classshape.PositionID.Valid = true
		} else {
			this.classshape.PositionID.Int64 = 0
			this.classshape.PositionID.Valid = true
		}
		if (this.classshape.GongStructID == undefined) {
			this.classshape.GongStructID = new NullInt64
		}
		if (this.classshape.GongStruct != undefined) {
			this.classshape.GongStructID.Int64 = this.classshape.GongStruct.ID
			this.classshape.GongStructID.Valid = true
		} else {
			this.classshape.GongStructID.Int64 = 0
			this.classshape.GongStructID.Valid = true
		}
		this.classshape.ShowNbInstances = this.ShowNbInstancesFormControl.value

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.classshape.Classdiagram_Classshapes_reverse != undefined) {
			if (this.classshape.Classdiagram_ClassshapesDBID == undefined) {
				this.classshape.Classdiagram_ClassshapesDBID = new NullInt64
			}
			this.classshape.Classdiagram_ClassshapesDBID.Int64 = this.classshape.Classdiagram_Classshapes_reverse.ID
			this.classshape.Classdiagram_ClassshapesDBID.Valid = true
			if (this.classshape.Classdiagram_ClassshapesDBID_Index == undefined) {
				this.classshape.Classdiagram_ClassshapesDBID_Index = new NullInt64
			}
			this.classshape.Classdiagram_ClassshapesDBID_Index.Valid = true
			this.classshape.Classdiagram_Classshapes_reverse = new ClassdiagramDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case ClassshapeDetailComponentState.UPDATE_INSTANCE:
				this.classshapeService.updateClassshape(this.classshape)
					.subscribe(classshape => {
						this.classshapeService.ClassshapeServiceChanged.next("update")
					});
				break;
			default:
				this.classshapeService.postClassshape(this.classshape).subscribe(classshape => {
					this.classshapeService.ClassshapeServiceChanged.next("post")
					this.classshape = new (ClassshapeDB) // reset fields
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

			dialogData.ID = this.classshape.ID!
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
			dialogData.ID = this.classshape.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "Classshape"
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
			ID: this.classshape.ID,
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
		if (this.classshape.Name == "") {
			this.classshape.Name = event.value.Name
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
