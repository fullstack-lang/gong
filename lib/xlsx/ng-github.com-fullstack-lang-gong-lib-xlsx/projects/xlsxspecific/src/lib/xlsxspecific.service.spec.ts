import { TestBed } from '@angular/core/testing';

import { XlsxspecificService } from './xlsxspecific.service';

describe('XlsxspecificService', () => {
  let service: XlsxspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(XlsxspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
