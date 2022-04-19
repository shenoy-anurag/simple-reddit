import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ContentpolicyComponent } from './contentpolicy.component';

describe('ContentpolicyComponent', () => {
  let component: ContentpolicyComponent;
  let fixture: ComponentFixture<ContentpolicyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ContentpolicyComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ContentpolicyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
