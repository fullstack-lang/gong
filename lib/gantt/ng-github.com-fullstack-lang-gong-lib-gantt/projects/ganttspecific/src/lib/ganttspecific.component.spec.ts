import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GanttspecificComponent } from './ganttspecific.component';

describe('GanttspecificComponent', () => {
  let component: GanttspecificComponent;
  let fixture: ComponentFixture<GanttspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GanttspecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GanttspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
