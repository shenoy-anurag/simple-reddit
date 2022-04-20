import { Component, OnInit, Inject } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ActivatedRoute, Router } from '@angular/router';
import { PostsService } from '../posts.service';
import { Storage } from '../storage';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';
import { EditpostComponent } from '../editpost/editpost.component';

@Component({
  selector: 'app-postpage',
  templateUrl: './postpage.component.html',
  styleUrls: ['./postpage.component.css']
})
export class PostpageComponent implements OnInit {

  constructor(private dialog: MatDialog, private router: Router, private route: ActivatedRoute, private service: PostsService, private snackbar: MatSnackBar, private matdialog: MatDialog) {
   }
  post_id: string;
  post: any;
  posts: any[];
  comments: any[];

  ngOnInit(): void {
    this.post_id = this.route.snapshot.paramMap.get('postID');
    this.getPosts();
    this.getCommentsForPost();
  }

  editPost(post: any) {
    console.log("EDITING A POST");
    if (Storage.isLoggedIn) {
      let dialogRef = this.dialog.open(EditpostComponent, {data: {post: post}});

      // Update page
      console.log("HELP");
      this.ngOnInit();
      console.log("test");
    }
    else {
      this.snackbar.open("Log In to edit posts", "Dismiss", { duration: 1500 });
    }
    
  }

  saveComment(comment_id: string) {
    if (Storage.isLoggedIn) {
      this.service.saveComment(Storage.username, comment_id).subscribe((response: any) => {
        if (response.status == 200) {
          this.snackbar.open("Comment saved", "Dismiss", { duration: 1500 });           
        }
      });
    }
    else {
      this.snackbar.open("Log in to save comments", "Dismiss", { duration: 1500 });
    }
  }

  upvoteComment(comment_id: string) {
    if (Storage.isLoggedIn) {
      this.service.voteComment(comment_id, Storage.username, "upvote").subscribe((response: any) => {
        if (response.status == 200 && response.message == "success") {
          // Refresh page
          this.getPosts();
          this.getCommentsForPost();
        }
      });
    }
    else {
      this.snackbar.open("Log in to upvote comments", "Dismiss", { duration: 1500 });
    }
  }

  downvoteComment(comment_id: string) {
    if (Storage.isLoggedIn) {
      this.service.voteComment(comment_id, Storage.username, "downvote").subscribe((response: any) => {
        if (response.status == 200 && response.message == "success") {
          // Refresh page
          this.getPosts();
          this.getCommentsForPost();
        }
      });
    }
    else {
      this.snackbar.open("Log in to upvote comments", "Dismiss", { duration: 1500 });
    }
  }

  getPosts(): void {
    this.service.getPosts().subscribe((response: any) => {
      if (response.status == 200) {
        this.posts = response.data.posts;
        this.posts.forEach(p => {
          if (p._id == this.post_id) {
            this.post = p;
          }
        });
      }
    });
  }

  addComment(post_id: string, body: string): void {
    if (Storage.isLoggedIn) {
      this.service.createComment(Storage.username, post_id, null, body).subscribe((response: any) => {
        if (response.status == 200) {
          this.snackbar.open("Comment added", "Dismiss", { duration: 500 });

          // refresh page to update
          this.getPosts();
          this.getCommentsForPost();
        }
      });
    }
    else {
      this.snackbar.open("Log in to vote on posts", "Dismiss", { duration: 1500 });
    }
  }
  
  getCommentsForPost(): void {
    // this.comments =
    this.service.getComments(this.post_id).subscribe((response: any) => {
      if (response.status == 200) {
        this.comments = response.data.comments;
      }
    })
  }

  downvotePost(id: string) {
    if (Storage.isLoggedIn) {
      this.service.votePost(id, Storage.username, -1).subscribe((response: any) => {
        if (response.status == 200) {
          this.getPosts();
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
    if (Storage.isLoggedIn) {
      this.service.deletePost(id).subscribe((response: any) => {
        console.log(response);
        if (response.status == 200) {
          this.snackbar.open("Post Deleted", "Dismiss", {duration: 1500 });

          // update posts
          this.getPosts

          // navigate to home
          this.router.navigate(['home']);
        }
      });
    }
    else {
      this.snackbar.open("You need to be logged in to delete posts", "Dismiss", {duration: 1500});
    }
  }

  togglePostSave(post_id: string) {
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
