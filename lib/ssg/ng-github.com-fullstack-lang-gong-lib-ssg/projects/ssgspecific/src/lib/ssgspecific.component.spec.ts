import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SsgspecificComponent } from './ssgspecific.component';

describe('SsgspecificComponent', () => {
  let component: SsgspecificComponent;
  let fixture: ComponentFixture<SsgspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SsgspecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SsgspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
