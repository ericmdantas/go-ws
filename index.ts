/// <reference path="node_modules/angular2/typings/browser.d.ts" />

import {bootstrap} from 'angular2/platform/browser';
import {
  Component,
  OnInit
} from 'angular2/core';

@Component({
  selector: 'chat',
  template: `
    <h1>!</h1>
    <form (submit)="go(msg)">
      <input type="text" [(ngModel)]="msg"/>
      <button type="submit">go</button>
    </form>


    <div>
      <ul>
        <li *ngFor="#m of msgs">{{m.message}}</li>
      </ul>
    </div>
  `
})
class Chat implements OnInit {
  msg: string = 'yo';
  msgs: string[] = [];
  ws: WebSocket;

  private _addMsg(msg: string) {
    this.msgs.push(JSON.parse(msg));
  }

  ngOnInit() {
    this.ws = new WebSocket('ws://localhost:3333/socket');

    this.ws.onmessage = ({data}) => {
      this._addMsg(data);
    }
  }

  go(msg: string) {
    let _msgStringified = JSON.stringify({message: msg});
    this.ws.send(_msgStringified);
    this.msg = '';
  }
}

bootstrap(Chat);
