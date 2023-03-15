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

    getTableOutlet(stackPath: string): string {
        return 'github_com_fullstack_lang_gong_test_go_table'
    }

    getAstructTablePath(stackPath: string): string {
        return 'github_com_fullstack_lang_gong_test_go-astructs/:GONG__StackPath'
    }

    getAstructTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getAstructTablePath(stackPath), component: AstructsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }

    addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
            this.getAstructTableRoute(stackPath),
            { path: 'github_com_fullstack_lang_gong_test_go-astruct-adder/:GONG__StackPath', component: AstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
            { path: 'github_com_fullstack_lang_gong_test_go-astruct-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: AstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
            { path: 'github_com_fullstack_lang_gong_test_go-astruct-detail/:id/:GONG__StackPath', component: AstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },

            { path: 'github_com_fullstack_lang_gong_test_go-astructbstruct2uses/:GONG__StackPath', component: AstructBstruct2UsesTableComponent, outlet: 'github_com_fullstack_lang_gong_test_go_table' },
            { path: 'github_com_fullstack_lang_gong_test_go-astructbstruct2use-adder/:GONG__StackPath', component: AstructBstruct2UseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
            { path: 'github_com_fullstack_lang_gong_test_go-astructbstruct2use-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: AstructBstruct2UseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
            { path: 'github_com_fullstack_lang_gong_test_go-astructbstruct2use-detail/:id/:GONG__StackPath', component: AstructBstruct2UseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },

            { path: 'github_com_fullstack_lang_gong_test_go-astructbstructuses/:GONG__StackPath', component: AstructBstructUsesTableComponent, outlet: 'github_com_fullstack_lang_gong_test_go_table' },
            { path: 'github_com_fullstack_lang_gong_test_go-astructbstructuse-adder/:GONG__StackPath', component: AstructBstructUseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
            { path: 'github_com_fullstack_lang_gong_test_go-astructbstructuse-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: AstructBstructUseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
            { path: 'github_com_fullstack_lang_gong_test_go-astructbstructuse-detail/:id/:GONG__StackPath', component: AstructBstructUseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },

            { path: 'github_com_fullstack_lang_gong_test_go-bstructs/:GONG__StackPath', component: BstructsTableComponent, outlet: 'github_com_fullstack_lang_gong_test_go_table' },
            { path: 'github_com_fullstack_lang_gong_test_go-bstruct-adder/:GONG__StackPath', component: BstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
            { path: 'github_com_fullstack_lang_gong_test_go-bstruct-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: BstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
            { path: 'github_com_fullstack_lang_gong_test_go-bstruct-detail/:id/:GONG__StackPath', component: BstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },

            { path: 'github_com_fullstack_lang_gong_test_go-dstructs/:GONG__StackPath', component: DstructsTableComponent, outlet: 'github_com_fullstack_lang_gong_test_go_table' },
            { path: 'github_com_fullstack_lang_gong_test_go-dstruct-adder/:GONG__StackPath', component: DstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
            { path: 'github_com_fullstack_lang_gong_test_go-dstruct-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: DstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
            { path: 'github_com_fullstack_lang_gong_test_go-dstruct-detail/:id/:GONG__StackPath', component: DstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
        ])
    }
}
