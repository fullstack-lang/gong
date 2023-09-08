import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Test2specificComponent } from './test2specific.component';

describe('Test2specificComponent', () => {
  let component: Test2specificComponent;
  let fixture: ComponentFixture<Test2specificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ Test2specificComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Test2specificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
