import { Component, OnInit, Inject, HostListener } from '@angular/core';
import { SignupService } from '../signup.service';
import { FormBuilder, FormGroup, Validator, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ProfileService } from '../profile.service';
import { Storage } from '../storage';
import { Router } from '@angular/router';

@Component({
  selector: 'app-newsubredditsform',
  templateUrl: './newsubredditsform.component.html',
  styleUrls: ['./newsubredditsform.component.css']
})
export class NewsubredditsformComponent implements OnInit {

  windowScrolled!: boolean;
  profile: any
  form: FormGroup = new FormGroup({});
  constructor(private router: Router, private signupService: SignupService,private service1: ProfileService , private fb: FormBuilder, private snackBar: MatSnackBar, private service: ProfileService) {
    this.form = this.fb.group({
      username: ['', [Validators.required]],
      name: ['', [Validators.required]],
      description: ['', [Validators.required]]
    })
  }
  
  @HostListener("window:scroll", [])

onWindowScroll() {
    if (window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop > 100) {
        this.windowScrolled = true;
    } 
   else if (this.windowScrolled && window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop < 10) {
        this.windowScrolled = false;
    }
}


scrollToTop() {
  (function smoothscroll() {
      var currentScroll = document.documentElement.scrollTop || document.body.scrollTop;
      if (currentScroll > 0) {
          window.requestAnimationFrame(smoothscroll);
          window.scrollTo(0, currentScroll - (currentScroll / 8));
      }
  })();
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
    // this.service.getProfile(Storage.username).subscribe((response: any) => {
      this.service1.getProfile(Storage.username).subscribe((response: any) => {
      console.log(response);
      console.log(response.status);
      
      if (response.status == 200) {
        // console.log(response.data.Profile.username);
        // console.log(response.data.Profile.email);
        this.profile = {
          "username": response.data.Profile.username,
          "email": response.data.Profile.email,
        }
      }
    });
  }


  createSubreddit(name: string, description: string)
  {
    // console.log("new subreddit: " + username + " " + name + " " + description);
    if (Storage.isLoggedIn) {
    this.signupService.createcommunity(this.profile.username, name, description).subscribe((response: any) => {
    console.log(response.status);
    console.log(response.message);
     if(response.status == 201 && response.message == "success"){
      this.snackBar.open("New subreddit created.", "Dismiss"), { duration: 2000 };

      // navigate to subreddits
      this.router.navigate(['subreddits']);
     }
     else if(response.status == 200 && response.message == "failure"){
       this.snackBar.open("Community with that name already exists","Dismiss"), { duration:4000};
     }
     else {
      // Something else is wrong
      this.snackBar.open("Something is wrong", "Alert Adminstration"), { duration: 2000 };
     }
     });
    }

  else{
    this.snackBar.open("Log in to vote on posts", "Dismiss"), { duration: 1500 };
  }
  }
}

