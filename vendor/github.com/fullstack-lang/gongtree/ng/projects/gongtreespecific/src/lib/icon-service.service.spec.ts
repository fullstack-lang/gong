import { TestBed } from '@angular/core/testing';

import { IconServiceService } from './icon-service.service';

describe('IconServiceService', () => {
  let service: IconServiceService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(IconServiceService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
