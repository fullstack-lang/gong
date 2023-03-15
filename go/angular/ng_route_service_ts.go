package angular

const NgRouteServiceTemplate = `import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports{{` + string(NgRouteServiceImports) + `}}

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
		return '{{PkgPathRootWithoutSlashes}}'
	}
	getTableOutlet(stackPath: string): string {
        return this.getPathRoot() + '_table'
    }
	getEditorOutlet(stackPath: string): string {
        return this.getPathRoot() + '_editor'
    }
	// insertion point for per gongstruct route/path getters{{` + string(NgRouteServiceGetters) + `}}


	addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
			// insertion point for all routes getter{{` + string(NgRouteServiceGetAllRoutes) + `}}
		])
	}
}
`

type NgRouteServiceServiceSubTemplate string

const (
	NgRouteServiceImports      NgRouteServiceServiceSubTemplate = "NgRouteServiceImports"
	NgRouteServiceGetters      NgRouteServiceServiceSubTemplate = "NgRouteServiceGetters"
	NgRouteServiceGetAllRoutes NgRouteServiceServiceSubTemplate = "NgRouteServiceGetAllRoutes"
)

var NgRouteServiceSubTemplateCode map[string]string = // new line
map[string]string{

	string(NgRouteServiceImports): `
import { {{Structname}}sTableComponent } from './{{structname}}s-table/{{structname}}s-table.component'
import { {{Structname}}DetailComponent } from './{{structname}}-detail/{{structname}}-detail.component'
`,

	string(NgRouteServiceGetters): `
    get{{Structname}}TablePath(): string {
        return this.getPathRoot() + '-{{structname}}s/:GONG__StackPath'
    }
    get{{Structname}}TableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.get{{Structname}}TablePath(), component: {{Structname}}sTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
	get{{Structname}}AdderPath(): string {
        return this.getPathRoot() + '-{{structname}}/:GONG__StackPath'
    }
    get{{Structname}}AdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.get{{Structname}}AdderPath(), component: {{Structname}}DetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
	get{{Structname}}AdderForUsePath(): string {
        return this.getPathRoot() + '-{{structname}}/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    get{{Structname}}AdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.get{{Structname}}AdderForUsePath(), component: {{Structname}}DetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
	get{{Structname}}DetailPath(): string {
        return this.getPathRoot() + '-{{structname}}-detail/:id/:GONG__StackPath'
    }
    get{{Structname}}DetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.get{{Structname}}DetailPath(), component: {{Structname}}DetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
`,

	string(NgRouteServiceGetAllRoutes): `
			this.get{{Structname}}TableRoute(stackPath),
			this.get{{Structname}}AdderRoute(stackPath),
			this.get{{Structname}}AdderForUseRoute(stackPath),
			this.get{{Structname}}DetailRoute(stackPath),
`,
}
