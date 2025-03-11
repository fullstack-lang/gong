import { TestBed } from '@angular/core/testing';

import { GanttspecificService } from './ganttspecific.service';

describe('GanttspecificService', () => {
  let service: GanttspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GanttspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
