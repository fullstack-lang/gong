import { TestBed } from '@angular/core/testing';

import { MouseEventService } from './mouse-event.service';

describe('MouseEventService', () => {
  let service: MouseEventService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MouseEventService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
