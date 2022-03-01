import { Component, OnInit } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { SignupService } from '../signup.service';


@Component({
  selector: 'app-signupform',
  templateUrl: './signupform.component.html',
  styleUrls: ['./signupform.component.css']
})
export class SignupformComponent implements OnInit {
  emailFormControl = new FormControl('', [Validators.required, Validators.email]);
  // usernameFormControl = new FormControl('', [this.chckUsername()])

  constructor(private signupService: SignupService) { }

  ngOnInit(): void {
  }

  checkUsername(username: string) {
    this.signupService.checkUsername(username).subscribe((response: any) => {
      if (response.status == 200 && response.message == "success") {
        return false;
      }
      else {
        return true;
      }
    });
  }

  getSignUp(first: string, last: string, username: string, email: string, password: string): void {
    console.log(`sign up attempt with: ${first} ${last} ${username} ${email} ${password}`);
    this.signupService.addNewAccount(email, username, password, first+ " " + last).subscribe((response: any) => {
      console.log(response);
    });
  }

  checkPassword(password: string, password2: string) {
    return password != password2;
  }
}
