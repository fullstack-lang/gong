import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongtreedatamodelComponent } from './gongtreedatamodel.component';

describe('GongtreedatamodelComponent', () => {
  let component: GongtreedatamodelComponent;
  let fixture: ComponentFixture<GongtreedatamodelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongtreedatamodelComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongtreedatamodelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
