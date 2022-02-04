import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SubredditsComponent } from './subreddits.component';

describe('SubredditsComponent', () => {
  let component: SubredditsComponent;
  let fixture: ComponentFixture<SubredditsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SubredditsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SubredditsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
