import { Pipe, PipeTransform } from '@angular/core';

import * as gongsvg from 'gongsvg'
import { Segment } from './draw.segments';

//Pipes and Pure Functions: If linkSegments is a custom pipe, make 
// sure it's designed to be impure if it depends on external state or data that 
// changes over time. Pure pipes won't re-evaluate when internal state or external data changes. 
// You can make a pipe impure by setting its pure property to false:
@Pipe({
  name: 'linkSegments',
  pure: false,
})
export class LinkSegmentsPipe implements PipeTransform {
  transform(link: gongsvg.LinkDB, map_Link_Segment: Map<gongsvg.LinkDB, Segment[]>): Segment[] {
    let res = map_Link_Segment.get(link)!
    return res
  }
}