export namespace models {
	
	export class CharBase {
	    Name: string;
	    Class: string;
	    Race: string;
	    Background: string;
	    Alignment: string;
	    Level: number;
	
	    static createFrom(source: any = {}) {
	        return new CharBase(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Class = source["Class"];
	        this.Race = source["Race"];
	        this.Background = source["Background"];
	        this.Alignment = source["Alignment"];
	        this.Level = source["Level"];
	    }
	}
	export class CharHealth {
	    MaxHp: number;
	    CurrHp: number;
	    TempHp: number;
	    HitDice: string;
	    ArmorClass: number;
	    Initiative: number;
	    Speed: number;
	    DeathSaves: number;
	    DeathFails: number;
	
	    static createFrom(source: any = {}) {
	        return new CharHealth(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.MaxHp = source["MaxHp"];
	        this.CurrHp = source["CurrHp"];
	        this.TempHp = source["TempHp"];
	        this.HitDice = source["HitDice"];
	        this.ArmorClass = source["ArmorClass"];
	        this.Initiative = source["Initiative"];
	        this.Speed = source["Speed"];
	        this.DeathSaves = source["DeathSaves"];
	        this.DeathFails = source["DeathFails"];
	    }
	}
	export class CharScores {
	    Str: number;
	    Dex: number;
	    Con: number;
	    Int: number;
	    Wis: number;
	    Cha: number;
	
	    static createFrom(source: any = {}) {
	        return new CharScores(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Str = source["Str"];
	        this.Dex = source["Dex"];
	        this.Con = source["Con"];
	        this.Int = source["Int"];
	        this.Wis = source["Wis"];
	        this.Cha = source["Cha"];
	    }
	}
	export class Skill {
	    Bonus: number;
	    Level: number;
	
	    static createFrom(source: any = {}) {
	        return new Skill(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Bonus = source["Bonus"];
	        this.Level = source["Level"];
	    }
	}
	export class CharSkills {
	    Acrobatics: Skill;
	    AnimalHandling: Skill;
	    Arcana: Skill;
	    Athletics: Skill;
	    Deception: Skill;
	    History: Skill;
	    Insight: Skill;
	    Intimidation: Skill;
	    Investigation: Skill;
	    Medicine: Skill;
	    Nature: Skill;
	    Perception: Skill;
	    Performance: Skill;
	    Persuasion: Skill;
	    Religion: Skill;
	    SleightOfHand: Skill;
	    Stealth: Skill;
	    Survival: Skill;
	
	    static createFrom(source: any = {}) {
	        return new CharSkills(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Acrobatics = this.convertValues(source["Acrobatics"], Skill);
	        this.AnimalHandling = this.convertValues(source["AnimalHandling"], Skill);
	        this.Arcana = this.convertValues(source["Arcana"], Skill);
	        this.Athletics = this.convertValues(source["Athletics"], Skill);
	        this.Deception = this.convertValues(source["Deception"], Skill);
	        this.History = this.convertValues(source["History"], Skill);
	        this.Insight = this.convertValues(source["Insight"], Skill);
	        this.Intimidation = this.convertValues(source["Intimidation"], Skill);
	        this.Investigation = this.convertValues(source["Investigation"], Skill);
	        this.Medicine = this.convertValues(source["Medicine"], Skill);
	        this.Nature = this.convertValues(source["Nature"], Skill);
	        this.Perception = this.convertValues(source["Perception"], Skill);
	        this.Performance = this.convertValues(source["Performance"], Skill);
	        this.Persuasion = this.convertValues(source["Persuasion"], Skill);
	        this.Religion = this.convertValues(source["Religion"], Skill);
	        this.SleightOfHand = this.convertValues(source["SleightOfHand"], Skill);
	        this.Stealth = this.convertValues(source["Stealth"], Skill);
	        this.Survival = this.convertValues(source["Survival"], Skill);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Character {
	    ID: number;
	    Base: CharBase;
	    Scores: CharScores;
	    Health: CharHealth;
	    Skills: CharSkills;
	    Inspiration: number;
	    Proficiency: number;
	
	    static createFrom(source: any = {}) {
	        return new Character(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Base = this.convertValues(source["Base"], CharBase);
	        this.Scores = this.convertValues(source["Scores"], CharScores);
	        this.Health = this.convertValues(source["Health"], CharHealth);
	        this.Skills = this.convertValues(source["Skills"], CharSkills);
	        this.Inspiration = source["Inspiration"];
	        this.Proficiency = source["Proficiency"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

