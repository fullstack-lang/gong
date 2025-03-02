import { TestBed } from '@angular/core/testing';

import { SvgEventService } from './svg-event.service';

describe('SvgEventService', () => {
  let service: SvgEventService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SvgEventService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
