// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { TableDB } from '../table-db'
import { TableService } from '../table.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// TableDetailComponent is initilizaed from different routes
// TableDetailComponentState detail different cases 
enum TableDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
}

@Component({
	selector: 'app-table-detail',
	templateUrl: './table-detail.component.html',
	styleUrls: ['./table-detail.component.css'],
})
export class TableDetailComponent implements OnInit {

	// insertion point for declarations
	HasFilteringFormControl: UntypedFormControl = new UntypedFormControl(false);
	HasColumnSortingFormControl: UntypedFormControl = new UntypedFormControl(false);
	HasPaginatorFormControl: UntypedFormControl = new UntypedFormControl(false);
	HasCheckableRowsFormControl: UntypedFormControl = new UntypedFormControl(false);
	HasSaveButtonFormControl: UntypedFormControl = new UntypedFormControl(false);
	CanDragDropRowsFormControl: UntypedFormControl = new UntypedFormControl(false);
	HasCloseButtonFormControl: UntypedFormControl = new UntypedFormControl(false);
	SavingInProgressFormControl: UntypedFormControl = new UntypedFormControl(false);

	// the TableDB of interest
	table: TableDB = new TableDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: TableDetailComponentState = TableDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private tableService: TableService,
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
			this.state = TableDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = TableDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getTable()

		// observable for changes in structs
		this.tableService.TableServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getTable()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getTable(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case TableDetailComponentState.CREATE_INSTANCE:
						this.table = new (TableDB)
						break;
					case TableDetailComponentState.UPDATE_INSTANCE:
						let table = frontRepo.Tables.get(this.id)
						console.assert(table != undefined, "missing table with id:" + this.id)
						this.table = table!
						break;
					// insertion point for init of association field
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
				this.HasFilteringFormControl.setValue(this.table.HasFiltering)
				this.HasColumnSortingFormControl.setValue(this.table.HasColumnSorting)
				this.HasPaginatorFormControl.setValue(this.table.HasPaginator)
				this.HasCheckableRowsFormControl.setValue(this.table.HasCheckableRows)
				this.HasSaveButtonFormControl.setValue(this.table.HasSaveButton)
				this.CanDragDropRowsFormControl.setValue(this.table.CanDragDropRows)
				this.HasCloseButtonFormControl.setValue(this.table.HasCloseButton)
				this.SavingInProgressFormControl.setValue(this.table.SavingInProgress)
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		this.table.HasFiltering = this.HasFilteringFormControl.value
		this.table.HasColumnSorting = this.HasColumnSortingFormControl.value
		this.table.HasPaginator = this.HasPaginatorFormControl.value
		this.table.HasCheckableRows = this.HasCheckableRowsFormControl.value
		this.table.HasSaveButton = this.HasSaveButtonFormControl.value
		this.table.CanDragDropRows = this.CanDragDropRowsFormControl.value
		this.table.HasCloseButton = this.HasCloseButtonFormControl.value
		this.table.SavingInProgress = this.SavingInProgressFormControl.value

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers

		switch (this.state) {
			case TableDetailComponentState.UPDATE_INSTANCE:
				this.tableService.updateTable(this.table, this.GONG__StackPath)
					.subscribe(table => {
						this.tableService.TableServiceChanged.next("update")
					});
				break;
			default:
				this.tableService.postTable(this.table, this.GONG__StackPath).subscribe(table => {
					this.tableService.TableServiceChanged.next("post")
					this.table = new (TableDB) // reset fields
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

			dialogData.ID = this.table.ID!
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
			dialogData.ID = this.table.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "Table"
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
			ID: this.table.ID,
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
		if (this.table.Name == "") {
			this.table.Name = event.value.Name
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
