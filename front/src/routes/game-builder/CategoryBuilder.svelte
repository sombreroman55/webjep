<script lang="ts">
	import type { Category } from '$lib/types';
	import ClueBuilder from './ClueBuilder.svelte';

	export let roundNumber: number;
	export let categoryNumber: number;
	export let category: Category;
</script>

<div>
	<h3>Category {categoryNumber}</h3>
	<label
		>Title: <input
			name="r{roundNumber}-c{categoryNumber}-title"
			bind:value={category.title}
		/></label
	>
	{#each category.clues as clue, i}
		<ClueBuilder {roundNumber} {categoryNumber} clueNumber={i} {clue} />
	{/each}
</div>
