import { Component, OnInit } from '@angular/core';
import { SignupService } from '../signup.service';
import { FormBuilder, FormGroup, Validator, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ProfileService } from '../profile.service';
import { Storage } from '../storage';

@Component({
  selector: 'app-newsubredditsform',
  templateUrl: './newsubredditsform.component.html',
  styleUrls: ['./newsubredditsform.component.css']
})
export class NewsubredditsformComponent implements OnInit {

  profile: any
  form: FormGroup = new FormGroup({});
  constructor(private signupService: SignupService, private fb: FormBuilder, private snackBar: MatSnackBar, private service: ProfileService) {
    this.form = this.fb.group({
      username: ['', [Validators.required]],
      name: ['', [Validators.required]],
      description: ['', [Validators.required]]
    })
  }
  
  get f() 
  {  
  return this.form.controls;
  }

  ngOnInit(): void 
  {
    this.getUsername();
  }

  getUsername() {
    this.service.getProfile(Storage.username).subscribe((response: any) => {
      console.log(response);
      console.log(response.status);
      
      if (response.status == 200) {
        console.log(response.data.Profile.username);
        this.profile = {
          "username": response.data.Profile.username,
        }
      }
    });
  }


  createSubreddit(username: string, name: string, description: string)
  {
    console.log("new subreddit: " + username + " " + name + " " + description);
    this.signupService.createcommunity(username, name, description).subscribe((response: any) => {
    console.log(response.status);
    console.log(response.message);
     if(response.status == 201 && response.message == "success"){
      this.snackBar.open("New subreddit created.", "Dismiss"), { duration: 2000 };
     }
     else if(response.status == 200 && response.message == "failure"){
       this.snackBar.open("Community with that name already exists","Dismiss"), { duration:4000};
     }
     else {
      // Something else is wrong
      this.snackBar.open("Something is wrong", "Alert Adminstration"), { duration: 2000 };
     }
    })
  }
}

