import { TestBed } from '@angular/core/testing';

import { GongtreedatamodelService } from './gongtreedatamodel.service';

describe('GongtreedatamodelService', () => {
  let service: GongtreedatamodelService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongtreedatamodelService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
