import { TestBed } from '@angular/core/testing';

import { SsgspecificService } from './ssgspecific.service';

describe('SsgspecificService', () => {
  let service: SsgspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SsgspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
