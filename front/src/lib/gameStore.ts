import { writable } from 'svelte/store';
import { browser } from '$app/environment';
import type { Clue, Game } from '$lib/types';

const defaultClue: Clue = {
	question: '',
	answer: '',
	value: 0
};

const defaultGame: Game = {
	rounds: []
};

export const currentGame = writable<Game>(
	(browser && JSON.parse(localStorage.getItem('currentGame'))) || defaultGame
);
export const currentClue = writable(defaultClue);
export const clueActive = writable(false);
export const currentPlayer = writable();

currentGame.subscribe((value) => {
	if (browser) localStorage.setItem('currentGame', JSON.stringify(value));
});

export function resetClue() {
	clueActive.set(false);
}
