/// <reference path="node_modules/angular2/typings/browser.d.ts" />

import {bootstrap} from 'angular2/platform/browser';
import {Component} from 'angular2/core';

@Component({
  selector: 'chat',
  template: `
    <h1>chat on</h1>
  `
})
class Chat {

}

bootstrap(Chat);
