import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DiagramSvgComponent } from './diagram-svg.component';

describe('DiagramSvgComponent', () => {
  let component: DiagramSvgComponent;
  let fixture: ComponentFixture<DiagramSvgComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [DiagramSvgComponent]
    });
    fixture = TestBed.createComponent(DiagramSvgComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
