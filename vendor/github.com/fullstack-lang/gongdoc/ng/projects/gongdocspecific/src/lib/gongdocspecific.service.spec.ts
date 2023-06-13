import { TestBed } from '@angular/core/testing';

import { GongdocspecificService } from './gongdocspecific.service';

describe('GongdocspecificService', () => {
  let service: GongdocspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongdocspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
