import { TestBed } from '@angular/core/testing';

import { SimspecificService } from './simspecific.service';

describe('SimspecificService', () => {
  let service: SimspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SimspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
