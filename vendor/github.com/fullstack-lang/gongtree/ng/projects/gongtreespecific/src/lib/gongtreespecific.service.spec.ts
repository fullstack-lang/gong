import { TestBed } from '@angular/core/testing';

import { GongtreespecificService } from './gongtreespecific.service';

describe('GongtreespecificService', () => {
  let service: GongtreespecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongtreespecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
