import { TestBed } from '@angular/core/testing';

import { GongrouterdatamodelService } from './gongrouterdatamodel.service';

describe('GongrouterdatamodelService', () => {
  let service: GongrouterdatamodelService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongrouterdatamodelService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
