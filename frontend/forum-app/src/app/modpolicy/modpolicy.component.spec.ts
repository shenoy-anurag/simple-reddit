import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ModpolicyComponent } from './modpolicy.component';

describe('ModpolicyComponent', () => {
  let component: ModpolicyComponent;
  let fixture: ComponentFixture<ModpolicyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ModpolicyComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ModpolicyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
