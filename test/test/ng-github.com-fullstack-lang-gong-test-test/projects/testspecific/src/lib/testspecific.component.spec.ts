import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TestspecificComponent } from './testspecific.component';

describe('TestspecificComponent', () => {
  let component: TestspecificComponent;
  let fixture: ComponentFixture<TestspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TestspecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TestspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
