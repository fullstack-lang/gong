import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
import { ButtonsTableComponent } from './buttons-table/buttons-table.component'
import { ButtonDetailComponent } from './button-detail/button-detail.component'

import { NodesTableComponent } from './nodes-table/nodes-table.component'
import { NodeDetailComponent } from './node-detail/node-detail.component'

import { TreesTableComponent } from './trees-table/trees-table.component'
import { TreeDetailComponent } from './tree-detail/tree-detail.component'


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
        return 'github_com_fullstack_lang_gongtree_go'
    }
    getTableOutlet(stackPath: string): string {
        return this.getPathRoot() + '_table' + '_' + stackPath
    }
    getEditorOutlet(stackPath: string): string {
        return this.getPathRoot() + '_editor' + '_' + stackPath
    }
    // insertion point for per gongstruct route/path getters
    getButtonTablePath(): string {
        return this.getPathRoot() + '-buttons/:GONG__StackPath'
    }
    getButtonTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getButtonTablePath(), component: ButtonsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getButtonAdderPath(): string {
        return this.getPathRoot() + '-button-adder/:GONG__StackPath'
    }
    getButtonAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getButtonAdderPath(), component: ButtonDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getButtonAdderForUsePath(): string {
        return this.getPathRoot() + '-button-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getButtonAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getButtonAdderForUsePath(), component: ButtonDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getButtonDetailPath(): string {
        return this.getPathRoot() + '-button-detail/:id/:GONG__StackPath'
    }
    getButtonDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getButtonDetailPath(), component: ButtonDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getNodeTablePath(): string {
        return this.getPathRoot() + '-nodes/:GONG__StackPath'
    }
    getNodeTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getNodeTablePath(), component: NodesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getNodeAdderPath(): string {
        return this.getPathRoot() + '-node-adder/:GONG__StackPath'
    }
    getNodeAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getNodeAdderPath(), component: NodeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getNodeAdderForUsePath(): string {
        return this.getPathRoot() + '-node-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getNodeAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getNodeAdderForUsePath(), component: NodeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getNodeDetailPath(): string {
        return this.getPathRoot() + '-node-detail/:id/:GONG__StackPath'
    }
    getNodeDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getNodeDetailPath(), component: NodeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getTreeTablePath(): string {
        return this.getPathRoot() + '-trees/:GONG__StackPath'
    }
    getTreeTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTreeTablePath(), component: TreesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getTreeAdderPath(): string {
        return this.getPathRoot() + '-tree-adder/:GONG__StackPath'
    }
    getTreeAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTreeAdderPath(), component: TreeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTreeAdderForUsePath(): string {
        return this.getPathRoot() + '-tree-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getTreeAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTreeAdderForUsePath(), component: TreeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTreeDetailPath(): string {
        return this.getPathRoot() + '-tree-detail/:id/:GONG__StackPath'
    }
    getTreeDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTreeDetailPath(), component: TreeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }



    addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
            // insertion point for all routes getter
            this.getButtonTableRoute(stackPath),
            this.getButtonAdderRoute(stackPath),
            this.getButtonAdderForUseRoute(stackPath),
            this.getButtonDetailRoute(stackPath),

            this.getNodeTableRoute(stackPath),
            this.getNodeAdderRoute(stackPath),
            this.getNodeAdderForUseRoute(stackPath),
            this.getNodeDetailRoute(stackPath),

            this.getTreeTableRoute(stackPath),
            this.getTreeAdderRoute(stackPath),
            this.getTreeAdderForUseRoute(stackPath),
            this.getTreeDetailRoute(stackPath),

        ])
    }
}
