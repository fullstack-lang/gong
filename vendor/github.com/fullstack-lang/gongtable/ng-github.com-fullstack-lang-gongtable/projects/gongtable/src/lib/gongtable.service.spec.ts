import { TestBed } from '@angular/core/testing';

import { GongtableService } from './gongtable.service';

describe('GongtableService', () => {
  let service: GongtableService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongtableService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
