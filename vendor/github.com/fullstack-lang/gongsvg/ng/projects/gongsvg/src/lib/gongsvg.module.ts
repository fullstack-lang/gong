import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatSortModule } from '@angular/material/sort'
import { MatPaginatorModule } from '@angular/material/paginator'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTreeModule } from '@angular/material/tree';
import { DragDropModule } from '@angular/cdk/drag-drop';

import { AngularSplitModule, SplitComponent } from 'angular-split';

import {
	NgxMatDatetimePickerModule,
	NgxMatNativeDateModule,
	NgxMatTimepickerModule
} from '@angular-material-components/datetime-picker';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent } from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';
import { GongstructSelectionService } from './gongstruct-selection.service'

// insertion point for imports 
import { AnimatesTableComponent } from './animates-table/animates-table.component'
import { AnimateSortingComponent } from './animate-sorting/animate-sorting.component'
import { AnimateDetailComponent } from './animate-detail/animate-detail.component'

import { CirclesTableComponent } from './circles-table/circles-table.component'
import { CircleSortingComponent } from './circle-sorting/circle-sorting.component'
import { CircleDetailComponent } from './circle-detail/circle-detail.component'

import { EllipsesTableComponent } from './ellipses-table/ellipses-table.component'
import { EllipseSortingComponent } from './ellipse-sorting/ellipse-sorting.component'
import { EllipseDetailComponent } from './ellipse-detail/ellipse-detail.component'

import { LayersTableComponent } from './layers-table/layers-table.component'
import { LayerSortingComponent } from './layer-sorting/layer-sorting.component'
import { LayerDetailComponent } from './layer-detail/layer-detail.component'

import { LinesTableComponent } from './lines-table/lines-table.component'
import { LineSortingComponent } from './line-sorting/line-sorting.component'
import { LineDetailComponent } from './line-detail/line-detail.component'

import { LinksTableComponent } from './links-table/links-table.component'
import { LinkSortingComponent } from './link-sorting/link-sorting.component'
import { LinkDetailComponent } from './link-detail/link-detail.component'

import { LinkAnchoredTextsTableComponent } from './linkanchoredtexts-table/linkanchoredtexts-table.component'
import { LinkAnchoredTextSortingComponent } from './linkanchoredtext-sorting/linkanchoredtext-sorting.component'
import { LinkAnchoredTextDetailComponent } from './linkanchoredtext-detail/linkanchoredtext-detail.component'

import { PathsTableComponent } from './paths-table/paths-table.component'
import { PathSortingComponent } from './path-sorting/path-sorting.component'
import { PathDetailComponent } from './path-detail/path-detail.component'

import { PointsTableComponent } from './points-table/points-table.component'
import { PointSortingComponent } from './point-sorting/point-sorting.component'
import { PointDetailComponent } from './point-detail/point-detail.component'

import { PolygonesTableComponent } from './polygones-table/polygones-table.component'
import { PolygoneSortingComponent } from './polygone-sorting/polygone-sorting.component'
import { PolygoneDetailComponent } from './polygone-detail/polygone-detail.component'

import { PolylinesTableComponent } from './polylines-table/polylines-table.component'
import { PolylineSortingComponent } from './polyline-sorting/polyline-sorting.component'
import { PolylineDetailComponent } from './polyline-detail/polyline-detail.component'

import { RectsTableComponent } from './rects-table/rects-table.component'
import { RectSortingComponent } from './rect-sorting/rect-sorting.component'
import { RectDetailComponent } from './rect-detail/rect-detail.component'

import { RectAnchoredRectsTableComponent } from './rectanchoredrects-table/rectanchoredrects-table.component'
import { RectAnchoredRectSortingComponent } from './rectanchoredrect-sorting/rectanchoredrect-sorting.component'
import { RectAnchoredRectDetailComponent } from './rectanchoredrect-detail/rectanchoredrect-detail.component'

import { RectAnchoredTextsTableComponent } from './rectanchoredtexts-table/rectanchoredtexts-table.component'
import { RectAnchoredTextSortingComponent } from './rectanchoredtext-sorting/rectanchoredtext-sorting.component'
import { RectAnchoredTextDetailComponent } from './rectanchoredtext-detail/rectanchoredtext-detail.component'

import { RectLinkLinksTableComponent } from './rectlinklinks-table/rectlinklinks-table.component'
import { RectLinkLinkSortingComponent } from './rectlinklink-sorting/rectlinklink-sorting.component'
import { RectLinkLinkDetailComponent } from './rectlinklink-detail/rectlinklink-detail.component'

import { SVGsTableComponent } from './svgs-table/svgs-table.component'
import { SVGSortingComponent } from './svg-sorting/svg-sorting.component'
import { SVGDetailComponent } from './svg-detail/svg-detail.component'

