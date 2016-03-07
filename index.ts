/// <reference path="node_modules/angular2/typings/browser.d.ts" />

import {bootstrap} from 'angular2/platform/browser';
import {
  Component,
  OnInit
} from 'angular2/core';

@Component({
  selector: 'chat',
  template: `
    <h1>chat on</h1>
    <input type="text" [(ngModel)]="msg" (keyup.enter)="go(msg)"/>
    <button type="submit" (submit)="go(msg)">go</button>

    <div>
      <ul>
        <li *ngFor="#m of msgs">{{m}}</li>
      </ul>
    </div>
  `
})
class Chat implements OnInit {
  msg: string = 'yo';
  msgs: string[] = [];
  ws: WebSocket;

  ngOnInit() {
    this.ws = new WebSocket('ws://localhost:3333');
  }

  go(msg: string) {
    console.log(msg);

    this.ws.send(msg);
  }
}

bootstrap(Chat);
