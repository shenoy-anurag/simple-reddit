import { Component, OnInit, Inject, HostListener } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { PostsService } from '../posts.service';
import { Storage } from '../storage';

@Component({
  selector: 'app-posts',
  templateUrl: './posts.component.html',
  styleUrls: ['./posts.component.css']
})

export class PostsComponent implements OnInit {
  windowScrolled!: boolean;
  title = "List of Posts";
  posts: any[] = [];
  comments = new Map([]);
  // comments: any[] = [];

  constructor(private router: Router, private service: PostsService, private snackbar: MatSnackBar) {}
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

        this.getAllComments(this.posts);
      }
      else {
        this.posts = []
      }
    });
  }

  ngOnInit(): void {
    this.getPosts();
  }

  // gets all the comments from a given post id
  getComments(post_id: string) {
    this.service.getComments(post_id).subscribe((response: any) => {
      if (response.status == 200) {
        return response.data;
      }
      else {
        return null;
      }
    });
  }

  // gets all the comments for all posts
  getAllComments(posts: any[]) {
    // loop through all posts
    posts.forEach((post) => {
      this.comments.set(post._id, this.getComments(post._id));
    });
  }

  gotoPost(post_id: string) {
    // Navigate to post page
    this.router.navigate(['/post/'+post_id]);
  }

  togglePostSave(post_id: string) {
    console.log("toggle save post id: " + post_id);
    if (Storage.isLoggedIn) {
      this.service.savePost(Storage.username, post_id).subscribe((response: any) => {
        console.log(response);
        if (response.status == 201) {
          this.snackbar.open("Post saved", "Dismiss", { duration: 500 });
        }
      });
    }
    else {
      this.snackbar.open("Log in to save posts", "Dismiss", { duration: 1500 });
    }
  }

  downvotePost(id: string) {
    if (Storage.isLoggedIn) {
      this.service.votePost(id, Storage.username, -1).subscribe((response: any) => {
        if (response.status == 200 && response.message == "success") {
          this.getPosts();
        }
      });
    }
    else {
      this.snackbar.open("Log in to vote on posts", "Dismiss"), { duration: 1500 };
    }
  }

  upvotePost(id: string) {
    if (Storage.isLoggedIn) {
      this.service.votePost(id, Storage.username, 1).subscribe((response: any) => {
        if (response.status == 200 && response.message == "success") {
          this.getPosts();
        }
      });
    }
    else {
      this.snackbar.open("Log in to vote on posts", "Dismiss"), { duration: 1500 };
    }
  }

  deletePost(id: string, title: string, postusername: string) { 
    console.log(id+","+title+","+postusername);
    if (Storage.isLoggedIn && postusername == Storage.username) {
      console.log("Deleting post: " + title + "id: " + id + " username: " + Storage.username);
      this.service.deletePost(id).subscribe((response: any) => {
        if (response.status == 200) {
          this.snackbar.open("Post Deleted", "Dismiss"), {duration: 1500 };

          // update posts
          this.getPosts
        }
      });
    }
    else if (postusername != Storage.username) {
      this.snackbar.open("You are not owner of this post.", "Dismiss"), {duration: 1500 };
    }
    else {
      this.snackbar.open("You need to be logged in to delete posts", "Dismiss"), {duration: 1500};
    }
  }
}
