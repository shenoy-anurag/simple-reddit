import { Component, OnInit } from '@angular/core';
import { FormControl, Validators, FormArray, FormBuilder, ValidatorFn, AbstractControl, Validator, ValidationErrors, FormGroup } from '@angular/forms';
import { SignupService } from '../signup.service';

export function passwordValidator(): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const value = control.value;

    if (!value) {
      return null;
    }

    const hasUpperCase = /[A-Z]+/.test(value);

    const hasLowerCase = /[a-z]+/.test(value);

    const hasNumeric = /[0-9]+/.test(value);

    const hasMatch = "password" == "password";

    const isPasswordValid = hasUpperCase && hasLowerCase && hasNumeric && hasMatch;

    return !isPasswordValid ? { passwordValid: true } : null;
  };
}

export function usernameValidator(signupService: any, username: any): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const value = control.value;

    if (!value) {
      return null;
    }

    signupService.checkUsername(username).subscribe((response: any) => {
      console.log(response);
      if (response.status == 200 && response.message == "success" && response.data == {"data": {"usernameAlreadyExists": true}}) {
        return null;
      }
      else {
        return { usernameValid: true };
      }
    });

    return null;
  };
}


@Component({
  selector: 'app-signupform',
  templateUrl: './signupform.component.html',
  styleUrls: ['./signupform.component.css']
})

export class SignupformComponent implements OnInit {
  form: FormGroup = new FormGroup({});

  constructor(private signupService: SignupService, private fb: FormBuilder) {
    this.form = this.fb.group({
      username: ['', [Validators.required, usernameValidator(signupService, this.username)]],
      password: ['', [Validators.required, passwordValidator(), Validators.minLength(6)]],
      email: ['', [Validators.required, Validators.email]]
    })
  }

  ngOnInit(): void {
  }

  get f() {
    return this.form.controls;
  }

  get username() {
    return this.form.controls['username']
  }

  get email() {
    return this.form.controls['email'];
  }

  get password() {
    return this.form.controls['password'];
  }

  get password2() {
    return this.form.controls['password2'];
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
    this.signupService.addNewAccount(email, username, password, first + " " + last).subscribe((response: any) => {
      console.log(response);
    });
  }
}

