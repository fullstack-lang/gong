import { AfterViewInit, ChangeDetectorRef, Component, ElementRef, OnInit, ViewChild } from '@angular/core';

@Component({
  selector: 'lib-text-width-calculator',
  templateUrl: './text-width-calculator.component.html',
  styleUrls: ['./text-width-calculator.component.css'],
  standalone: true,
})
export class TextWidthCalculatorComponent implements OnInit, AfterViewInit {
  @ViewChild('measureElement') measureElement!: ElementRef;

  ngAfterViewInit() {
    // console.log("measure text width", "ngAfterViewInit")
    // Use this method to measure the text width
  }

  ngOnInit() {
    this.changeDetectorRef.detach()
    this.changeDetectorRef.detectChanges()
  }

  constructor(
    private changeDetectorRef: ChangeDetectorRef,
  ) {

  }

  measureTextWidth(text: string): number {
    // console.log("measure text width")
    this.changeDetectorRef.detectChanges()
    const element = this.measureElement.nativeElement;
    element.textContent = text;
    return element.offsetWidth
  }

  measureTextHeight(text: string): number {
    // console.log("measure text height")
    const element = this.measureElement.nativeElement;
    element.textContent = text;
    return element.offsetHeight
  }
}
