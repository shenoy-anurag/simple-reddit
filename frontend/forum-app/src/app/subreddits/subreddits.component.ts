import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Storage } from '../storage';
import { SubredditsService } from '../subreddits.service';
import { SignupService } from '../signup.service';

@Component({
  selector: 'app-subreddits',
  templateUrl: './subreddits.component.html',
  styleUrls: ['./subreddits.component.css']
})
export class SubredditsComponent implements OnInit {
  subreddits: any[] = [];
  constructor(private service: SubredditsService, private snackbar: MatSnackBar) {
  }

  ngOnInit(): void {
    this.getCommunities();
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
        });
      }
    }
    else {
      this.snackbar.open("You are not logged in. Please log in to delete communities.", "Dismiss", {duration: 1500});
    }
  }


// upvoteCommunity(id: string) {
//   if (Storage.isLoggedIn) {
//     this.service.voteCommunity(id, Storage.username, 1).subscribe((response: any) => {
//       if (response.status == 200) {
//         this.getCommunities();
//         this.snackbar.open("Upvote Succesfull", "Dismiss", { duration: 1500 });
//       }
//       else{
//         this.snackbar.open("Vote Unsuccesfull", "Dismiss", { duration: 1500 });
//       }
//     });
//   }
//   else {
//     this.snackbar.open("Log in to vote on posts", "Dismiss", { duration: 1500 });
//   }
// }

  

// downvoteCommunity(id: string) {
//   if (Storage.isLoggedIn) {
//     this.service.voteCommunity(id, Storage.username, -1).subscribe((response: any) => {
//       if (response.status == 200) {
//         this.getCommunities();
//         this.snackbar.open("Downvote Succesfull", "Dismiss", { duration: 1500 });
//       }
//      else{
//         this.snackbar.open("Vote Unsuccesfull", "Dismiss", { duration: 1500 });
//       }
//     });
//   }
//   else {
//     this.snackbar.open("Log in to vote on posts", "Dismiss", { duration: 1500 });
//   }
// }

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
