import { Pipe, PipeTransform } from '@angular/core';

import * as gongsvg from 'gongsvg'
import { Segment } from './draw.segments';

@Pipe({ name: 'linkSegments' })
export class LinkSegmentsPipe implements PipeTransform {
  transform(link: gongsvg.LinkDB, map_Link_Segment: Map<gongsvg.LinkDB, Segment[]>): Segment[] {
    let res = map_Link_Segment.get(link)!
    return res
  }
}