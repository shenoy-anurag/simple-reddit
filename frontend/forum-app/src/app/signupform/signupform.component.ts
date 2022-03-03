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

export function passwordConValidator(password: any): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const value = control.value;
    console.log("password: " + password + " " + "pass2: " + value);
    if (!value) {
      return null;
    }

    console.log(password);
    console.log(value);
    const isPasswordValid = password == value;

    return !isPasswordValid ? { passwordValid: true } : null;
  };
}

export function usernameValidator(signupService: any): ValidatorFn {
  let temp: any = true;
  return (control: AbstractControl): ValidationErrors | null => {
    const value = control.value;
    console.log("this is the value " + value)

    if (!value) {
      return null;
    }

    signupService.checkUsername(value).subscribe((response: any) => {
      console.log(response)
      if (response.status == 200 && response.message == "success" && response.data.usernameAlreadyExists == true) {
        console.log("BREAK");
        temp = null;
      }
      else {
        console.log("VALID");
        temp = { usernameValid: true };
      }
    });

    return temp;
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
      firstname: ['', [Validators.required]],
      lastname: ['', [Validators.required]],
      username: ['', [Validators.required, usernameValidator(signupService)]],
      password: ['', [Validators.required, passwordValidator(), Validators.minLength(6)]],
      password2: ['', [Validators.required, passwordConValidator(this.password)]],
      email: ['', [Validators.required, Validators.email]]
    })
  }

  ngOnInit(): void {
  }

  get f() {
    return this.form.controls;
  }

  get password() {
    return this.form.controls['password'];
  }

  getSignUp(first: string, last: string, username: string, email: string, password: string): void {
    console.log(`sign up attempt with: ${first} ${last} ${username} ${email} ${password}`);
    this.signupService.addNewAccount(email, username, password, first + " " + last).subscribe((response: any) => {
      console.log(response);
    });
  }
}

