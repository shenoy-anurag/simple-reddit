import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-contentpolicy',
  templateUrl: './contentpolicy.component.html',
  styleUrls: ['./contentpolicy.component.css']
})
export class ContentpolicyComponent implements OnInit {

  constructor() { }

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
