import { TestBed } from '@angular/core/testing';

import { SubredditsService } from './subreddits.service';

describe('SubredditsService', () => {
  let service: SubredditsService;

  it('Get Subreddits should return json data', () => {
    const result = service.getSubreddits();
    expect(result).toBeDefined;
  })

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SubredditsService);
  });
});
