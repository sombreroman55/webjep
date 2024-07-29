<script lang="ts">
	import type { Category, Clue } from '$lib/types';
	import { currentClue, clueActive } from '$lib/gameStore';
    import "$lib/assets/common.css";

	function setCurrentClue(clue: Clue) {
		clueActive.set(true);
		currentClue.set(clue);
	}

	export let category: Category;
</script>

<div class="category">
	<div class="jep-blue shadow category-title box-center">
		<h2>{category.title}</h2>
	</div>
	{#each category.clues as clue}
		<div class="jep-blue shadow board-square box-center">
			{#if !clue.answered}
				<div on:click={() => setCurrentClue(clue)}>
					<b>${clue.value}</b>
				</div>
			{/if}
		</div>
	{/each}
</div>

<style>
	.category {
		display: grid;
		grid-template-rows: 1fr 1fr 1fr 1fr 1fr 1fr;
	}

	.category-title {
		color: white;
		border-right: 2px solid black;
		border-left: 2px solid black;
		border-top: 4px solid black;
		border-bottom: 4px solid black;
	}

	.board-square {
		border: 2px solid black;
		font-size: 72px;
		color: #ffcc00;
	}
</style>
