// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { AnimateDB } from '../animate-db'
import { AnimateService } from '../animate.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { CircleDB } from '../circle-db'
import { EllipseDB } from '../ellipse-db'
import { LineDB } from '../line-db'
import { LinkAnchoredTextDB } from '../linkanchoredtext-db'
import { PathDB } from '../path-db'
import { PolygoneDB } from '../polygone-db'
import { PolylineDB } from '../polyline-db'
import { RectDB } from '../rect-db'
import { RectAnchoredTextDB } from '../rectanchoredtext-db'
import { TextDB } from '../text-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// AnimateDetailComponent is initilizaed from different routes
// AnimateDetailComponentState detail different cases 
enum AnimateDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Circle_Animations_SET,
	CREATE_INSTANCE_WITH_ASSOCIATION_Ellipse_Animates_SET,
	CREATE_INSTANCE_WITH_ASSOCIATION_Line_Animates_SET,
	CREATE_INSTANCE_WITH_ASSOCIATION_LinkAnchoredText_Animates_SET,
	CREATE_INSTANCE_WITH_ASSOCIATION_Path_Animates_SET,
	CREATE_INSTANCE_WITH_ASSOCIATION_Polygone_Animates_SET,
	CREATE_INSTANCE_WITH_ASSOCIATION_Polyline_Animates_SET,
	CREATE_INSTANCE_WITH_ASSOCIATION_Rect_Animations_SET,
	CREATE_INSTANCE_WITH_ASSOCIATION_RectAnchoredText_Animates_SET,
	CREATE_INSTANCE_WITH_ASSOCIATION_Text_Animates_SET,
}

@Component({
	selector: 'app-animate-detail',
	templateUrl: './animate-detail.component.html',
	styleUrls: ['./animate-detail.component.css'],
})
export class AnimateDetailComponent implements OnInit {

	// insertion point for declarations

	// the AnimateDB of interest
	animate: AnimateDB = new AnimateDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: AnimateDetailComponentState = AnimateDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private animateService: AnimateService,
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
			this.state = AnimateDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = AnimateDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "Animations":
						// console.log("Animate" + " is instanciated with back pointer to instance " + this.id + " Circle association Animations")
						this.state = AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Circle_Animations_SET
						break;
					case "Animates":
						// console.log("Animate" + " is instanciated with back pointer to instance " + this.id + " Ellipse association Animates")
						this.state = AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Ellipse_Animates_SET
						break;
					case "Animates":
						// console.log("Animate" + " is instanciated with back pointer to instance " + this.id + " Line association Animates")
						this.state = AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Line_Animates_SET
						break;
					case "Animates":
						// console.log("Animate" + " is instanciated with back pointer to instance " + this.id + " LinkAnchoredText association Animates")
						this.state = AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_LinkAnchoredText_Animates_SET
						break;
					case "Animates":
						// console.log("Animate" + " is instanciated with back pointer to instance " + this.id + " Path association Animates")
						this.state = AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Path_Animates_SET
						break;
					case "Animates":
						// console.log("Animate" + " is instanciated with back pointer to instance " + this.id + " Polygone association Animates")
						this.state = AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Polygone_Animates_SET
						break;
					case "Animates":
						// console.log("Animate" + " is instanciated with back pointer to instance " + this.id + " Polyline association Animates")
						this.state = AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Polyline_Animates_SET
						break;
					case "Animations":
						// console.log("Animate" + " is instanciated with back pointer to instance " + this.id + " Rect association Animations")
						this.state = AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Rect_Animations_SET
						break;
					case "Animates":
						// console.log("Animate" + " is instanciated with back pointer to instance " + this.id + " RectAnchoredText association Animates")
						this.state = AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_RectAnchoredText_Animates_SET
						break;
					case "Animates":
						// console.log("Animate" + " is instanciated with back pointer to instance " + this.id + " Text association Animates")
						this.state = AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Text_Animates_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getAnimate()

