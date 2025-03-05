import { TestBed } from '@angular/core/testing';

import { SplitspecificService } from './splitspecific.service';

describe('SplitspecificService', () => {
  let service: SplitspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SplitspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
