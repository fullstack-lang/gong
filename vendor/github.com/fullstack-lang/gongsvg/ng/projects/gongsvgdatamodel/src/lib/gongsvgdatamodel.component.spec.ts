import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongsvgdatamodelComponent } from './gongsvgdatamodel.component';

describe('GongsvgdatamodelComponent', () => {
  let component: GongsvgdatamodelComponent;
  let fixture: ComponentFixture<GongsvgdatamodelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongsvgdatamodelComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongsvgdatamodelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
