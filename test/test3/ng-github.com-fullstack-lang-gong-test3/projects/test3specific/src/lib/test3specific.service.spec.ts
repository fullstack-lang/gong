import { TestBed } from '@angular/core/testing';

import { Test3specificService } from './test3specific.service';

describe('Test3specificService', () => {
  let service: Test3specificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(Test3specificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
