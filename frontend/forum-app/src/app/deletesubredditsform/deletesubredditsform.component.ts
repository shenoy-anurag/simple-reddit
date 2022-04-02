import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { SignupService } from '../signup.service';
import { Storage } from '../storage';

@Component({
  selector: 'app-deletesubredditsform',
  templateUrl: './deletesubredditsform.component.html',
  styleUrls: ['./deletesubredditsform.component.css']
})

export class DeletesubredditsformComponent implements OnInit {
  form: FormGroup = new FormGroup({});
  constructor(private signupService: SignupService, public snackBar: MatSnackBar, public fb: FormBuilder) {
    this.form = this.fb.group({
      username: ['', [Validators.required]],
      name: ['', [Validators.required]]
    })
   }
  
  get f() {
    return this.form.controls;
  }

  ngOnInit(): void {
  }

  getLoggedUsername() {
    if (Storage.isLoggedIn) {
      return Storage.username;
    }
    else {
      return "";
    }
  }

  deletesubreddit(username: string, name: string) {
    this.signupService.deletecommunity(username, name).subscribe((response: any) => {
      console.log(response);
      console.log(response.status);
      console.log(response.message);
      if(response.status == 200 && response.message == "success"){
        this.snackBar.open("Subreddit deleted.", "Dismiss"), { duration: 2000 };
      }
      else if(response.status == 400 && response.message == "failure"){
        this.snackBar.open("No such Community Exists", "Dismiss"), 
        {
          duration: 2000
        }
      }
      else if(response.status == 401) {
        this.snackBar.open("No Communities Owned by User", "Dismiss"), 
        {
          duration: 2000
        }
      }
      else {
        // Something else is wrong
        this.snackBar.open("Something is wrong", "Alert Adminstration"), { duration: 2000 };
      }
    })
  }
}
