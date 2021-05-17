import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { AclasssTableComponent } from './aclasss-table/aclasss-table.component'
import { AclassDetailComponent } from './aclass-detail/aclass-detail.component'
import { AclassPresentationComponent } from './aclass-presentation/aclass-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'aclasss', component: AclasssTableComponent, outlet: 'table' },
	{ path: 'aclass-adder', component: AclassDetailComponent, outlet: 'editor' },
	{ path: 'aclass-adder/:id/:association', component: AclassDetailComponent, outlet: 'editor' },
	{ path: 'aclass-detail/:id', component: AclassDetailComponent, outlet: 'editor' },
	{ path: 'aclass-presentation/:id', component: AclassPresentationComponent, outlet: 'presentation' },
	{ path: 'aclass-presentation-special/:id', component: AclassPresentationComponent, outlet: 'aclasspres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
