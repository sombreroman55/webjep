<script lang="ts">
	import type { Round } from '$lib/types';
	import CategoryView from './CategoryView.svelte';
	import ClueView from './ClueView.svelte';
	import { currentClue, clueActive } from '$lib/gameStore';

	export let round: Round;
    const gridStyle = `display:grid; grid-template-columns: repeat(${round.categories.length}, 1fr);`;
</script>

<div>
	{#if $clueActive}
		<ClueView clue={$currentClue} />
	{:else}
    <div class="game-board" style={gridStyle}>
		{#each round.categories as category}
			<CategoryView {category} />
		{/each}
  </div>
	{/if}
</div>

<style>
	.game-board {
		width: 100%;
		height: 100%;
	}
</style>
