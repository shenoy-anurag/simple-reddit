import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { PostsComponent } from './posts/posts.component';
import { PostsService } from './posts.service';
import { SubredditsService } from './subreddits.service';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgMaterialModule } from './ng-material/ng-material.module';
import { FormsModule } from '@angular/forms';
import { SignupformComponent } from './signupform/signupform.component';
import { AppRoutingModule} from './app-routing.module';
import { HttpClientModule } from '@angular/common/http';
import { NavbarComponent } from './navbar/navbar.component';
import { LoginComponent } from './login/login.component';
import { HomeComponent } from './home/home.component';
import { ProfileComponent } from './profile/profile.component';
import { SubredditsComponent } from './subreddits/subreddits.component';
import { DeleteuserformComponent } from './deleteuserform/deleteuserform.component';
import { NewpostformComponent } from './newpostform/newpostform.component';
import { NewsubredditsformComponent } from './newsubredditsform/newsubredditsform.component';
import { DeletesubredditsformComponent } from './deletesubredditsform/deletesubredditsform.component';
import { TermsandconditionsComponent } from './termsandconditions/termsandconditions.component';
import { PostpageComponent } from './postpage/postpage.component';
import { CommunitypageComponent } from './communitypage/communitypage.component';
import { PrivacypolicyComponent } from './privacypolicy/privacypolicy.component';
import { ContentpolicyComponent } from './contentpolicy/contentpolicy.component';
import { ModpolicyComponent } from './modpolicy/modpolicy.component';


@NgModule({
  declarations: [
    AppComponent,
    PostsComponent,
    SignupformComponent,
    NavbarComponent,
    LoginComponent,
    HomeComponent,
    ProfileComponent,
    SubredditsComponent,
    DeleteuserformComponent,
    NewpostformComponent,
    NewsubredditsformComponent,
    DeletesubredditsformComponent,
    TermsandconditionsComponent,
    PostpageComponent,
    CommunitypageComponent,
    PrivacypolicyComponent,
    ContentpolicyComponent,
    ModpolicyComponent,
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    FormsModule,
    NgMaterialModule,
    AppRoutingModule,
    HttpClientModule,
  ],
  exports: [
    NavbarComponent,
    PostsComponent,
  ],
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
  providers: [PostsService, SubredditsService],
  bootstrap: [AppComponent]
})
export class AppModule { }
