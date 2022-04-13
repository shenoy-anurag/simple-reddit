import { Component, OnInit, Inject, HostListener } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { PostsService } from '../posts.service';
import { Storage } from '../storage';
import { DOCUMENT } from '@angular/common';

@Component({
  selector: 'app-posts',
  templateUrl: './posts.component.html',
  styleUrls: ['./posts.component.css']
})
export class PostsComponent implements OnInit {
  windowScrolled!: boolean;
  title = "List of Posts";
  posts: any[] = [];

  constructor(private router: Router, private service: PostsService, private snackbar: MatSnackBar, @Inject(DOCUMENT) private document: Document) {
  }
  @HostListener("window:scroll", [])
  
  onWindowScroll() {
    if (window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop > 100) {
        this.windowScrolled = true;
    } 
   else if (this.windowScrolled && window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop < 10) {
        this.windowScrolled = false;
    }
}


scrollToTop() {
  (function smoothscroll() {
      var currentScroll = document.documentElement.scrollTop || document.body.scrollTop;
      if (currentScroll > 0) {
          window.requestAnimationFrame(smoothscroll);
          window.scrollTo(0, currentScroll - (currentScroll / 8));
      }
  })();
}


  getPosts() {
    this.service.getPosts().subscribe((response: any) => {
      if (response.status == 200) {
        this.posts = response.data.posts;
      }
      else {
        this.posts = []
      }
    });
  }

  ngOnInit(): void {
    this.getPosts();
  }

  downvotePost(id: string) {
    if (Storage.isLoggedIn) {
      this.service.votePost(id, Storage.username, -1).subscribe((response: any) => {
        if (response.status == 200) {
          this.getPosts();
          this.snackbar.open("Downvote Succesfull", "Dismiss", { duration: 1500 });
        }

        if (response.status == 500) {
          this.snackbar.open("Downvote Unsuccesfull", "Dismiss", { duration: 1500 });
        }
      });
    }
    else {
      this.snackbar.open("Log in to vote on posts", "Dismiss", { duration: 1500 });
    }
  }

  upvotePost(id: string) {
    if (Storage.isLoggedIn) {
      this.service.votePost(id, Storage.username, 1).subscribe((response: any) => {
        if (response.status == 200) {
          this.getPosts();
        }
      });
    }
    else {
      this.snackbar.open("Log in to vote on posts", "Dismiss", { duration: 1500 });
    }
  }

  deletePost(id: string, title: string, postusername: string) { 
    console.log(id+","+title+","+postusername);
    if (Storage.isLoggedIn && postusername == Storage.username) {
      console.log("Deleting post: " + title + "id: " + id + " username: " + Storage.username);
      this.service.deletePost(id).subscribe((response: any) => {
        if (response.status == 200) {
          this.snackbar.open("Post Deleted", "Dismiss", {duration: 1500 });

          // update posts
          this.getPosts
        }
      });
    }
    else if (postusername != Storage.username) {
      this.snackbar.open("You are not owner of this post.", "Dismiss", {duration: 1500 });
    }
    else {
      this.snackbar.open("You need to be logged in to delete posts", "Dismiss", {duration: 1500});
    }
  }
}
