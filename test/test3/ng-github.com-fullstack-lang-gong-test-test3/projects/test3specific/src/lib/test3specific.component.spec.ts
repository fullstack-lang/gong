import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Test3specificComponent } from './test3specific.component';

describe('Test3specificComponent', () => {
  let component: Test3specificComponent;
  let fixture: ComponentFixture<Test3specificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Test3specificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Test3specificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
