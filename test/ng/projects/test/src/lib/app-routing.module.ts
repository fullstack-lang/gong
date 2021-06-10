import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { AclasssTableComponent } from './aclasss-table/aclasss-table.component'
import { AclassDetailComponent } from './aclass-detail/aclass-detail.component'
import { AclassPresentationComponent } from './aclass-presentation/aclass-presentation.component'

import { BclasssTableComponent } from './bclasss-table/bclasss-table.component'
import { BclassDetailComponent } from './bclass-detail/bclass-detail.component'
import { BclassPresentationComponent } from './bclass-presentation/bclass-presentation.component'

import { DclasssTableComponent } from './dclasss-table/dclasss-table.component'
import { DclassDetailComponent } from './dclass-detail/dclass-detail.component'
import { DclassPresentationComponent } from './dclass-presentation/dclass-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github.com.fullstack-lang.gong.test.go-aclasss', component: AclasssTableComponent, outlet: 'table' },
	{ path: 'github.com.fullstack-lang.gong.test.go-aclass-adder', component: AclassDetailComponent, outlet: 'editor' },
	{ path: 'github.com.fullstack-lang.gong.test.go-aclass-adder/:id/:association', component: AclassDetailComponent, outlet: 'editor' },
	{ path: 'github.com.fullstack-lang.gong.test.go-aclass-detail/:id', component: AclassDetailComponent, outlet: 'editor' },
	{ path: 'github.com.fullstack-lang.gong.test.go-aclass-presentation/:id', component: AclassPresentationComponent, outlet: 'presentation' },
	{ path: 'github.com.fullstack-lang.gong.test.go-aclass-presentation-special/:id', component: AclassPresentationComponent, outlet: 'aclasspres' },

	{ path: 'github.com.fullstack-lang.gong.test.go-bclasss', component: BclasssTableComponent, outlet: 'table' },
	{ path: 'github.com.fullstack-lang.gong.test.go-bclass-adder', component: BclassDetailComponent, outlet: 'editor' },
	{ path: 'github.com.fullstack-lang.gong.test.go-bclass-adder/:id/:association', component: BclassDetailComponent, outlet: 'editor' },
	{ path: 'github.com.fullstack-lang.gong.test.go-bclass-detail/:id', component: BclassDetailComponent, outlet: 'editor' },
	{ path: 'github.com.fullstack-lang.gong.test.go-bclass-presentation/:id', component: BclassPresentationComponent, outlet: 'presentation' },
	{ path: 'github.com.fullstack-lang.gong.test.go-bclass-presentation-special/:id', component: BclassPresentationComponent, outlet: 'bclasspres' },

	{ path: 'github.com.fullstack-lang.gong.test.go-dclasss', component: DclasssTableComponent, outlet: 'table' },
	{ path: 'github.com.fullstack-lang.gong.test.go-dclass-adder', component: DclassDetailComponent, outlet: 'editor' },
	{ path: 'github.com.fullstack-lang.gong.test.go-dclass-adder/:id/:association', component: DclassDetailComponent, outlet: 'editor' },
	{ path: 'github.com.fullstack-lang.gong.test.go-dclass-detail/:id', component: DclassDetailComponent, outlet: 'editor' },
	{ path: 'github.com.fullstack-lang.gong.test.go-dclass-presentation/:id', component: DclassPresentationComponent, outlet: 'presentation' },
	{ path: 'github.com.fullstack-lang.gong.test.go-dclass-presentation-special/:id', component: DclassPresentationComponent, outlet: 'dclasspres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
