<script lang="ts">
	import type { Game, Round } from '$lib/Clues';
	import type Player from '$lib/Player';
	import BoardView from './BoardView.svelte';
	import PlayerBarView from './PlayerBarView.svelte';

	export let game: Game;
	let round: Round = game.rounds[0];
	let players: Player[] = [];

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
		grid-template-rows: 4fr 1fr;
	}
</style>
