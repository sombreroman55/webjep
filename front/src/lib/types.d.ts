export interface Clue {
	question: string;
	answer: string;
	value?: number;
	answered?: boolean;
}

export interface Category {
	title: string;
	clues: Clue[];
}

export interface Round {
	categories: Category[];
}

export interface FinalRound {
	category: string;
	question: string;
	answer: string;
}

export interface Game {
	rounds: Round[];
	finalRound?: FinalRound;
}

export interface GameSettings {
	playFinalJep: boolean;
	baseClueValue: number;
	interClueInterval: number;
	interRoundMultiplier: number;
	numberOfRounds: number;
	categoriesPerRound: number;
	cluesPerCategory: number;
}

export interface Session {
	game: Game;
	settings: GameSettings;
}

export interface Player {
	name: string;
	score: number;
	isHost: boolean;
	mediaEnabled: boolean;
}
