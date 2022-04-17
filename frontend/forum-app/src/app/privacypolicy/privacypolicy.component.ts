import { Component, OnInit, Inject, HostListener } from '@angular/core';
import { Storage } from '../storage';
import { Router } from '@angular/router';
import { DOCUMENT } from '@angular/common';
@Component({
  selector: 'app-privacypolicy',
  templateUrl: './privacypolicy.component.html',
  styleUrls: ['./privacypolicy.component.css']
})
export class PrivacypolicyComponent implements OnInit {

  windowScrolled!: boolean;

  constructor(@Inject(DOCUMENT) private document: Document) { }

  @HostListener("window:scroll", [])

onWindowScroll() {
    if (window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop > 100) {
        this.windowScrolled = true;
    } 
   else if (this.windowScrolled && window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop < 10) {
        this.windowScrolled = false;
    }
}


scrollToTop() {
  (function smoothscroll() {
      var currentScroll = document.documentElement.scrollTop || document.body.scrollTop;
      if (currentScroll > 0) {
          window.requestAnimationFrame(smoothscroll);
          window.scrollTo(0, currentScroll - (currentScroll / 8));
      }
  })();
}


  copyMessage(val: string){
    const selBox = document.createElement('textarea');
    selBox.style.position = 'fixed';
    selBox.style.left = '0';
    selBox.style.top = '0';
    selBox.style.opacity = '0';
    selBox.value = val;
    document.body.appendChild(selBox);
    selBox.focus();
    selBox.select();
    document.execCommand('copy');
    document.body.removeChild(selBox);
  }
  
  public showMyMessage = false
  
  showMessageSoon1() {
    setTimeout(() => {
      this.showMyMessage = true
    }, 1000)
  }
  
  showMessageSoon2() {
    setTimeout(() => {
      this.showMyMessage = true
    }, 1500)
  }
  
  showMessageSoon3() {
    setTimeout(() => {
      this.showMyMessage = true
    }, 2000)
  }
  
  showMessageSoon4() {
    setTimeout(() => {
      this.showMyMessage = true
    }, 1500)
  }

  ngOnInit(): void {
  }

}
