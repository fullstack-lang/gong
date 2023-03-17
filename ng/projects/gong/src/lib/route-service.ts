import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
import { GongBasicFieldsTableComponent } from './gongbasicfields-table/gongbasicfields-table.component'
import { GongBasicFieldDetailComponent } from './gongbasicfield-detail/gongbasicfield-detail.component'

import { GongEnumsTableComponent } from './gongenums-table/gongenums-table.component'
import { GongEnumDetailComponent } from './gongenum-detail/gongenum-detail.component'

import { GongEnumValuesTableComponent } from './gongenumvalues-table/gongenumvalues-table.component'
import { GongEnumValueDetailComponent } from './gongenumvalue-detail/gongenumvalue-detail.component'

import { GongLinksTableComponent } from './gonglinks-table/gonglinks-table.component'
import { GongLinkDetailComponent } from './gonglink-detail/gonglink-detail.component'

import { GongNotesTableComponent } from './gongnotes-table/gongnotes-table.component'
import { GongNoteDetailComponent } from './gongnote-detail/gongnote-detail.component'

import { GongStructsTableComponent } from './gongstructs-table/gongstructs-table.component'
import { GongStructDetailComponent } from './gongstruct-detail/gongstruct-detail.component'

import { GongTimeFieldsTableComponent } from './gongtimefields-table/gongtimefields-table.component'
import { GongTimeFieldDetailComponent } from './gongtimefield-detail/gongtimefield-detail.component'

import { MetasTableComponent } from './metas-table/metas-table.component'
import { MetaDetailComponent } from './meta-detail/meta-detail.component'

import { MetaReferencesTableComponent } from './metareferences-table/metareferences-table.component'
import { MetaReferenceDetailComponent } from './metareference-detail/metareference-detail.component'

import { ModelPkgsTableComponent } from './modelpkgs-table/modelpkgs-table.component'
import { ModelPkgDetailComponent } from './modelpkg-detail/modelpkg-detail.component'

import { PointerToGongStructFieldsTableComponent } from './pointertogongstructfields-table/pointertogongstructfields-table.component'
import { PointerToGongStructFieldDetailComponent } from './pointertogongstructfield-detail/pointertogongstructfield-detail.component'

