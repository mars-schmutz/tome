export interface CharacterBase {
  name: string;
  class: string;
  race: string;
  background: string;
  alignment: string;
}

export interface LedgerEntry {
  id: string;
  type: string;
  amount: number;
  source: string;
  description: string;
  timestamp: string;
}
