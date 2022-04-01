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
  constructor(public signupService: SignupService, public snackBar: MatSnackBar, public fb: FormBuilder) {
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

  createdeletesubreddit(username: string, name: string) {
    this.signupService.deletecommunity(username, name).subscribe((response: any) => {
      console.log(response);
      if(response.status == 200 && response.message == "success"){
        this.snackBar.open("Subreddit deleted."), { duration: 2000 };
       }
      else {
        // Something else is wrong
        this.snackBar.open("Something is wrong", "Alert Adminstration"), { duration: 2000 };
      }
       })

  }
}
