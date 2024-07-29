<script lang="ts">
	import type { Game } from '$lib/types';
	import FinalRoundBuilder from './FinalRoundBuilder.svelte';
	import RoundBuilder from './RoundBuilder.svelte';

	export let game: Game;

	function downloadGame() {
		var a = document.createElement('a');
		var gameData = JSON.stringify(game);
		var blob = new Blob([gameData], { type: 'application/json' });
		var url = URL.createObjectURL(blob);
		a.setAttribute('href', url);
		a.setAttribute('download', 'clues.json');
		a.click();
	}
</script>

{#each game.rounds as round, i}
	<RoundBuilder {round} roundNumber={i + 1} />
	<hr />
{/each}
<FinalRoundBuilder final={game.finalRound} />
<button on:click={downloadGame}>Download Clues</button>
