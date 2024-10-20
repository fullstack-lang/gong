import { TestBed } from '@angular/core/testing';

import { Test3Service } from './test3.service';

describe('Test3Service', () => {
  let service: Test3Service;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(Test3Service);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
