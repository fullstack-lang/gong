// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { RectDB } from '../rect-db'
import { RectService } from '../rect.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { LayerDB } from '../layer-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// RectDetailComponent is initilizaed from different routes
// RectDetailComponentState detail different cases 
enum RectDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Layer_Rects_SET,
}

@Component({
	selector: 'app-rect-detail',
	templateUrl: './rect-detail.component.html',
	styleUrls: ['./rect-detail.component.css'],
})
export class RectDetailComponent implements OnInit {

	// insertion point for declarations
	IsSelectableFormControl: UntypedFormControl = new UntypedFormControl(false);
	IsSelectedFormControl: UntypedFormControl = new UntypedFormControl(false);
	CanHaveLeftHandleFormControl: UntypedFormControl = new UntypedFormControl(false);
	HasLeftHandleFormControl: UntypedFormControl = new UntypedFormControl(false);
	CanHaveRightHandleFormControl: UntypedFormControl = new UntypedFormControl(false);
	HasRightHandleFormControl: UntypedFormControl = new UntypedFormControl(false);
	CanHaveTopHandleFormControl: UntypedFormControl = new UntypedFormControl(false);
	HasTopHandleFormControl: UntypedFormControl = new UntypedFormControl(false);
	CanHaveBottomHandleFormControl: UntypedFormControl = new UntypedFormControl(false);
	HasBottomHandleFormControl: UntypedFormControl = new UntypedFormControl(false);
	CanMoveHorizontalyFormControl: UntypedFormControl = new UntypedFormControl(false);
	CanMoveVerticalyFormControl: UntypedFormControl = new UntypedFormControl(false);

	// the RectDB of interest
	rect: RectDB = new RectDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: RectDetailComponentState = RectDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private rectService: RectService,
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
			this.state = RectDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = RectDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "Rects":
						// console.log("Rect" + " is instanciated with back pointer to instance " + this.id + " Layer association Rects")
						this.state = RectDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Layer_Rects_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getRect()

		// observable for changes in structs
		this.rectService.RectServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getRect()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getRect(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case RectDetailComponentState.CREATE_INSTANCE:
						this.rect = new (RectDB)
						break;
					case RectDetailComponentState.UPDATE_INSTANCE:
						let rect = frontRepo.Rects.get(this.id)
						console.assert(rect != undefined, "missing rect with id:" + this.id)
						this.rect = rect!
						break;
					// insertion point for init of association field
					case RectDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Layer_Rects_SET:
						this.rect = new (RectDB)
						this.rect.Layer_Rects_reverse = frontRepo.Layers.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
				this.IsSelectableFormControl.setValue(this.rect.IsSelectable)
				this.IsSelectedFormControl.setValue(this.rect.IsSelected)
				this.CanHaveLeftHandleFormControl.setValue(this.rect.CanHaveLeftHandle)
				this.HasLeftHandleFormControl.setValue(this.rect.HasLeftHandle)
				this.CanHaveRightHandleFormControl.setValue(this.rect.CanHaveRightHandle)
				this.HasRightHandleFormControl.setValue(this.rect.HasRightHandle)
				this.CanHaveTopHandleFormControl.setValue(this.rect.CanHaveTopHandle)
				this.HasTopHandleFormControl.setValue(this.rect.HasTopHandle)
				this.CanHaveBottomHandleFormControl.setValue(this.rect.CanHaveBottomHandle)
				this.HasBottomHandleFormControl.setValue(this.rect.HasBottomHandle)
				this.CanMoveHorizontalyFormControl.setValue(this.rect.CanMoveHorizontaly)
				this.CanMoveVerticalyFormControl.setValue(this.rect.CanMoveVerticaly)
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		this.rect.IsSelectable = this.IsSelectableFormControl.value
		this.rect.IsSelected = this.IsSelectedFormControl.value
		this.rect.CanHaveLeftHandle = this.CanHaveLeftHandleFormControl.value
		this.rect.HasLeftHandle = this.HasLeftHandleFormControl.value
		this.rect.CanHaveRightHandle = this.CanHaveRightHandleFormControl.value
		this.rect.HasRightHandle = this.HasRightHandleFormControl.value
		this.rect.CanHaveTopHandle = this.CanHaveTopHandleFormControl.value
		this.rect.HasTopHandle = this.HasTopHandleFormControl.value
		this.rect.CanHaveBottomHandle = this.CanHaveBottomHandleFormControl.value
		this.rect.HasBottomHandle = this.HasBottomHandleFormControl.value
		this.rect.CanMoveHorizontaly = this.CanMoveHorizontalyFormControl.value
		this.rect.CanMoveVerticaly = this.CanMoveVerticalyFormControl.value

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.rect.Layer_Rects_reverse != undefined) {
			if (this.rect.Layer_RectsDBID == undefined) {
				this.rect.Layer_RectsDBID = new NullInt64
			}
			this.rect.Layer_RectsDBID.Int64 = this.rect.Layer_Rects_reverse.ID
			this.rect.Layer_RectsDBID.Valid = true
			if (this.rect.Layer_RectsDBID_Index == undefined) {
				this.rect.Layer_RectsDBID_Index = new NullInt64
			}
			this.rect.Layer_RectsDBID_Index.Valid = true
			this.rect.Layer_Rects_reverse = new LayerDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case RectDetailComponentState.UPDATE_INSTANCE:
				this.rectService.updateRect(this.rect, this.GONG__StackPath)
					.subscribe(rect => {
						this.rectService.RectServiceChanged.next("update")
					});
				break;
			default:
				this.rectService.postRect(this.rect, this.GONG__StackPath).subscribe(rect => {
					this.rectService.RectServiceChanged.next("post")
					this.rect = new (RectDB) // reset fields
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

			dialogData.ID = this.rect.ID!
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
			dialogData.ID = this.rect.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "Rect"
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
			ID: this.rect.ID,
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
		if (this.rect.Name == "") {
			this.rect.Name = event.value.Name
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
