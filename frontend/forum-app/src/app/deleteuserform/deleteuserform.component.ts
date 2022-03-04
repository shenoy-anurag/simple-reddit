import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { SignupService } from '../signup.service';

@Component({
  selector: 'app-deleteuserform',
  templateUrl: './deleteuserform.component.html',
  styleUrls: ['./deleteuserform.component.css']
})
export class DeleteuserformComponent implements OnInit {

  form: FormGroup = new FormGroup({});
  constructor(private signupService: SignupService, private snackBar: MatSnackBar, private fb: FormBuilder) {
    this.form = this.fb.group({
      username: ['', [Validators.required]],
      password: ['', [Validators.required]]
    })
   }

  ngOnInit(): void {
  }

  get f() {
    return this.form.controls;
  }
  deleteUser(useranme: string) {

  }

  checkAuth(username: string, password: string) {
    this.signupService.checkLogIn(username, password).subscribe((response: any) => {

      if (response.status == 200 && response.message == "success") {
        this.snackBar.open(username + " deleted", "Dismiss");
        // this.signupService.deleteUser(username).subscribe((response: any) => {
        // });
      }
      else if (response.status == 200 && response.message == "failure" && response.data.data == 'Incorrect Credentials') {
        // Prompt user, incorrect login
        this.snackBar.open("Incorrect Credentials", "Dismiss", { duration: 2000 });
      }
      else {
        // Something else is wrong
        this.snackBar.open("Something is wrong", "Alert Adminstration"), { duration: 2000 };
      }
    })
  }
}
