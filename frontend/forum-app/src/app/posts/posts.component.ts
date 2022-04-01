import { Component, OnInit } from '@angular/core';
import { PostsService } from '../posts.service';

@Component({
  selector: 'app-posts',
  templateUrl: './posts.component.html',
  styleUrls: ['./posts.component.css']
})
export class PostsComponent implements OnInit {
  title = "List of Posts";
  posts: any[] = [];

  constructor(private service: PostsService) {
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

  deletePost(owner: string, title: string) { 
    console.log("Deleting post: " + title);
  }
}
