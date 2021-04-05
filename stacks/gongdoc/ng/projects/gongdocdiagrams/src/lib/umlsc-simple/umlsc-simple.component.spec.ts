import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { UmlscSimpleComponent } from './umlsc-simple.component';

describe('UmlscSimpleComponent', () => {
  let component: UmlscSimpleComponent;
  let fixture: ComponentFixture<UmlscSimpleComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ UmlscSimpleComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(UmlscSimpleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
