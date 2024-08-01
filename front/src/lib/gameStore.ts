import { writable } from "svelte/store";
import type { Clue } from "$lib/types";

const defaultClue: Clue = {
  question: "",
  answer: "",
  value: 0
};

export const currentClue = writable(defaultClue);
export const clueActive = writable(false);
export const currentPlayer = writable();

export function resetClue() {
    clueActive.set(false);
}
