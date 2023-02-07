// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { GongStructShapeDB } from '../gongstructshape-db'
import { GongStructShapeService } from '../gongstructshape.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { ClassdiagramDB } from '../classdiagram-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// GongStructShapeDetailComponent is initilizaed from different routes
// GongStructShapeDetailComponentState detail different cases 
enum GongStructShapeDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Classdiagram_GongStructShapes_SET,
}

@Component({
	selector: 'app-gongstructshape-detail',
	templateUrl: './gongstructshape-detail.component.html',
	styleUrls: ['./gongstructshape-detail.component.css'],
})
export class GongStructShapeDetailComponent implements OnInit {

	// insertion point for declarations
	ShowNbInstancesFormControl: UntypedFormControl = new UntypedFormControl(false);
	IsSelectedFormControl: UntypedFormControl = new UntypedFormControl(false);

	// the GongStructShapeDB of interest
	gongstructshape: GongStructShapeDB = new GongStructShapeDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: GongStructShapeDetailComponentState = GongStructShapeDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	constructor(
		private gongstructshapeService: GongStructShapeService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private activatedRoute: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.activatedRoute.params.subscribe(params => {
			this.onChangedActivatedRoute()
		});
	}
	onChangedActivatedRoute(): void {

		// compute state
		this.id = +this.activatedRoute.snapshot.paramMap.get('id')!;
		this.originStruct = this.activatedRoute.snapshot.paramMap.get('originStruct')!;
		this.originStructFieldName = this.activatedRoute.snapshot.paramMap.get('originStructFieldName')!;

		const association = this.activatedRoute.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = GongStructShapeDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = GongStructShapeDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "GongStructShapes":
						// console.log("GongStructShape" + " is instanciated with back pointer to instance " + this.id + " Classdiagram association GongStructShapes")
						this.state = GongStructShapeDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Classdiagram_GongStructShapes_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getGongStructShape()

		// observable for changes in structs
		this.gongstructshapeService.GongStructShapeServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getGongStructShape()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getGongStructShape(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case GongStructShapeDetailComponentState.CREATE_INSTANCE:
						this.gongstructshape = new (GongStructShapeDB)
						break;
					case GongStructShapeDetailComponentState.UPDATE_INSTANCE:
						let gongstructshape = frontRepo.GongStructShapes.get(this.id)
						console.assert(gongstructshape != undefined, "missing gongstructshape with id:" + this.id)
						this.gongstructshape = gongstructshape!
						break;
					// insertion point for init of association field
					case GongStructShapeDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Classdiagram_GongStructShapes_SET:
						this.gongstructshape = new (GongStructShapeDB)
						this.gongstructshape.Classdiagram_GongStructShapes_reverse = frontRepo.Classdiagrams.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
				this.ShowNbInstancesFormControl.setValue(this.gongstructshape.ShowNbInstances)
				this.IsSelectedFormControl.setValue(this.gongstructshape.IsSelected)
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		if (this.gongstructshape.PositionID == undefined) {
			this.gongstructshape.PositionID = new NullInt64
		}
		if (this.gongstructshape.Position != undefined) {
			this.gongstructshape.PositionID.Int64 = this.gongstructshape.Position.ID
			this.gongstructshape.PositionID.Valid = true
		} else {
			this.gongstructshape.PositionID.Int64 = 0
			this.gongstructshape.PositionID.Valid = true
		}
		this.gongstructshape.ShowNbInstances = this.ShowNbInstancesFormControl.value
		this.gongstructshape.IsSelected = this.IsSelectedFormControl.value

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.gongstructshape.Classdiagram_GongStructShapes_reverse != undefined) {
			if (this.gongstructshape.Classdiagram_GongStructShapesDBID == undefined) {
				this.gongstructshape.Classdiagram_GongStructShapesDBID = new NullInt64
			}
			this.gongstructshape.Classdiagram_GongStructShapesDBID.Int64 = this.gongstructshape.Classdiagram_GongStructShapes_reverse.ID
			this.gongstructshape.Classdiagram_GongStructShapesDBID.Valid = true
			if (this.gongstructshape.Classdiagram_GongStructShapesDBID_Index == undefined) {
				this.gongstructshape.Classdiagram_GongStructShapesDBID_Index = new NullInt64
			}
			this.gongstructshape.Classdiagram_GongStructShapesDBID_Index.Valid = true
			this.gongstructshape.Classdiagram_GongStructShapes_reverse = new ClassdiagramDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case GongStructShapeDetailComponentState.UPDATE_INSTANCE:
				this.gongstructshapeService.updateGongStructShape(this.gongstructshape)
					.subscribe(gongstructshape => {
						this.gongstructshapeService.GongStructShapeServiceChanged.next("update")
					});
				break;
			default:
				this.gongstructshapeService.postGongStructShape(this.gongstructshape).subscribe(gongstructshape => {
					this.gongstructshapeService.GongStructShapeServiceChanged.next("post")
					this.gongstructshape = new (GongStructShapeDB) // reset fields
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

			dialogData.ID = this.gongstructshape.ID!
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
			dialogData.ID = this.gongstructshape.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "GongStructShape"
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
			ID: this.gongstructshape.ID,
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
		if (this.gongstructshape.Name == "") {
			this.gongstructshape.Name = event.value.Name
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
