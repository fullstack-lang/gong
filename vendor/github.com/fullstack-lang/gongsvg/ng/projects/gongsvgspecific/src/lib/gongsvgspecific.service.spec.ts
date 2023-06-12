import { TestBed } from '@angular/core/testing';

import { GongsvgspecificService } from './gongsvgspecific.service';

describe('GongsvgspecificService', () => {
  let service: GongsvgspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongsvgspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
