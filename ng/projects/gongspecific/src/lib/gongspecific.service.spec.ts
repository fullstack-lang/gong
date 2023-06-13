import { TestBed } from '@angular/core/testing';

import { GongspecificService } from './gongspecific.service';

describe('GongspecificService', () => {
  let service: GongspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
