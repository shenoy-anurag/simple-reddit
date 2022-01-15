import { Component, OnInit } from '@angular/core';
import { PostsService } from '../posts.service';

@Component({
  selector: 'app-posts',
  template: `
        <h2>{{title}}</h2>
        <ul>
             <li *ngFor="let post of posts">
                 {{ post }}
             </li>
        </ul>
    `,
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

}
