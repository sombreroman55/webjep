<script lang="ts">
	import type { Round } from '$lib/Clues';
	import CategoryView from './CategoryView.svelte';

	export let round: Round;
</script>

<div class="game-board">
	{#each round.categories as category}
		<CategoryView {category} />
	{/each}
</div>

<style>
	.game-board {
		display: grid;
		grid-template-columns: 1fr 1fr 1fr 1fr 1fr 1fr;
	}
</style>
