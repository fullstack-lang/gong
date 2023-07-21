import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongrouterdatamodelComponent } from './gongrouterdatamodel.component';

describe('GongrouterdatamodelComponent', () => {
  let component: GongrouterdatamodelComponent;
  let fixture: ComponentFixture<GongrouterdatamodelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongrouterdatamodelComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongrouterdatamodelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
