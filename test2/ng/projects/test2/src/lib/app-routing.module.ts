import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { AclasssTableComponent } from './aclasss-table/aclasss-table.component'
import { AclassDetailComponent } from './aclass-detail/aclass-detail.component'
import { AclassPresentationComponent } from './aclass-presentation/aclass-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'githubcomfullstacklanggongtest2go-aclasss', component: AclasssTableComponent, outlet: 'githubcomfullstacklanggongtest2gotable' },
	{ path: 'githubcomfullstacklanggongtest2go-aclass-adder', component: AclassDetailComponent, outlet: 'githubcomfullstacklanggongtest2goeditor' },
	{ path: 'githubcomfullstacklanggongtest2go-aclass-adder/:id/:association', component: AclassDetailComponent, outlet: 'githubcomfullstacklanggongtest2goeditor' },
	{ path: 'githubcomfullstacklanggongtest2go-aclass-detail/:id', component: AclassDetailComponent, outlet: 'githubcomfullstacklanggongtest2goeditor' },
	{ path: 'githubcomfullstacklanggongtest2go-aclass-presentation/:id', component: AclassPresentationComponent, outlet: 'githubcomfullstacklanggongtest2gopresentation' },
	{ path: 'githubcomfullstacklanggongtest2go-aclass-presentation-special/:id', component: AclassPresentationComponent, outlet: 'githubcomfullstacklanggongtest2goaclasspres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
