import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongdatamodelComponent } from './gongdatamodel.component';

describe('GongdatamodelComponent', () => {
  let component: GongdatamodelComponent;
  let fixture: ComponentFixture<GongdatamodelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongdatamodelComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongdatamodelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
