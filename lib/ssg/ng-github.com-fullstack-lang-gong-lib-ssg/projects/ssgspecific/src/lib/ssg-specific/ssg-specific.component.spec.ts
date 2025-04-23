import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SsgSpecificComponent } from './ssg-specific.component';

describe('SsgSpecificComponent', () => {
  let component: SsgSpecificComponent;
  let fixture: ComponentFixture<SsgSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SsgSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SsgSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
