import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoadSpecificComponent } from './load-specific.component';

describe('LoadSpecificComponent', () => {
  let component: LoadSpecificComponent;
  let fixture: ComponentFixture<LoadSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LoadSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LoadSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
