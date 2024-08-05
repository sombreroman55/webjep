import type { Clue, Category, Round, FinalRound, Game, GameSettings } from '$lib/types';

export function newGame(settings: GameSettings): Game {
	const rounds: Round[] = [];
	for (let i = 0; i < settings.numberOfRounds; i++) {
		const categories: Category[] = [];
		for (let j = 0; j < settings.categoriesPerRound; j++) {
			const clues: Clue[] = [];
			for (let k = 0; k < settings.cluesPerCategory; k++) {
				const clue: Clue = {
					question: '',
					answer: ''
				};
				clues.push(clue);
			}
			const category: Category = {
				title: '',
				clues: clues
			};
			categories.push(category);
		}
		const round: Round = {
			categories: categories
		};
		rounds.push(round);
	}

	const game: Game = {
		rounds: rounds
	};

	if (settings.playFinalJep) {
		const finalJep: FinalRound = {
			category: '',
			question: '',
			answer: ''
		};
		game.finalRound = finalJep;
	}

	return game;
}
