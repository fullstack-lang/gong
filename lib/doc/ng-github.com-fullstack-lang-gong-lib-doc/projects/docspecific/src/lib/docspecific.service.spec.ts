import { TestBed } from '@angular/core/testing';

import { DocspecificService } from './docspecific.service';

describe('DocspecificService', () => {
  let service: DocspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(DocspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
