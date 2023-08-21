import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongtabledatamodelComponent } from './gongtabledatamodel.component';

describe('GongtabledatamodelComponent', () => {
  let component: GongtabledatamodelComponent;
  let fixture: ComponentFixture<GongtabledatamodelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongtabledatamodelComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongtabledatamodelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
