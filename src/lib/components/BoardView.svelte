<script lang="ts">
	import type { Round } from '$lib/Clues';
	import CategoryView from './CategoryView.svelte';
    import ClueView from './ClueView.svelte';
	import { currentClue, clueActive } from '$lib/gameStore';

	export let round: Round;
</script>

<div class="game-board">
	{#if $clueActive}
      <ClueView clue={$currentClue} />
    {:else}
		{#each round.categories as category}
			<CategoryView {category} />
		{/each}
	{/if}
</div>

<style>
	.game-board {
        width: 100%;
        height: 100%;
		display: grid;
		grid-template-columns: 1fr 1fr 1fr 1fr 1fr 1fr;
	}
</style>
