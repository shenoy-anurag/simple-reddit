import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { SignupService } from '../signup.service';

@Component({
  selector: 'app-deletesubredditsform',
  templateUrl: './deletesubredditsform.component.html',
  styleUrls: ['./deletesubredditsform.component.css']
})
export class DeletesubredditsformComponent implements OnInit {

  form: FormGroup = new FormGroup({});
  constructor(private signupService: SignupService, private snackBar: MatSnackBar, private fb: FormBuilder) {
    this.form = this.fb.group({
      username: ['', [Validators.required]],
      name: ['', [Validators.required]]
    })
   }


  ngOnInit(): void {
  }

  deletesubreddit(useranme: string,name: string) {

  }
}
