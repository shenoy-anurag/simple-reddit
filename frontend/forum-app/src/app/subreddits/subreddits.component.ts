import { Component, OnInit } from '@angular/core';
import { SubredditsService } from '../subreddits.service';

@Component({
  selector: 'app-subreddits',
  templateUrl: './subreddits.component.html',
  styleUrls: ['./subreddits.component.css']
})
export class SubredditsComponent implements OnInit {
  subreddits;
  constructor(service: SubredditsService) { 
    this.subreddits = service.getSubreddits();
  }

  ngOnInit(): void {
  }

}
