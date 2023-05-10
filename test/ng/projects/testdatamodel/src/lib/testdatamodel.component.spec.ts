import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TestdatamodelComponent } from './testdatamodel.component';

describe('TestdatamodelComponent', () => {
  let component: TestdatamodelComponent;
  let fixture: ComponentFixture<TestdatamodelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TestdatamodelComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TestdatamodelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
