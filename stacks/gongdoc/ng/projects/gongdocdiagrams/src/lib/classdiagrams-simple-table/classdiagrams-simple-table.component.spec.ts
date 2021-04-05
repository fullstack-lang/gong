import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ClassdiagramsSimpleTableComponent } from './classdiagrams-simple-table.component';

describe('ClassdiagramsSimpleTableComponent', () => {
  let component: ClassdiagramsSimpleTableComponent;
  let fixture: ComponentFixture<ClassdiagramsSimpleTableComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ClassdiagramsSimpleTableComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ClassdiagramsSimpleTableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
