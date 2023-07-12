import { Component, Input } from '@angular/core';

@Component({
  selector: 'lib-splitter-new',
  templateUrl: './splitter-new.component.html',
  styleUrls: ['./splitter-new.component.css']
})
export class SplitterNewComponent {
  @Input() GONG__StackPath: string = ""

}
