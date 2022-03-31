import { Component, OnInit } from '@angular/core';
import { SignupService } from '../signup.service';
import { FormBuilder, FormGroup, Validator, Validators } from '@angular/forms';

@Component({
  selector: 'app-newsubredditsform',
  templateUrl: './newsubredditsform.component.html',
  styleUrls: ['./newsubredditsform.component.css']
})
export class NewsubredditsformComponent implements OnInit {

  form: FormGroup = new FormGroup({});
  constructor(private signupService: SignupService, private fb: FormBuilder) {
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
    // console.log("new subreddit: " + name + " " + description);
    this.signupService.createcommunity(user_id, name, description).subscribe((response: any) => {
    console.log(response);
    })
  }
}

