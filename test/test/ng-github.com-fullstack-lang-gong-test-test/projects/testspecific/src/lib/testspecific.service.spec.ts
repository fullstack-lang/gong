import { TestBed } from '@angular/core/testing';

import { TestspecificService } from './testspecific.service';

describe('TestspecificService', () => {
  let service: TestspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TestspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
