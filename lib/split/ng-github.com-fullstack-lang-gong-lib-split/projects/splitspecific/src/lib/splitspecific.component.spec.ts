import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SplitspecificComponent } from './splitspecific.component';

describe('SplitspecificComponent', () => {
  let component: SplitspecificComponent;
  let fixture: ComponentFixture<SplitspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SplitspecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SplitspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
