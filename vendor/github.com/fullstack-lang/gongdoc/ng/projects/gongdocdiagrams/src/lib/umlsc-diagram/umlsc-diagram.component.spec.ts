import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { UmlscDiagramComponent } from './umlsc-diagram.component';

describe('UmlscDiagramComponent', () => {
  let component: UmlscDiagramComponent;
  let fixture: ComponentFixture<UmlscDiagramComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ UmlscDiagramComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(UmlscDiagramComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
