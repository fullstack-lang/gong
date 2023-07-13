import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
import { OutletsTableComponent } from './outlets-table/outlets-table.component'
import { OutletDetailComponent } from './outlet-detail/outlet-detail.component'


@Injectable({
    providedIn: 'root'
})
export class RouteService {
    private routes: Routes = []

    constructor(private router: Router) { }

    public addRoutes(newRoutes: Routes): void {
        const existingRoutes = this.router.config
        this.routes = this.router.config

        for (let newRoute of newRoutes) {
            if (!existingRoutes.includes(newRoute)) {
                this.routes.push(newRoute)
            }
        }
        this.router.resetConfig(this.routes)
    }

    getPathRoot(): string {
        return 'github_com_fullstack_lang_gongrouter_go'
    }
    getTableOutlet(stackPath: string): string {
        return this.getPathRoot() + '_table' + '_' + stackPath
    }
    getEditorOutlet(stackPath: string): string {
        return this.getPathRoot() + '_editor' + '_' + stackPath
    }
    // insertion point for per gongstruct route/path getters
    getOutletTablePath(): string {
        return this.getPathRoot() + '-outlets/:GONG__StackPath'
    }
    getOutletTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getOutletTablePath(), component: OutletsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getOutletAdderPath(): string {
        return this.getPathRoot() + '-outlet-adder/:GONG__StackPath'
    }
    getOutletAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getOutletAdderPath(), component: OutletDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getOutletAdderForUsePath(): string {
        return this.getPathRoot() + '-outlet-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getOutletAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getOutletAdderForUsePath(), component: OutletDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getOutletDetailPath(): string {
        return this.getPathRoot() + '-outlet-detail/:id/:GONG__StackPath'
    }
    getOutletDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getOutletDetailPath(), component: OutletDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }



    addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
            // insertion point for all routes getter
            this.getOutletTableRoute(stackPath),
            this.getOutletAdderRoute(stackPath),
            this.getOutletAdderForUseRoute(stackPath),
            this.getOutletDetailRoute(stackPath),

        ])
    }
}
