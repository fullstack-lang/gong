import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ToomuchslocsspecificComponent } from './toomuchslocsspecific.component';

describe('ToomuchslocsspecificComponent', () => {
  let component: ToomuchslocsspecificComponent;
  let fixture: ComponentFixture<ToomuchslocsspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ToomuchslocsspecificComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(ToomuchslocsspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
