<script lang="ts">
	import type { Session } from '$lib/types';
	import FinalRoundBuilder from './FinalRoundBuilder.svelte';
	import RoundBuilder from './RoundBuilder.svelte';

	export let session: Session;

	function downloadGame() {
		var a = document.createElement('a');
		var gameData = JSON.stringify(session);
		var blob = new Blob([gameData], { type: 'application/json' });
		var url = URL.createObjectURL(blob);
		a.setAttribute('href', url);
		a.setAttribute('download', 'game.json');
		a.click();
	}
</script>

{#each session.game.rounds as round, i}
	<RoundBuilder {round} roundNumber={i + 1} />
	<hr />
{/each}
{#if session.game.finalRound}
	<FinalRoundBuilder final={session.game.finalRound} />
{/if}

<button on:click={downloadGame}>Download Clues</button>
