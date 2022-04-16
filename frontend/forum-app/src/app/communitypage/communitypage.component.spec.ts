import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CommunitypageComponent } from './communitypage.component';

describe('CommunitypageComponent', () => {
  let component: CommunitypageComponent;
  let fixture: ComponentFixture<CommunitypageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CommunitypageComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CommunitypageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
