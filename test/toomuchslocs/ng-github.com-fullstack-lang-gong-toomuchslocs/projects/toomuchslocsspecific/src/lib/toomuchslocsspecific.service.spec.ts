import { TestBed } from '@angular/core/testing';

import { ToomuchslocsspecificService } from './toomuchslocsspecific.service';

describe('ToomuchslocsspecificService', () => {
  let service: ToomuchslocsspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ToomuchslocsspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
