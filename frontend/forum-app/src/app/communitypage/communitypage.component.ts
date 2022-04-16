import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { SubredditsService } from '../subreddits.service';

@Component({
  selector: 'app-communitypage',
  templateUrl: './communitypage.component.html',
  styleUrls: ['./communitypage.component.css']
})
export class CommunitypageComponent implements OnInit {

  constructor(private route: ActivatedRoute, private service: SubredditsService) { }
  subreddit: any;
  subreddits: any[];
  communityID: string;
  ngOnInit(): void {
    this.communityID = this.route.snapshot.paramMap.get('communityID');
    this.getCommunities();
  }

  getCommunities() {
    this.service.getSubreddits().subscribe((response: any) => {
      if (response.status == 200) {
        this.subreddits = response.data.communities;

        this.subreddits.forEach(s => {
          if (s._id == this.communityID) {
            this.subreddit = s;
          }
        });
      }
      else {
        this.subreddits = []
      }
    });
  }
}
