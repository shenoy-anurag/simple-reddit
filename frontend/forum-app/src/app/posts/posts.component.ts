import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { PostsService } from '../posts.service';
import { Storage } from '../storage';

@Component({
  selector: 'app-posts',
  templateUrl: './posts.component.html',
  styleUrls: ['./posts.component.css']
})
export class PostsComponent implements OnInit {
  title = "List of Posts";
  posts: any[] = [];

  constructor(private service: PostsService, private snackbar: MatSnackBar) {
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

  deletePost(id: string, title: string) { 
    console.log("Deleting post: " + title + "id: " + id);
    if (Storage.isLoggedIn) {
      this.service.deletePost(id).subscribe((response: any) => {
        if (response.status == 200) {
          this.snackbar.open("Post Deleted", "Dismiss", {duration: 1500 });
        }
      });
    }
    else {
      this.snackbar.open("You need to be logged in to delete posts", "Dismiss", {duration: 1500});
    }
  }
}
