import { PopupMessageComponent } from './popup-message/popup-message.component';
import { SignupformComponent } from './signupform/signupform.component';
import { HeaderComponent } from './header/header.component';
import { NgModule, Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';


const routes: Routes = [
  {path: 'header', component: HeaderComponent},
  {path: 'signupform', component: SignupformComponent},
  {path: 'popup-message', component: PopupMessageComponent},
  {path: '', redirectTo: '/header', pathMatch:'full'},

];



@NgModule({
  declarations: [],
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
