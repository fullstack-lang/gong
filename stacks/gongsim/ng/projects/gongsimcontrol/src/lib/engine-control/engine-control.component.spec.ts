import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EngineControlComponent } from './engine-control.component';

describe('EngineControlComponent', () => {
  let component: EngineControlComponent;
  let fixture: ComponentFixture<EngineControlComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EngineControlComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EngineControlComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
