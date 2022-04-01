import { Component, OnInit } from '@angular/core';
import { SubredditsService } from '../subreddits.service';

@Component({
  selector: 'app-subreddits',
  templateUrl: './subreddits.component.html',
  styleUrls: ['./subreddits.component.css']
})
export class SubredditsComponent implements OnInit {
  subreddits: any[] = [];
  constructor(private service: SubredditsService) {
  }

  ngOnInit(): void {
    this.getCommunities();
  }

  getCommunities() {
    this.service.getSubreddits().subscribe((response: any) => {
      if (response.status == 200) {
        this.subreddits = response.data.communities;
      }
      else {
        this.subreddits = []
      }
    });
  }

}
