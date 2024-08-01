<script lang="ts">
	import type { Category, Clue } from '$lib/types';
	import { currentClue, clueActive } from '$lib/gameStore';
	import '$lib/assets/common.css';

	function setCurrentClue(clue: Clue) {
		currentClue.set(clue);
		clueActive.set(true);
	}

	export let category: Category;
	const gridStyle = `display: grid; grid-template-rows: repeat(${category.clues.length + 1}, 1fr);`;
</script>

<div class="category" style={gridStyle}>
	<div class="jep-blue category-title box-center">
		{category.title.toUpperCase()}
	</div>
	{#each category.clues as clue}
		{#if !clue.answered}
			<div class="jep-blue shadow board-square box-center" on:click={() => setCurrentClue(clue)}>
				<b>${clue.value}</b>
			</div>
		{:else}
			<div class="jep-blue board-square"></div>
		{/if}
	{/each}
</div>

<style>
	@font-face {
		font-family: 'swiss-911';
		font-style: normal;
		font-weight: normal;
		src: url('/fonts/swiss911/swiss911.woff') format('woff');
	}

	.category {
		font-family: 'swiss-911';
	}

	.category-title {
		color: white;
		font-size: 48px;
		border-right: 2px solid black;
		border-left: 2px solid black;
		border-top: 4px solid black;
		border-bottom: 4px solid black;
	}

	.board-square {
		border: 2px solid black;
		letter-spacing: 2px;
		font-size: 72px;
		color: #ffcc00;
	}
</style>