import { TextsTableComponent } from './texts-table/texts-table.component'
import { TextSortingComponent } from './text-sorting/text-sorting.component'
import { TextDetailComponent } from './text-detail/text-detail.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		AnimatesTableComponent,
		AnimateSortingComponent,
		AnimateDetailComponent,

		CirclesTableComponent,
		CircleSortingComponent,
		CircleDetailComponent,

		EllipsesTableComponent,
		EllipseSortingComponent,
		EllipseDetailComponent,

		LayersTableComponent,
		LayerSortingComponent,
		LayerDetailComponent,

		LinesTableComponent,
		LineSortingComponent,
		LineDetailComponent,

		LinksTableComponent,
		LinkSortingComponent,
		LinkDetailComponent,

		LinkAnchoredTextsTableComponent,
		LinkAnchoredTextSortingComponent,
		LinkAnchoredTextDetailComponent,

		PathsTableComponent,
		PathSortingComponent,
		PathDetailComponent,

		PointsTableComponent,
		PointSortingComponent,
		PointDetailComponent,

		PolygonesTableComponent,
		PolygoneSortingComponent,
		PolygoneDetailComponent,

		PolylinesTableComponent,
		PolylineSortingComponent,
		PolylineDetailComponent,

		RectsTableComponent,
		RectSortingComponent,
		RectDetailComponent,

		RectAnchoredRectsTableComponent,
		RectAnchoredRectSortingComponent,
		RectAnchoredRectDetailComponent,

		RectAnchoredTextsTableComponent,
		RectAnchoredTextSortingComponent,
		RectAnchoredTextDetailComponent,

		RectLinkLinksTableComponent,
		RectLinkLinkSortingComponent,
		RectLinkLinkDetailComponent,

		SVGsTableComponent,
		SVGSortingComponent,
		SVGDetailComponent,

		TextsTableComponent,
		TextSortingComponent,
		TextDetailComponent,


		SplitterComponent,
		SidebarComponent
	],
	imports: [
		FormsModule,
		ReactiveFormsModule,
		CommonModule,
		RouterModule,

		AppRoutingModule,

		MatSliderModule,
		MatSelectModule,
		MatFormFieldModule,
		MatInputModule,
		MatDatepickerModule,
		MatTableModule,
		MatSortModule,
		MatPaginatorModule,
		MatCheckboxModule,
		MatButtonModule,
		MatIconModule,
		MatToolbarModule,
		MatExpansionModule,
		MatListModule,
		MatDialogModule,
		MatGridListModule,
		MatTreeModule,
		DragDropModule,

		NgxMatDatetimePickerModule,
		NgxMatNativeDateModule,
		NgxMatTimepickerModule,

		AngularSplitModule,
	],
	exports: [
		// insertion point for declarations 
		AnimatesTableComponent,
		AnimateSortingComponent,
		AnimateDetailComponent,

		CirclesTableComponent,
		CircleSortingComponent,
		CircleDetailComponent,

		EllipsesTableComponent,
		EllipseSortingComponent,
		EllipseDetailComponent,

		LayersTableComponent,
		LayerSortingComponent,
		LayerDetailComponent,

		LinesTableComponent,
		LineSortingComponent,
		LineDetailComponent,

		LinksTableComponent,
		LinkSortingComponent,
		LinkDetailComponent,

		LinkAnchoredTextsTableComponent,
		LinkAnchoredTextSortingComponent,
		LinkAnchoredTextDetailComponent,

		PathsTableComponent,
		PathSortingComponent,
		PathDetailComponent,

		PointsTableComponent,
		PointSortingComponent,
		PointDetailComponent,

		PolygonesTableComponent,
		PolygoneSortingComponent,
		PolygoneDetailComponent,

		PolylinesTableComponent,
		PolylineSortingComponent,
		PolylineDetailComponent,

		RectsTableComponent,
		RectSortingComponent,
		RectDetailComponent,

		RectAnchoredRectsTableComponent,
		RectAnchoredRectSortingComponent,
		RectAnchoredRectDetailComponent,

		RectAnchoredTextsTableComponent,
		RectAnchoredTextSortingComponent,
		RectAnchoredTextDetailComponent,

		RectLinkLinksTableComponent,
		RectLinkLinkSortingComponent,
		RectLinkLinkDetailComponent,

		SVGsTableComponent,
		SVGSortingComponent,
		SVGDetailComponent,

		TextsTableComponent,
		TextSortingComponent,
		TextDetailComponent,


		SplitterComponent,
		SidebarComponent,

	],
	providers: [
		GongstructSelectionService,
		{
			provide: MatDialogRef,
			useValue: {}
		},
	],
})
export class GongsvgModule { }
