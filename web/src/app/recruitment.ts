export class Recruitment {
    public Classes: Array<WoWClass>;

    constructor() {
        this.Classes = [];

        this.Classes.push(new WoWClass('DeathKnight'));
        this.Classes.push(new WoWClass('DemonHunter'));
        this.Classes.push(new WoWClass('Druid'));
        this.Classes.push(new WoWClass('Hunter'));
        this.Classes.push(new WoWClass('Mage'));
        this.Classes.push(new WoWClass('Monk'));
        this.Classes.push(new WoWClass('Paladin'));
        this.Classes.push(new WoWClass('Priest'));
        this.Classes.push(new WoWClass('Rogue'));
        this.Classes.push(new WoWClass('Shaman'));
        this.Classes.push(new WoWClass('Warlock'));
        this.Classes.push(new WoWClass('Warrior'));
    }

    public Need(): boolean {
        return true;
    }

    public GetClass(name: string): WoWClass {
        for(let c of this.Classes) {
            if(c.Name == name) {
                return c;
            }
        }
    }

    public SetStatus(className: string, specName: string) {
        if(specName == "*") {
            this.GetClass(className).NeedAll();
            return;
        }
        this.GetClass(className).GetSpec(specName).Need = true;
    }
}

export class Spec {
    public Name: string;
    public Need: boolean;

    constructor(name: string) {
        this.Name = name;
        this.Need = false;
    }
}

export class WoWClass {
    public Name: string;
    public Specs: Array<Spec>;

    public GetSpec(name: string): Spec {
        for(let s of this.Specs) {
            if(s.Name == name) {
                return s;
            }
        }
    }

    public Need(): boolean {
        for(let spec of this.Specs) {
            if(spec.Need) {
                return true;
            }
        } 
    }

    public NeedAll() {
        for(let spec of this.Specs) {
            spec.Need = true;
        } 
    }

    constructor(name: string) {
        this.Specs = [];
        switch(name) {
            case 'DeathKnight':
                this.Name = name;
                this.Specs.push(new Spec('Blood'));
                this.Specs.push(new Spec('Frost'));
                this.Specs.push(new Spec('Unholy'));
                break;
            case 'DemonHunter':
                this.Name = name;
                this.Specs.push(new Spec('Vengeance'));
                this.Specs.push(new Spec('Havoc'));
                break;
            case 'Druid':
                this.Name = name;
                this.Specs.push(new Spec('Balance'));
                this.Specs.push(new Spec('Feral'));
                this.Specs.push(new Spec('Guardian'));
                this.Specs.push(new Spec('Restoration'));
                break;
            case 'Hunter':
                this.Name = name;
                this.Specs.push(new Spec('Beast Mastery'));
                this.Specs.push(new Spec('Marksmanship'));
                this.Specs.push(new Spec('Survival'));
                break;
            case 'Mage':
                this.Name = name;
                this.Specs.push(new Spec('Arcane'));
                this.Specs.push(new Spec('Fire'));
                this.Specs.push(new Spec('Frost'));
                break;
            case 'Monk':
                this.Name = name;
                this.Specs.push(new Spec('Brewmaster'));
                this.Specs.push(new Spec('Mistweaver'));
                this.Specs.push(new Spec('Windwalker'));
                break;
            case 'Paladin':
                this.Name = name;
                this.Specs.push(new Spec('Holy'));
                this.Specs.push(new Spec('Protection'));
                this.Specs.push(new Spec('Retribution'));
                break;
            case 'Priest':
                this.Name = name;
                this.Specs.push(new Spec('Discipline'));
                this.Specs.push(new Spec('Holy'));
                this.Specs.push(new Spec('Shadow'));
                break;
            case 'Rogue':
                this.Name = name;
                this.Specs.push(new Spec('Assassination'));
                this.Specs.push(new Spec('Outlaw'));
                this.Specs.push(new Spec('Subtlety'));
                break;
            case 'Shaman':
                this.Name = name;
                this.Specs.push(new Spec('Elemental'));
                this.Specs.push(new Spec('Enhancement'));
                this.Specs.push(new Spec('Restoration'));
                break;
            case 'Warlock':
                this.Name = name;
                this.Specs.push(new Spec('Affliction'));
                this.Specs.push(new Spec('Demonology'));
                this.Specs.push(new Spec('Destruction'));
                break;
            case 'Warrior':
                this.Name = name;
                this.Specs.push(new Spec('Arms'));
                this.Specs.push(new Spec('Fury'));
                this.Specs.push(new Spec('Protection'));
                break;
            default:
                break;
        }
    }
}
