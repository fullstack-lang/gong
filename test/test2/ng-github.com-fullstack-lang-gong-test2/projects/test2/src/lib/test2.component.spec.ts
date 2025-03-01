import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Test2Component } from './test2.component';

describe('Test2Component', () => {
  let component: Test2Component;
  let fixture: ComponentFixture<Test2Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Test2Component]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(Test2Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
