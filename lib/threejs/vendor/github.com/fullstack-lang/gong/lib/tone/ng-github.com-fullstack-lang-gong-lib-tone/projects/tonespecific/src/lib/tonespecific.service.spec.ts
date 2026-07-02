import { TestBed } from '@angular/core/testing';

import { TonespecificService } from './tonespecific.service';

describe('TonespecificService', () => {
  let service: TonespecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TonespecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
