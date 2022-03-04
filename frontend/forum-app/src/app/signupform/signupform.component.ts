import { Component, OnInit } from '@angular/core';
import { FormControl, Validators, FormArray, FormBuilder, ValidatorFn, AbstractControl, Validator, ValidationErrors, FormGroup } from '@angular/forms';
import { SignupService } from '../signup.service';
import { MatSnackBar, MatSnackBarModule } from '@angular/material/snack-bar';

export function checkIfMatchingPasswords(passwordKey: string, passwordConfirmationKey: string) {
  return (group: FormGroup) => {
    let passwordInput = group.controls[passwordKey],
        passwordConfirmationInput = group.controls[passwordConfirmationKey];
    if (passwordInput.value !== passwordConfirmationInput.value) {
      return passwordConfirmationInput.setErrors({notEquivalent: true})
    }
    else {
        return passwordConfirmationInput.setErrors(null);
    }
  }
}


export function passwordValidator(): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const value = control.value;

    if (!value) {
      return null;
    }

    const hasUpperCase = /[A-Z]+/.test(value);

    const hasLowerCase = /[a-z]+/.test(value);

    const hasNumeric = /[0-9]+/.test(value);

    const isPasswordValid = hasUpperCase && hasLowerCase && hasNumeric;

    return !isPasswordValid ? { passwordValid: true } : null;
  };
}

export function ConfirmedValidator(password: any): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const value = control.value;

    if (!value) {
      return null;
    }
    
    console.log(password);
    console.log(value);
    console.log(password == value);

    return password == value ? { passwordValid: true } : null;
  };
}

export function usernameValidator(signupService: any): ValidatorFn {
  let temp: any = null;
  return (control: AbstractControl): ValidationErrors | null => {
    const value = control.value;

    if (!value) {
      return null;
    }

    signupService.checkUsername(value).subscribe((response: any) => {
      if (response.data.usernameAlreadyExists == false) {
        temp = true;
      }
      else if (response.status == 200 && response.message == "success" && response.data.usernameAlreadyExists == true) {
        temp = false;
      }
    });
    return temp ? { usernameValid: true } : null;
  };
}


@Component({
  selector: 'app-signupform',
  templateUrl: './signupform.component.html',
  styleUrls: ['./signupform.component.css']
})

export class SignupformComponent implements OnInit {
  form: FormGroup = new FormGroup({});

  constructor(private snackBar: MatSnackBar, private signupService: SignupService, private fb: FormBuilder) {
    this.form = this.fb.group({
      firstname: ['', [Validators.required]],
      lastname: ['', [Validators.required]],
      username: ['', [Validators.required, usernameValidator(signupService)]],
      password: ['', [Validators.required, passwordValidator(), Validators.minLength(6)]],
      password2: ['', [Validators.required]],
      email: ['', [Validators.required, Validators.email]]
    },
    {
      validator: checkIfMatchingPasswords('password', 'password2')
    })
  }

  ngOnInit(): void {
  }

  get f() {
    return this.form.controls;
  }

  get password() {
    console.log("GETTING: "+ this.form.get('password'));
    return this.form.get('password');
  }

  // get password2() {
  //   return this.form.controls['password2'];
  // }

  getSignUp(first: string, last: string, username: string, email: string, password: string): void {
    console.log(`sign up attempt with: ${first} ${last} ${username} ${email} ${password}`);
    this.signupService.addNewAccount(email, username, password, first + " " + last).subscribe((response: any) => {
      console.log(response);
      if (response.status == 201 && response.message == "success") {
        this.snackBar.open("Sign Up Successfull", "Dismiss", { duration: 1000});
      }
      
      else if (response.status == 200 && response.data.usernameAlreadyExists == true) {
          this.snackBar.open("Sign Up Failed. Username is already taken.", "Dismiss", { duration: 2000});
      }
    });
  }
}

