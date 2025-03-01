import { TestBed } from '@angular/core/testing';

import { SliderspecificService } from './sliderspecific.service';

describe('SliderspecificService', () => {
  let service: SliderspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SliderspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