import { SliceOfPointerToGongStructFieldsTableComponent } from './sliceofpointertogongstructfields-table/sliceofpointertogongstructfields-table.component'
import { SliceOfPointerToGongStructFieldDetailComponent } from './sliceofpointertogongstructfield-detail/sliceofpointertogongstructfield-detail.component'


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
        return 'github_com_fullstack_lang_gong_go'
    }
    getTableOutlet(stackPath: string): string {
        return this.getPathRoot() + '_table' + '_' + stackPath
    }
    getEditorOutlet(stackPath: string): string {
        return this.getPathRoot() + '_editor' + '_' + stackPath
    }
    // insertion point for per gongstruct route/path getters
    getGongBasicFieldTablePath(): string {
        return this.getPathRoot() + '-gongbasicfields/:GONG__StackPath'
    }
    getGongBasicFieldTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongBasicFieldTablePath(), component: GongBasicFieldsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getGongBasicFieldAdderPath(): string {
        return this.getPathRoot() + '-gongbasicfield-adder/:GONG__StackPath'
    }
    getGongBasicFieldAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongBasicFieldAdderPath(), component: GongBasicFieldDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongBasicFieldAdderForUsePath(): string {
        return this.getPathRoot() + '-gongbasicfield-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getGongBasicFieldAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongBasicFieldAdderForUsePath(), component: GongBasicFieldDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongBasicFieldDetailPath(): string {
        return this.getPathRoot() + '-gongbasicfield-detail/:id/:GONG__StackPath'
    }
    getGongBasicFieldDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongBasicFieldDetailPath(), component: GongBasicFieldDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getGongEnumTablePath(): string {
        return this.getPathRoot() + '-gongenums/:GONG__StackPath'
    }
    getGongEnumTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongEnumTablePath(), component: GongEnumsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getGongEnumAdderPath(): string {
        return this.getPathRoot() + '-gongenum-adder/:GONG__StackPath'
    }
    getGongEnumAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongEnumAdderPath(), component: GongEnumDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongEnumAdderForUsePath(): string {
        return this.getPathRoot() + '-gongenum-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getGongEnumAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongEnumAdderForUsePath(), component: GongEnumDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongEnumDetailPath(): string {
        return this.getPathRoot() + '-gongenum-detail/:id/:GONG__StackPath'
    }
    getGongEnumDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongEnumDetailPath(), component: GongEnumDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getGongEnumValueTablePath(): string {
        return this.getPathRoot() + '-gongenumvalues/:GONG__StackPath'
    }
    getGongEnumValueTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongEnumValueTablePath(), component: GongEnumValuesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getGongEnumValueAdderPath(): string {
        return this.getPathRoot() + '-gongenumvalue-adder/:GONG__StackPath'
    }
    getGongEnumValueAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongEnumValueAdderPath(), component: GongEnumValueDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongEnumValueAdderForUsePath(): string {
        return this.getPathRoot() + '-gongenumvalue-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getGongEnumValueAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongEnumValueAdderForUsePath(), component: GongEnumValueDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongEnumValueDetailPath(): string {
        return this.getPathRoot() + '-gongenumvalue-detail/:id/:GONG__StackPath'
    }
    getGongEnumValueDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongEnumValueDetailPath(), component: GongEnumValueDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getGongLinkTablePath(): string {
        return this.getPathRoot() + '-gonglinks/:GONG__StackPath'
    }
    getGongLinkTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongLinkTablePath(), component: GongLinksTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getGongLinkAdderPath(): string {
        return this.getPathRoot() + '-gonglink-adder/:GONG__StackPath'
    }
    getGongLinkAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongLinkAdderPath(), component: GongLinkDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongLinkAdderForUsePath(): string {
        return this.getPathRoot() + '-gonglink-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getGongLinkAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongLinkAdderForUsePath(), component: GongLinkDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongLinkDetailPath(): string {
        return this.getPathRoot() + '-gonglink-detail/:id/:GONG__StackPath'
    }
    getGongLinkDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongLinkDetailPath(), component: GongLinkDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getGongNoteTablePath(): string {
        return this.getPathRoot() + '-gongnotes/:GONG__StackPath'
    }
    getGongNoteTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongNoteTablePath(), component: GongNotesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getGongNoteAdderPath(): string {
        return this.getPathRoot() + '-gongnote-adder/:GONG__StackPath'
    }
    getGongNoteAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongNoteAdderPath(), component: GongNoteDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongNoteAdderForUsePath(): string {
        return this.getPathRoot() + '-gongnote-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getGongNoteAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongNoteAdderForUsePath(), component: GongNoteDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongNoteDetailPath(): string {
        return this.getPathRoot() + '-gongnote-detail/:id/:GONG__StackPath'
    }
    getGongNoteDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongNoteDetailPath(), component: GongNoteDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getGongStructTablePath(): string {
        return this.getPathRoot() + '-gongstructs/:GONG__StackPath'
    }
    getGongStructTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongStructTablePath(), component: GongStructsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getGongStructAdderPath(): string {
        return this.getPathRoot() + '-gongstruct-adder/:GONG__StackPath'
    }
    getGongStructAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongStructAdderPath(), component: GongStructDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongStructAdderForUsePath(): string {
        return this.getPathRoot() + '-gongstruct-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getGongStructAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongStructAdderForUsePath(), component: GongStructDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongStructDetailPath(): string {
        return this.getPathRoot() + '-gongstruct-detail/:id/:GONG__StackPath'
    }
    getGongStructDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongStructDetailPath(), component: GongStructDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getGongTimeFieldTablePath(): string {
        return this.getPathRoot() + '-gongtimefields/:GONG__StackPath'
    }
    getGongTimeFieldTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongTimeFieldTablePath(), component: GongTimeFieldsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getGongTimeFieldAdderPath(): string {
        return this.getPathRoot() + '-gongtimefield-adder/:GONG__StackPath'
    }
    getGongTimeFieldAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongTimeFieldAdderPath(), component: GongTimeFieldDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongTimeFieldAdderForUsePath(): string {
        return this.getPathRoot() + '-gongtimefield-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getGongTimeFieldAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongTimeFieldAdderForUsePath(), component: GongTimeFieldDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongTimeFieldDetailPath(): string {
        return this.getPathRoot() + '-gongtimefield-detail/:id/:GONG__StackPath'
    }
    getGongTimeFieldDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongTimeFieldDetailPath(), component: GongTimeFieldDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getMetaTablePath(): string {
        return this.getPathRoot() + '-metas/:GONG__StackPath'
    }
    getMetaTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMetaTablePath(), component: MetasTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getMetaAdderPath(): string {
        return this.getPathRoot() + '-meta-adder/:GONG__StackPath'
    }
    getMetaAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMetaAdderPath(), component: MetaDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getMetaAdderForUsePath(): string {
        return this.getPathRoot() + '-meta-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getMetaAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMetaAdderForUsePath(), component: MetaDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getMetaDetailPath(): string {
        return this.getPathRoot() + '-meta-detail/:id/:GONG__StackPath'
    }
    getMetaDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMetaDetailPath(), component: MetaDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getMetaReferenceTablePath(): string {
        return this.getPathRoot() + '-metareferences/:GONG__StackPath'
    }
    getMetaReferenceTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMetaReferenceTablePath(), component: MetaReferencesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getMetaReferenceAdderPath(): string {
        return this.getPathRoot() + '-metareference-adder/:GONG__StackPath'
    }
    getMetaReferenceAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMetaReferenceAdderPath(), component: MetaReferenceDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getMetaReferenceAdderForUsePath(): string {
        return this.getPathRoot() + '-metareference-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getMetaReferenceAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMetaReferenceAdderForUsePath(), component: MetaReferenceDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getMetaReferenceDetailPath(): string {
        return this.getPathRoot() + '-metareference-detail/:id/:GONG__StackPath'
    }
    getMetaReferenceDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMetaReferenceDetailPath(), component: MetaReferenceDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getModelPkgTablePath(): string {
        return this.getPathRoot() + '-modelpkgs/:GONG__StackPath'
    }
    getModelPkgTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getModelPkgTablePath(), component: ModelPkgsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getModelPkgAdderPath(): string {
        return this.getPathRoot() + '-modelpkg-adder/:GONG__StackPath'
    }
    getModelPkgAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getModelPkgAdderPath(), component: ModelPkgDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getModelPkgAdderForUsePath(): string {
        return this.getPathRoot() + '-modelpkg-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getModelPkgAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getModelPkgAdderForUsePath(), component: ModelPkgDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getModelPkgDetailPath(): string {
        return this.getPathRoot() + '-modelpkg-detail/:id/:GONG__StackPath'
    }
    getModelPkgDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getModelPkgDetailPath(), component: ModelPkgDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getPointerToGongStructFieldTablePath(): string {
        return this.getPathRoot() + '-pointertogongstructfields/:GONG__StackPath'
    }
    getPointerToGongStructFieldTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPointerToGongStructFieldTablePath(), component: PointerToGongStructFieldsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getPointerToGongStructFieldAdderPath(): string {
        return this.getPathRoot() + '-pointertogongstructfield-adder/:GONG__StackPath'
    }
    getPointerToGongStructFieldAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPointerToGongStructFieldAdderPath(), component: PointerToGongStructFieldDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getPointerToGongStructFieldAdderForUsePath(): string {
        return this.getPathRoot() + '-pointertogongstructfield-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getPointerToGongStructFieldAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPointerToGongStructFieldAdderForUsePath(), component: PointerToGongStructFieldDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getPointerToGongStructFieldDetailPath(): string {
        return this.getPathRoot() + '-pointertogongstructfield-detail/:id/:GONG__StackPath'
    }
    getPointerToGongStructFieldDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPointerToGongStructFieldDetailPath(), component: PointerToGongStructFieldDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getSliceOfPointerToGongStructFieldTablePath(): string {
        return this.getPathRoot() + '-sliceofpointertogongstructfields/:GONG__StackPath'
    }
    getSliceOfPointerToGongStructFieldTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getSliceOfPointerToGongStructFieldTablePath(), component: SliceOfPointerToGongStructFieldsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getSliceOfPointerToGongStructFieldAdderPath(): string {
        return this.getPathRoot() + '-sliceofpointertogongstructfield-adder/:GONG__StackPath'
    }
    getSliceOfPointerToGongStructFieldAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getSliceOfPointerToGongStructFieldAdderPath(), component: SliceOfPointerToGongStructFieldDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getSliceOfPointerToGongStructFieldAdderForUsePath(): string {
        return this.getPathRoot() + '-sliceofpointertogongstructfield-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getSliceOfPointerToGongStructFieldAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getSliceOfPointerToGongStructFieldAdderForUsePath(), component: SliceOfPointerToGongStructFieldDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getSliceOfPointerToGongStructFieldDetailPath(): string {
        return this.getPathRoot() + '-sliceofpointertogongstructfield-detail/:id/:GONG__StackPath'
    }
    getSliceOfPointerToGongStructFieldDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getSliceOfPointerToGongStructFieldDetailPath(), component: SliceOfPointerToGongStructFieldDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }



    addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
            // insertion point for all routes getter
            this.getGongBasicFieldTableRoute(stackPath),
            this.getGongBasicFieldAdderRoute(stackPath),
            this.getGongBasicFieldAdderForUseRoute(stackPath),
            this.getGongBasicFieldDetailRoute(stackPath),

            this.getGongEnumTableRoute(stackPath),
            this.getGongEnumAdderRoute(stackPath),
            this.getGongEnumAdderForUseRoute(stackPath),
            this.getGongEnumDetailRoute(stackPath),

            this.getGongEnumValueTableRoute(stackPath),
            this.getGongEnumValueAdderRoute(stackPath),
            this.getGongEnumValueAdderForUseRoute(stackPath),
            this.getGongEnumValueDetailRoute(stackPath),

            this.getGongLinkTableRoute(stackPath),
            this.getGongLinkAdderRoute(stackPath),
            this.getGongLinkAdderForUseRoute(stackPath),
            this.getGongLinkDetailRoute(stackPath),

            this.getGongNoteTableRoute(stackPath),
            this.getGongNoteAdderRoute(stackPath),
            this.getGongNoteAdderForUseRoute(stackPath),
            this.getGongNoteDetailRoute(stackPath),

            this.getGongStructTableRoute(stackPath),
            this.getGongStructAdderRoute(stackPath),
            this.getGongStructAdderForUseRoute(stackPath),
            this.getGongStructDetailRoute(stackPath),

            this.getGongTimeFieldTableRoute(stackPath),
            this.getGongTimeFieldAdderRoute(stackPath),
            this.getGongTimeFieldAdderForUseRoute(stackPath),
            this.getGongTimeFieldDetailRoute(stackPath),

            this.getMetaTableRoute(stackPath),
            this.getMetaAdderRoute(stackPath),
            this.getMetaAdderForUseRoute(stackPath),
            this.getMetaDetailRoute(stackPath),

            this.getMetaReferenceTableRoute(stackPath),
            this.getMetaReferenceAdderRoute(stackPath),
            this.getMetaReferenceAdderForUseRoute(stackPath),
            this.getMetaReferenceDetailRoute(stackPath),

            this.getModelPkgTableRoute(stackPath),
            this.getModelPkgAdderRoute(stackPath),
            this.getModelPkgAdderForUseRoute(stackPath),
            this.getModelPkgDetailRoute(stackPath),

            this.getPointerToGongStructFieldTableRoute(stackPath),
            this.getPointerToGongStructFieldAdderRoute(stackPath),
            this.getPointerToGongStructFieldAdderForUseRoute(stackPath),
            this.getPointerToGongStructFieldDetailRoute(stackPath),

            this.getSliceOfPointerToGongStructFieldTableRoute(stackPath),
            this.getSliceOfPointerToGongStructFieldAdderRoute(stackPath),
            this.getSliceOfPointerToGongStructFieldAdderForUseRoute(stackPath),
            this.getSliceOfPointerToGongStructFieldDetailRoute(stackPath),

        ])
    }
}
