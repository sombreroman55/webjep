<script lang="ts">
	import type { Game, Round } from '$lib/Clues';
	import type Player from '$lib/Player';
	import BoardView from './BoardView.svelte';
	import PlayerBarView from './PlayerBarView.svelte';

	export let game: Game;
	let round: Round = game.rounds[0];
	let players: Player[] = [
		{ name: 'fart', score: 123 },
		{ name: 'shart', score: 234 },
		{ name: 'mart', score: 345 }
	];

	function nextRound() {
		round = game.rounds[1];
	}
</script>

<div class="game-view">
	<BoardView {round} />
	<PlayerBarView {players} />
</div>

<style>
	.game-view {
		display: grid;
		grid-template-rows: 80% 20%;
	}
</style>
