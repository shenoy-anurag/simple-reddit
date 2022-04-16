import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ActivatedRoute } from '@angular/router';
import { PostsService } from '../posts.service';
import { Storage } from '../storage';

@Component({
  selector: 'app-postpage',
  templateUrl: './postpage.component.html',
  styleUrls: ['./postpage.component.css']
})
export class PostpageComponent implements OnInit {

  constructor(private route: ActivatedRoute, private service: PostsService, private snackbar: MatSnackBar) { }
  post_id: string;
  post: any;
  posts: any[];

  ngOnInit(): void {
    this.post_id = this.route.snapshot.paramMap.get('postID');
    console.log("post id: ", this.post_id);
    this.getPosts();

    console.log("test");
  }

  getPosts(): void {
    this.service.getPosts().subscribe((response: any) => {
      if (response.status == 200) {
        this.posts = response.data.posts;
        
        this.posts.forEach(p => {
          console.log(p._id);

          if (p._id == this.post_id) {
            this.post = p;
          }
        });
      }
    });
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
}
