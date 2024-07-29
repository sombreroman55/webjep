<script lang="ts">
	import type { Game, Round, Player } from '$lib/types';
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
      return round.categories.every(c => (c.clues.every(q => q.answered)));
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
        width: 98vw;
        height: 98vh;
		display: grid;
		grid-template-rows: 4fr 1fr;
	}
</style>
