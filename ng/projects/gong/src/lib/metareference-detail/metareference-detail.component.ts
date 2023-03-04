// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { MetaReferenceDB } from '../metareference-db'
import { MetaReferenceService } from '../metareference.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { MetaDB } from '../meta-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// MetaReferenceDetailComponent is initilizaed from different routes
// MetaReferenceDetailComponentState detail different cases 
enum MetaReferenceDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Meta_MetaReferences_SET,
}

@Component({
	selector: 'app-metareference-detail',
	templateUrl: './metareference-detail.component.html',
	styleUrls: ['./metareference-detail.component.css'],
})
export class MetaReferenceDetailComponent implements OnInit {

	// insertion point for declarations

	// the MetaReferenceDB of interest
	metareference: MetaReferenceDB = new MetaReferenceDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: MetaReferenceDetailComponentState = MetaReferenceDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private metareferenceService: MetaReferenceService,
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
			this.state = MetaReferenceDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = MetaReferenceDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "MetaReferences":
						// console.log("MetaReference" + " is instanciated with back pointer to instance " + this.id + " Meta association MetaReferences")
						this.state = MetaReferenceDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Meta_MetaReferences_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getMetaReference()

		// observable for changes in structs
		this.metareferenceService.MetaReferenceServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getMetaReference()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getMetaReference(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case MetaReferenceDetailComponentState.CREATE_INSTANCE:
						this.metareference = new (MetaReferenceDB)
						break;
					case MetaReferenceDetailComponentState.UPDATE_INSTANCE:
						let metareference = frontRepo.MetaReferences.get(this.id)
						console.assert(metareference != undefined, "missing metareference with id:" + this.id)
						this.metareference = metareference!
						break;
					// insertion point for init of association field
					case MetaReferenceDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Meta_MetaReferences_SET:
						this.metareference = new (MetaReferenceDB)
						this.metareference.Meta_MetaReferences_reverse = frontRepo.Metas.get(this.id)!
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

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.metareference.Meta_MetaReferences_reverse != undefined) {
			if (this.metareference.Meta_MetaReferencesDBID == undefined) {
				this.metareference.Meta_MetaReferencesDBID = new NullInt64
			}
			this.metareference.Meta_MetaReferencesDBID.Int64 = this.metareference.Meta_MetaReferences_reverse.ID
			this.metareference.Meta_MetaReferencesDBID.Valid = true
			if (this.metareference.Meta_MetaReferencesDBID_Index == undefined) {
				this.metareference.Meta_MetaReferencesDBID_Index = new NullInt64
			}
			this.metareference.Meta_MetaReferencesDBID_Index.Valid = true
			this.metareference.Meta_MetaReferences_reverse = new MetaDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case MetaReferenceDetailComponentState.UPDATE_INSTANCE:
				this.metareferenceService.updateMetaReference(this.metareference, this.GONG__StackPath)
					.subscribe(metareference => {
						this.metareferenceService.MetaReferenceServiceChanged.next("update")
					});
				break;
			default:
				this.metareferenceService.postMetaReference(this.metareference, this.GONG__StackPath).subscribe(metareference => {
					this.metareferenceService.MetaReferenceServiceChanged.next("post")
					this.metareference = new (MetaReferenceDB) // reset fields
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

			dialogData.ID = this.metareference.ID!
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
			dialogData.ID = this.metareference.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "MetaReference"
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
			ID: this.metareference.ID,
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
		if (this.metareference.Name == "") {
			this.metareference.Name = event.value.Name
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
