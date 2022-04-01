import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validator, Validators } from '@angular/forms';
import { Title } from '@angular/platform-browser';
import { SignupService } from '../signup.service';

@Component({
  selector: 'app-newpostform',
  templateUrl: './newpostform.component.html',
  styleUrls: ['./newpostform.component.css']
})
export class NewpostformComponent implements OnInit {

  foods: any[] = [
    {value: 'steak-0', viewValue: 'Steak'},
    {value: 'pizza-1', viewValue: 'Pizza'},
    {value: 'tacos-2', viewValue: 'Tacos'},
  ];

  form: FormGroup = new FormGroup({});
  constructor(private signupService: SignupService, private fb: FormBuilder) {
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

  createPost(username: string, title: string, body: string) {
    console.log("new post: " + title + " " + body);
    this.signupService.createPost(username, "621d4aef4d5510eeee3ad715", title, body).subscribe((response: any) => {
      console.log(response);
    })
  }

}
