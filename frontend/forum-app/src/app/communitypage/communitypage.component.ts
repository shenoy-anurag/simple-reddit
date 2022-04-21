import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ActivatedRoute, Router } from '@angular/router';
import { Storage } from '../storage';
import { SubredditsService } from '../subreddits.service';

@Component({
  selector: 'app-communitypage',
  templateUrl: './communitypage.component.html',
  styleUrls: ['./communitypage.component.css']
})
export class CommunitypageComponent implements OnInit {

  constructor(private router: Router, private route: ActivatedRoute, private service: SubredditsService, private snackbar: MatSnackBar) { }
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

  deleteCommunity(title: string, communityUser: string) {
    // check if user is logged in
    if (Storage.isLoggedIn) {
      if (communityUser != Storage.username) {
        this.snackbar.open("You are not the community owner. You do not have permission to delete this community.", "Dismiss", { duration: 1500 });
      }
      else {
        console.log(Storage.username + "," + title);
        this.service.deleteSubreddit(Storage.username, title).subscribe((response: any) => {
          if (response.status == 200) {
            this.snackbar.open(title + " has been deleted.", "Dismiss", { duration: 1500 });          }

            // update communities
            this.getCommunities();
            
            this.router.navigate(['subreddits']);
        });
      }
    }
    else {
      this.snackbar.open("You are not logged in. Please log in to delete communities.", "Dismiss", {duration: 1500});
    }
  }

  subscribeToCommunity(communityName: string) {
    if (Storage.isLoggedIn) {
      this.service.subscribeToSubreddit(Storage.username, communityName).subscribe((response: any) => {
        if (response.status == 201 && response.message == "success") {
          this.snackbar.open("Subscribed to " + communityName, "Ok", { duration: 1500 });

          // update communities
          this.getCommunities();

          
        }
      });
    }
    else {
      this.snackbar.open("Log in to subscribe", "Dismiss", {duration: 1500});
    }
  }
}