		// observable for changes in structs
		this.animateService.AnimateServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getAnimate()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getAnimate(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case AnimateDetailComponentState.CREATE_INSTANCE:
						this.animate = new (AnimateDB)
						break;
					case AnimateDetailComponentState.UPDATE_INSTANCE:
						let animate = frontRepo.Animates.get(this.id)
						console.assert(animate != undefined, "missing animate with id:" + this.id)
						this.animate = animate!
						break;
					// insertion point for init of association field
					case AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Circle_Animations_SET:
						this.animate = new (AnimateDB)
						this.animate.Circle_Animations_reverse = frontRepo.Circles.get(this.id)!
						break;
					case AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Ellipse_Animates_SET:
						this.animate = new (AnimateDB)
						this.animate.Ellipse_Animates_reverse = frontRepo.Ellipses.get(this.id)!
						break;
					case AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Line_Animates_SET:
						this.animate = new (AnimateDB)
						this.animate.Line_Animates_reverse = frontRepo.Lines.get(this.id)!
						break;
					case AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_LinkAnchoredText_Animates_SET:
						this.animate = new (AnimateDB)
						this.animate.LinkAnchoredText_Animates_reverse = frontRepo.LinkAnchoredTexts.get(this.id)!
						break;
					case AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Path_Animates_SET:
						this.animate = new (AnimateDB)
						this.animate.Path_Animates_reverse = frontRepo.Paths.get(this.id)!
						break;
					case AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Polygone_Animates_SET:
						this.animate = new (AnimateDB)
						this.animate.Polygone_Animates_reverse = frontRepo.Polygones.get(this.id)!
						break;
					case AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Polyline_Animates_SET:
						this.animate = new (AnimateDB)
						this.animate.Polyline_Animates_reverse = frontRepo.Polylines.get(this.id)!
						break;
					case AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Rect_Animations_SET:
						this.animate = new (AnimateDB)
						this.animate.Rect_Animations_reverse = frontRepo.Rects.get(this.id)!
						break;
					case AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_RectAnchoredText_Animates_SET:
						this.animate = new (AnimateDB)
						this.animate.RectAnchoredText_Animates_reverse = frontRepo.RectAnchoredTexts.get(this.id)!
						break;
					case AnimateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Text_Animates_SET:
						this.animate = new (AnimateDB)
						this.animate.Text_Animates_reverse = frontRepo.Texts.get(this.id)!
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
		if (this.animate.Circle_Animations_reverse != undefined) {
			if (this.animate.Circle_AnimationsDBID == undefined) {
				this.animate.Circle_AnimationsDBID = new NullInt64
			}
			this.animate.Circle_AnimationsDBID.Int64 = this.animate.Circle_Animations_reverse.ID
			this.animate.Circle_AnimationsDBID.Valid = true
			if (this.animate.Circle_AnimationsDBID_Index == undefined) {
				this.animate.Circle_AnimationsDBID_Index = new NullInt64
			}
			this.animate.Circle_AnimationsDBID_Index.Valid = true
			this.animate.Circle_Animations_reverse = new CircleDB // very important, otherwise, circular JSON
		}
		if (this.animate.Ellipse_Animates_reverse != undefined) {
			if (this.animate.Ellipse_AnimatesDBID == undefined) {
				this.animate.Ellipse_AnimatesDBID = new NullInt64
			}
			this.animate.Ellipse_AnimatesDBID.Int64 = this.animate.Ellipse_Animates_reverse.ID
			this.animate.Ellipse_AnimatesDBID.Valid = true
			if (this.animate.Ellipse_AnimatesDBID_Index == undefined) {
				this.animate.Ellipse_AnimatesDBID_Index = new NullInt64
			}
			this.animate.Ellipse_AnimatesDBID_Index.Valid = true
			this.animate.Ellipse_Animates_reverse = new EllipseDB // very important, otherwise, circular JSON
		}
		if (this.animate.Line_Animates_reverse != undefined) {
			if (this.animate.Line_AnimatesDBID == undefined) {
				this.animate.Line_AnimatesDBID = new NullInt64
			}
			this.animate.Line_AnimatesDBID.Int64 = this.animate.Line_Animates_reverse.ID
			this.animate.Line_AnimatesDBID.Valid = true
			if (this.animate.Line_AnimatesDBID_Index == undefined) {
				this.animate.Line_AnimatesDBID_Index = new NullInt64
			}
			this.animate.Line_AnimatesDBID_Index.Valid = true
			this.animate.Line_Animates_reverse = new LineDB // very important, otherwise, circular JSON
		}
		if (this.animate.LinkAnchoredText_Animates_reverse != undefined) {
			if (this.animate.LinkAnchoredText_AnimatesDBID == undefined) {
				this.animate.LinkAnchoredText_AnimatesDBID = new NullInt64
			}
			this.animate.LinkAnchoredText_AnimatesDBID.Int64 = this.animate.LinkAnchoredText_Animates_reverse.ID
			this.animate.LinkAnchoredText_AnimatesDBID.Valid = true
			if (this.animate.LinkAnchoredText_AnimatesDBID_Index == undefined) {
				this.animate.LinkAnchoredText_AnimatesDBID_Index = new NullInt64
			}
			this.animate.LinkAnchoredText_AnimatesDBID_Index.Valid = true
			this.animate.LinkAnchoredText_Animates_reverse = new LinkAnchoredTextDB // very important, otherwise, circular JSON
		}
		if (this.animate.Path_Animates_reverse != undefined) {
			if (this.animate.Path_AnimatesDBID == undefined) {
				this.animate.Path_AnimatesDBID = new NullInt64
			}
			this.animate.Path_AnimatesDBID.Int64 = this.animate.Path_Animates_reverse.ID
			this.animate.Path_AnimatesDBID.Valid = true
			if (this.animate.Path_AnimatesDBID_Index == undefined) {
				this.animate.Path_AnimatesDBID_Index = new NullInt64
			}
			this.animate.Path_AnimatesDBID_Index.Valid = true
			this.animate.Path_Animates_reverse = new PathDB // very important, otherwise, circular JSON
		}
		if (this.animate.Polygone_Animates_reverse != undefined) {
			if (this.animate.Polygone_AnimatesDBID == undefined) {
				this.animate.Polygone_AnimatesDBID = new NullInt64
			}
			this.animate.Polygone_AnimatesDBID.Int64 = this.animate.Polygone_Animates_reverse.ID
			this.animate.Polygone_AnimatesDBID.Valid = true
			if (this.animate.Polygone_AnimatesDBID_Index == undefined) {
				this.animate.Polygone_AnimatesDBID_Index = new NullInt64
			}
			this.animate.Polygone_AnimatesDBID_Index.Valid = true
			this.animate.Polygone_Animates_reverse = new PolygoneDB // very important, otherwise, circular JSON
		}
		if (this.animate.Polyline_Animates_reverse != undefined) {
			if (this.animate.Polyline_AnimatesDBID == undefined) {
				this.animate.Polyline_AnimatesDBID = new NullInt64
			}
			this.animate.Polyline_AnimatesDBID.Int64 = this.animate.Polyline_Animates_reverse.ID
			this.animate.Polyline_AnimatesDBID.Valid = true
			if (this.animate.Polyline_AnimatesDBID_Index == undefined) {
				this.animate.Polyline_AnimatesDBID_Index = new NullInt64
			}
			this.animate.Polyline_AnimatesDBID_Index.Valid = true
			this.animate.Polyline_Animates_reverse = new PolylineDB // very important, otherwise, circular JSON
		}
		if (this.animate.Rect_Animations_reverse != undefined) {
			if (this.animate.Rect_AnimationsDBID == undefined) {
				this.animate.Rect_AnimationsDBID = new NullInt64
			}
			this.animate.Rect_AnimationsDBID.Int64 = this.animate.Rect_Animations_reverse.ID
			this.animate.Rect_AnimationsDBID.Valid = true
			if (this.animate.Rect_AnimationsDBID_Index == undefined) {
				this.animate.Rect_AnimationsDBID_Index = new NullInt64
			}
			this.animate.Rect_AnimationsDBID_Index.Valid = true
			this.animate.Rect_Animations_reverse = new RectDB // very important, otherwise, circular JSON
		}
		if (this.animate.RectAnchoredText_Animates_reverse != undefined) {
			if (this.animate.RectAnchoredText_AnimatesDBID == undefined) {
				this.animate.RectAnchoredText_AnimatesDBID = new NullInt64
			}
			this.animate.RectAnchoredText_AnimatesDBID.Int64 = this.animate.RectAnchoredText_Animates_reverse.ID
			this.animate.RectAnchoredText_AnimatesDBID.Valid = true
			if (this.animate.RectAnchoredText_AnimatesDBID_Index == undefined) {
				this.animate.RectAnchoredText_AnimatesDBID_Index = new NullInt64
			}
			this.animate.RectAnchoredText_AnimatesDBID_Index.Valid = true
			this.animate.RectAnchoredText_Animates_reverse = new RectAnchoredTextDB // very important, otherwise, circular JSON
		}
		if (this.animate.Text_Animates_reverse != undefined) {
			if (this.animate.Text_AnimatesDBID == undefined) {
				this.animate.Text_AnimatesDBID = new NullInt64
			}
			this.animate.Text_AnimatesDBID.Int64 = this.animate.Text_Animates_reverse.ID
			this.animate.Text_AnimatesDBID.Valid = true
			if (this.animate.Text_AnimatesDBID_Index == undefined) {
				this.animate.Text_AnimatesDBID_Index = new NullInt64
			}
			this.animate.Text_AnimatesDBID_Index.Valid = true
			this.animate.Text_Animates_reverse = new TextDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case AnimateDetailComponentState.UPDATE_INSTANCE:
				this.animateService.updateAnimate(this.animate, this.GONG__StackPath)
					.subscribe(animate => {
						this.animateService.AnimateServiceChanged.next("update")
					});
				break;
			default:
				this.animateService.postAnimate(this.animate, this.GONG__StackPath).subscribe(animate => {
					this.animateService.AnimateServiceChanged.next("post")
					this.animate = new (AnimateDB) // reset fields
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

			dialogData.ID = this.animate.ID!
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
			dialogData.ID = this.animate.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "Animate"
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
			ID: this.animate.ID,
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
		if (this.animate.Name == "") {
			this.animate.Name = event.value.Name
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
