import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SignupformComponent } from './signupform.component';

describe('SignupformComponent', () => {
  let component: SignupformComponent;
  let fixture: ComponentFixture<SignupformComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SignupformComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SignupformComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
