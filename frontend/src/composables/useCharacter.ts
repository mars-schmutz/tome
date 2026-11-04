import { ref } from "vue";
import {
  CreateCharacter,
  ListCharacters,
  SelectCharacter,
  SetName,
} from "../../wailsjs/go/main/CharacterApi";
import { models } from "../../wailsjs/go/models";

export const character = ref<models.Character | null>(null);
export const characters = ref<models.CharacterSummary[]>([]);
export const error = ref<string>("");

const run = async (call: () => Promise<models.Character>): Promise<boolean> => {
  try {
    character.value = await call();
    error.value = "";
    return true;
  } catch (err) {
    error.value = String(err);
    return false;
  }
};

export const loadList = async () => {
  try {
    characters.value = await ListCharacters();
  } catch (err) {
    error.value = String(err);
  }
};
export const setName = async (name: string) => {
  if (await run(() => SetName(name))) {
    await loadList();
  }
};
export const pick = (id: number) => run(() => SelectCharacter(id));
export const createNew = async () => {
  if (await run(() => CreateCharacter())) {
    await loadList();
  }
};
