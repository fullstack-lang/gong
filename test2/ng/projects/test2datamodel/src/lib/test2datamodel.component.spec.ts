import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Test2datamodelComponent } from './test2datamodel.component';

describe('Test2datamodelComponent', () => {
  let component: Test2datamodelComponent;
  let fixture: ComponentFixture<Test2datamodelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ Test2datamodelComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Test2datamodelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
