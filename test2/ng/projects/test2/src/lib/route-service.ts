import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
import { DummysTableComponent } from './dummys-table/dummys-table.component'
import { DummyDetailComponent } from './dummy-detail/dummy-detail.component'


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
        return 'github_com_fullstack_lang_gong_test2_go'
    }
    getTableOutlet(stackPath: string): string {
        return this.getPathRoot() + '_table' + '_' + stackPath
    }
    getEditorOutlet(stackPath: string): string {
        return this.getPathRoot() + '_editor' + '_' + stackPath
    }
    // insertion point for per gongstruct route/path getters
    getDummyTablePath(): string {
        return this.getPathRoot() + '-dummys/:GONG__StackPath'
    }
    getDummyTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDummyTablePath(), component: DummysTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getDummyAdderPath(): string {
        return this.getPathRoot() + '-dummy-adder/:GONG__StackPath'
    }
    getDummyAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDummyAdderPath(), component: DummyDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getDummyAdderForUsePath(): string {
        return this.getPathRoot() + '-dummy-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getDummyAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDummyAdderForUsePath(), component: DummyDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getDummyDetailPath(): string {
        return this.getPathRoot() + '-dummy-detail/:id/:GONG__StackPath'
    }
    getDummyDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDummyDetailPath(), component: DummyDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }



    addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
            // insertion point for all routes getter
            this.getDummyTableRoute(stackPath),
            this.getDummyAdderRoute(stackPath),
            this.getDummyAdderForUseRoute(stackPath),
            this.getDummyDetailRoute(stackPath),

        ])
    }
}
