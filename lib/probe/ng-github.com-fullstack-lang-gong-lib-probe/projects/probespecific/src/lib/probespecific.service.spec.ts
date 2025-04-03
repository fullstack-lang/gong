import { TestBed } from '@angular/core/testing';

import { ProbespecificService } from './probespecific.service';

describe('ProbespecificService', () => {
  let service: ProbespecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ProbespecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
