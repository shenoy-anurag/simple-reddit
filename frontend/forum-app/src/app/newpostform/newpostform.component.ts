import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validator, Validators } from '@angular/forms';
import { Title } from '@angular/platform-browser';
import { SignupService } from '../signup.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-newpostform',
  templateUrl: './newpostform.component.html',
  styleUrls: ['./newpostform.component.css']
})
export class NewpostformComponent implements OnInit {

  communities: any[] = [
    {value: '6247263303a4c16c6d6470de', viewValue: 'Sociology'},
    {value: '6247263303a4c16c6d6470de', viewValue: 'Pizza'},
    {value: '6247263303a4c16c6d6470de', viewValue: 'Tacos'},
  ];
  selectedCommunity: string = "";
  form: FormGroup = new FormGroup({});
  constructor(private signupService: SignupService, private fb: FormBuilder, private snackBar: MatSnackBar) {
    this.form = this.fb.group({
      username: ['', [Validators.required]],
      title: ['', [Validators.required]],
      body: ['', [Validators.required]],
    })
  }

  get f() {
    return this.form.controls;
  }

  ngOnInit(): void {
  }

  createPost(username: string, community: string, title: string, body: string) {
    console.log("new post: " + title + " " + "community: " + community + " " + body);
    this.signupService.createPost(username, community, title, body).subscribe((response: any) => {
      console.log(response);
      if(response.status == 200 && response.message == "success"){
        this.snackBar.open("New post created."), { duration: 3000 };
       }
      else {
        // Something else is wrong
        this.snackBar.open("Something is wrong", "Alert Adminstration"), { duration: 3000 };
      }
    })
  }

}
