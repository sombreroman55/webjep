<script lang="ts">
	import type { Round } from '$lib/Clues';
	import CategoryBuilder from './CategoryBuilder.svelte';

	export let round: Round;
	export let roundNumber: number;
</script>

<div>
  <h2>Round {roundNumber}</h2>
  {#each round.categories as category, i}
    <CategoryBuilder category={category} roundNumber={roundNumber} categoryNumber={i+1} />
  {/each}
</div>
