import { TestBed } from '@angular/core/testing';

import { Test2datamodelService } from './test2datamodel.service';

describe('Test2datamodelService', () => {
  let service: Test2datamodelService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(Test2datamodelService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
