import { TestBed } from '@angular/core/testing';

import { SvgspecificService } from './svgspecific.service';

describe('SvgspecificService', () => {
  let service: SvgspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SvgspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
