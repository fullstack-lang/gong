import { TestBed } from '@angular/core/testing';

import { LoadspecificService } from './loadspecific.service';

describe('LoadspecificService', () => {
  let service: LoadspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(LoadspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
