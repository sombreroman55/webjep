<script lang="ts">
	import type { Category } from '$lib/Clues';
	import ClueBuilder from './ClueBuilder.svelte';

	export let category: Category;
	export let categoryNumber: number;
</script>

<div>
  <h3>Category {categoryNumber}</h3>
  <label>Title: <input bind:value={category.title}/> </label>
  {#each category.clues as clue}
    <ClueBuilder clue={clue} />
  {/each}
</div>
