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
	{ path: '{{structname}}s', component: {{Structname}}sTableComponent, outlet: 'table' },
	{ path: '{{structname}}-adder', component: {{Structname}}DetailComponent, outlet: 'editor' },
	{ path: '{{structname}}-adder/:id/:association', component: {{Structname}}DetailComponent, outlet: 'editor' },
	{ path: '{{structname}}-detail/:id', component: {{Structname}}DetailComponent, outlet: 'editor' },
	{ path: '{{structname}}-presentation/:id', component: {{Structname}}PresentationComponent, outlet: 'presentation' },
	{ path: '{{structname}}-presentation-special/:id', component: {{Structname}}PresentationComponent, outlet: '{{structname}}pres' },
`,
}
