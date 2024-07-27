import { writable } from "svelte/store";
import type { Clue } from "$lib/Clues";

const defaultClue: Clue = {
  question: "",
  answer: "",
  value: 0
};

export const currentClue = writable(defaultClue);
export const clueActive = writable(false);

export function resetClue() {
    currentClue.set(defaultClue);
    clueActive.set(false);
}
