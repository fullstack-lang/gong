import { TestBed } from '@angular/core/testing';

import { Test2specificService } from './test2specific.service';

describe('Test2specificService', () => {
  let service: Test2specificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(Test2specificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
