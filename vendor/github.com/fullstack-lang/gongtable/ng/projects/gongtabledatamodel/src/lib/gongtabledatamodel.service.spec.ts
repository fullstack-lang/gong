import { TestBed } from '@angular/core/testing';

import { GongtabledatamodelService } from './gongtabledatamodel.service';

describe('GongtabledatamodelService', () => {
  let service: GongtabledatamodelService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongtabledatamodelService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
