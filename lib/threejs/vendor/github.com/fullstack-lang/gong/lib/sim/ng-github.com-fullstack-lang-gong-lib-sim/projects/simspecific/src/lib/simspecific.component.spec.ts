import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SimspecificComponent } from './simspecific.component';

describe('SimspecificComponent', () => {
  let component: SimspecificComponent;
  let fixture: ComponentFixture<SimspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SimspecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SimspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
