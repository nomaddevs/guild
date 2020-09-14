import { Component } from '@angular/core';

import { Recruitment } from './recruitment';

declare var $ : any;
//import * as $ from 'jquery';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  public lfm: Recruitment;

  constructor() {
  	this.lfm = new Recruitment();
    this.lfm.SetStatus('DeathKnight', 'Frost');
    this.lfm.SetStatus('DeathKnight', 'Unholy');
  	this.lfm.SetStatus('DemonHunter', 'Havoc');
  	this.lfm.SetStatus('Druid', 'Restoration');
  	this.lfm.SetStatus('Hunter', 'Beast Mastery');
  	this.lfm.SetStatus('Mage', '*');
  	this.lfm.SetStatus('Monk', 'Mistweaver');
    this.lfm.SetStatus('Paladin', 'Retribution');
    this.lfm.SetStatus('Paladin', 'Holy');
  	this.lfm.SetStatus('Priest', '*');
    this.lfm.SetStatus('Rogue', 'Subtlety');
  	this.lfm.SetStatus('Shaman', '*');
  	this.lfm.SetStatus('Warlock', '*');
  	this.lfm.SetStatus('Warrior', '*');
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
