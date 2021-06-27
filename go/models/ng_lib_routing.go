package models

const NgRoutingTemplate = `import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports{{` + string(NgRoutingImports) + `}}

const routes: Routes = [ // insertion point for routes declarations{{` + string(NgRoutingDeclarations) + `}}
];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
`

type NgRoutingServiceSubTemplate string

const (
	NgRoutingDeclarations NgRoutingServiceSubTemplate = "NgRoutingIndivDecls"
	NgRoutingImports      NgRoutingServiceSubTemplate = "NgRoutingImports"
)

var NgRoutingSubTemplateCode map[string]string = // new line
map[string]string{

	string(NgRoutingImports): `
import { {{Structname}}sTableComponent } from './{{structname}}s-table/{{structname}}s-table.component'
import { {{Structname}}DetailComponent } from './{{structname}}-detail/{{structname}}-detail.component'
import { {{Structname}}PresentationComponent } from './{{structname}}-presentation/{{structname}}-presentation.component'
`,

	string(NgRoutingDeclarations): `
	{ path: '{{PkgPathRootWithoutSlashes}}-{{structname}}s', component: {{Structname}}sTableComponent, outlet: '{{PkgPathRootWithoutSlashes}}_table' },
	{ path: '{{PkgPathRootWithoutSlashes}}-{{structname}}-adder', component: {{Structname}}DetailComponent, outlet: '{{PkgPathRootWithoutSlashes}}_editor' },
	{ path: '{{PkgPathRootWithoutSlashes}}-{{structname}}-adder/:id/:originStruct/:originStructFieldName', component: {{Structname}}DetailComponent, outlet: '{{PkgPathRootWithoutSlashes}}_editor' },
	{ path: '{{PkgPathRootWithoutSlashes}}-{{structname}}-detail/:id', component: {{Structname}}DetailComponent, outlet: '{{PkgPathRootWithoutSlashes}}_editor' },
	{ path: '{{PkgPathRootWithoutSlashes}}-{{structname}}-presentation/:id', component: {{Structname}}PresentationComponent, outlet: '{{PkgPathRootWithoutSlashes}}_presentation' },
	{ path: '{{PkgPathRootWithoutSlashes}}-{{structname}}-presentation-special/:id', component: {{Structname}}PresentationComponent, outlet: '{{PkgPathRootWithoutSlashes}}{{structname}}pres' },
`,
}
