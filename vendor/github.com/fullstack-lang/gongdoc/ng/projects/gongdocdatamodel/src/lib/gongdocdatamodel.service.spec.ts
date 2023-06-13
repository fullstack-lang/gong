import { TestBed } from '@angular/core/testing';

import { GongdocdatamodelService } from './gongdocdatamodel.service';

describe('GongdocdatamodelService', () => {
  let service: GongdocdatamodelService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongdocdatamodelService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
