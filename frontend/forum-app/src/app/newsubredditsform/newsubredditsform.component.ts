import { Component, OnInit } from '@angular/core';
import { SignupService } from '../signup.service';
import { FormBuilder, FormGroup, Validator, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-newsubredditsform',
  templateUrl: './newsubredditsform.component.html',
  styleUrls: ['./newsubredditsform.component.css']
})
export class NewsubredditsformComponent implements OnInit {

  form: FormGroup = new FormGroup({});
  constructor(private signupService: SignupService, private fb: FormBuilder, private snackBar: MatSnackBar) {
    this.form = this.fb.group({
      user_id: ['', [Validators.required]],
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
  }

  createSubreddit(user_id: string, name: string, description: string)
  {
<<<<<<< HEAD
    console.log("new subreddit: " + name + " " + description + "by " + user_id);
    this.signupService.createcommunity(user_id, name, description).subscribe((response: any) => {
      console.log(response);
    })
=======
    console.log("new subreddit: " + user_id + " " + name + " " + description);
    this.signupService.createcommunity(user_id, name, description).subscribe((response: any) => {
    console.log(response);
     if(response.status == 200 && response.message == "success"){
      this.snackBar.open("New subreddit created."), { duration: 2000 };
     }
    else {
      // Something else is wrong
      this.snackBar.open("Something is wrong", "Alert Adminstration"), { duration: 2000 };
    }
     })
>>>>>>> main
  }
}

