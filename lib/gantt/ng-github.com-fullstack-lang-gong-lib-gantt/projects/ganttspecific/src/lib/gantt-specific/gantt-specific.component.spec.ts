import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GanttSpecificComponent } from './gantt-specific.component';

describe('GanttSpecificComponent', () => {
  let component: GanttSpecificComponent;
  let fixture: ComponentFixture<GanttSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GanttSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GanttSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
