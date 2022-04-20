import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { PostsService } from '../posts.service';
import { Storage } from '../storage';
import { SubredditsService } from '../subreddits.service';

@Component({
  selector: 'app-editpost',
  templateUrl: './editpost.component.html',
  styleUrls: ['./editpost.component.css']
})
export class EditpostComponent implements OnInit {

  constructor(private snackbar: MatSnackBar, private fb: FormBuilder, private postservice: PostsService, private subredditservice: SubredditsService, @Inject(MAT_DIALOG_DATA) public data: any) {
    this.form = this.fb.group({
      username: ['', [Validators.required]],
      title: ['', [Validators.required]],
      body: ['', [Validators.required]],
      community: ['', [Validators.required]],
    })
   }

  communities: any[] = [];
  selectedCommunity: string = "";
  form: FormGroup = new FormGroup({});

  ngOnInit(): void {
    this.getCommunities();
  }
  
  editPost(post_id: string, community: string, title: string, body: string) {
    if (Storage.isLoggedIn) {
      this.postservice.editPost(post_id, Storage.username, title, body).subscribe((response: any) => {
        if (response.status == 200) {
          this.snackbar.open("Post Edited", "Dismiss", { duration: 1500 });    
        }
      });
    }
    else {
      this.snackbar.open("Please log in to edit posts", "Dismiss", { duration: 1500 });
    }
  }

  get f() {
    return this.form.controls;
  }

  getCommunities() {
    this.subredditservice.getSubreddits().subscribe((response: any) => {
      if (response.status == 200) {
        this.communities = response.data.communities;
      }
    });
  }

}
