import { TestBed } from '@angular/core/testing';

import { GongdatamodelService } from './gongdatamodel.service';

describe('GongdatamodelService', () => {
  let service: GongdatamodelService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongdatamodelService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
