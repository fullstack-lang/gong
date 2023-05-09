import { TestBed } from '@angular/core/testing';

import { TestdatamodelService } from './testdatamodel.service';

describe('TestdatamodelService', () => {
  let service: TestdatamodelService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TestdatamodelService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
