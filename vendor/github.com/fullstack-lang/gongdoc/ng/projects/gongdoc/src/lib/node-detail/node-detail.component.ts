// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { NodeDB } from '../node-db'
import { NodeService } from '../node.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { TreeDB } from '../tree-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// NodeDetailComponent is initilizaed from different routes
// NodeDetailComponentState detail different cases 
enum NodeDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Node_Children_SET,
	CREATE_INSTANCE_WITH_ASSOCIATION_Tree_RootNodes_SET,
}

@Component({
	selector: 'app-node-detail',
	templateUrl: './node-detail.component.html',
	styleUrls: ['./node-detail.component.css'],
})
export class NodeDetailComponent implements OnInit {

	// insertion point for declarations
	IsExpandedFormControl: UntypedFormControl = new UntypedFormControl(false);
	HasCheckboxButtonFormControl: UntypedFormControl = new UntypedFormControl(false);
	IsCheckedFormControl: UntypedFormControl = new UntypedFormControl(false);
	IsCheckboxDisabledFormControl: UntypedFormControl = new UntypedFormControl(false);
	IsInEditModeFormControl: UntypedFormControl = new UntypedFormControl(false);

	// the NodeDB of interest
	node: NodeDB = new NodeDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: NodeDetailComponentState = NodeDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private nodeService: NodeService,
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
			this.state = NodeDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = NodeDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "Children":
						// console.log("Node" + " is instanciated with back pointer to instance " + this.id + " Node association Children")
						this.state = NodeDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Node_Children_SET
						break;
					case "RootNodes":
						// console.log("Node" + " is instanciated with back pointer to instance " + this.id + " Tree association RootNodes")
						this.state = NodeDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Tree_RootNodes_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getNode()

		// observable for changes in structs
		this.nodeService.NodeServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getNode()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getNode(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case NodeDetailComponentState.CREATE_INSTANCE:
						this.node = new (NodeDB)
						break;
					case NodeDetailComponentState.UPDATE_INSTANCE:
						let node = frontRepo.Nodes.get(this.id)
						console.assert(node != undefined, "missing node with id:" + this.id)
						this.node = node!
						break;
					// insertion point for init of association field
					case NodeDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Node_Children_SET:
						this.node = new (NodeDB)
						this.node.Node_Children_reverse = frontRepo.Nodes.get(this.id)!
						break;
					case NodeDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Tree_RootNodes_SET:
						this.node = new (NodeDB)
						this.node.Tree_RootNodes_reverse = frontRepo.Trees.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
				this.IsExpandedFormControl.setValue(this.node.IsExpanded)
				this.HasCheckboxButtonFormControl.setValue(this.node.HasCheckboxButton)
				this.IsCheckedFormControl.setValue(this.node.IsChecked)
				this.IsCheckboxDisabledFormControl.setValue(this.node.IsCheckboxDisabled)
				this.IsInEditModeFormControl.setValue(this.node.IsInEditMode)
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		this.node.IsExpanded = this.IsExpandedFormControl.value
		this.node.HasCheckboxButton = this.HasCheckboxButtonFormControl.value
		this.node.IsChecked = this.IsCheckedFormControl.value
		this.node.IsCheckboxDisabled = this.IsCheckboxDisabledFormControl.value
		this.node.IsInEditMode = this.IsInEditModeFormControl.value

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.node.Node_Children_reverse != undefined) {
			if (this.node.Node_ChildrenDBID == undefined) {
				this.node.Node_ChildrenDBID = new NullInt64
			}
			this.node.Node_ChildrenDBID.Int64 = this.node.Node_Children_reverse.ID
			this.node.Node_ChildrenDBID.Valid = true
			if (this.node.Node_ChildrenDBID_Index == undefined) {
				this.node.Node_ChildrenDBID_Index = new NullInt64
			}
			this.node.Node_ChildrenDBID_Index.Valid = true
			this.node.Node_Children_reverse = new NodeDB // very important, otherwise, circular JSON
		}
		if (this.node.Tree_RootNodes_reverse != undefined) {
			if (this.node.Tree_RootNodesDBID == undefined) {
				this.node.Tree_RootNodesDBID = new NullInt64
			}
			this.node.Tree_RootNodesDBID.Int64 = this.node.Tree_RootNodes_reverse.ID
			this.node.Tree_RootNodesDBID.Valid = true
			if (this.node.Tree_RootNodesDBID_Index == undefined) {
				this.node.Tree_RootNodesDBID_Index = new NullInt64
			}
			this.node.Tree_RootNodesDBID_Index.Valid = true
			this.node.Tree_RootNodes_reverse = new TreeDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case NodeDetailComponentState.UPDATE_INSTANCE:
				this.nodeService.updateNode(this.node, this.GONG__StackPath)
					.subscribe(node => {
						this.nodeService.NodeServiceChanged.next("update")
					});
				break;
			default:
				this.nodeService.postNode(this.node, this.GONG__StackPath).subscribe(node => {
					this.nodeService.NodeServiceChanged.next("post")
					this.node = new (NodeDB) // reset fields
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

			dialogData.ID = this.node.ID!
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
			dialogData.ID = this.node.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "Node"
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
			ID: this.node.ID,
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
		if (this.node.Name == "") {
			this.node.Name = event.value.Name
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
