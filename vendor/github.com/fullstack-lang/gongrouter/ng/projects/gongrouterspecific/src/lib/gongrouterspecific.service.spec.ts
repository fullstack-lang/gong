import { TestBed } from '@angular/core/testing';

import { GongrouterspecificService } from './gongrouterspecific.service';

describe('GongrouterspecificService', () => {
  let service: GongrouterspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongrouterspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
