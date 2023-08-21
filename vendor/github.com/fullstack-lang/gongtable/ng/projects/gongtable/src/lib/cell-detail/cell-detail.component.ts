// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { CellDB } from '../cell-db'
import { CellService } from '../cell.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { RowDB } from '../row-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// CellDetailComponent is initilizaed from different routes
// CellDetailComponentState detail different cases 
enum CellDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Row_Cells_SET,
}

@Component({
	selector: 'app-cell-detail',
	templateUrl: './cell-detail.component.html',
	styleUrls: ['./cell-detail.component.css'],
})
export class CellDetailComponent implements OnInit {

	// insertion point for declarations

	// the CellDB of interest
	cell: CellDB = new CellDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: CellDetailComponentState = CellDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private cellService: CellService,
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
			this.state = CellDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = CellDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "Cells":
						// console.log("Cell" + " is instanciated with back pointer to instance " + this.id + " Row association Cells")
						this.state = CellDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Row_Cells_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getCell()

		// observable for changes in structs
		this.cellService.CellServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getCell()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getCell(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case CellDetailComponentState.CREATE_INSTANCE:
						this.cell = new (CellDB)
						break;
					case CellDetailComponentState.UPDATE_INSTANCE:
						let cell = frontRepo.Cells.get(this.id)
						console.assert(cell != undefined, "missing cell with id:" + this.id)
						this.cell = cell!
						break;
					// insertion point for init of association field
					case CellDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Row_Cells_SET:
						this.cell = new (CellDB)
						this.cell.Row_Cells_reverse = frontRepo.Rows.get(this.id)!
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
		if (this.cell.CellStringID == undefined) {
			this.cell.CellStringID = new NullInt64
		}
		if (this.cell.CellString != undefined) {
			this.cell.CellStringID.Int64 = this.cell.CellString.ID
			this.cell.CellStringID.Valid = true
		} else {
			this.cell.CellStringID.Int64 = 0
			this.cell.CellStringID.Valid = true
		}
		if (this.cell.CellFloat64ID == undefined) {
			this.cell.CellFloat64ID = new NullInt64
		}
		if (this.cell.CellFloat64 != undefined) {
			this.cell.CellFloat64ID.Int64 = this.cell.CellFloat64.ID
			this.cell.CellFloat64ID.Valid = true
		} else {
			this.cell.CellFloat64ID.Int64 = 0
			this.cell.CellFloat64ID.Valid = true
		}
		if (this.cell.CellIntID == undefined) {
			this.cell.CellIntID = new NullInt64
		}
		if (this.cell.CellInt != undefined) {
			this.cell.CellIntID.Int64 = this.cell.CellInt.ID
			this.cell.CellIntID.Valid = true
		} else {
			this.cell.CellIntID.Int64 = 0
			this.cell.CellIntID.Valid = true
		}
		if (this.cell.CellBoolID == undefined) {
			this.cell.CellBoolID = new NullInt64
		}
		if (this.cell.CellBool != undefined) {
			this.cell.CellBoolID.Int64 = this.cell.CellBool.ID
			this.cell.CellBoolID.Valid = true
		} else {
			this.cell.CellBoolID.Int64 = 0
			this.cell.CellBoolID.Valid = true
		}
		if (this.cell.CellIconID == undefined) {
			this.cell.CellIconID = new NullInt64
		}
		if (this.cell.CellIcon != undefined) {
			this.cell.CellIconID.Int64 = this.cell.CellIcon.ID
			this.cell.CellIconID.Valid = true
		} else {
			this.cell.CellIconID.Int64 = 0
			this.cell.CellIconID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.cell.Row_Cells_reverse != undefined) {
			if (this.cell.Row_CellsDBID == undefined) {
				this.cell.Row_CellsDBID = new NullInt64
			}
			this.cell.Row_CellsDBID.Int64 = this.cell.Row_Cells_reverse.ID
			this.cell.Row_CellsDBID.Valid = true
			if (this.cell.Row_CellsDBID_Index == undefined) {
				this.cell.Row_CellsDBID_Index = new NullInt64
			}
			this.cell.Row_CellsDBID_Index.Valid = true
			this.cell.Row_Cells_reverse = new RowDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case CellDetailComponentState.UPDATE_INSTANCE:
				this.cellService.updateCell(this.cell, this.GONG__StackPath)
					.subscribe(cell => {
						this.cellService.CellServiceChanged.next("update")
					});
				break;
			default:
				this.cellService.postCell(this.cell, this.GONG__StackPath).subscribe(cell => {
					this.cellService.CellServiceChanged.next("post")
					this.cell = new (CellDB) // reset fields
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

			dialogData.ID = this.cell.ID!
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
			dialogData.ID = this.cell.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "Cell"
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
			ID: this.cell.ID,
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
		if (this.cell.Name == "") {
			this.cell.Name = event.value.Name
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
