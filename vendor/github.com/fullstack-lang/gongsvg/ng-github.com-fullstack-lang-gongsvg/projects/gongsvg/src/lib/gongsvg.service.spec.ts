import { TestBed } from '@angular/core/testing';

import { GongsvgService } from './gongsvg.service';

describe('GongsvgService', () => {
  let service: GongsvgService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongsvgService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
