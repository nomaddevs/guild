import { Component } from '@angular/core';

import { Recruitment } from './recruitment';

declare var $ : any;

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  public lfm: Recruitment;

  constructor() {
  	this.lfm = new Recruitment();
    this.lfm.SetStatus('Druid', 'Balance');
    this.lfm.SetStatus('Druid', 'Feral');
    this.lfm.SetStatus('Monk', 'Mistweaver');
    this.lfm.SetStatus('Monk', 'Windwalker');
    this.lfm.SetStatus('Mage', '*');
    this.lfm.SetStatus('Paladin', 'Holy');
    this.lfm.SetStatus('Paladin', 'Retribution');
    this.lfm.SetStatus('Warrior', 'Arms');
    this.lfm.SetStatus('Warrior', 'Fury');
    this.lfm.SetStatus('DemonHunter', 'Havoc');
  }

  public about() {
    $('#collapse-about').collapse('show');
    $('#collapse-recruitment').collapse('hide');
    $('#collapse-faq').collapse('hide');
  }

  public recruitment() {
    $('#collapse-about').collapse('hide');
    $('#collapse-recruitment').collapse('show');
    $('#collapse-faq').collapse('hide');
  }

  public faq() {
    $('#collapse-about').collapse('hide');
    $('#collapse-recruitment').collapse('hide');
    $('#collapse-faq').collapse('show');
  }
}
