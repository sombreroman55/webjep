<script lang="ts">
	import type { Game, Round } from '$lib/Clues';
	import type Player from '$lib/Player';
	import BoardView from './BoardView.svelte';
	import PlayerBarView from './PlayerBarView.svelte';

	export let game: Game;
	export let players: Player[];
	let finalRound: boolean = false;
	let roundNum = 0;
	let round: Round = game.rounds[0];

	$: if (noCluesLeft()) {
		nextRound();
	}

	function noCluesLeft(): boolean {
		return true;
	}

	function nextRound() {
		roundNum++;
		if (roundNum < 2) round = game.rounds[roundNum];
		else finalRound = true;
	}
</script>

<div class="game-view">
	{#if !finalRound}
		<BoardView {round} />
	{:else}
		<BoardView {round} />
	{/if}
	<PlayerBarView {players} />
</div>

<style>
	.game-view {
		display: grid;
		grid-template-rows: 4fr 1fr;
	}
</style>
