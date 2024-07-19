import { TestBed } from '@angular/core/testing';

import { Test2Service } from './test2.service';

describe('Test2Service', () => {
  let service: Test2Service;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(Test2Service);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
