<script lang="ts">
	import type { Game, Round, Player } from '$lib/types';
	import BoardView from './BoardView.svelte';
	import FinalRoundView from './FinalRoundView.svelte';
	import GameOverView from './GameOverView.svelte';
	import PlayerBarView from './PlayerBarView.svelte';

	export let game: Game;
	export let players: Player[];
	let finalRound = false;
	let allCluesUsed = false;
	let gameOver = false;
	let roundNum = 0;
	let round: Round = game.rounds[roundNum];

	$: allCluesUsed = round.categories.every((c) => c.clues.every((q) => q.answered));
	$: if (allCluesUsed) {
		console.log('All clues used!');
		nextRound();
	} else {
		console.log('Is this getting called????');
	}

	function nextRound() {
		roundNum++;
		console.log(roundNum);
		if (roundNum < game.rounds.length) {
			round = game.rounds[roundNum];
		} else {
			if (game.finalRound) {
				finalRound = true;
			} else {
              gameOver = true;
			}
		}
	}

</script>

<div class="game-view">
	{#if gameOver}
		<GameOverView winner={players.reduce((winner, player) => winner.score > player.score ? winner : player)} />
	{:else if game.finalRound && finalRound}
		<FinalRoundView clue={game.finalRound} />
	{:else}
		<BoardView {round} />
	{/if}
	<PlayerBarView {players} />
</div>

<style>
	.game-view {
		width: 100vw;
		height: 100vh;
		display: grid;
		grid-template-rows: 2fr 1fr;
	}
</style>
