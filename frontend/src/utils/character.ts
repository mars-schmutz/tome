import { models } from "../../wailsjs/go/models";

export type Ability = keyof models.CharScores;
export const abilities: readonly Ability[] = [
  "Str",
  "Dex",
  "Con",
  "Int",
  "Wis",
  "Cha",
];
