import { TestBed } from '@angular/core/testing';

import { GongService } from './gong.service';

describe('GongService', () => {
  let service: GongService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
