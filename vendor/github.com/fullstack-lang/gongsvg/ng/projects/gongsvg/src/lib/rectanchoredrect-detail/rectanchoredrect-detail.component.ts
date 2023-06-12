// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { RectAnchoredRectDB } from '../rectanchoredrect-db'
import { RectAnchoredRectService } from '../rectanchoredrect.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { RectAnchorTypeSelect, RectAnchorTypeList } from '../RectAnchorType'
import { RectDB } from '../rect-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// RectAnchoredRectDetailComponent is initilizaed from different routes
// RectAnchoredRectDetailComponentState detail different cases 
enum RectAnchoredRectDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Rect_RectAnchoredRects_SET,
}

@Component({
	selector: 'app-rectanchoredrect-detail',
	templateUrl: './rectanchoredrect-detail.component.html',
	styleUrls: ['./rectanchoredrect-detail.component.css'],
})
export class RectAnchoredRectDetailComponent implements OnInit {

	// insertion point for declarations
	RectAnchorTypeList: RectAnchorTypeSelect[] = []
	WidthFollowRectFormControl: UntypedFormControl = new UntypedFormControl(false);
	HeightFollowRectFormControl: UntypedFormControl = new UntypedFormControl(false);

	// the RectAnchoredRectDB of interest
	rectanchoredrect: RectAnchoredRectDB = new RectAnchoredRectDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: RectAnchoredRectDetailComponentState = RectAnchoredRectDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private rectanchoredrectService: RectAnchoredRectService,
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
			this.state = RectAnchoredRectDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = RectAnchoredRectDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "RectAnchoredRects":
						// console.log("RectAnchoredRect" + " is instanciated with back pointer to instance " + this.id + " Rect association RectAnchoredRects")
						this.state = RectAnchoredRectDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Rect_RectAnchoredRects_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getRectAnchoredRect()

		// observable for changes in structs
		this.rectanchoredrectService.RectAnchoredRectServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getRectAnchoredRect()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.RectAnchorTypeList = RectAnchorTypeList
	}

	getRectAnchoredRect(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case RectAnchoredRectDetailComponentState.CREATE_INSTANCE:
						this.rectanchoredrect = new (RectAnchoredRectDB)
						break;
					case RectAnchoredRectDetailComponentState.UPDATE_INSTANCE:
						let rectanchoredrect = frontRepo.RectAnchoredRects.get(this.id)
						console.assert(rectanchoredrect != undefined, "missing rectanchoredrect with id:" + this.id)
						this.rectanchoredrect = rectanchoredrect!
						break;
					// insertion point for init of association field
					case RectAnchoredRectDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Rect_RectAnchoredRects_SET:
						this.rectanchoredrect = new (RectAnchoredRectDB)
						this.rectanchoredrect.Rect_RectAnchoredRects_reverse = frontRepo.Rects.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
				this.WidthFollowRectFormControl.setValue(this.rectanchoredrect.WidthFollowRect)
				this.HeightFollowRectFormControl.setValue(this.rectanchoredrect.HeightFollowRect)
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		this.rectanchoredrect.WidthFollowRect = this.WidthFollowRectFormControl.value
		this.rectanchoredrect.HeightFollowRect = this.HeightFollowRectFormControl.value

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.rectanchoredrect.Rect_RectAnchoredRects_reverse != undefined) {
			if (this.rectanchoredrect.Rect_RectAnchoredRectsDBID == undefined) {
				this.rectanchoredrect.Rect_RectAnchoredRectsDBID = new NullInt64
			}
			this.rectanchoredrect.Rect_RectAnchoredRectsDBID.Int64 = this.rectanchoredrect.Rect_RectAnchoredRects_reverse.ID
			this.rectanchoredrect.Rect_RectAnchoredRectsDBID.Valid = true
			if (this.rectanchoredrect.Rect_RectAnchoredRectsDBID_Index == undefined) {
				this.rectanchoredrect.Rect_RectAnchoredRectsDBID_Index = new NullInt64
			}
			this.rectanchoredrect.Rect_RectAnchoredRectsDBID_Index.Valid = true
			this.rectanchoredrect.Rect_RectAnchoredRects_reverse = new RectDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case RectAnchoredRectDetailComponentState.UPDATE_INSTANCE:
				this.rectanchoredrectService.updateRectAnchoredRect(this.rectanchoredrect, this.GONG__StackPath)
					.subscribe(rectanchoredrect => {
						this.rectanchoredrectService.RectAnchoredRectServiceChanged.next("update")
					});
				break;
			default:
				this.rectanchoredrectService.postRectAnchoredRect(this.rectanchoredrect, this.GONG__StackPath).subscribe(rectanchoredrect => {
					this.rectanchoredrectService.RectAnchoredRectServiceChanged.next("post")
					this.rectanchoredrect = new (RectAnchoredRectDB) // reset fields
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

			dialogData.ID = this.rectanchoredrect.ID!
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
			dialogData.ID = this.rectanchoredrect.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "RectAnchoredRect"
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
			ID: this.rectanchoredrect.ID,
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
		if (this.rectanchoredrect.Name == "") {
			this.rectanchoredrect.Name = event.value.Name
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
