import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validator, Validators } from '@angular/forms';
import { SubredditsService } from '../subreddits.service';
import { SignupService } from '../signup.service';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Storage } from '../storage';
import { ProfileService } from '../profile.service';

@Component({
  selector: 'app-newpostform',
  templateUrl: './newpostform.component.html',
  styleUrls: ['./newpostform.component.css']
})
export class NewpostformComponent implements OnInit {
  communities: any[] = [];
  profile: any;
  // communities: any[] = [
  //   {value: '6247263303a4c16c6d6470de', viewValue: 'Sociology'},
  //   {value: '6247263303a4c16c6d6470de', viewValue: 'Pizza'},
  //   {value: '6247263303a4c16c6d6470de', viewValue: 'Tacos'},
  // ];
  selectedCommunity: string = "";
  form: FormGroup = new FormGroup({});
  constructor(private service1: SubredditsService, private service2: ProfileService, private signupService: SignupService, private fb: FormBuilder, private snackBar: MatSnackBar) {
    this.form = this.fb.group({
      username: ['', [Validators.required]],
      title: ['', [Validators.required]],
      body: ['', [Validators.required]],
    })
  }

  getCommunities() {
    this.service1.getSubreddits().subscribe((response: any) => {
      console.log(response.data.communities);
      if (response.status == 200) {
        this.communities = response.data.communities;
      }
      else {
      }
    });

    this.service2.getProfile(Storage.username).subscribe((response: any) => {
      console.log(response);
      console.log(response.status);
      
      if (response.status == 200) {
        console.log(response.data.Profile.username);
        console.log(response.data.Profile.email);
        this.profile = {
          "username": response.data.Profile.username,
          "email": response.data.Profile.email,
        }
      }
    });
  }

  get f() {
    return this.form.controls;
  }

  ngOnInit(): void {
    this.getCommunities();
  }

  createPost(community: string, title: string, body: string) {
    // console.log("testing the code")
    console.log("new post: " + title + " " + "community: " + community + " " + body);
    if (Storage.isLoggedIn) {
    this.signupService.createPost(this.profile.username, community, title, body).subscribe((response: any) => {
      console.log(response);
      if(response.status == 201 && response.message == "success"){
        this.snackBar.open("New post created.", "Dismiss", { duration: 1500 });
       }
      else {
        // Something else is wrong
        this.snackBar.open("Failed to create new post", "Dismiss", { duration: 1500 });
      }
    });
  }
  else {
    this.snackBar.open("Log in to vote on posts", "Dismiss", { duration: 1500 });
  }
  }
}
