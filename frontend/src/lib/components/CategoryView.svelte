<script lang="ts">
	import type { Category } from '$lib/Clues';
	// import ClueView from "./ClueView.svelte";

	// TODO: When you click a clue value, open a Clue View that covers the screen

	export let category: Category;
</script>

<div class="category">
	<div>{category.title}</div>
	{#each category.clues as clue}
		<div>${clue.value}</div>
	{/each}
</div>

<style>
	.category {
		display: grid;
	}
</style>
