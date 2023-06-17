import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
import { ClassdiagramsTableComponent } from './classdiagrams-table/classdiagrams-table.component'
import { ClassdiagramDetailComponent } from './classdiagram-detail/classdiagram-detail.component'

import { DiagramPackagesTableComponent } from './diagrampackages-table/diagrampackages-table.component'
import { DiagramPackageDetailComponent } from './diagrampackage-detail/diagrampackage-detail.component'

import { FieldsTableComponent } from './fields-table/fields-table.component'
import { FieldDetailComponent } from './field-detail/field-detail.component'

import { GongEnumShapesTableComponent } from './gongenumshapes-table/gongenumshapes-table.component'
import { GongEnumShapeDetailComponent } from './gongenumshape-detail/gongenumshape-detail.component'

import { GongEnumValueEntrysTableComponent } from './gongenumvalueentrys-table/gongenumvalueentrys-table.component'
import { GongEnumValueEntryDetailComponent } from './gongenumvalueentry-detail/gongenumvalueentry-detail.component'

import { GongStructShapesTableComponent } from './gongstructshapes-table/gongstructshapes-table.component'
import { GongStructShapeDetailComponent } from './gongstructshape-detail/gongstructshape-detail.component'

import { LinksTableComponent } from './links-table/links-table.component'
import { LinkDetailComponent } from './link-detail/link-detail.component'

import { NoteShapesTableComponent } from './noteshapes-table/noteshapes-table.component'
import { NoteShapeDetailComponent } from './noteshape-detail/noteshape-detail.component'

import { NoteShapeLinksTableComponent } from './noteshapelinks-table/noteshapelinks-table.component'
import { NoteShapeLinkDetailComponent } from './noteshapelink-detail/noteshapelink-detail.component'

import { PositionsTableComponent } from './positions-table/positions-table.component'
import { PositionDetailComponent } from './position-detail/position-detail.component'

import { UmlStatesTableComponent } from './umlstates-table/umlstates-table.component'
import { UmlStateDetailComponent } from './umlstate-detail/umlstate-detail.component'

import { UmlscsTableComponent } from './umlscs-table/umlscs-table.component'
import { UmlscDetailComponent } from './umlsc-detail/umlsc-detail.component'

