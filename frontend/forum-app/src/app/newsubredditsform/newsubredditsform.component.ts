import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validator, Validators } from '@angular/forms';
import { Title } from '@angular/platform-browser';
import { SignupService } from '../signup.service';


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
      description: ['', [Validators.required]],
    })
  }
    get f() {
      return this.form.controls;
    }

    createSubreddit()
    {
        
    }

   ngOnInit(): void {
   }

  }

