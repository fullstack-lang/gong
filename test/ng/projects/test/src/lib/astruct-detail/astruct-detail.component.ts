// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { AstructDB } from '../astruct-db'
import { AstructService } from '../astruct.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { AEnumTypeSelect, AEnumTypeList } from '../AEnumType'
import { BEnumTypeSelect, BEnumTypeList } from '../BEnumType'
import { CEnumTypeIntSelect, CEnumTypeIntList } from '../CEnumTypeInt'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// AstructDetailComponent is initilizaed from different routes
// AstructDetailComponentState detail different cases 
enum AstructDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Astruct_Anarrayofa_SET,
}

@Component({
	selector: 'app-astruct-detail',
	templateUrl: './astruct-detail.component.html',
	styleUrls: ['./astruct-detail.component.css'],
})
export class AstructDetailComponent implements OnInit {

	// insertion point for declarations
	BooleanfieldFormControl = new FormControl(false);
	AEnumTypeList: AEnumTypeSelect[] = []
	BEnumTypeList: BEnumTypeSelect[] = []
	CEnumTypeIntList: CEnumTypeIntSelect[] = []
	AnotherbooleanfieldFormControl = new FormControl(false);
	Duration1_Hours: number = 0
	Duration1_Minutes: number = 0
	Duration1_Seconds: number = 0

	// the AstructDB of interest
	astruct: AstructDB = new AstructDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: AstructDetailComponentState = AstructDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	constructor(
		private astructService: AstructService,
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
			this.state = AstructDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = AstructDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "Anarrayofa":
						// console.log("Astruct" + " is instanciated with back pointer to instance " + this.id + " Astruct association Anarrayofa")
						this.state = AstructDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Astruct_Anarrayofa_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getAstruct()

		// observable for changes in structs
		this.astructService.AstructServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getAstruct()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.AEnumTypeList = AEnumTypeList
		this.BEnumTypeList = BEnumTypeList
		this.CEnumTypeIntList = CEnumTypeIntList
	}

	getAstruct(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case AstructDetailComponentState.CREATE_INSTANCE:
						this.astruct = new (AstructDB)
						break;
					case AstructDetailComponentState.UPDATE_INSTANCE:
						let astruct = frontRepo.Astructs.get(this.id)
						console.assert(astruct != undefined, "missing astruct with id:" + this.id)
						this.astruct = astruct!
						break;
					// insertion point for init of association field
					case AstructDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Astruct_Anarrayofa_SET:
						this.astruct = new (AstructDB)
						this.astruct.Astruct_Anarrayofa_reverse = frontRepo.Astructs.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
				this.BooleanfieldFormControl.setValue(this.astruct.Booleanfield)
				this.AnotherbooleanfieldFormControl.setValue(this.astruct.Anotherbooleanfield)
				// computation of Hours, Minutes, Seconds for Duration1
				this.Duration1_Hours = Math.floor(this.astruct.Duration1 / (3600 * 1000 * 1000 * 1000))
				this.Duration1_Minutes = Math.floor(this.astruct.Duration1 % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.Duration1_Seconds = this.astruct.Duration1 % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		this.astruct.Booleanfield = this.BooleanfieldFormControl.value
		this.astruct.Anotherbooleanfield = this.AnotherbooleanfieldFormControl.value
		this.astruct.Duration1 =
			this.Duration1_Hours * (3600 * 1000 * 1000 * 1000) +
			this.Duration1_Minutes * (60 * 1000 * 1000 * 1000) +
			this.Duration1_Seconds * (1000 * 1000 * 1000)
		if (this.astruct.AssociationtobID == undefined) {
			this.astruct.AssociationtobID = new NullInt64
		}
		if (this.astruct.Associationtob != undefined) {
			this.astruct.AssociationtobID.Int64 = this.astruct.Associationtob.ID
			this.astruct.AssociationtobID.Valid = true
		} else {
			this.astruct.AssociationtobID.Int64 = 0
			this.astruct.AssociationtobID.Valid = true
		}
		if (this.astruct.Anotherassociationtob_2ID == undefined) {
			this.astruct.Anotherassociationtob_2ID = new NullInt64
		}
		if (this.astruct.Anotherassociationtob_2 != undefined) {
			this.astruct.Anotherassociationtob_2ID.Int64 = this.astruct.Anotherassociationtob_2.ID
			this.astruct.Anotherassociationtob_2ID.Valid = true
		} else {
			this.astruct.Anotherassociationtob_2ID.Int64 = 0
			this.astruct.Anotherassociationtob_2ID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.astruct.Astruct_Anarrayofa_reverse != undefined) {
			if (this.astruct.Astruct_AnarrayofaDBID == undefined) {
				this.astruct.Astruct_AnarrayofaDBID = new NullInt64
			}
			this.astruct.Astruct_AnarrayofaDBID.Int64 = this.astruct.Astruct_Anarrayofa_reverse.ID
			this.astruct.Astruct_AnarrayofaDBID.Valid = true
			if (this.astruct.Astruct_AnarrayofaDBID_Index == undefined) {
				this.astruct.Astruct_AnarrayofaDBID_Index = new NullInt64
			}
			this.astruct.Astruct_AnarrayofaDBID_Index.Valid = true
			this.astruct.Astruct_Anarrayofa_reverse = new AstructDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case AstructDetailComponentState.UPDATE_INSTANCE:
				this.astructService.updateAstruct(this.astruct)
					.subscribe(astruct => {
						this.astructService.AstructServiceChanged.next("update")
					});
				break;
			default:
				this.astructService.postAstruct(this.astruct).subscribe(astruct => {
					this.astructService.AstructServiceChanged.next("post")
					this.astruct = new (AstructDB) // reset fields
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

			dialogData.ID = this.astruct.ID!
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
			dialogData.ID = this.astruct.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "Astruct"
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
			ID: this.astruct.ID,
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
		if (this.astruct.Name == undefined) {
			this.astruct.Name = event.value.Name
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
