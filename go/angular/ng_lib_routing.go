package angular

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
`,

	string(NgRoutingDeclarations): `
	{ path: '{{PkgPathRootWithoutSlashes}}-{{structname}}s/:GONG__StackPath', component: {{Structname}}sTableComponent, outlet: '{{PkgPathRootWithoutSlashes}}_table' },
	{ path: '{{PkgPathRootWithoutSlashes}}-{{structname}}-adder/:GONG__StackPath', component: {{Structname}}DetailComponent, outlet: '{{PkgPathRootWithoutSlashes}}_editor' },
	{ path: '{{PkgPathRootWithoutSlashes}}-{{structname}}-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: {{Structname}}DetailComponent, outlet: '{{PkgPathRootWithoutSlashes}}_editor' },
	{ path: '{{PkgPathRootWithoutSlashes}}-{{structname}}-detail/:id', component: {{Structname}}DetailComponent, outlet: '{{PkgPathRootWithoutSlashes}}_editor' },
`,
}
