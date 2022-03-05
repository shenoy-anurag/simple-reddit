import { Component, OnInit } from '@angular/core';
import { PostsService } from '../posts.service';

@Component({
  selector: 'app-posts',
  templateUrl: './posts.component.html',
  styleUrls: ['./posts.component.css']
})
export class PostsComponent implements OnInit {
  title = "List of Posts";
  posts;
  constructor(service: PostsService) {
    this.posts = service.getPosts();
  }

  ngOnInit(): void {
  }

  deletePost(owner: string, title: string) { 
    console.log("Deleting post: " + title);
  }
}
