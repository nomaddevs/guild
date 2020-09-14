import { Component } from '@angular/core';

import { Recruitment } from './recruitment';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  public lfm: Recruitment;

  constructor() {
  	this.lfm = new Recruitment();
  	this.lfm.SetStatus('DeathKnight', '*');
  	this.lfm.SetStatus('DemonHunter', '*');
  	this.lfm.SetStatus('Druid', '*');
  	this.lfm.SetStatus('Hunter', '*');
  	this.lfm.SetStatus('Mage', '*');
  	this.lfm.SetStatus('Monk', 'Mistweaver');
  	this.lfm.SetStatus('Paladin', '*');
  	this.lfm.SetStatus('Priest', '*');
  	this.lfm.SetStatus('Shaman', '*');
  	this.lfm.SetStatus('Warlock', '*');
  	this.lfm.SetStatus('Warrior', '*');
  }
}
