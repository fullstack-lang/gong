import { TestBed } from '@angular/core/testing';

import { RectangleEventService } from './rectangle-event.service';

describe('RectangleEventService', () => {
  let service: RectangleEventService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(RectangleEventService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
