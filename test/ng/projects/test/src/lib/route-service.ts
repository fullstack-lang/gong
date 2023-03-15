import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
import { AstructsTableComponent } from './astructs-table/astructs-table.component'
import { AstructDetailComponent } from './astruct-detail/astruct-detail.component'

import { AstructBstruct2UsesTableComponent } from './astructbstruct2uses-table/astructbstruct2uses-table.component'
import { AstructBstruct2UseDetailComponent } from './astructbstruct2use-detail/astructbstruct2use-detail.component'

import { AstructBstructUsesTableComponent } from './astructbstructuses-table/astructbstructuses-table.component'
import { AstructBstructUseDetailComponent } from './astructbstructuse-detail/astructbstructuse-detail.component'

import { BstructsTableComponent } from './bstructs-table/bstructs-table.component'
import { BstructDetailComponent } from './bstruct-detail/bstruct-detail.component'

import { DstructsTableComponent } from './dstructs-table/dstructs-table.component'
import { DstructDetailComponent } from './dstruct-detail/dstruct-detail.component'


@Injectable({
    providedIn: 'root'
})
export class RouteService {
    private routes: Routes = []

    constructor(private router: Router) { }

    public addRoutes(routes: Routes): void {
        this.routes.push(...routes)
        this.router.resetConfig(this.routes)
    }

	getPathRoot() : string {
		return 'github_com_fullstack_lang_gong_test_go'
	}
	getTableOutlet(stackPath: string): string {
        return this.getPathRoot() + '_table'
    }
	getEditorOutlet(stackPath: string): string {
        return this.getPathRoot() + '_editor'
    }
	// insertion point for per gongstruct route/path getters
    getAstructTablePath(): string {
        return this.getPathRoot() + '-astructs/:GONG__StackPath'
    }
    getAstructTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getAstructTablePath(), component: AstructsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
	getAstructAdderPath(): string {
        return this.getPathRoot() + '-astruct/:GONG__StackPath'
    }
    getAstructAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getAstructAdderPath(), component: AstructDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
	getAstructAdderForUsePath(): string {
        return this.getPathRoot() + '-astruct/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getAstructAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getAstructAdderForUsePath(), component: AstructDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
	getAstructDetailPath(): string {
        return this.getPathRoot() + '-astruct-detail/:id/:GONG__StackPath'
    }
    getAstructDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getAstructDetailPath(), component: AstructDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getAstructBstruct2UseTablePath(): string {
        return this.getPathRoot() + '-astructbstruct2uses/:GONG__StackPath'
    }
    getAstructBstruct2UseTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getAstructBstruct2UseTablePath(), component: AstructBstruct2UsesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
	getAstructBstruct2UseAdderPath(): string {
        return this.getPathRoot() + '-astructbstruct2use/:GONG__StackPath'
    }
    getAstructBstruct2UseAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getAstructBstruct2UseAdderPath(), component: AstructBstruct2UseDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
	getAstructBstruct2UseAdderForUsePath(): string {
        return this.getPathRoot() + '-astructbstruct2use/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getAstructBstruct2UseAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getAstructBstruct2UseAdderForUsePath(), component: AstructBstruct2UseDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
	getAstructBstruct2UseDetailPath(): string {
        return this.getPathRoot() + '-astructbstruct2use-detail/:id/:GONG__StackPath'
    }
    getAstructBstruct2UseDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getAstructBstruct2UseDetailPath(), component: AstructBstruct2UseDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getAstructBstructUseTablePath(): string {
        return this.getPathRoot() + '-astructbstructuses/:GONG__StackPath'
    }
    getAstructBstructUseTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getAstructBstructUseTablePath(), component: AstructBstructUsesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
	getAstructBstructUseAdderPath(): string {
        return this.getPathRoot() + '-astructbstructuse/:GONG__StackPath'
    }
    getAstructBstructUseAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getAstructBstructUseAdderPath(), component: AstructBstructUseDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
	getAstructBstructUseAdderForUsePath(): string {
        return this.getPathRoot() + '-astructbstructuse/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getAstructBstructUseAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getAstructBstructUseAdderForUsePath(), component: AstructBstructUseDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
	getAstructBstructUseDetailPath(): string {
        return this.getPathRoot() + '-astructbstructuse-detail/:id/:GONG__StackPath'
    }
    getAstructBstructUseDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getAstructBstructUseDetailPath(), component: AstructBstructUseDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getBstructTablePath(): string {
        return this.getPathRoot() + '-bstructs/:GONG__StackPath'
    }
    getBstructTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getBstructTablePath(), component: BstructsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
	getBstructAdderPath(): string {
        return this.getPathRoot() + '-bstruct/:GONG__StackPath'
    }
    getBstructAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getBstructAdderPath(), component: BstructDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
	getBstructAdderForUsePath(): string {
        return this.getPathRoot() + '-bstruct/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getBstructAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getBstructAdderForUsePath(), component: BstructDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
	getBstructDetailPath(): string {
        return this.getPathRoot() + '-bstruct-detail/:id/:GONG__StackPath'
    }
    getBstructDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getBstructDetailPath(), component: BstructDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getDstructTablePath(): string {
        return this.getPathRoot() + '-dstructs/:GONG__StackPath'
    }
    getDstructTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDstructTablePath(), component: DstructsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
	getDstructAdderPath(): string {
        return this.getPathRoot() + '-dstruct/:GONG__StackPath'
    }
    getDstructAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDstructAdderPath(), component: DstructDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
	getDstructAdderForUsePath(): string {
        return this.getPathRoot() + '-dstruct/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getDstructAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDstructAdderForUsePath(), component: DstructDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
	getDstructDetailPath(): string {
        return this.getPathRoot() + '-dstruct-detail/:id/:GONG__StackPath'
    }
    getDstructDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDstructDetailPath(), component: DstructDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }



	addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
			// insertion point for all routes getter
			this.getAstructTableRoute(stackPath),
			this.getAstructAdderRoute(stackPath),
			this.getAstructAdderForUseRoute(stackPath),
			this.getAstructDetailRoute(stackPath),

			this.getAstructBstruct2UseTableRoute(stackPath),
			this.getAstructBstruct2UseAdderRoute(stackPath),
			this.getAstructBstruct2UseAdderForUseRoute(stackPath),
			this.getAstructBstruct2UseDetailRoute(stackPath),

			this.getAstructBstructUseTableRoute(stackPath),
			this.getAstructBstructUseAdderRoute(stackPath),
			this.getAstructBstructUseAdderForUseRoute(stackPath),
			this.getAstructBstructUseDetailRoute(stackPath),

			this.getBstructTableRoute(stackPath),
			this.getBstructAdderRoute(stackPath),
			this.getBstructAdderForUseRoute(stackPath),
			this.getBstructDetailRoute(stackPath),

			this.getDstructTableRoute(stackPath),
			this.getDstructAdderRoute(stackPath),
			this.getDstructAdderForUseRoute(stackPath),
			this.getDstructDetailRoute(stackPath),

		])
	}
}
