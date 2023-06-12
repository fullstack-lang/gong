import { NgModule } from '@angular/core';

import { MatButtonModule } from '@angular/material/button';


import { GongsvgspecificComponent } from './gongsvgspecific.component';
import { LayerComponent } from './layer/layer.component';
import { RectComponent } from './rect/rect.component';
import { TextComponent } from './text/text.component';
import { CircleComponent } from './circle/circle.component';
import { LineComponent } from './line/line.component';
import { EllipseComponent } from './ellipse/ellipse.component';
import { PolylineComponent } from './polyline/polyline.component';
import { PathComponent } from './path/path.component';
import { PolygoneComponent } from './polygone/polygone.component';

import { CommonModule } from '@angular/common';

import { GongsvgModule } from 'gongsvg';
import { SvgComponent } from './svg/svg.component';
import { LinkComponent } from './link/link.component';
import { RectLinkLinkComponent } from './rect-link-link/rect-link-link.component';



@NgModule({
  declarations: [
    GongsvgspecificComponent,
    LayerComponent,
    RectComponent,
    TextComponent,
    CircleComponent,
    LineComponent,
    EllipseComponent,
    PolylineComponent,
    PathComponent,
    PolygoneComponent,
    SvgComponent,
    LinkComponent,
    RectLinkLinkComponent
  ],
  imports: [
    CommonModule,

    MatButtonModule,

    GongsvgModule
  ],
  exports: [
    GongsvgspecificComponent,
    SvgComponent,
  ]
})
export class GongsvgspecificModule { }
