import { Component, Injectable, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-gong-splitter',
  templateUrl: './splitter.component.html',
  styleUrls: ['./splitter.component.css'],
})
export class SplitterComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  constructor() { }

  ngOnInit(): void {
    console.log("Splitter: " + this.GONG__StackPath)
  }
}
