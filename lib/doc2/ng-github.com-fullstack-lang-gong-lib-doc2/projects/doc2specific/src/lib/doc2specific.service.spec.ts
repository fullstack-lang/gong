import { TestBed } from '@angular/core/testing';

import { Doc2specificService } from './doc2specific.service';

describe('Doc2specificService', () => {
  let service: Doc2specificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(Doc2specificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
