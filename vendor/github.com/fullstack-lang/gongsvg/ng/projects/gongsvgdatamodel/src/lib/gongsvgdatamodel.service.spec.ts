import { TestBed } from '@angular/core/testing';

import { GongsvgdatamodelService } from './gongsvgdatamodel.service';

describe('GongsvgdatamodelService', () => {
  let service: GongsvgdatamodelService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongsvgdatamodelService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