import { VerticesTableComponent } from './vertices-table/vertices-table.component'
import { VerticeDetailComponent } from './vertice-detail/vertice-detail.component'


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
        return 'github_com_fullstack_lang_gongdoc_go'
    }
    getTableOutlet(stackPath: string): string {
        return this.getPathRoot() + '_table' + '_' + stackPath
    }
    getEditorOutlet(stackPath: string): string {
        return this.getPathRoot() + '_editor' + '_' + stackPath
    }
    // insertion point for per gongstruct route/path getters
    getClassdiagramTablePath(): string {
        return this.getPathRoot() + '-classdiagrams/:GONG__StackPath'
    }
    getClassdiagramTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getClassdiagramTablePath(), component: ClassdiagramsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getClassdiagramAdderPath(): string {
        return this.getPathRoot() + '-classdiagram-adder/:GONG__StackPath'
    }
    getClassdiagramAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getClassdiagramAdderPath(), component: ClassdiagramDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getClassdiagramAdderForUsePath(): string {
        return this.getPathRoot() + '-classdiagram-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getClassdiagramAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getClassdiagramAdderForUsePath(), component: ClassdiagramDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getClassdiagramDetailPath(): string {
        return this.getPathRoot() + '-classdiagram-detail/:id/:GONG__StackPath'
    }
    getClassdiagramDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getClassdiagramDetailPath(), component: ClassdiagramDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getDiagramPackageTablePath(): string {
        return this.getPathRoot() + '-diagrampackages/:GONG__StackPath'
    }
    getDiagramPackageTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDiagramPackageTablePath(), component: DiagramPackagesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getDiagramPackageAdderPath(): string {
        return this.getPathRoot() + '-diagrampackage-adder/:GONG__StackPath'
    }
    getDiagramPackageAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDiagramPackageAdderPath(), component: DiagramPackageDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getDiagramPackageAdderForUsePath(): string {
        return this.getPathRoot() + '-diagrampackage-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getDiagramPackageAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDiagramPackageAdderForUsePath(), component: DiagramPackageDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getDiagramPackageDetailPath(): string {
        return this.getPathRoot() + '-diagrampackage-detail/:id/:GONG__StackPath'
    }
    getDiagramPackageDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDiagramPackageDetailPath(), component: DiagramPackageDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFieldTablePath(): string {
        return this.getPathRoot() + '-fields/:GONG__StackPath'
    }
    getFieldTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFieldTablePath(), component: FieldsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFieldAdderPath(): string {
        return this.getPathRoot() + '-field-adder/:GONG__StackPath'
    }
    getFieldAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFieldAdderPath(), component: FieldDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFieldAdderForUsePath(): string {
        return this.getPathRoot() + '-field-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFieldAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFieldAdderForUsePath(), component: FieldDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFieldDetailPath(): string {
        return this.getPathRoot() + '-field-detail/:id/:GONG__StackPath'
    }
    getFieldDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFieldDetailPath(), component: FieldDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getGongEnumShapeTablePath(): string {
        return this.getPathRoot() + '-gongenumshapes/:GONG__StackPath'
    }
    getGongEnumShapeTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongEnumShapeTablePath(), component: GongEnumShapesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getGongEnumShapeAdderPath(): string {
        return this.getPathRoot() + '-gongenumshape-adder/:GONG__StackPath'
    }
    getGongEnumShapeAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongEnumShapeAdderPath(), component: GongEnumShapeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongEnumShapeAdderForUsePath(): string {
        return this.getPathRoot() + '-gongenumshape-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getGongEnumShapeAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongEnumShapeAdderForUsePath(), component: GongEnumShapeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongEnumShapeDetailPath(): string {
        return this.getPathRoot() + '-gongenumshape-detail/:id/:GONG__StackPath'
    }
    getGongEnumShapeDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongEnumShapeDetailPath(), component: GongEnumShapeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getGongEnumValueEntryTablePath(): string {
        return this.getPathRoot() + '-gongenumvalueentrys/:GONG__StackPath'
    }
    getGongEnumValueEntryTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongEnumValueEntryTablePath(), component: GongEnumValueEntrysTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getGongEnumValueEntryAdderPath(): string {
        return this.getPathRoot() + '-gongenumvalueentry-adder/:GONG__StackPath'
    }
    getGongEnumValueEntryAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongEnumValueEntryAdderPath(), component: GongEnumValueEntryDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongEnumValueEntryAdderForUsePath(): string {
        return this.getPathRoot() + '-gongenumvalueentry-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getGongEnumValueEntryAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongEnumValueEntryAdderForUsePath(), component: GongEnumValueEntryDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongEnumValueEntryDetailPath(): string {
        return this.getPathRoot() + '-gongenumvalueentry-detail/:id/:GONG__StackPath'
    }
    getGongEnumValueEntryDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongEnumValueEntryDetailPath(), component: GongEnumValueEntryDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getGongStructShapeTablePath(): string {
        return this.getPathRoot() + '-gongstructshapes/:GONG__StackPath'
    }
    getGongStructShapeTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongStructShapeTablePath(), component: GongStructShapesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getGongStructShapeAdderPath(): string {
        return this.getPathRoot() + '-gongstructshape-adder/:GONG__StackPath'
    }
    getGongStructShapeAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongStructShapeAdderPath(), component: GongStructShapeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongStructShapeAdderForUsePath(): string {
        return this.getPathRoot() + '-gongstructshape-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getGongStructShapeAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongStructShapeAdderForUsePath(), component: GongStructShapeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongStructShapeDetailPath(): string {
        return this.getPathRoot() + '-gongstructshape-detail/:id/:GONG__StackPath'
    }
    getGongStructShapeDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongStructShapeDetailPath(), component: GongStructShapeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getLinkTablePath(): string {
        return this.getPathRoot() + '-links/:GONG__StackPath'
    }
    getLinkTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLinkTablePath(), component: LinksTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getLinkAdderPath(): string {
        return this.getPathRoot() + '-link-adder/:GONG__StackPath'
    }
    getLinkAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLinkAdderPath(), component: LinkDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getLinkAdderForUsePath(): string {
        return this.getPathRoot() + '-link-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getLinkAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLinkAdderForUsePath(), component: LinkDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getLinkDetailPath(): string {
        return this.getPathRoot() + '-link-detail/:id/:GONG__StackPath'
    }
    getLinkDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLinkDetailPath(), component: LinkDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getNoteShapeTablePath(): string {
        return this.getPathRoot() + '-noteshapes/:GONG__StackPath'
    }
    getNoteShapeTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getNoteShapeTablePath(), component: NoteShapesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getNoteShapeAdderPath(): string {
        return this.getPathRoot() + '-noteshape-adder/:GONG__StackPath'
    }
    getNoteShapeAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getNoteShapeAdderPath(), component: NoteShapeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getNoteShapeAdderForUsePath(): string {
        return this.getPathRoot() + '-noteshape-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getNoteShapeAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getNoteShapeAdderForUsePath(), component: NoteShapeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getNoteShapeDetailPath(): string {
        return this.getPathRoot() + '-noteshape-detail/:id/:GONG__StackPath'
    }
    getNoteShapeDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getNoteShapeDetailPath(), component: NoteShapeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getNoteShapeLinkTablePath(): string {
        return this.getPathRoot() + '-noteshapelinks/:GONG__StackPath'
    }
    getNoteShapeLinkTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getNoteShapeLinkTablePath(), component: NoteShapeLinksTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getNoteShapeLinkAdderPath(): string {
        return this.getPathRoot() + '-noteshapelink-adder/:GONG__StackPath'
    }
    getNoteShapeLinkAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getNoteShapeLinkAdderPath(), component: NoteShapeLinkDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getNoteShapeLinkAdderForUsePath(): string {
        return this.getPathRoot() + '-noteshapelink-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getNoteShapeLinkAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getNoteShapeLinkAdderForUsePath(), component: NoteShapeLinkDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getNoteShapeLinkDetailPath(): string {
        return this.getPathRoot() + '-noteshapelink-detail/:id/:GONG__StackPath'
    }
    getNoteShapeLinkDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getNoteShapeLinkDetailPath(), component: NoteShapeLinkDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getPositionTablePath(): string {
        return this.getPathRoot() + '-positions/:GONG__StackPath'
    }
    getPositionTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPositionTablePath(), component: PositionsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getPositionAdderPath(): string {
        return this.getPathRoot() + '-position-adder/:GONG__StackPath'
    }
    getPositionAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPositionAdderPath(), component: PositionDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getPositionAdderForUsePath(): string {
        return this.getPathRoot() + '-position-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getPositionAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPositionAdderForUsePath(), component: PositionDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getPositionDetailPath(): string {
        return this.getPathRoot() + '-position-detail/:id/:GONG__StackPath'
    }
    getPositionDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPositionDetailPath(), component: PositionDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getUmlStateTablePath(): string {
        return this.getPathRoot() + '-umlstates/:GONG__StackPath'
    }
    getUmlStateTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getUmlStateTablePath(), component: UmlStatesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getUmlStateAdderPath(): string {
        return this.getPathRoot() + '-umlstate-adder/:GONG__StackPath'
    }
    getUmlStateAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getUmlStateAdderPath(), component: UmlStateDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getUmlStateAdderForUsePath(): string {
        return this.getPathRoot() + '-umlstate-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getUmlStateAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getUmlStateAdderForUsePath(), component: UmlStateDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getUmlStateDetailPath(): string {
        return this.getPathRoot() + '-umlstate-detail/:id/:GONG__StackPath'
    }
    getUmlStateDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getUmlStateDetailPath(), component: UmlStateDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getUmlscTablePath(): string {
        return this.getPathRoot() + '-umlscs/:GONG__StackPath'
    }
    getUmlscTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getUmlscTablePath(), component: UmlscsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getUmlscAdderPath(): string {
        return this.getPathRoot() + '-umlsc-adder/:GONG__StackPath'
    }
    getUmlscAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getUmlscAdderPath(), component: UmlscDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getUmlscAdderForUsePath(): string {
        return this.getPathRoot() + '-umlsc-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getUmlscAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getUmlscAdderForUsePath(), component: UmlscDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getUmlscDetailPath(): string {
        return this.getPathRoot() + '-umlsc-detail/:id/:GONG__StackPath'
    }
    getUmlscDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getUmlscDetailPath(), component: UmlscDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getVerticeTablePath(): string {
        return this.getPathRoot() + '-vertices/:GONG__StackPath'
    }
    getVerticeTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getVerticeTablePath(), component: VerticesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getVerticeAdderPath(): string {
        return this.getPathRoot() + '-vertice-adder/:GONG__StackPath'
    }
    getVerticeAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getVerticeAdderPath(), component: VerticeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getVerticeAdderForUsePath(): string {
        return this.getPathRoot() + '-vertice-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getVerticeAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getVerticeAdderForUsePath(), component: VerticeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getVerticeDetailPath(): string {
        return this.getPathRoot() + '-vertice-detail/:id/:GONG__StackPath'
    }
    getVerticeDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getVerticeDetailPath(), component: VerticeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }



    addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
            // insertion point for all routes getter
            this.getClassdiagramTableRoute(stackPath),
            this.getClassdiagramAdderRoute(stackPath),
            this.getClassdiagramAdderForUseRoute(stackPath),
            this.getClassdiagramDetailRoute(stackPath),

            this.getDiagramPackageTableRoute(stackPath),
            this.getDiagramPackageAdderRoute(stackPath),
            this.getDiagramPackageAdderForUseRoute(stackPath),
            this.getDiagramPackageDetailRoute(stackPath),

            this.getFieldTableRoute(stackPath),
            this.getFieldAdderRoute(stackPath),
            this.getFieldAdderForUseRoute(stackPath),
            this.getFieldDetailRoute(stackPath),

            this.getGongEnumShapeTableRoute(stackPath),
            this.getGongEnumShapeAdderRoute(stackPath),
            this.getGongEnumShapeAdderForUseRoute(stackPath),
            this.getGongEnumShapeDetailRoute(stackPath),

            this.getGongEnumValueEntryTableRoute(stackPath),
            this.getGongEnumValueEntryAdderRoute(stackPath),
            this.getGongEnumValueEntryAdderForUseRoute(stackPath),
            this.getGongEnumValueEntryDetailRoute(stackPath),

            this.getGongStructShapeTableRoute(stackPath),
            this.getGongStructShapeAdderRoute(stackPath),
            this.getGongStructShapeAdderForUseRoute(stackPath),
            this.getGongStructShapeDetailRoute(stackPath),

            this.getLinkTableRoute(stackPath),
            this.getLinkAdderRoute(stackPath),
            this.getLinkAdderForUseRoute(stackPath),
            this.getLinkDetailRoute(stackPath),

            this.getNoteShapeTableRoute(stackPath),
            this.getNoteShapeAdderRoute(stackPath),
            this.getNoteShapeAdderForUseRoute(stackPath),
            this.getNoteShapeDetailRoute(stackPath),

            this.getNoteShapeLinkTableRoute(stackPath),
            this.getNoteShapeLinkAdderRoute(stackPath),
            this.getNoteShapeLinkAdderForUseRoute(stackPath),
            this.getNoteShapeLinkDetailRoute(stackPath),

            this.getPositionTableRoute(stackPath),
            this.getPositionAdderRoute(stackPath),
            this.getPositionAdderForUseRoute(stackPath),
            this.getPositionDetailRoute(stackPath),

            this.getUmlStateTableRoute(stackPath),
            this.getUmlStateAdderRoute(stackPath),
            this.getUmlStateAdderForUseRoute(stackPath),
            this.getUmlStateDetailRoute(stackPath),

            this.getUmlscTableRoute(stackPath),
            this.getUmlscAdderRoute(stackPath),
            this.getUmlscAdderForUseRoute(stackPath),
            this.getUmlscDetailRoute(stackPath),

            this.getVerticeTableRoute(stackPath),
            this.getVerticeAdderRoute(stackPath),
            this.getVerticeAdderForUseRoute(stackPath),
            this.getVerticeDetailRoute(stackPath),

        ])
    }
}
