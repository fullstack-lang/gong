import { TestBed } from '@angular/core/testing';

import { GanttService } from './gantt.service';

describe('GanttService', () => {
  let service: GanttService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GanttService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
