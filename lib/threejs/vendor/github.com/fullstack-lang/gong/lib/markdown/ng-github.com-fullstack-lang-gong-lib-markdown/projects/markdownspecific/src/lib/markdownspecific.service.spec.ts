import { TestBed } from '@angular/core/testing';

import { MarkdownspecificService } from './markdownspecific.service';

describe('MarkdownspecificService', () => {
  let service: MarkdownspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MarkdownspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
