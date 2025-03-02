import { TestBed } from '@angular/core/testing';

import { SizeTrackerService } from './size-tracker.service';

describe('SizeTrackerService', () => {
  let service: SizeTrackerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SizeTrackerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
