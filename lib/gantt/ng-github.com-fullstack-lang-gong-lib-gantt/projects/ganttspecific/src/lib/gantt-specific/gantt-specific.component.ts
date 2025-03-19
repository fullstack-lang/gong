import { Component, Input } from '@angular/core';

import * as gantt from '../../../../gantt/src/public-api'

import { SvgSpecificComponent } from '../../../../../../../svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svgspecific/src/lib/svg-specific/svg-specific.component'


@Component({
  selector: 'lib-gantt-specific',
  imports: [
    SvgSpecificComponent
  ],
  templateUrl: './gantt-specific.component.html',
  styleUrl: './gantt-specific.component.css'
})
export class GanttSpecificComponent {


  @Input() Name: string = ""


  SvgStackName = gantt.GanttStacksNames.SvgStackName
}
