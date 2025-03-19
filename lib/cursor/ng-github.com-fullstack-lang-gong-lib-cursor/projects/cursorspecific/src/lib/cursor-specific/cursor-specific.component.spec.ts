import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CursorSpecificComponent } from './cursor-specific.component';

describe('CursorSpecificComponent', () => {
  let component: CursorSpecificComponent;
  let fixture: ComponentFixture<CursorSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CursorSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CursorSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
