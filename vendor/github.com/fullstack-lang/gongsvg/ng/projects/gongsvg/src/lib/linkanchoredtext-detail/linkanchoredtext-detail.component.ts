// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { LinkAnchoredTextDB } from '../linkanchoredtext-db'
import { LinkAnchoredTextService } from '../linkanchoredtext.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { LinkDB } from '../link-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// LinkAnchoredTextDetailComponent is initilizaed from different routes
// LinkAnchoredTextDetailComponentState detail different cases 
enum LinkAnchoredTextDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Link_TextAtArrowEnd_SET,
	CREATE_INSTANCE_WITH_ASSOCIATION_Link_TextAtArrowStart_SET,
}

@Component({
	selector: 'app-linkanchoredtext-detail',
	templateUrl: './linkanchoredtext-detail.component.html',
	styleUrls: ['./linkanchoredtext-detail.component.css'],
})
export class LinkAnchoredTextDetailComponent implements OnInit {

	// insertion point for declarations

	// the LinkAnchoredTextDB of interest
	linkanchoredtext: LinkAnchoredTextDB = new LinkAnchoredTextDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: LinkAnchoredTextDetailComponentState = LinkAnchoredTextDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private linkanchoredtextService: LinkAnchoredTextService,
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
			this.state = LinkAnchoredTextDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = LinkAnchoredTextDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "TextAtArrowEnd":
						// console.log("LinkAnchoredText" + " is instanciated with back pointer to instance " + this.id + " Link association TextAtArrowEnd")
						this.state = LinkAnchoredTextDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Link_TextAtArrowEnd_SET
						break;
					case "TextAtArrowStart":
						// console.log("LinkAnchoredText" + " is instanciated with back pointer to instance " + this.id + " Link association TextAtArrowStart")
						this.state = LinkAnchoredTextDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Link_TextAtArrowStart_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getLinkAnchoredText()

		// observable for changes in structs
		this.linkanchoredtextService.LinkAnchoredTextServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getLinkAnchoredText()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getLinkAnchoredText(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case LinkAnchoredTextDetailComponentState.CREATE_INSTANCE:
						this.linkanchoredtext = new (LinkAnchoredTextDB)
						break;
					case LinkAnchoredTextDetailComponentState.UPDATE_INSTANCE:
						let linkanchoredtext = frontRepo.LinkAnchoredTexts.get(this.id)
						console.assert(linkanchoredtext != undefined, "missing linkanchoredtext with id:" + this.id)
						this.linkanchoredtext = linkanchoredtext!
						break;
					// insertion point for init of association field
					case LinkAnchoredTextDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Link_TextAtArrowEnd_SET:
						this.linkanchoredtext = new (LinkAnchoredTextDB)
						this.linkanchoredtext.Link_TextAtArrowEnd_reverse = frontRepo.Links.get(this.id)!
						break;
					case LinkAnchoredTextDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Link_TextAtArrowStart_SET:
						this.linkanchoredtext = new (LinkAnchoredTextDB)
						this.linkanchoredtext.Link_TextAtArrowStart_reverse = frontRepo.Links.get(this.id)!
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
		if (this.linkanchoredtext.Link_TextAtArrowEnd_reverse != undefined) {
			if (this.linkanchoredtext.Link_TextAtArrowEndDBID == undefined) {
				this.linkanchoredtext.Link_TextAtArrowEndDBID = new NullInt64
			}
			this.linkanchoredtext.Link_TextAtArrowEndDBID.Int64 = this.linkanchoredtext.Link_TextAtArrowEnd_reverse.ID
			this.linkanchoredtext.Link_TextAtArrowEndDBID.Valid = true
			if (this.linkanchoredtext.Link_TextAtArrowEndDBID_Index == undefined) {
				this.linkanchoredtext.Link_TextAtArrowEndDBID_Index = new NullInt64
			}
			this.linkanchoredtext.Link_TextAtArrowEndDBID_Index.Valid = true
			this.linkanchoredtext.Link_TextAtArrowEnd_reverse = new LinkDB // very important, otherwise, circular JSON
		}
		if (this.linkanchoredtext.Link_TextAtArrowStart_reverse != undefined) {
			if (this.linkanchoredtext.Link_TextAtArrowStartDBID == undefined) {
				this.linkanchoredtext.Link_TextAtArrowStartDBID = new NullInt64
			}
			this.linkanchoredtext.Link_TextAtArrowStartDBID.Int64 = this.linkanchoredtext.Link_TextAtArrowStart_reverse.ID
			this.linkanchoredtext.Link_TextAtArrowStartDBID.Valid = true
			if (this.linkanchoredtext.Link_TextAtArrowStartDBID_Index == undefined) {
				this.linkanchoredtext.Link_TextAtArrowStartDBID_Index = new NullInt64
			}
			this.linkanchoredtext.Link_TextAtArrowStartDBID_Index.Valid = true
			this.linkanchoredtext.Link_TextAtArrowStart_reverse = new LinkDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case LinkAnchoredTextDetailComponentState.UPDATE_INSTANCE:
				this.linkanchoredtextService.updateLinkAnchoredText(this.linkanchoredtext, this.GONG__StackPath)
					.subscribe(linkanchoredtext => {
						this.linkanchoredtextService.LinkAnchoredTextServiceChanged.next("update")
					});
				break;
			default:
				this.linkanchoredtextService.postLinkAnchoredText(this.linkanchoredtext, this.GONG__StackPath).subscribe(linkanchoredtext => {
					this.linkanchoredtextService.LinkAnchoredTextServiceChanged.next("post")
					this.linkanchoredtext = new (LinkAnchoredTextDB) // reset fields
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

			dialogData.ID = this.linkanchoredtext.ID!
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
			dialogData.ID = this.linkanchoredtext.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "LinkAnchoredText"
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
			ID: this.linkanchoredtext.ID,
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
		if (this.linkanchoredtext.Name == "") {
			this.linkanchoredtext.Name = event.value.Name
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
